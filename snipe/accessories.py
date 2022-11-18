class Accessories:
    def get_accessories(self, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/accessories

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
        query_opts = {
            "limit": "",
            "offset": "",
            "search": "",
            "order_number": "",
            "sort": "",
            "order": "",
            "expand": "",
        }
        query = self.opt_sorter(query_opts, **kwargs)

        return self._requester("GET", "accessories", params=query)

    def create_accessory(self, name: str, qty: int, category_id: int, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/accessories-1
        """
        opts = {
            "supplier_id": "",
            "manufacturer_id": "",
            "location_id": "",
            "company_id": "",
            "category_id": category_id,
            "model_number": "",
            "purchase_date": "",
            "purchase_cost": "",
            "order_number": "",
            "qty": qty,
            "name": name,
        }
        payload = self.opt_sorter(opts, **kwargs)

        return self._requester("POST", "accessories", json=payload)

    def get_accessory(self, accessory_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/accessoriesid
        """
        return self._requester("GET", f"accessories/{accessory_id}")

    def update_accessory(
        self,
        method: str,
        accessory_id: str,
        name: str,
        qty: int,
        category_id: int,
        **kwargs,
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/accessoriesid-1
        https://snipe-it.readme.io/reference/accessoriesid-2

        method can be PATCH or PUT
        """
        opts = {
            "supplier_id": "",
            "manufacturer_id": "",
            "location_id": "",
            "company_id": "",
            "category_id": category_id,
            "model_number": "",
            "purchase_date": "",
            "purchase_cost": "",
            "order_number": "",
            "qty": qty,
            "name": name,
        }
        payload = self.opt_sorter(opts, **kwargs)

        if method.upper() not in ["PATCH", "PUT"]:
            raise self.InvalidArg(f"{method} must be either PATCH or PUT")

        return self._requester(
            method.upper(), f"accessories/{accessory_id}", json=payload
        )

    def delete_accessory(self, accessory_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/accessoriesid-3
        """
        return self._requester("DELETE", f"accessories/{accessory_id}")

    def get_co_accessory(self, accessory_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/accessoriesidcheckedout
        """
        return self._requester("GET", f"accessories/{accessory_id}/checkedout")

    def checkin_accessory(self, accessory_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/accessoriesidcheckin
        """
        return self._requester("POST", f"accessories/{accessory_id}/checkout")

    def checkout_accessory(
        self, accessory_id: str, assigned_to: int, note=None
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/accessoriesidcheckedout-1
        """
        payload = {"assigned_to": assigned_to, "note": note}

        if payload["note"] is None:
            del payload["note"]

        return self._requester(
            "POST", f"accessories/{accessory_id}/checkout", json=payload
        )
