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

    def get_general_info(self, arg: str, info: dict) -> dict:
        if arg not in self.api_values["general"]:
            raise InvalidAPIArg(
                f"{arg} not acceptable search paramater. Valid options are: {self.api_values['general']}"
            )

        return info["general"][arg]

    def get_hardware_info(self, arg: str, info: dict) -> dict:
        if arg not in self.api_values["hardware"]:
            raise InvalidAPIArg(
                f"{arg} not acceptable search paramater. Valid options are: {self.api_values['hardware']}"
            )

        return info["hardware"][arg]

    def get_location_info(self, arg: str, info: dict) -> dict:
        if arg not in self.api_values["location"]:
            raise InvalidAPIArg(
                f"{arg} not acceptable search paramater. Valid options are: {self.api_values['location']}"
            )

        return info["location"][arg]

    def get_purchasing_info(self, arg: str, info: dict) -> dict:
        if arg not in self.api_values["purchasing"]:
            raise InvalidAPIArg(
                f"{arg} not acceptable search paramater. Valid options are: {self.api_values['purchasing']}"
            )

        return info["purchasing"][arg]

    def get_security_info(self, arg: str, info: dict) -> dict:
        if arg not in self.api_values["security"]:
            raise InvalidAPIArg(
                f"{arg} not acceptable search paramater. Valid options are: {self.api_values['security']}"
            )

        return info["security"][arg]

    def get_software_info(self, arg: str, info: dict) -> dict:
        if arg not in self.api_values["software"]:
            raise InvalidAPIArg(
                f"{arg} not acceptable search paramater. Valid options are: {self.api_values['software']}"
            )

        return info["software"][arg]

    def get_group_account_info(self, arg: str, info: dict) -> dict:
        if arg not in self.api_values["groups_accounts"]:
            raise InvalidAPIArg(
                f"{arg} not acceptable search paramater. Valid options are: {self.api_values['groups_accounts']}"
            )

        return info["groups_accounts"][arg]

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
