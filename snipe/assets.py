class Assets:
    def get_all_hardware(self, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/hardware-list

        Sortable Columns

        Field Name	        Description
        id	                Asset ids
        name	            Asset names
        asset_tag	        Asset unique asset tag
        serial	            Asset serial numbers
        model	            Name of the associated asset model
        model_number	    Model number of the associated asset model
        last_checkout	    Date the asset was last checked out
        category	        Name of the asset model's associated category
        manufacturer	    Name of the asset's manufacturer
        notes	            Asset notes
        expected_checkin	Expected checkin date
        order_number	    Order number associated with the asset
        companyName	        Company name associated with the asset, if applicable
        location	        Name of the location of the asset, either the default Ready to Deploy location, or the assigned user's location
        image	            Name of the optional image file associated with the asset
        status_label	    Name of the status label associated with the asset
        assigned_to	        Name of the person the asset is assigned to
        created_at	        Date the asset was created
        purchase_date	    Date the asset was purchased
        purchase_cost	    Purchase cost of the asset
        """
        results = []
        limit = 500

        res = self._requester(
            "GET", "hardware", params={"offset": 0, "limit": limit}, **kwargs
        )
        [results.append(i) for i in res["rows"]]
        self.log.debug(f"total number of assets in snipe: {res['total']}")
        offset_range = self._offsetter(total=res["total"], limit=limit)
        for offset in offset_range:
            res = self._requester(
                "GET",
                "hardware",
                params={"offset": offset["offset"], "limit": offset["limit"]},
                **kwargs,
            )
            [results.append(i) for i in res["rows"]]

        return results

    def create_asset(
        self, asset_tag: str, status_id: int, model_id: int, **kwargs
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/hardware-create
        """
        opts = {
            "asset_tag": asset_tag,
            "status_id": status_id,
            "model_id": model_id,
            "name": "",
            "serial": "",
            "purchase_date": "",
            "purchase_cost": "",
            "order_number": "",
            "notes": "",
            "archived": False,
            "warranty_months": "",
            "depreciate": False,
            "supplier_id": "",
            "requestable": False,
            "rtd_location_id": "",
            "last_audit_date": "",
            "location_id": "",
        }
        payload = self._opt_sorter(opts, **kwargs)
        self.log.debug(payload)

        return self._requester("POST", "hardware", json=payload)

    def get_asset_by_id(self, asset_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/hardware-by-id
        """
        return self._requester("GET", f"hardware/{asset_id}")

    def get_asset_by_tag(self, asset_tag: str) -> dict:
        """
        https://snipe-it.readme.io/reference/hardware-by-asset-tag
        """
        return self._requester("GET", f"hardware/bytag/{asset_tag}")

    def get_asset_by_serial(self, serial: str) -> dict:
        """
        https://snipe-it.readme.io/reference/hardware-by-serial
        """
        return self._requester("GET", f"hardware/byserial/{serial}")

    def update_asset(
        self, asset_id: int, asset_tag: str, status_id: int, model_id: int, **kwargs
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/hardware-update
        """
        opts = {
            "last_checkout": "",
            "assigned_to": "",
            "company_id": "",
            "serial": "",
            "warranty_months": "",
            "purchase_cost": "",
            "purchase_date": "",
            "requestable": "",
            "archived": "",
            "rtd_location_id": "",
            "name": "",
            "location_id": "",
            "image": "",
            "asset_tag": asset_tag,
            "status_id": status_id,
            "model_id": model_id,
            "notes": "",
            "order_number": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        return self._requester("PUT", f"hardware/{asset_id}", json=payload)

    def patch_asset(self, asset_id: int, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/hardware-partial-update
        """
        opts = {
            "last_checkout": "",
            "assigned_to": "",
            "company_id": "",
            "serial": "",
            "warranty_months": "",
            "purchase_cost": "",
            "purchase_date": "",
            "requestable": "",
            "archived": "",
            "rtd_location_id": "",
            "name": "",
            "location_id": "",
            "image": "",
            "asset_tag": "",
            "status_id": "",
            "model_id": "",
            "notes": "",
            "order_number": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        return self._requester("PATCH", f"hardware/{asset_id}", json=payload)

    def delete_asset(self, asset_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/hardware-delete
        """
        return self._requester("DELETE", f"hardware/{asset_id}")

    def checkout_asset(
        self, asset_id: int, status_id: int, checkout_to_type: str, **kwargs
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/hardware-checkout
        """
        if checkout_to_type not in ["user", "asset", "location"]:
            raise self.InvalidArg(
                f"{checkout_to_type} must be either asset, user, or location"
            )

        opts = {
            "status_id": status_id,
            "checkout_to_type": checkout_to_type,
            "assigned_user": "",
            "assigned_asset": "",
            "assigned_location": "",
            "expected_checkin": "",
            "checkout_at": "",
            "name": "",
            "note": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        return self._requester("POST", f"hardware/{asset_id}/checkout", json=payload)

    def checkin_asset(self, asset_id: int, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/hardware-checkin
        """
        opts = {"note": "", "location": ""}
        payload = self._opt_sorter(opts, **kwargs)

        return self._requester("POST", f"hardware/{asset_id}/checkin", json=payload)

    def mark_asset_audited(self, asset_tag: str, location_id=None) -> dict:
        """
        https://snipe-it.readme.io/reference/hardwareaudit
        """
        opts = {"asset_tag": asset_tag, "location": location_id}

        if opts["location"] is None:
            del opts["location"]

        return self._requester("POST", "hardware/audit", json=opts)

    def get_audit_assets(self) -> dict:
        """
        https://snipe-it.readme.io/reference/hardwareauditdue
        """
        return self._requester("GET", "hardware/audit/due")

    def get_overdue_assets(self) -> dict:
        """
        https://snipe-it.readme.io/reference/hardwareauditoverdue
        """
        return self._requester("GET", "hardware/audit/overdue")

    def get_asset_licenses(
        self,
        asset_id: int,
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/hardwareidlicenses
        """
        return self._requester("GET", f"hardware/{asset_id}/licenses")
