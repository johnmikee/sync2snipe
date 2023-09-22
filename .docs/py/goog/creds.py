import base64
import json

from google.oauth2 import service_account
from googleapiclient.discovery import build

from common import Auth


class Creds:
    def __init__(
        self,
    ):
        self.config = Auth(env_var=self.env_vars).get_config(
            "GOOGLE_USER", "GOOGLE_AUTH_USER", "GOOGLE_DOMAIN", "SERVICE_CREDS"
        )
        self.domain = self.config.GOOGLE_DOMAIN
        self.test = self._credential_grabber()
        credentials = service_account.Credentials.from_service_account_info(
            self._credential_grabber(), scopes=self.scopes
        )
        self.credentials = credentials
        self.user_credentials = self.credentials.with_subject(
            self._user_validator(self.config.GOOGLE_USER)
        )
        self.delegated_credentials = self.credentials.with_subject(
            self._user_validator(self.config.GOOGLE_AUTH_USER)
        )

        self.admin_service = self._admin()

    def _credential_grabber(self) -> json:
        """
        the SERVICE_CREDS should be exported as base64 encoded string.
        this can be created by `cat /path/to/service_account.json | base64`
        """
        return json.loads(base64.b64decode(self.config.SERVICE_CREDS))

    def _user_validator(self, user: str) -> str:
        """
        make sure the user ends in the domain
        """
        if not user.endswith(self.domain):
            updated = f"{user.strip()}{self.domain}"
            self.log.debug(f"changing {user} to {updated}")

            return updated

        return user

    def _admin(self):
        return build(
            "admin",
            "directory_v1",
            credentials=self.delegated_credentials,
            cache_discovery=False,
        )
