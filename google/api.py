from google.oauth2 import service_account
from googleapiclient.discovery import build


class API:
    def __init__(self):
        credentials = service_account.Credentials.from_service_account_file(
            self.service_account_file, scopes=self.scopes
        )

        self.credentials = credentials
        self.user_credentials = self.credentials.with_subject(str(self.google_user))
        self.delegated_credentials = self.credentials.with_subject(
            self.google_auth_user
        )

        self.admin_service = self._admin()

    def _admin(self):
        return build(
            "admin",
            "directory_v1",
            credentials=self.delegated_credentials,
            cache_discovery=False,
        )
