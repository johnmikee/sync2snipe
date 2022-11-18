class Components:
    def get_components(self, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/components

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
            "sort": "",
            "order": "",
            "expand": "",
        }
        query = self._opt_sorter(query_opts, **kwargs)

        return self._requester("GET", "components", params=query)

    def create_component(self, name: str, qty: int, category_id: int, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/components-1
        """
        opts = {
            "name": name,
            "qty": qty,
            "location_id": "",
            "category_id": category_id,
            "company_id": "",
            "order_number": "",
            "purchase_date": "",
            "purchase_cost": "",
            "min_amt": "",
            "serial": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        return self._requester("POST", "components", json=payload)

    def get_component(self, component_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/componentsid
        """
        return self._requester("GET", f"components/{component_id}")

    def delete_component(self, component_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/componentsid-3
        """
        return self._requester("DELETE", f"components/{component_id}")

    def update_component(
        self,
        method: str,
        component_id: str,
        name: str,
        qty: int,
        category_id: int,
        **kwargs,
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/components-1
        """
        opts = {
            "name": name,
            "qty": qty,
            "location_id": "",
            "category_id": category_id,
            "company_id": "",
            "order_number": "",
            "purchase_date": "",
            "purchase_cost": "",
            "min_amt": "",
            "serial": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        if method.upper() not in ["PATCH", "PUT"]:
            raise self.InvalidArg(f"{method} must be either PATCH or PUT")

        return self._requester(
            method.upper(), f"components/{component_id}", json=payload
        )

    def get_checkedout_components(self, component_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/componentsidassets
        """
        return self._requester("DELETE", f"components/{component_id}/assets")

    def checkout_components(
        self, component_id: str, assigned_to: int, assigned_qty: int
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/componentsidcheckout
        """
        payload = {"assigned_to": assigned_to, "assigned_qty": assigned_qty}

        return self._requester(
            "POST", f"components/{component_id}/checkout", json=payload
        )

    def checkin_components(self, component_id: str, checkin_qty: int) -> dict:
        """
        https://snipe-it.readme.io/reference/componentsidcheckin
        """
        payload = {"checkin_qty": checkin_qty}

        return self._requester(
            "POST", f"components/{component_id}/checkin", json=payload
        )
