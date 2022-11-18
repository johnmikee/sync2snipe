import requests


class Users:
    def _user_group_template_static(
        self, name: str, smart_status: str = "false", notify_change: str = "false"
    ) -> str:
        return f"""
    <user_group>
        <name>{name}</name>
        <is_smart>{smart_status}</is_smart>
        <is_notify_on_change>{notify_change}</is_notify_on_change>
        <site>
            <id>0</id>
            <name>None</name>
        </site>
        <users>
        </users>
    </user_group>
        """.strip()

    def get_all_users(self) -> dict:
        return self._requester("GET", "users")

    def get_all_user_groups(self) -> dict:
        return self._requester("GET", "usergroups")

    def get_user_group_by_id(self, group_id: str) -> dict:
        return self._requester("GET", f"/usergroups/id/{group_id}")

    def get_user_by_id(self, user_id: str, raw=False) -> requests.Response:
        return self._requester("GET", f"/users/id/{user_id}", raw=raw)

    def get_user_by_name(self, name: str, raw: bool = False) -> requests.Response:
        return self._requester("GET", f"/users/name/{name}", raw=raw)

    def create_user(self, name: str, full_name: str, email: str) -> requests.Response:
        body = f"""
        <user>
            <name>{name}</name>
            <full_name>{full_name}</full_name>
            <email>{email}</email>
            <sites>
                <site>
                    <id>-1</id>
                    <name>None</name>
                </site>
            </sites>
        </user>
        """.strip()

        return self._requester("POST", "/users/id/0", xml_return=True, data=body)

    def create_static_user_group(self, name: str) -> requests.Response:
        return self._requester(
            "POST",
            "/usergroups/id/0",
            xml_return=True,
            data=self._user_group_template_static(name),
        )

    def users_to_group(self, group_id: str, data: dict) -> requests.Response:
        """
        the data body can add or delete users
        add: {"user_group": {"user_additions": {"user": []}}}
        delete: {"user_group": {"user_deletions": {"user": []}}}
        """
        return self._requester(
            "POST", f"/usergroups/id/{group_id}", xml_return=True, data=data
        )
