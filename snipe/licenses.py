class Licenses:
    def get_licenses(self, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/licenses
        """
        query_opts = {
            "limit": "",
            "offset": "",
            "search": "",
            "order_number": "",
            "sort": "",
            "order": "",
            "expand": "",
        }
        query = self._opt_sorter(query_opts, **kwargs)

        return self._requester("GET", "licenses", params=query)

    def create_license(self, name: str, seats: int, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/testinput
        """
        opts = {
            "name": name,
            "seats": seats,
            "company_id": "",
            "expiration_date": "",
            "license_email": "",
            "license_name": "",
            "serial": "",
            "maintained": False,
            "manufacturer_id": "",
            "notes": "",
            "order_number": "",
            "purchase_cost": "",
            "purchase_date": "",
            "purchase_order": "",
            "reassignable": False,
            "supplier_id": "",
            "termination_date": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        return self._requester("POST", "licenses", json=payload)

    def get_license(self, license_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/licensesid
        """
        return self._requester("GET", f"licenses/{license_id}")

    def delete_license(self, license_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/licensesid-3
        """
        return self._requester("DELETE", f"licenses/{license_id}")

    def update_license(self, method: str, license_id: str, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/licensesid-1
        https://snipe-it.readme.io/reference/licensesid-2
        """
        opts = {
            "name": "",
            "company_id": "",
            "expiration_date": "",
            "license_email": "",
            "license_name": "",
            "serial": "",
            "maintained": False,
            "manufacturer_id": "",
            "notes": "",
            "order_number": "",
            "purchase_cost": "",
            "purchase_date": "",
            "purchase_order": "",
            "reassignable": False,
            "seats": "",
            "supplier_id": "",
            "termination_date": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        if method.upper() not in ["PATCH", "PUT"]:
            raise self.InvalidArg(f"{method} must be either PATCH or PUT")

        return self._requester(method.upper(), f"licenses/{license_id}", json=payload)
