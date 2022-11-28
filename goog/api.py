from common import get_logger

from .chromedevices import ChromeDevices
from .creds import Creds
from .groups import Groups
from .mobile import Mobile
from .users import Users


class Google(Creds, ChromeDevices, Groups, Mobile, Users):
    def __init__(
        self,
        env_vars=True,
        level=None,
        scopes=[],
    ):
        self.env_vars = env_vars
        self.log = get_logger(level=level)
        self.scopes = scopes

        super(Google, self).__init__()
