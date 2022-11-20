class Categories:
    def get_categories(self, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/categories-1
        """
        query_opts = {"limit": "", "offset": "", "search": "", "sort": "", "order": ""}
        query = self._opt_sorter(query_opts, **kwargs)

        return self._requester("GET", "categories", params=query)

    def create_category(self, name: str, category_type: str, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/categories-2
        """
        opts = {
            "use_default_eula": "",
            "require_acceptance": "",
            "checkin_email": "",
            "name": name,
            "category_type": category_type,
        }
        payload = self.opt_parser(opts, **kwargs)

        return self._requester("GET", "categories", json=payload)

    def get_category(self, category_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/categoriesid-3
        """
        return self._requester("GET", f"categories/{category_id}")

    def delete_category(self, category_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/categoriesid-2
        """
        return self._requester("DELETE", f"categories/{category_id}")

    def update_category(
        self, method: str, category_id: int, name: str, category_type: str, **kwargs
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/categoriesid
        https://snipe-it.readme.io/reference/categoriesid-1
        """
        opts = {
            "use_default_eula": "",
            "require_acceptance": "",
            "checkin_email": "",
            "name": name,
            "category_type": category_type,
        }
        payload = self.opt_parser(opts, **kwargs)

        if method.upper() not in ["PATCH", "PUT"]:
            raise self.InvalidArg(f"{method} must be either PATCH or PUT")

        return self._requester(
            method.upper(), f"categories/{category_id}", json=payload
        )
