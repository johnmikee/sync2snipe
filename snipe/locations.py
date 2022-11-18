class Locations:
    def get_locations(self, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/locations
        """
        query_opts = {"limit": "", "offset": "", "search": "", "sort": "", "order": ""}
        query = self._opt_sorter(query_opts, **kwargs)

        return self._requester("GET", "locations", params=query)

    def create_location(self, name: str, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/locations-2
        """
        opts = {
            "name": name,
            "address": "",
            "address2": "",
            "state": "",
            "country": "",
            "zip": "",
            "ldap_ou": "",
            "parent_id": "",
            "currency": "",
            "manager_id": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        return self._requester("POST", "locations", json=payload)

    def get_location_detail(self, location_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/locations-1
        """
        return self._requester("GET", f"locations/{location_id}")

    def delete_location(self, location_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/locationsid-2
        """
        return self._requester("DELETE", f"locations/{location_id}")

    def update_location(
        self, method: str, location_id: str, name: str, **kwargs
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/locations-3
        https://snipe-it.readme.io/reference/locations-4

        method can be PUT or PATCH
        """
        opts = {
            "name": name,
            "address": "",
            "address2": "",
            "state": "",
            "country": "",
            "zip": "",
            "ldap_ou": "",
            "parent_id": "",
            "currency": "",
            "manager_id": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        if method.upper() not in ["PATCH", "PUT"]:
            raise self.InvalidArg(f"{method} must be either PATCH or PUT")

        return self._requester(method.upper(), f"locations/{icense_id}", json=payload)
