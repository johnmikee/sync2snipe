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
    sync chromeos devices to snipe and optionally write the asset tags back
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
        runtimeargs.add_argument(
            "--do-not-update-google",
            help="Does not update Google with the asset tags stored in Snipe.",
            action="store_false",
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
        asset_info = [i for i in machine_dict if i["serial"] == serial]
        if asset_info:
            return asset_info[0]

        return asset_info

    def create_new_google_asset(self, new_assets: list) -> None:
        for asset in new_assets:
            if self.args.dryrun:
                self.log.info(f"would be creating a new asset: {asset}")
                continue

            self.log.info(f"creating a new asset: {asset}")
            res = self.snipe.create_asset(
                asset_tag=asset["annotatedAssetId"],
                status_id=self.snipe.default_status,
                serial=asset["serialNumber"],
                model_id=self.snipe.model_numbers[asset["model"]],
            )
            self.log.debug(
                f"response from creating asset: {res['messages']} : {res['payload']}"
            )
            if res.get("status") == "success":
                self.created_assets[asset["serialNumber"]] = res["payload"]["id"]

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

            res = self.snipe.create_manufacturer(name=manufacturer)
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
                    m
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
                        category_id = {self.config.google.manufacturer_mapping.google.category_id}
                        manufacturer_id = {manufacturer_id}
                        """
                        )
                        continue

                    res = self.snipe.create_models(
                        name=model,
                        model_number=model,
                        category_id=self.config.google.manufacturer_mapping.google.category_id,
                        manufacturer_id=manufacturer_id,
                    )

                    self.log.info(f"response from creating {model}: {res}")
                    if res.get("status") == "success":
                        self.snipe.model_numbers[model] = res["payload"]["id"]
            else:
                self.log.debug(f"need to create manufacturer for {model}")

    def generate_asset_tags(
        self, asset_type: str, missing_machines: dict
    ) -> list[dict]:
        """
        the device ID in chrome is a fairly long string.
        we take the last 12 of the ID for the asset tag.
        """
        self.log.info("checking what we need to generate asset tags for")

        new_assets = []
        for device in missing_machines:
            if hasattr(self.config.snipe_it, "asset_tag"):
                # this needs looking at
                asset_tag = self.config.snipe_it.asset_tag

            else:
                if asset_type == "chrome_os":
                    asset_tag = f'google-chromeos-id-{device.split("-")[-1]}'

            self.log.debug(
                f"adding asset tag {asset_tag} to {missing_machines[device]['serialNumber']}"
            )
            missing_machines[device].update({"annotatedAssetId": asset_tag})

            new_assets.append(missing_machines[device])

        return new_assets

    def update_google_tags(self, update_tags: list[dict]):
        """
        format the data how the google api wants it
        """
        for update in update_tags:
            if not self.args.do_not_update_google:
                if self.args.dryrun:
                    self.log.debug(f"would be updating asset tag on {update['serial']}")
                    continue

                self.goog.patch_chromeos_device(
                    device_id=update["machine_id"], annotatedAssetId=update["asset_tag"]
                )

            self.log.debug(
                f"not updating {update['serial']} because the args said nope"
            )

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
        self.log.debug(f"following machines exist in google: {google_serials}")

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
        new_assets = self.generate_asset_tags(
            asset_type=asset_type, missing_machines=updates
        )
        # create missing machines
        self.create_new_google_asset(new_assets=new_assets)

        # get info to sync asset tags
        tag_updates = self.asset_tag_sync(
            asset_key="annotatedAssetId",
            existing_snipe_serials=existing_snipe_serials,
            serial_key="serialNumber",
            snipe_machines=snipe_machines,
            source="google",
            update_dict=machines,
        )

        # sync time
        self.update_google_tags(tag_updates)

    def check_users(self, snipe_assets: dict, google_assets: dict) -> None:
        """
        assets: dictionary of all snipe machines
        update_dict: dictionary of all google machines
        """
        asset_updates = {
            "force_update": [],
            "no_user": [],
            "user_mismatch": [],
        }

        for item in google_assets:
            snipe_machine = google_assets[item]["serialNumber"]
            asset_info = self._asset_by_serial(
                serial=snipe_machine, machine_dict=snipe_assets
            )

            if not asset_info:
                self.log.debug(f"could not get info in snipe for {snipe_machine}")
                continue

            if asset_info["status_label"]["status_meta"] not in [
                "deployable",
                "deployed",
            ]:
                self.log.info(
                    f'{snipe_machine} is not in a state we can checkout or update: {asset["status_label"]["status_meta"]}'
                )
                continue

            if google_assets[item][self.config.google.user_key] is None:
                # there are cases when the annotatedUser is None. For this,
                # we fall back to recent users.
                #
                # there can be multiple users in recentUsers - I dont have a
                # good solution for which one is the right one, so we default
                # to the first in the list.
                if google_assets[item]["recentUsers"][0]["email"] is None:
                    self.log.debug(f"cannot find user for {snipe_machine}")
                    continue
                else:
                    user = google_assets[item]["recentUsers"][0]["email"]
            else:
                user = google_assets[item][self.config.google.user_key]

            google_user = self._user_email_validator(user)

            # check if the user exists
            ok, snipe_uid = self.user_checker(
                google_user,
                create=False,
            )

            if not ok:
                self.log.info(
                    f"could not get user id for {google_user} in snipe - cannot check out the asset"
                )
                continue

            self.log.debug(f"{google_user} exists in snipe with ID-{snipe_uid}")

            update_info = {
                "asset_id": asset_info["id"],
                "status_id": self.snipe.default_status,
                "checkout_to_type": "user",
                "assigned_user": snipe_uid,
                "serial": snipe_machine,
            }

            if self.args.users_force:
                self.log.debug("forcing user sync")
                asset_updates["force_update"].append(update_info)
                continue

            if asset_info["assigned_to"] is None:
                self.log.debug(
                    f'{snipe_machine} not checked out to anyone: {asset_info["assigned_to"]}'
                )
                asset_updates["no_user"].append(update_info)
                continue

            if asset_info["assigned_to"]["username"] != google_user:
                self.log.debug(
                    f'{snipe_machine} is checked out to {asset_info["assigned_to"]["username"]} in snipe and {google_user} in google'
                )
                asset_updates["user_mismatch"].append(update_info)
                continue

        self.log.debug(f"asset updates: {asset_updates}")

        if self.args.dryrun:
            for key, value in asset_updates.items():
                if value:
                    for asset in value:
                        self.log.info(
                            f"Would be checking out {asset['serial']} to user {asset['assigned_user']}. Reason: {key}"
                        )

        return asset_updates


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
user_updates = g.check_users(g.snipe.get_all_hardware(), res)
g.checkout_assets(assets=user_updates)
# g.modeler(machines=res)
# g.creator(asset_type="chrome_os", machines=res)

# g.modeler(machines=res)
