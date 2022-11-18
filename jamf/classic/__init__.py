from .api_items import api_values
from .general import General
from .groups import Groups
from .mobile import Mobile
from .users import Users

from common import get_logger
from jamf.auth import JamfAuth


class Classic(
    JamfAuth,
    General,
    Groups,
    Mobile,
    Users,
):
    def __init__(
        self,
        ci=False,
        level=None,
    ):
        self.api_values = api_values
        self.ci = ci
        self.classic = True
        self.classic_url = "JSSResource"
        self.computer_group_ids = None
        self.mobile_group_ids = None
        self.log = get_logger(level=level)
        self.uapi = False
        super(Classic, self).__init__()
