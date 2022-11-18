class Maintenances:
    def get_asset_maintenances(self, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/maintenances
        """
        query_opts = {
            "limit": "",
            "offset": "",
            "search": "",
            "sort": "",
            "order": "",
            "asset_id": "",
        }
        query = self._opt_sorter(query_opts, **kwargs)

        return self._requester("GET", "maintenances", params=query)

    def create_maintenance(
        self, title: str, asset_id: int, supplier_id: int, start_date: str, **kwargs
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/maintenances-1
        """
        asset_maintenance_type = [
            "Maintenance",
            "Repair",
            "PAT Test",
            "Upgrade",
            "Hardware Support",
            "Software Support",
        ]

        if kwargs.get("asset_maintenance_type"):
            if kwargs["asset_maintenance_type"] not in asset_maintenance_type:
                raise self.InvalidArg(
                    f"{kwargs['asset_maintenance_type']} must be one of {asset_maintenance_type}"
                )
        opts = {
            "title": title,
            "asset_id": asset_id,
            "supplier_id": supplier_id,
            "is_warranty": False,
            "cost": "",
            "notes": "",
            "start_date": start_date,
            "completion_date": "",
            "asset_maintenance_type": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        return self._requester("POST", "maintenances", json=payload)

    def delete_maintenance(self, maintenance_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/maintenancesid
        """
        return self._requester("DELETE", f"maintenances/{maintenance_id}")
