import argparse
import json
import os
from concurrent.futures import ThreadPoolExecutor
from types import SimpleNamespace

import configparser

from common import get_logger
from snipe import Snipe


# pylint: disable=E1101
class ToSnipe:
    def __init__(self):
        self.level = None
        self.log = self._log_leveler()
        self.apple_manufacturer_id = 1
        self.snipe = Snipe(
            env_vars=self.env_vars,
            level=self.level,
        )
        self.created_assets = {}
        self.models = self.get_models()
        self.users = self.get_users()

    def _log_leveler(self):
        if self.args.verbose:
            self.level = "INFO"
        elif self.args.debug:
            self.level = "DEBUG"
        else:
            self.level = "WARNING"

        return get_logger(
            level=self.level,
            log_file=self.args.log_to_file,
            log_location=self.args.log_file,
            disable_requests_logging=self.args.disable_requests_logging,
        )

    def _get_snipe_args(self):
        runtimeargs = argparse.ArgumentParser()
        runtimeargs.add_argument(
            "--create-missing-users",
            help="Create users if they are missing in Snipe-IT.",
            action="store_true",
        )
        runtimeargs.add_argument(
            "-d",
            "--debug",
            help="Sets logging to include additional DEBUG messages.",
            action="store_true",
        )
        runtimeargs.add_argument(
            "--disable-requests-logging",
            help="In debug logging disables the requests library logs",
            action="store_true",
        )
        runtimeargs.add_argument(
            "--do_not_verify_ssl",
            help="Skips SSL verification for all requests. Helpful when you use self-signed certificate.",
            action="store_false",
        )
        runtimeargs.add_argument(
            "--dryrun",
            help="This checks your config and tries to contact both the JAMFPro and Snipe-it instances, but exits before updating or syncing any assets.",
            action="store_true",
        )
        runtimeargs.add_argument(
            "-f",
            "--force",
            help="Updates the Snipe asset with information every time, despite what the timestamps indicate.",
            action="store_true",
        )
        runtimeargs.add_argument(
            "-l",
            "--log-to-file",
            help="Output log results to file.",
            action="store_true",
        )
        runtimeargs.add_argument(
            "-lf",
            "--log-file",
            help="location of log file. defaults to sync2snipe-$date.log.",
            default="",
            type=str,
        )
        runtimeargs.add_argument(
            "-rt",
            "--run-tests",
            help="Run startup tests to see if hosts are reachable.",
            action="store_true",
        )
        runtimeargs.add_argument(
            "-s",
            "--settings-file",
            help="location of settings file. defaults to settings.json.",
            default="settings.json",
            type=str,
        )
        runtimeargs.add_argument(
            "-v",
            "--verbose",
            help="Sets the logging level to INFO and gives you a better idea of what the script is doing.",
            action="store_true",
        )

        return runtimeargs

    def load_conf(self, conf_path: str) -> SimpleNamespace:
        with open(conf_path, "r", encoding="utf-8") as c:
            c.seek(0)
            return json.load(c, object_hook=lambda x: SimpleNamespace(**x))

    def validate_conf(self) -> configparser.ConfigParser:
        """
        Find a valid settings.conf file.
        """
        valid_opts = [
            "/opt/stuff2snipe/settings.json",
            "/etc/stuff2snipe/settings.json",
        ]

        if self.args.settings_file != "settings.json":
            if os.path.exists(self.args.settings_file):
                return self.load_conf(self.args.settings_file)

            self.log.info(f"{self.args.settings_file} does not exist. checking default")
        else:
            if os.path.exists(self.args.settings_file):
                return self.load_conf(self.args.settings_file)

        self.log.info(
            f"{self.args.settings_file} does not exist. checking remaining options."
        )

        for opt in valid_opts:
            if os.path.exists(opt):
                return self.load_conf(opt)

        raise SystemExit(self.log.info("no valid settings file could be found."))

    def _asset_by_serial(self, serial: str, machine_dict: dict) -> dict:
        asset_info = {k: v for k, v in machine_dict.items() if v["serial"] == serial}

        return asset_info

    def _asset_creator(self, asset: dict) -> None:
        self.log.info(f"creating a new asset: {asset}")
        res = self.snipe.create_asset(
            asset_tag=asset["asset_tag"],
            status_id=asset["status_id"],
            model_id=asset["model_id"],
            **{
                k: v
                for k, v in asset.items()
                if k not in ["asset_tag", "status_id", "model_id"]
            },
        )
        self.created_assets[asset["serial"]] = res["payload"]["id"]
        self.log.debug(
            f"response from creating asset: {res['messages']} : {res['payload']}"
        )

    def _checkout_snipe_assets(self, asset_id: int, assigned_user: str) -> None:
        res = self.snipe.checkout_asset(
            asset_id=asset_id,
            status_id=self.snipe.default_status,
            checkout_to_type="user",
            assigned_user=assigned_user,
        )
        self.log.debug(f"response from checking out device: {res}")

    def _model_creator(self, model: dict) -> None:
        res = self.snipe.create_models(
            name=model["name"],
            model_number=model["model_number"],
            category_id=model["category_id"],
            manufacturer_id=self.snipe.model_numbers[model["model"]],
        )
        self.log.info(f"response from creating {model['name']}: {res}")

    def _update_snipe_machines(self, serial: str, machine_id: int, **kwargs) -> None:
        res = self.snipe.patch_asset(asset_id=machine_id, **kwargs)

        self.log.debug(f"response from updating {serial}: {res}")

    def asset_tag_sync(
        self,
        asset_key: str,
        existing_snipe_serials: list,
        serial_key: str,
        snipe_machines: dict,
        source: str,
        update_dict: dict,
    ) -> list[dict]:
        """
        Update/Sync the Snipe Asset Tag Number back to the source
        """
        updates = []
        for serial in existing_snipe_serials:
            snipe_machine = [i for i in snipe_machines if serial == i["serial"]][0]
            asset_dict = {
                k: v for k, v in update_dict.items() if v[serial_key] == serial
            }

            if not asset_dict.keys():
                continue

            asset_info = asset_dict[next(iter(asset_dict))]

            if not (asset_info.get(asset_key)) or (
                asset_info.get(asset_key) != snipe_machine["asset_tag"]
            ):
                if source == "jamf":
                    machine_id = asset_info["general"]["id"]

                if source == "google":
                    machine_id = list(asset_dict.keys())[0]

                self.log.info(
                    f"{source} doesn't have the same asset tag as snipe, {snipe_machine['asset_tag']}, so we'll update it."
                )

                if snipe_machine.get("asset_tag"):
                    if self.args.dryrun:
                        self.log.info(
                            f"would be updating {serial} in {source} with the asset tag {asset_info[asset_key]} => {snipe_machine['asset_tag']}"
                        )
                        continue

                    updates.append(
                        {
                            "asset_tag": snipe_machine["asset_tag"],
                            "machine_id": machine_id,
                            "serial": serial,
                        }
                    )

        return updates

    def checkout_assets(self, assets: dict) -> None:
        if not self.args.dryrun:
            for key, values in assets.items():
                if values:
                    for asset in values:
                        with ThreadPoolExecutor(max_workers=10) as executor:
                            for asset in values:
                                if key == "no_user":
                                    executor.submit(
                                        self._checkout_snipe_assets,
                                        asset["asset_id"],
                                        asset["assigned_user"],
                                    )
                                    continue

                                executor.submit(
                                    self._update_snipe_machines(
                                        asset["serial"],
                                        asset["asset_id"],
                                        assigned_to=asset["assigned_user"],
                                        **{
                                            k: v
                                            for k, v in asset.items()
                                            if k
                                            not in [
                                                "serial",
                                                "asset_id",
                                                "assigned_user",
                                            ]
                                        },
                                    )
                                )

    def create_new_asset(self, new_assets: list) -> None:
        with ThreadPoolExecutor(max_workers=10) as executor:
            for asset in new_assets:
                if self.args.dryrun:
                    self.log.info(f"would be creating a new asset: {asset}")
                    continue
                executor.submit(self._asset_creator, asset)

    def create_new_snipe_models(
        self, models: list[str], model_info: dict = {}, manufacturer_id=False
    ) -> None:
        """
        create all the new models after gathering machine info
        """
        new_models = []
        for model in models:
            if model != "":
                self.log.info(f"Could not find a model ID in snipe for: {model}")

                if manufacturer_id is False:
                    manufacturer_id = self.apple_manufacturer_id

                if not model_info:
                    model_number = model
                else:
                    model_number = model_info[model]

                new_model = {
                    "category_id": self.config.snipe_it.computer_model_category_id,
                    "manufacturer_id": manufacturer_id,
                    "name": model,
                    "model_number": model_number,
                }

                if hasattr(self.config.snipe_it, "computer_custom_fieldset_id"):
                    fieldset_split = self.config.snipe_it.computer_custom_fieldset_id
                    new_model["fieldset_id"] = fieldset_split

                if self.args.dryrun:
                    self.log.info(f"would be creating new model: {new_model}")
                    continue

                new_models.append(new_model)

        with ThreadPoolExecutor(max_workers=10) as executor:
            for model in new_models:
                executor.submit(self._model_creator, model)

    def get_models(self) -> dict:
        """
        get a list of known models from Snipe-IT
        """
        self.log.info("Getting a list of computer models currently in Snipe-IT.")

        snipe_models = self.snipe.get_models()
        self.log.debug(
            f"Parsing the snipe model results for models with model numbers: {snipe_models['rows']}"
        )

        for model in snipe_models["rows"]:
            if model["model_number"] == "":
                self.log.debug(
                    f"The model, {model['name']}, did not have a model number. Skipping."
                )
                continue

            self.snipe.model_numbers[model["model_number"]] = model["id"]

        self.log.info(
            f"Our list of models has {len(self.snipe.model_numbers)} entries."
        )
        self.log.debug(
            f"""
            Here's the list of the models and their id's that we were able to collect:
            {self.snipe.model_numbers}
            """
        )

        return snipe_models

    def get_users(self) -> dict:
        """
        get a list of users from Snipe-IT
        """
        users = {}

        self.log.info("Getting a list of users currently in Snipe-IT.")

        snipe_users = self.snipe.get_users()
        self.log.debug("Parsing the snipe users.")

        for user in snipe_users:
            users[user["id"]] = user["username"]

        return users

    def update_existing_snipe(self, updates: dict) -> None:
        if updates:
            with ThreadPoolExecutor(max_workers=10) as executor:
                for machine_id, vals in updates.items():
                    if vals["payload"]:
                        self.log.debug(
                            f"payload for {vals['serial']} : {vals['payload']}"
                        )
                        if self.args.dryrun:
                            self.log.debug(
                                f"would be updating snipe machine(serial: {vals['serial']} - id: {machine_id}) with: {vals['payload']}"
                            )
                            continue
                        executor.submit(
                            self._update_snipe_machines,
                            vals["serial"],
                            machine_id,
                            **vals["payload"],
                        )
        else:
            self.log.info("no machines listed to update")

    def update_new_asset_info(self, asset_type: str, new_assets: list) -> list[dict]:
        update_info = []
        for asset in new_assets:
            try:
                new_asset = {
                    "asset_tag": asset["tag"],
                    "model_id": self.snipe.model_numbers[
                        asset["info"]["model_identifier"]
                    ],
                    "name": asset["info"]["name"],
                    "status_id": self.snipe.default_status,
                    "serial": asset["info"]["serial"],
                }

            # TODO: handle this better
            except Exception as e:
                self.log.info(f"encountered error putting together new asset: {e}")
                continue

            if asset["info"]["serial"] == "Not Available":
                self.log.warning(
                    """
                    The serial number is not available in JAMF.
                    This is normal for DEP enrolled devices that have not yet checked in for the first time.
                    Since there's no serial number yet, we'll skip it for now.
                    """
                )
                continue

            new_asset.update(
                self._attribute_adder(asset_info=asset["info"], asset_type=asset_type)
            )

            if self.args.auto_incrementing:
                self.log.debug(f"popping asset_tag off: {new_asset}")
                new_asset.pop("asset_tag", None)

            update_info.append(new_asset)

        return update_info

    def user_checker(self, user: str, create: bool = True, **kwargs) -> tuple[bool:str]:
        """
        this will check if the user exists in snipe.
        if create is True the user will be created if they do not exist.
        """
        exists = [self.users[i] for i in self.users if self.users[i] == user]
        if exists:
            return True, [i for i in self.users if self.users[i] == user][0]

        if create:
            res = self.snipe.create_user(
                first_name=kwargs["first_name"],
                last_name=kwargs["last_name"],
                username=user,
            )
            if res.get("status") == "success":
                return True, res["payload"]["id"]

        return False, ""
