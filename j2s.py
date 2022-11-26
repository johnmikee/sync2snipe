import argparse

from concurrent.futures import ThreadPoolExecutor
from types import SimpleNamespace

from jamf import UAPI, Classic
from tosnipe import ToSnipe

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
            "--auto-incrementing",
            help="You can use this if you have auto-incrementing enabled in your snipe instance to utilize that instead of adding the Jamf ID for the asset tag.",
            action="store_true",
        )
        runtimeargs.add_argument(
            "--do-not-update-jamf",
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
            "-uf",
            "--users-force",
            help="Checks out the item to the user specified in Jamf no matter what",
            action="store_true",
        )
        user_opts.add_argument(
            "-ui",
            "--users-inverse",
            help="Checks out the item to the current user in Jamf if it's already deployed",
            action="store_true",
        )
        user_opts.add_argument(
            "-uae",
            "--users-are-emails",
            help="This flag will append your domain to the username.",
            action="store_true",
        )
        runtimeargs.add_argument(
            "-ued",
            "--users-email-domain",
            help="Domain to append onto user names.",
            default="",
            type=str,
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
        check the config file for correct mappings
        """
        self.config = self.validate_conf()

        SETTINGS_CORRECT = True

        if not hasattr(self.config, "jamf"):
            self.log.debug(
                """
                No jamf config specified. This may be fine depending on flags passed.
                """
            )
            if self.args.users or self.args.users_force or self.args.users_inverse:
                self.log.error(
                    """
                    You've chosen to check out assets to users in some capacity using a cmdline switch,
                    but not specified how you want to search Snipe IT for the users from Jamf.
                    Make sure you have a 'user_mapping' section in your settings file under the jamf options.
                    """
                )
                SETTINGS_CORRECT = False

        if not SETTINGS_CORRECT:
            raise SystemExit("Invalid settings. Please validate and try again.")

        # Check the config file for valid jamf subsets.
        self.log.debug("Checking the settings.json file for valid Jamf options.")

        if self.args.users or self.args.users_force or self.args.users_inverse:
            if self.args.computers:
                try:
                    valid_mapping = [
                        self.config.jamf.computers_mapping.user_name.key in VALID_SUBSET
                    ][0]
                except AttributeError as e:
                    self.log.info(
                        f"Could not find a valid mapping for computer users: {e}"
                    )

            if self.args.mobiles:
                try:
                    valid_mapping = [
                        self.config.jamf.mobiles_mapping.user_name.key in VALID_SUBSET
                    ][0]
                except AttributeError as e:
                    self.log.info(
                        f"Could not find a valid mapping for mobile users: {e}"
                    )

            if valid_mapping:
                self.log.info(f"Found subset {valid_mapping}: Acceptable")

        if self.args.users_are_emails:
            if self.args.users_email_domain == "":
                raise Exception(
                    "If users are emails is set you must provide the domain."
                )

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
            mapping = self.config.jamf.computers_mapping
        if asset_type == "mobile":
            mapping = self.config.jamf.mobiles_mapping

        self.log.debug(f"checking for match on {mapping.key}")
        if mapping.key not in self.jamf.api_values:
            self.log.info(
                f"{mapping.key} is not a valid option. Set this to one of the following: {list(self.jamf.api_values.keys())}"
            )
            return attributes

        jamf_value = asset_info[mapping.key]

        if mapping.value == "extension_attributes":
            for attribute in jamf_value:
                if attribute.get("id") == mapping.value:
                    jamf_value = attribute["value"]
                    self.log.debug(f"match on {mapping.value}: {jamf_value}")
        else:
            jamf_value = asset_info[mapping.key][mapping.value]
            self.log.debug(f"match on {mapping.key}:{mapping.value}: {jamf_value}")

        attributes[mapping.key] = jamf_value

        return attributes

    def _user_email_validator(self, user: str) -> str:
        """
        in certain cases the jamf user will not match the snipe user
        if the user is synced from an IdP where the user login is their
        email. this will append the domain to the user string if needed.
        """
        if not user.endswith(self.args.users_email_domain):
            fixed_user = f"{user}{self.args.users_email_domain}"
            self.log.debug(f"changing {user} to {fixed_user}")

            return fixed_user

        return user

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
            except Exception as e:
                self.log.debug(
                    f"Skipping the payload, because the jamf key we're mapping to doesn't exist: {e}"
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

    def sync_jamf_tags(self, machine_type: str, updates: dict):
        with ThreadPoolExecutor(max_workers=10) as executor:
            for update in updates:
                executor.submit(
                    self._sync_asset_tags,
                    update["asset_tag"],
                    update["machine_id"],
                    machine_type,
                    update["serial"],
                )

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

    def checkout_to_users(
        self, machine_type: str, assets: list, update_dict: dict, create: bool = False
    ) -> dict[str : list[dict]]:

        # these are very likely the same mapping in most cases
        # but flexibility ya know?
        if machine_type == "computer":
            user_mapping = self.config.jamf.computers_mapping.user_name
            if create:
                full_name_mapping = self.config.jamf.computers_mapping.full_name
        if machine_type == "mobile":
            user_mapping = self.config.jamf.mobiles_mapping.user_name
            if create:
                full_name_mapping = self.config.jamf.mobiles_mapping.full_name

        asset_updates = {
            "force_update": [],
            "no_user": [],
            "user_mismatch": [],
        }

        for asset in assets:
            snipe_machine = asset
            asset_info = self._asset_by_serial(
                serial=asset["serial"], machine_dict=update_dict
            )
            if not asset_info.keys():
                continue

            if asset["status_label"]["status_meta"] not in [
                "deployable",
                "deployed",
            ]:
                self.log.info(
                    f'{asset["serial"]} is not in a state we can checkout or update: {asset["status_label"]["status_meta"]}'
                )
                continue

            asset_info = asset_info[next(iter(asset_info))]
            jamf_user = self._user_email_validator(
                asset_info[user_mapping.key][user_mapping.value]
            )

            # check if the user exists
            create_info = {}
            if create:
                """
                this assumes the real name is in a "Pied Piper" format.
                If that is not the case you may need to switch this logic out.
                """
                create_user = asset_info[full_name_mapping.key][
                    full_name_mapping.value
                ].split(" ")
                create_info = {
                    "first_name": create_user[0],
                    "last_name": create_user[1],
                }

            ok, snipe_uid = self.user_checker(
                jamf_user, create=self.args.create_missing_users, **create_info
            )
            if not ok:
                self.log.info(
                    f"could not get user id for {jamf_user} in snipe - cannot check out the asset"
                )
                continue

            self.log.debug(f"{jamf_user} exists in snipe with ID-{snipe_uid}")

            update_info = {
                "asset_id": asset["id"],
                "status_id": self.snipe.default_status,
                "checkout_to_type": "user",
                "assigned_user": snipe_uid,
                "serial": asset["serial"],
            }

            if self.args.users_force:
                self.log.debug("forcing user sync")
                asset_updates["force_update"].append(update_info)
                continue

            if snipe_machine["assigned_to"] is None:
                self.log.debug(
                    f'{snipe_machine["serial"]} not checked out to anyone: {snipe_machine["assigned_to"]}'
                )
                asset_updates["no_user"].append(update_info)
                continue

            if snipe_machine["assigned_to"]["username"] != jamf_user:
                self.log.debug(
                    f'{snipe_machine["serial"]} is checked out to {snipe_machine["assigned_to"]["username"]} in snipe and {jamf_user} in jamf'
                )
                asset_updates["user_mismatch"].append(update_info)
                continue

        self.log.debug(f"asset updates: {asset_updates}")

        if self.args.dryrun:
            for key, value in asset_updates.items():
                if value:
                    for asset in value:
                        self.log.info(
                            f"Would be checking out new item {asset['serial']} to user {jamf_user}:{snipe_uid}. Reason: {key}"
                        )

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
        # pylint: disable=R1718
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
        tag_updates = self.asset_tag_sync(
            asset_key="asset_tag",
            existing_snipe_serials=existing_snipe_serials,
            serial_key="serial",
            snipe_machines=snipe_machines,
            source="jamf",
            update_dict=update_dict,
        )
        self.sync_jamf_tags(machine_type=machine_type, updates=tag_updates)

        # checkout to users if args specify it or if there is a need to update
        prepared_assets = self.checkout_to_users(
            assets=snipe_machines,
            machine_type=machine_type,
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
    if j2s.args.computers:
        j2s.update_machines(machines=computers, machine_type="computer")
    # if j2s.args.mobiles:
    #     j2s.update_machines(machines=mobiles, machine_type="mobile")
