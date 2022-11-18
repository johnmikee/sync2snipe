from .errors import InvalidComputerGroupID


class Groups:
    def _computer_group_parser(self, computer_group: dict, **kwargs) -> list:
        valid_keys = ["id", "name", "alt_mac_address", "serial_number"]

        for k in kwargs.keys():
            if k not in valid_keys:
                kwargs.pop(k)

        response = []
        for item in computer_group["computer_group"]["computers"]:
            for key in list(item):
                if key not in list(kwargs.keys()):
                    del item[key]

            response.append(item)

        return response

    def _computer_group_template_static(
        self, name: str, group_id: int, serial: str
    ) -> str:
        return f"""
        <computer_group>
            <id>{group_id}</id>
            <name>{name}</name>
            <computer_additions>
                <computer>
                    <serial_number>{serial}</serial_number>
                </computer>
            </computer_additions>
        </computer_group>
        """.strip()

    def add_machine_to_group(self, name: str, group_id: int, serial: str):
        data = self._computer_group_template_static(
            name=name,
            group_id=group_id,
            serial=serial,
        )
        return self._requester(
            "PUT", f"/computergroups/id/{group_id}", data=data, xml_return=True
        )

    def get_all_computer_groups(self) -> dict:
        """
        the first time this runs it will update the class
        with a list of group_ids. that list will be used to
        validate subsequent calls to the computergroups endpoint.
        """
        res = self._requester("GET", "computergroups")

        if self.computer_group_ids is None:
            self.computer_group_ids = [
                i["id"] for i in res["computer_groups"] if res.get("computer_groups")
            ]

        return res

    def get_computers_by_group(
        self, group_id: str, sorted: bool = True
    ) -> (list | dict):
        if self.computer_group_ids is None:
            self.get_all_computer_groups()

        if group_id not in self.computer_group_ids:
            raise InvalidComputerGroupID(
                f"{group_id} is not valid. Existing group_ids are: {self.computer_group_ids}"
            )

        res = self._requester("GET", f"/computergroups/id/{group_id}")

        if sorted:
            return self._computer_group_parser(
                computer_group=res, id=True, serial_number=True
            )

        return res
