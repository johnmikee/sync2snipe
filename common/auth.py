import os
import sys
import getpass
import keyring

from .errors import InvalidCreds

from .logger import get_logger

from types import SimpleNamespace


class Auth:
    def __init__(self, keychain_username="sync2snipe", env_var=False):
        self.env_var = env_var
        self.log = get_logger()
        self.keychain_username = keychain_username

    def _get_key(self, cred_type: str) -> str:
        try:
            api_key = keyring.get_password(
                service_name=cred_type, username=self.keychain_username
            )
        except keyring.errors.NoKeyringError:
            self.log.info("No keyring found")

        key_name = cred_type + "_sync2snipe"
        self.log.debug(f" - Getting creds for: {key_name}")

        if sys.platform == "darwin" and api_key:
            self.log.debug("Using key from keyring on macOS")
            return keyring.get_password(
                service_name=cred_type, username=self.keychain_username
            )
        else:
            print(f"** No value set for: {cred_type}")
            return self._set_key(cred_type)

    def _set_key(self, cred_type: str) -> str:
        print(f"Please enter the value for: {cred_type}")
        api_key = getpass.getpass()
        keyring.set_password(
            service_name=cred_type, username=self.keychain_username, password=api_key
        )

        return api_key

    def get_config(self, *args, accept_none: bool = False) -> SimpleNamespace:
        config = {}

        if self.env_var:
            for arg in args:
                config.update({arg: os.environ.get(arg)})
        else:
            # default to keyring
            for arg in args:
                config.update({arg: self._get_key(arg)})

        # make sure the values arent empty
        values = [i for i in list(config.values())]
        if None in values:
            if not accept_none:
                raise InvalidCreds(
                    f"could not get all values for arguements passed: {config}"
                )

        return SimpleNamespace(**config)
