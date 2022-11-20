import re


class StatusLabels:
    def _hex_validator(self, color: hex) -> dict:
        """
        make sure color is valid hex code
        """
        match = re.search(r"^#(?:[0-9a-fA-F]{3}){1,2}$", color)
        if not match:
            self.log.debug(f"{color} is not a valid hex code")
            return False
        return True

    def get_status_labels(self, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/statuslabels
        """
        query_opts = {"limit": "", "offset": "", "search": "", "sort": "", "order": ""}
        query = self._opt_sorter(query_opts, **kwargs)

        return self._requester("GET", "statuslabels", params=query)

    def create_status_label(self, name: str, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/statuslabels-1
        """
        if kwargs.get("color"):
            if not self._hex_validator(kwargs["color"]):
                del kwargs["color"]

        opts = {
            "type": "",
            "show_in_nav": False,
            "default_label": False,
            "name": name,
            "notes": "",
            "color": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        return self._requester("POST", "statuslabels", json=payload)

    def get_status_label(self, label_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/statuslabelsid
        """
        return self._requester("GET", f"statuslabels/{label_id}")

    def delete_status_label(self, label_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/statuslabelsid-1
        """
        return self._requester("DELETE", f"statuslabels/{label_id}")

    def update_status_label(
        self, method: str, label_id: int, name: str, label_type: str, **kwargs
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/statuslabelsid-2
        https://snipe-it.readme.io/reference/statuslabelsid-3
        """
        if kwargs.get("color"):
            if not self._hex_validator(kwargs["color"]):
                del kwargs["color"]

        opts = {
            "type": label_type,
            "show_in_nav": False,
            "default_label": False,
            "name": name,
            "notes": "",
            "color": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        if method.upper() not in ["PATCH", "PUT"]:
            raise self.InvalidArg(f"{method} must be either PATCH or PUT")

        return self._requester(method.upper(), f"statuslabels/{label_id}", json=payload)

    def get_assets_by_label(self, label_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/statuslabelsidassetlist
        """
        return self._requester("GET", f"statuslabels/{label_id}/assetlist")
