import os

from common import get_logger

from .api import API
from .errors import ServiceFileNotFound, MissingConfig
from .chromedevices import ChromeDevices
from .groups import Groups
from .mobile import Mobile
from .users import Users


class Google(API, ChromeDevices, Groups, Mobile, Users):
    def __init__(
        self,
        domain=None,
        google_user=None,
        google_auth_user=None,
        level=None,
        service_file=None,
    ):
        self.domain = domain
        self.log = get_logger(level=level)
        self.google_user = self._user_validator(google_user)
        self.google_auth_user = self._user_validator(google_auth_user)
        self.service_account_file = self._check_service_file(service_file)
        self.verify = self._config_checker()

        super(Google, self).__init__()

    def _check_service_file(self, service_file):
        if not os.path.exists(service_file):
            raise ServiceFileNotFound(
                f"Check the path for your google service json file. {service_file} cannot be found"
            )

        return os.path.abspath(service_file)

    def _config_checker(self) -> None:
        required_config = [self.google_auth_user, self.google_user, self.domain]
        if any(elem is None for elem in required_config):
            raise MissingConfig(
                f"All three of the following values must be passed. google_user:{self.google_user}, google_auth_user:{self.google_auth_user}, domain:{self.domain}"
            )

    def _user_validator(self, user: str) -> str:
        """
        make sure the user ends in the domain
        """
        if not user.endswith(self.domain):
            updated = f"{user.strip()}{self.domain}"
            self.log.debug(f"changing {user} to {updated}")

            return updated

        return user


__all__ = ["Google"]
