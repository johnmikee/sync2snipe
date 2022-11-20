import argparse
import json

from concurrent.futures import ThreadPoolExecutor
from types import SimpleNamespace

from tosnipe import ToSnipe
from jamf import UAPI, Classic


VALID_SUBSET = [
    "general",
    "location",
    "purchasing",
    "peripherals",
    "hardware",
    "certificates",
    "software",
    "extension_attributes",
    "groups_accounts",
    "iphones",
    "configuration_profiles",
]

COMPUTER_INFO = {}
MOBILE_INFO = {}


class Jamf2Snipe(ToSnipe):
    def __init__(self, env_vars=False):
        self.env_vars = env_vars
        self.args = self._get_jamf2_args()
        super().__init__()

        self.config = self.set_conf()
        self.jamf = Classic(ci=self.env_vars, level=self.level)
        self.jamf_uapi = UAPI(ci=self.env_vars, level=self.level)
        self.start_up_test = self.run_tests()

    def _get_jamf2_args(self) -> argparse.Namespace:
        runtimeargs = self._get_snipe_args()

        runtimeargs.add_argument(
            "--auto_incrementing",
            help="You can use this if you have auto-incrementing enabled in your snipe instance to utilize that instead of adding the Jamf ID for the asset tag.",
            action="store_true",
        )
        runtimeargs.add_argument(
            "--do_not_update_jamf",
            help="Does not update Jamf with the asset tags stored in Snipe.",
            action="store_false",
        )
        user_opts = runtimeargs.add_mutually_exclusive_group()
        user_opts.add_argument(
            "-u",
            "--users",
            help="Checks out the item to the current user in Jamf if it's not already deployed",
            action="store_true",
        )
        user_opts.add_argument(
            "-ui",
            "--users_inverse",
            help="Checks out the item to the current user in Jamf if it's already deployed",
            action="store_true",
        )
        user_opts.add_argument(
            "-uf",
            "--users_force",
            help="Checks out the item to the user specified in Jamf no matter what",
            action="store_true",
        )
        runtimeargs.add_argument(
            "-uns",
            "--users_no_search",
            help="Doesn't search for any users if the specified fields in Jamf and Snipe don't match. (case insensitive)",
            action="store_true",
        )
        type_opts = runtimeargs.add_mutually_exclusive_group()
        type_opts.add_argument(
            "-m",
            "--mobiles",
            help="Runs against the Jamf mobiles endpoint only.",
            action="store_true",
        )
        type_opts.add_argument(
            "-c",
            "--computers",
            help="Runs against the Jamf computers endpoint only.",
            action="store_true",
        )

        return runtimeargs.parse_args()

    def set_conf(self) -> SimpleNamespace:
        """
        check the config file for correct mappings and
        do some tests to see if the old settings.conf file is in use
        """
        self.config = self.validate_conf()

        SETTINGS_CORRECT = True
        if hasattr(self.config, "api-mapping"):
            self.log.error(
                """
                Looks like you're using the old method for api-mapping.
                Please use computers_mapping and mobile_mapping.
                """
            )
            SETTINGS_CORRECT = False

        if not hasattr(self.config, "user_mapping") and (
            self.args.users or self.args.users_force or self.args.users_inverse
        ):
            self.log.error(
                """
                You've chosen to check out assets to users in some capacity using a cmdline switch,
                but not specified how you want to search Snipe IT for the users from Jamf.
                Make sure you have a 'user-mapping' section in your settings.conf file.
                """
            )
            SETTINGS_CORRECT = False

        if not SETTINGS_CORRECT:
            raise SystemExit("Invalid settings. Please validate and try again.")

        # Check the config file for valid jamf subsets.
        # This is based off the JAMF API and if it's not right we can't map fields over to SNIPE properly.
        self.log.debug(
            "Checking the settings.conf file for valid JAMF subsets of the JAMF API so mapping can occur properly."
        )

        valid_mappings = [
            i for i in self.config.computers_mapping.name if i in VALID_SUBSET
        ]
        if valid_mappings:
            self.log.info(f"Found subset {valid_mappings}: Acceptable")
        else:
            self.log.error(
                f"""
                Found invalid subset: {self.config.computers_mapping.name} in the settings.conf file.
                This is not in the acceptable list of subsets - check your settings.conf.

                Valid subsets are: {', '.join(VALID_SUBSET)}
                """
            )
            raise SystemExit("Invalid Subset found in settings.conf")

        return self.validate_conf()

    def run_tests(self) -> bool:
        if self.args.run_tests:
            # Report if we're verifying SSL or not.
            self.log.info(f"SSL Verification is set to: {self.args.do_not_verify_ssl}")

            # Do some tests to see if the hosts are up.
            self.log.info("Running tests to see if hosts are up.")
            try:
                SNIPE_UP = (
                    True
                    if self.snipe.requester(
                        method="GET",
                        url=self.snipe.base_url,
                        override_url=True,
                        raw=True,
                    ).status_code
                    == 200
                    else False
                )
            except Exception as e:
                self.log.exception(e)
                SNIPE_UP = False

            """
            some self hosted servers may not respond to this type of request.
            in the config we can pass the option to skip the tests. long term,
            this should be refactored to a different test to address these
            edge cases instead of skipping it.
            """
            JAMF_UP = True
            if hasattr(self.config.jamf, "skip_test"):
                if self.config.jamf.skip_test is not True:
                    try:
                        JAMF_UP = (
                            True
                            if self.jamf.requester(
                                method="GET",
                                url=self.jamf.base_url,
                                override_url=True,
                                raw=True,
                            ).status_code
                            in (200, 401)
                            else False
                        )
                    except Exception as e:
                        self.log.exception(e)
                        JAMF_UP = False

            if not SNIPE_UP:
                self.log.error(
                    "Snipe-IT looks like it is down from here. Please check your config in the settings.conf file, or your instance."
                )
                return False

            self.log.info(
                "We were able to get a good response from your Snipe-IT instance."
            )

            if not JAMF_UP:
                self.log.error(
                    """
                    JAMFPro looks down from here.
                    Please check the your config in the settings.conf file, or your hosted JAMFPro instance.
                    """
                )
                return False

            self.log.info(
                "We were able to get a good response from your JAMFPro instance."
            )

            ##TODO: Write actual tests.
            self.log.info("Finished running our tests.")

        return True

    def _attribute_adder(self, asset_info: dict, asset_type: str) -> dict:
        attributes = {}
        if asset_type == "computer":
            mapping = self.config.computers_mapping.__dict__
        if asset_type == "mobile":
            mapping = self.config.mobile_mapping.__dict__

        for name, values in mapping.items():
            self.log.debug(f"checking for match on {name}")
            for i, item in enumerate(values):
                # key off the top level value.
                # thats where the remaining items in the list will potentially be
                if i == 0:
                    jamf_value = asset_info[item]

                else:
                    if values[0] == "extension_attributes":
                        for attribute in jamf_value:
                            if attribute["id"] == item:
                                jamf_value = attribute["value"]
                                self.log.debug(f"match on {item}: {jamf_value}")
                    else:
                        self.log.debug(item)
                        self.log.debug(jamf_value)
                        jamf_value = jamf_value[item]
                        self.log.debug(f"match on {item}: {jamf_value}")

                    attributes[item] = jamf_value

        return attributes

    def _gather_all_jamf_machine_info(
        self, machine_id: int, machine_type: str, update_dict: str
    ) -> None:

        valid_opts = {
            "computer": {
                "endpoint": "computers",
                "device_type": "computer",
                "top_level_key": "hardware",
                "report_time": "report_date",
            },
            "mobile": {
                "endpoint": "mobiledevices",
                "device_type": "mobile_device",
                "top_level_key": "general",
                "report_time": "last_inventory_update",
            },
        }
        if machine_type.lower() not in valid_opts:
            self.log.debug(f"{machine_type.lower()} not in {list(valid_opts.keys())}")
            return

        res = self.jamf.get_device_info(
            machine_id=machine_id,
            device_type=valid_opts[machine_type]["device_type"],
            endpoint=valid_opts[machine_type]["endpoint"],
        )

        # computers and mobile store the model, model_identifier, and inventory update time in different places.
        # figure out which we are querying and set the key based off that.
        top_level_key = valid_opts[machine_type]["top_level_key"]
        report_time_key = valid_opts[machine_type]["report_time"]
        # add the keys we need access to routinely.
        update_dict[machine_id] = {
            "model_identifier": res[top_level_key]["model_identifier"],
            "model": res[top_level_key]["model"],
            "asset_tag": res["general"]["asset_tag"],
            "name": res["general"]["name"],
            "serial": res["general"]["serial_number"],
            "report_date": res["general"][report_time_key],
        }

        # add the remaining top level keys from the jamf response.

        [
            update_dict[machine_id].update({item: res[item]})
            for item in res
            if not update_dict[machine_id].get(item)
        ]

    def _sync_asset_tags(
        self, asset_tag: str, machine_id: str, machine_type: str, serial: str
    ) -> None:
        res = self.jamf.update_asset_tag(
            asset_tag=asset_tag,
            machine_id=machine_id,
            machine_type=machine_type,
        )
        self.log.debug(f"response from updating {serial} in jamf: {res.text}")

    def asset_tag_sync(
        self,
        existing_snipe_serials: list,
        machine_type: str,
        snipe_machines: dict,
        update_dict: dict,
    ) -> list[dict]:
        """
        Update/Sync the Snipe Asset Tag Number back to jamf
        The user arg below is set to false if it's called, so this would fail if the user called it.
        """
        updates = []
        for serial in existing_snipe_serials:
            snipe_machine = [i for i in snipe_machines if serial == i["serial"]][0]
            asset_info = {k: v for k, v in update_dict.items() if v["serial"] == serial}

            if not asset_info.keys():
                continue

            asset_info = asset_info[next(iter(asset_info))]

            if (
                asset_info["asset_tag"] != snipe_machine["asset_tag"]
            ) and self.args.do_not_update_jamf:
                self.log.info(
                    f"jamf doesn't have the same asset tag, {asset_info['asset_tag']}, as snipe, {snipe_machine['asset_tag']}, so we'll update it."
                )
                if snipe_machine["asset_tag"]:
                    if self.args.dryrun:
                        self.log.info(
                            f"would be updating {serial} in jamf with the asset tag {asset_info['asset_tag']} => {snipe_machine['asset_tag']}"
                        )
                        continue
                    updates.append(
                        {
                            "asset_tag": snipe_machine["asset_tag"],
                            "machine_id": asset_info["general"]["id"],
                            "serial": serial,
                        }
                    )

        with ThreadPoolExecutor(max_workers=10) as executor:
            for update in updates:
                executor.submit(
                    self._sync_asset_tags,
                    update["asset_tag"],
                    update["machine_id"],
                    machine_type,
                    update["serial"],
                )

    def existing_update_check(
        self, asset_type: str, existing_snipe_serials: list, snipe_machines: dict
    ) -> dict[int:[dict]]:
        self.log.info("checking which machines in snipe need to be updated")

        updates = {}
        for serial in existing_snipe_serials:
            snipe_machine = [i for i in snipe_machines if serial == i["serial"]][0]
            asset_info = {
                k: v for k, v in COMPUTER_INFO.items() if v["serial"] == serial
            }

            if not asset_info.keys():
                continue

            snipe_id = snipe_machine["id"]
            updates[snipe_id] = {"serial": serial, "payload": {}}

            asset_info = asset_info[next(iter(asset_info))]
            jamf_update_time = asset_info["report_date"]
            snipe_update_time = snipe_machine["updated_at"]["datetime"]

            if self.args.force:
                self.log.debug(
                    f"forcing the update on {serial} regardless of the timestamps below."
                )
            if jamf_update_time > snipe_update_time:
                self.log.debug(
                    f"Updating {serial} in snipe because jamf has a more recent timestamp: {jamf_update_time} > {snipe_update_time}"
                )
            try:
                payload = self._attribute_adder(
                    asset_info=asset_info, asset_type=asset_type
                )
                self.log.debug(f"generated payload: {payload}")
            except Exception:
                self.log.debug(
                    "Skipping the payload, because the jamf key we're mapping to doesn't exist"
                )
                continue

            # Need to check that we're not needlessly updating the asset.
            # If it's a custom value it'll fail the first section and send it to except section that will parse custom sections.
            try:
                for key, val in payload.items():
                    if snipe_machine[key] != val:
                        self.log.debug(f"{snipe_machine[key]} != {val}")
                        updates[snipe_id]["payload"].update(payload)

                    else:
                        self.log.debug(
                            "Skipping the payload, because it already exits."
                        )
            except:
                self.log.debug(
                    "the key lookup failed, which means it's a custom field. parsing those to see if it needs to be updated or not."
                )
                needs_update = False
                for key, value in payload.items():
                    for custom_field in snipe_machine["custom_fields"]:
                        if snipe_machine["custom_fields"][custom_field]["field"] == key:
                            if snipe_machine["custom_fields"][custom_field][
                                "value"
                            ] != str(value):
                                self.log.debug(
                                    f"found the field, and the value needs to be updated from {snipe_machine['custom_fields'][custom_field]['value']} to {value}"
                                )
                                needs_update = True
                    if needs_update:
                        updates[snipe_id]["payload"].update(payload)
                    else:
                        self.log.debug(
                            "Skipping the payload because it already exists or the snipe key we're mapping to doesn't."
                        )

        return updates

    def generate_asset_tag(
        self, asset_type: str, serials: str, update_dict: dict
    ) -> list[dict]:
        self.log.info("checking what we need to generate asset tags for")

        new_assets = []

        for serial in serials:
            asset_info = self._asset_by_serial(serial=serial, machine_dict=update_dict)
            # there is only one key in this dictionary.
            # key off that to get the value dict info.
            if not asset_info.keys():
                continue

            asset_info = asset_info[next(iter(asset_info))]

            asset_tag = asset_info["asset_tag"]
            if asset_tag == "":
                asset_tag = None
                self.log.debug(
                    f"no asset tag found in jamf for {serial}, checking settings.conf for alternative specified field."
                )
                if hasattr(self.config.snipe_it, "asset_tag"):
                    # this needs looking at
                    asset_tag = self.config.snipe_it.asset_tag

                # strings are falsy so this will work for None or ""
                if not asset_tag:
                    if asset_type == "computer":
                        asset_tag = f'jamfid-{[k for k, v in COMPUTER_INFO.items() if v["serial"] == serial][0]}'
                    if asset_type == "mobile":
                        asset_tag = f'jamfid-m-{[k for k, v in MOBILE_INFO.items() if v["serial"] == serial][0]}'
            else:
                self.log.info(f"Asset tag found in Jamf, setting it to: {asset_tag}")

            new_assets.append({"info": asset_info, "tag": asset_tag})

        return new_assets

    def get_active_ids(self) -> tuple[dict, dict]:
        """
        Get the IDS of all active assets.
        """
        self.log.info("Getting a list of all active ids.")

        if (
            hasattr(self.config.jamf, "computer_group_id")
            and self.config.jamf.computer_group_id
        ):
            self.log.info("Getting list of computers from jamf by computer group id.")
            jamf_computers = self.jamf.get_computers_by_group(
                group_id=self.config.jamf.computer_group_id
            )
        else:
            jamf_computers = self.jamf.get_all_machines()

        if (
            hasattr(self.config.jamf, "mobile_group_id")
            and self.config.jamf.mobile_group_id
        ):
            self.log.info("Getting list of mobiles from jamf by mobile group id.")
            jamf_mobiles = self.jamf.get_mobile_by_group(
                self.config.jamf.mobile_group_id
            )
        else:
            jamf_mobiles = self.jamf_uapi.get_mobile_devices()

        return jamf_computers, jamf_mobiles

    def prepare_asset_checkout(
        self, assets: list, snipe_machines: dict, update_dict: dict
    ) -> list[dict]:

        mapping = self.config.user_mapping.jamf_api_field

        if len(mapping) != 2:
            self.log.error("the jamf user mapping must be a list with two values.")
            # prune this list?
            self.log.debug(
                f"acceptable keys and values are: {json.dumps(self.jamf.api_values, sort_keys=True, indent=4)}"
            )
            return None

        key = mapping[0]
        value = mapping[1]

        asset_updates = []

        for asset in assets:
            asset_info = self._asset_by_serial(
                serial=asset["serial"], machine_dict=update_dict
            )
            snipe_machine = [
                i for i in snipe_machines if asset["serial"] == i["serial"]
            ]

            if (
                not (
                    (self.args.users or self.args.users_inverse)
                    and (snipe_machine["assigned_to"] is None) == self.args.users
                )
                or self.args.users_force
            ):
                return None

            if snipe_machine["status_label"]["status_meta"] in (
                "deployable",
                "deployed",
            ):
                if asset_info.get(key):
                    if value not in asset_info[key]:
                        self.log.info(
                            f"{value} not present for the device in {asset_info['all'][key]}. not checking it out."
                        )
                        continue

                    if self.args.dryrun:
                        self.log.info(
                            f"would be checking out new item {asset_info['general']['name']} to user {asset_info[key][value]}"
                        )
                        continue

                    asset_updates.append(
                        {
                            "asset_id": self.created_assets[asset["serial"]],
                            "status_id": self.snipe.default_status,
                            "checkout_to_type": "user",
                            "assigned_user": asset_info[key][value],
                        }
                    )
                    self.log.info(
                        f"checking out new item {asset_info['general']['name']} to user {asset_info[key][value]}"
                    )

        self.log.debug(f"asset updates: {asset_updates}")

        return asset_updates

    def update_machines(self, machines: list[dict], machine_type: str):
        update_type = {
            "computer": COMPUTER_INFO,
            "mobile": MOBILE_INFO,
        }
        update_dict = update_type[machine_type]

        machine_ids = [m["id"] for m in machines]

        # quickly grab all the machine info so we dont have to loop through each machine
        self.log.info(f"Getting all info for {machine_type}'s from jamf")
        with ThreadPoolExecutor(max_workers=10) as executor:
            for mid in machine_ids:
                executor.submit(
                    self._gather_all_jamf_machine_info, mid, machine_type, update_dict
                )

        # create a dictionary of model info
        machine_identifiers = {
            update_dict[i]["model_identifier"]: update_dict[i]["model"]
            for i in update_dict
        }
        # grab whats already in snipe
        existing_identifiers = list(self.snipe.model_numbers.keys())
        # see whats missing
        diff = [
            i for i in list(machine_identifiers.keys()) if i not in existing_identifiers
        ]
        # create what needs to be
        self.create_new_snipe_models(models=diff, model_info=machine_identifiers)

        # get a list of existing snipe machines
        snipe_machines = self.snipe.get_all_hardware()

        # make sure they are unique
        existing_snipe_serials = list(set([i["serial"] for i in snipe_machines]))

        self.log.debug(f"following machines exist in snipe: {existing_snipe_serials}")
        jamf_serials = [update_dict[i]["serial"] for i in update_dict]
        self.log.debug(f"following machines exist in jamf: {jamf_serials}")

        # find out which are missing
        missing_machines = [
            i for i in jamf_serials if i not in existing_snipe_serials and len(i)
        ]
        if not missing_machines:
            self.log.info("all machines in jamf exist in snipe")
        else:
            self.log.info(
                f"following machines need to be created in snipe: {missing_machines}"
            )
        # generate, update, and create new assets.
        new_assets = self.generate_asset_tag(
            serials=missing_machines, asset_type=machine_type, update_dict=update_dict
        )
        new_asset_tags = self.update_new_asset_info(
            asset_type=machine_type, new_assets=new_assets
        )
        self.create_new_asset(new_assets=new_asset_tags)
        # get the asset tags sync'd up
        self.asset_tag_sync(
            existing_snipe_serials=existing_snipe_serials,
            snipe_machines=snipe_machines,
            machine_type=machine_type,
            update_dict=update_dict,
        )

        # checkout to users if args specify it
        prepared_assets = self.prepare_asset_checkout(
            assets=new_asset_tags,
            snipe_machines=snipe_machines,
            update_dict=update_dict,
        )
        self.checkout_assets(assets=prepared_assets)

        # only update the existing machines if jamf has more recent info
        updates = self.existing_update_check(
            asset_type=machine_type,
            existing_snipe_serials=existing_snipe_serials,
            snipe_machines=snipe_machines,
        )
        self.update_existing_snipe(updates=updates)


if __name__ == "__main__":
    j2s = Jamf2Snipe()
    computers, mobiles = j2s.get_active_ids()
    j2s.update_machines(machines=computers, machine_type="computer")
    j2s.update_machines(machines=mobiles, machine_type="mobile")
