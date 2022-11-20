import requests

from .logger import get_logger


class Requester:
    def __init__(self):
        self.log = get_logger(level="DEBUG")

    def requester(
        self,
        method: str,
        url: str,
        basic_auth: bool = False,
        raw: bool = False,
        override_url: bool = False,
        **kwargs,
    ):
        if "headers" in kwargs:
            self.headers.update(kwargs["headers"])
            kwargs.pop("headers")

        if override_url:
            url = url
        else:
            url = self.url_handler(self.base_url, url_type="base") + self.url_handler(
                url=url, url_type="endpoint"
            )

        req = requests.Request(
            method,
            url,
            headers=self.headers,
            **kwargs,
        )
        if basic_auth:
            req.auth = requests.auth.HTTPBasicAuth(self.api_user, self.api_key)

        prepped = req.prepare()
        try:
            response = self.session.send(prepped)
        except requests.exceptions.HTTPError as err:
            self.log.info(f"Ran into error with requester: {err}")
            return

        if raw:
            return response

        return response.json()

    def url_handler(self, url: str, url_type: str) -> str:
        """
        ensure the base url ends with a /

        trim any endpoint urls to join the base url
        """
        if url_type == "base":
            if not url.endswith("/"):
                url = url.strip() + "/"
        if url_type == "endpoint":
            if url.startswith("/"):
                url = url[1:]

        return url
