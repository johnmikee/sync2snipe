import datetime

import requests

from common import Auth, Requester, get_logger

from .formatter import Formatter


class JamfAuth(Requester, Formatter):
    def __init__(
        self,
    ):
        self.config = Auth().get_config("JAMF_USER", "JAMF_TOKEN", "JAMF_URL")
        self.api_user = self.config.JAMF_USER
        self.base_url = self._url_merger()
        self.api_key = self.config.JAMF_TOKEN
        self.headers = {"Accept": "application/json"}
        self.log = get_logger()
        self.token = None
        self.token_ttl = None
        self.session = requests.Session()
        super(JamfAuth, self).__init__()

    def _url_merger(self) -> str:
        if self.uapi:
            endpoint_type = ""

        if self.classic:
            endpoint_type = self.classic_url

        return self.url_handler(
            url=self.url_handler(self.config.JAMF_URL, "base") + endpoint_type,
            url_type="base",
        )

    def auth_request(self) -> None:
        self.base_url = self.url_handler(self.config.JAMF_URL, "base")
        # check if we have a token
        if not self.token:
            self.log.debug("token time")
            self.get_token()

        # make sure the expiration has not passed
        if not self.token_ttl:
            self.refresh_token()

    def get_token(self) -> bool:
        res = self.requester(
            "POST", "api/v1/auth/token", basic_auth=True, headers=self.headers
        )
        self.log.debug(f"result from getting a token: {res}")

        if res.get("token"):
            self.token = res["token"]
            self.update_headers_token()
        else:
            self.log.info("Could not obtain token")
            return False

        if res.get("expires"):
            self.log.debug(f"token expiration is: {res['expires']}")
            self.token_ttl = res["expires"]
        else:
            self.log.info("Could not obtain token")
            return False

        self.check_token_expiry()
        return True

    def refresh_token(self) -> None:
        res = self.requester(
            "POST",
            "api/v1/auth/keep-alive",
            headers={"Authorization": f"Bearer {self.token}"},
        )
        self.token = res["token"]
        self.update_headers_token()
        if res.get("expires"):
            self.log.debug(f"token expiration is: {res['expires']}")
            self.token_ttl = res["expires"]

    def check_token_expiry(self) -> None:
        expiration = self.token_ttl[:-1]

        deadline = datetime.datetime.strptime(
            expiration.split(".")[0], "%Y-%m-%dT%H:%M:%S"
        )
        if deadline > datetime.datetime.utcnow():
            self.log.debug("refreshing token")
            self.refresh_token()

    def update_headers_token(self) -> None:
        self.log.debug("updating headers with bearer token")
        self.headers.update({"Authorization": f"Bearer {self.token}"})

    def _requester(
        self, method: str, endpoint: str, **kwargs
    ) -> (dict | requests.Response):
        self.auth_request()
        exisiting_headers = self.headers

        if kwargs.get("xml_return"):
            self.headers.update({"Content-Type": "text/xml"})

        self.base_url = self.url_handler(self.config.JAMF_URL, "base")
        if self.classic:
            self.base_url = self.url_handler(
                url=self.base_url + "JSSResource", url_type="base"
            )

        res = self.requester(method, endpoint, **kwargs)

        self.headers = exisiting_headers

        return res
