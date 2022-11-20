class Consumables:
    def get_consumables(self, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/consumables

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
            "category_id": "",
            "company_id": "",
            "manufacturer_id": "",
        }
        query = self._opt_sorter(query_opts, **kwargs)

        return self._requester("GET", "consumables", params=query)

    def create_consumable(
        self, name: str, qty: int, category_id: int, **kwargs
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/consumables-1
        """
        opts = {
            "name": name,
            "qty": qty,
            "category_id": category_id,
            "company_id": "",
            "order_number": "",
            "manufacturer_id": "",
            "location_id": "",
            "requestable": False,
            "purchase_date": "",
            "min_amt": "",
            "model_number": "",
            "item_no": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        return self._requester("POST", "consumables", json=payload)

    def get_consumable(self, consumbable_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/consumablesid
        """
        return self._requester("GET", f"consumables/{consumbable_id}")

    def delete_consumable(self, consumbable_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/consumablesid-2
        """
        return self._requester("DELETE", f"consumables/{consumbable_id}")

    def update_consumable(
        self,
        method: str,
        consumbable_id: int,
        name: str,
        qty: int,
        category_id: int,
        **kwargs,
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/consumables-2
        """
        opts = {
            "name": name,
            "qty": qty,
            "category_id": category_id,
            "company_id": "",
            "order_number": "",
            "manufacturer_id": "",
            "location_id": "",
            "requestable": False,
            "purchase_date": "",
            "min_amt": "",
            "model_number": "",
            "item_no": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        if method not in ["PATCH", "PUT"]:
            raise self.InvalidArg(f"{method} must be either PATCH or PUT")
        return self._requester(method, f"consumables/{consumbable_id}", json=payload)

    def checkin_consumbable(self, consumbable_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/consumablesidcheckin
        """
        return self._requester("POST", f"consumables/{consumbable_id}/checkin")

    def checkout_consumbable(self, consumbable_id: int, assigned_to=None) -> dict:
        """
        https://snipe-it.readme.io/reference/consumablesidcheckout
        """
        payload = {"assigned_to": assigned_to}
        return self._requester(
            "POST", f"consumables/{consumbable_id}/checkout", json=payload
        )
