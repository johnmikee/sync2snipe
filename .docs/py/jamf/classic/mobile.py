from .errors import InvalidMobileGroupID


class Mobile:
    def get_all_mobile_groups(self) -> dict:
        """
        the first time this runs it will update the class
        with a list of mobile_group_ids. that list will be used to
        validate subsequent calls to the computergroups endpoint.
        """
        res = self._requester("GET", "mobiledevicegroups")

        if self.mobile_group_ids is None:
            self.mobile_group_ids = [
                i["id"]
                for i in res["mobile_device_groups"]
                if res.get("mobile_device_groups")
            ]

        return res

    def get_mobile_by_group(self, group_id: int) -> dict:
        if self.mobile_group_ids is None:
            self.get_all_mobile_groups()

        if group_id not in self.mobile_group_ids:
            raise InvalidMobileGroupID(
                f"{group_id} is not valid. Existing group_ids are: {self.mobile_group_ids}"
            )
        return self._requester("GET", f"mobiledevicegroups/id/{group_id}")[
            "mobile_device_group"
        ]

    def get_mobile_info(self, machine_id: str) -> dict:
        return self._requester("GET", f"/mobiledevices/id/{machine_id}")[
            "mobile_device"
        ]
