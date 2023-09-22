from .api_items import arg_checker
from .errors import InvalidAPIArg


class General:
    def get_all_machines(self, **kwargs) -> dict:
        """
        classicapi/doc/#/computers/findComputers
        returns a list of dicts
        ex:
        {'id': 123, 'name': 'john doe macbook'}
        """
        res = self._requester("GET", "computers", **kwargs)
        return res["computers"]

    def get_device_info(self, machine_id: int, device_type: str, endpoint: str) -> dict:
        """
        returns device info for computers or mobiles

        valid endpoint options:
            - computers, mobiledevices
        valid device_type options:
            - computer, mobile_device

        """
        return self._requester("GET", f"/{endpoint}/id/{machine_id}")[device_type]

    def general_info(self, arg: str, info: dict) -> dict:
        return arg_checker(key="general", value=arg, info=info)

    def hardware_info(self, arg: str, info: dict) -> dict:
        return arg_checker(key="hardware", value=arg, info=info)

    def location_info(self, arg: str, info: dict) -> dict:
        return arg_checker(key="location", value=arg, info=info)

    def purchasing_info(self, arg: str, info: dict) -> dict:
        return arg_checker(key="purchasing", value=arg, info=info)

    def security_info(self, arg: str, info: dict) -> dict:
        return arg_checker(key="security", value=arg, info=info)

    def software_info(self, arg: str, info: dict) -> dict:
        return arg_checker(key="software", value=arg, info=info)

    def group_account_info(self, arg: str, info: dict) -> dict:
        return arg_checker(key="groups_accounts", value=arg, info=info)

    def update_machine(
        self,
        machine_id: str,
        data: dict,
    ) -> dict:
        res = self._requester(
            "PUT",
            f"/computers/id/{machine_id}",
            data=self._dict_to_xml(data),
            raw=True,
        )

        return res

    def update_asset_tag(
        self, asset_tag: int, machine_id: int, machine_type: str
    ) -> dict:
        tag_types = {
            "computer": {"tag": "computer", "endpoint": "computers"},
            "mobile": {"tag": "mobile_device", "endpoint": "mobiledevices"},
        }
        tag = tag_types[machine_type]["tag"]
        endpoint = tag_types[machine_type]["endpoint"]

        data = f"""<{tag}>
            <general>
            <asset_tag>{asset_tag}</asset_tag>
            </general>
            </{tag}>
            """.strip()

        res = self._requester(
            "PUT",
            f"/{endpoint}/id/{machine_id}",
            data=data,
            raw=True,
        )

        return res
