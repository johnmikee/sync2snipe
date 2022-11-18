import requests

from .accessories import Accessories
from .api import API
from .assets import Assets
from .categories import Categories
from .companies import Companies
from .consumables import Consumables
from .components import Components
from .departments import Departments
from .fields import Fields, FieldSets
from .licenses import Licenses
from .locations import Locations
from .maintenances import Maintenances
from .manufacturers import Manufacturers
from .models import Models
from .reports import Reports
from .status_labels import StatusLabels
from .users import Users

from .errors import InvalidArg

from common import Auth, Requester, get_logger


class Snipe(
    Requester,
    API,
    Accessories,
    Assets,
    Categories,
    Companies,
    Consumables,
    Components,
    Departments,
    Fields,
    FieldSets,
    Licenses,
    Locations,
    Maintenances,
    Manufacturers,
    Models,
    Reports,
    StatusLabels,
    Users,
):
    def __init__(
        self,
        users_no_search=True,
        default_status=2,
        level="INFO",
    ):
        self.config = Auth().get_config("SNIPE_TOKEN", "SNIPE_URL")
        self.default_status = default_status
        self.base_url = self.url_handler(self.config.SNIPE_URL, "base") + "api/v1/"
        self.token = self.config.SNIPE_TOKEN
        self.headers = {
            "Accept": "application/json",
            "Content-Type": "application/json",
            "Authorization": f"Bearer {self.config.SNIPE_TOKEN}",
        }
        self.InvalidArg = InvalidArg
        self.session = requests.Session()
        self.first_snipe_call = None
        self.log = get_logger(level=level)
        self.model_numbers = {}
        self.users_no_search = users_no_search
        super(Snipe, self).__init__()


__all__ = [
    "Snipe",
]
