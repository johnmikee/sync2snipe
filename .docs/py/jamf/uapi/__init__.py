from common import Requester, get_logger
from jamf.auth import JamfAuth

from .category import Category
from .computer import Computers
from .mdm import MDM
from .mobile_devices import Mobile
from .utils import Utils


class UAPI(
    Utils,
    JamfAuth,
    Requester,
    Category,
    Computers,
    Mobile,
    MDM,
):
    def __init__(
        self,
        ci=False,
        level=None,
    ):
        self.ci = ci
        self.classic = False
        self.log = get_logger(level=level)
        self.uapi = True
        self.uapi_url = ""
        super(UAPI, self).__init__()
