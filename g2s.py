import argparse

from types import SimpleNamespace
from concurrent.futures import ThreadPoolExecutor

from goog import Google
from tosnipe import ToSnipe


CHROME_OPTS = [
    "osVersion",
    "platformVersion",
    "firmwareVersion",
    "macAddress",
    "bootMode",
    "lastEnrollmentTime",
    "firstEnrollmentTime",
    "orgUnitPath",
    "orgUnitId",
    "recentUsers",
    "ethernetMacAddress",
    "activeTimeRanges",
    "tpmVersionInfo",
    "cpuStatusReports",
    "systemRamTotal",
    "systemRamFreeReports",
    "diskVolumeReports",
    "autoUpdateExpiration",
    "cpuInfo",
]


class Google2Snipe(ToSnipe):
    """
    Currently, this is a one way sync.

    Information is pulled from Google Admin and then updated in Snipe.
    """

    def __init__(self, scopes: list = [], env_vars=False):
        self.env_vars = env_vars
        self.args = self._get_google2_args()

        super().__init__()

        self.config = self.set_conf()
        self.goog = Google(
            env_vars=env_vars,
            level=None,
            scopes=scopes,
        )

    def _get_google2_args(self) -> argparse.Namespace:
        runtimeargs = self._get_snipe_args()

        runtimeargs.add_argument(
            "-cos",
            "--chrome-os",
            help="Runs against the ChromeOS devices only.",
            action="store_true",
        )

        return runtimeargs.parse_args()

    def set_conf(self) -> SimpleNamespace:
        """
        check the config file for correct mappings
        """
        config = self.validate_conf()

        if not hasattr(config, "google"):
            self.log.debug("no config passed for Google")
            # TODO: when more options are added remove this
            return config

        if hasattr(config.google, "chromeos_attributes"):
            if not any(
                i for i in config.google.chromeos_attributes if i in CHROME_OPTS
            ):
                self.log.warning(
                    f"""
                The options in the settings file for chromeos devices are not valid.
                Valid options are {CHROME_OPTS}
                """
                )

        return config

    def _asset_by_serial(self, serial: str, machine_dict: dict) -> dict:
        asset_info = {
            k: v for k, v in machine_dict.items() if v["serialNumber"] == serial
        }

        return asset_info

    def _google_asset_creator(self, asset: dict) -> None:
        self.log.info(f"creating a new asset: {asset}")
        res = self.snipe.create_asset(
            asset_tag=asset["annotatedUser"],
            status_id=self.snipe.default_status,
            model_id=self.snipe.model_numbers[asset["model"]],
        )
        self.log.debug(
            f"response from creating asset: {res['messages']} : {res['payload']}"
        )
        if res.get("status") == "success":
            self.created_assets[asset["serial"]] = res["payload"]["id"]

    def create_new_gogle_asset(self, new_assets: list) -> None:
        with ThreadPoolExecutor(max_workers=10) as executor:
            for asset in new_assets:
                if self.args.dryrun:
                    self.log.info(f"would be creating a new asset: {asset}")
                    continue
                executor.submit(self._chrome_asset_creator, asset)

    def get_chromeos_devices(self) -> dict:
        # get some values
        return self.goog.filter_chromeos_devices(
            opts=self.config.google.chromeos_attributes
        )

    def create_chromeos_manufacturers(
        self, existing: dict, manufacturers: list[str]
    ) -> dict:
        for manufacturer in manufacturers:
            manufacturer = manufacturer.split(" ")[0]
            if self.args.dryrun:
                self.log.info(f"would be creating {manufacturer}")
                continue

            res = self.snipe.create_manufacturer(manufacturer=manufacturer)
            self.log.info(f"results from creating {manufacturer}: {res}")
            if res.get("status") == "success":
                existing.update({res["payload"]["id"]: manufacturer})

        return existing

    def create_chromeos_models(
        self, create_models: list[str], existing_manufacturers: dict
    ):
        for model in create_models:
            if model.split(" ")[0] in existing_manufacturers.values():
                manufacturer_id = [
                    existing_manufacturers[m]
                    for m in existing_manufacturers
                    if existing_manufacturers[m]
                    in [model.split(" ")[0] for model in create_models]
                ][0]

                if manufacturer_id:
                    if self.args.dryrun:
                        self.log.info(
                            f"""
                        would be creating a new model:
                        name = {model}
                        model_number = {model}
                        category_id = {self.config.snipe_it.computer_model_category_id}
                        manufacturer_id = {manufacturer_id}
                        """
                        )
                        continue

                    res = self.snipe.create_models(
                        name=model,
                        model_number=model,
                        category_id=self.config.snipe_it.computer_model_category_id,
                        manufacturer_id=manufacturer_id,
                    )

                    self.log.info(f"response from creating {model}: {res}")
                    if res.get("status") == "success":
                        self.snipe.model_numbers[model] = res["payload"]["id"]

    def generate_asset_tags(self, asset_type: str, update_dict: dict) -> list[dict]:
        self.log.info("checking what we need to generate asset tags for")
        new_assets = []
        for device in update_dict:
            if hasattr(self.config.snipe_it, "asset_tag"):
                # this needs looking at
                asset_tag = self.config.snipe_it.asset_tag

            else:
                if asset_type == "chrome_os":
                    asset_tag = f"google-chromeos-id-{device}"

            self.log.debug(
                f"adding asset tag {asset_tag} to {update_dict[device]['serialNumber']}"
            )
            update_dict[device].update({"annotatedAssetId": asset_tag})
            new_assets.append(device)

        return new_assets

    def validate_asset_tags(self, machine_info: dict) -> dict:
        no_tags = []
        for machine_id, values in machine_info.items():
            # check if the item has an asset tag.
            # if none has been set nothing will be returned
            if values.get("annotatedAssetId"):
                continue

            no_tags.append({machine_id: values})

        return no_tags

    def modeler(self, machines: dict) -> None:
        # get all the models
        machine_identifiers = list(set([machines[i]["model"] for i in machines]))
        # grab whats already in snipe
        existing_identifiers = list(self.snipe.model_numbers.keys())
        # see whats missing
        diff = [i for i in machine_identifiers if i not in existing_identifiers]
        # grab what manufactureres are in Snipe
        existing_manufacturers = {
            i["id"]: i["name"] for i in self.snipe.get_manufacturers()["rows"]
        }
        # this is not full-proof logic. what I have found is that models are stored
        # in a format that is $Manufacturer Model Info. with that we will try to match the first
        # item in the model string to an existing manufacturer.
        #
        # optionally, if the settings specify manufacturer_create_assume=true we will create the
        # manufacturer based on the logic above.
        create_models = []
        create_manufacturers = []
        for item in diff:
            contains = [
                i for i in existing_manufacturers.values() if item.split(" ")[0] in i
            ]
            if contains:
                self.log.debug(f"Manufacturer for {item} exists: {contains}")
                create_models.append(item)
            else:
                if self.config.google.manufacturer_create_assume:
                    create_manufacturers.append(item)

        manufacturer_ids = self.create_chromeos_manufacturers(
            existing=existing_manufacturers, manufacturers=create_manufacturers
        )
        self.create_chromeos_models(
            create_models=create_models,
            existing_manufacturers=manufacturer_ids,
        )

    def creator(self, asset_type: str, machines: dict):
        # get all the serials
        google_serials = [machines[i]["serialNumber"] for i in machines]

        # get a list of existing snipe machines
        snipe_machines = self.snipe.get_all_hardware()

        # make sure they are unique
        # pylint: disable=R1718
        existing_snipe_serials = list(set([i["serial"] for i in snipe_machines]))

        self.log.debug(f"following machines exist in snipe: {existing_snipe_serials}")
        self.log.debug(f"following machines exist in jamf: {google_serials}")

        # find out which are missing
        missing_machines = [
            i for i in google_serials if i not in existing_snipe_serials and len(i)
        ]
        if not missing_machines:
            self.log.info("all machines in google exist in snipe")
        else:
            self.log.info(
                f"following machines need to be created in snipe: {missing_machines}"
            )

        updates = {
            i: machines[i]
            for i in machines
            if machines[i]["serialNumber"] in missing_machines
        }

        # see which machines are missing asset tags
        missing_google_tags = self.validate_asset_tags(machine_info=machines)
        # generate asset tags
        new_tags = self.generate_asset_tags(
            asset_type=asset_type, update_dict=missing_google_tags
        )
        # create missing machines
        new_machines = self.create_new_google_asset(new_assets=new_tags)


g = Google2Snipe(
    scopes=[
        "https://www.googleapis.com/auth/admin.directory.user",
        "https://www.googleapis.com/auth/admin.directory.group",
        "https://www.googleapis.com/auth/admin.directory.user.security",
        "https://www.googleapis.com/auth/admin.directory.device.mobile",
        "https://www.googleapis.com/auth/admin.directory.device.chromeos",
    ],
)

res = g.get_chromeos_devices()

print(res)
# g.modeler(machines=res)
