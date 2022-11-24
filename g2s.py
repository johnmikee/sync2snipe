import argparse

from types import SimpleNamespace

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

    def get_chromeos_devices(self) -> dict:
        # get some values
        return self.goog.filter_chromeos_devices(
            opts=self.config.google.chromeos_attributes
        )

    def validate_asset_tags(self):
        pass

    def update_machines(self, machines: dict):
        # get all the serials
        machine_serials = [machines[i]["serialNumber"] for i in machines]

        # get all the models
        machine_identifiers = list(set([machines[i]["model"] for i in machines]))
        self.log.info(machine_identifiers)
        # grab whats already in snipe
        existing_identifiers = list(self.snipe.model_numbers.keys())
        self.log.info(existing_identifiers)
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
        for item in diff:
            contains = [
                i for i in existing_manufacturers.values() if item.split(" ")[0] in i
            ]
            if contains:
                self.log.info(contains)

        self.log.info(diff)
        self.log.info(existing_manufacturers)


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
g.update_machines(machines=res)
