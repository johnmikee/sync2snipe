class Models:
    def get_models(self, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/models
        """
        query_opts = {"limit": "", "offset": "", "search": "", "sort": "", "order": ""}
        query = self._opt_sorter(query_opts, **kwargs)

        return self._requester("GET", "models", params=query)

    def create_models(
        self,
        name: str,
        model_number: str,
        category_id: int,
        manufacturer_id: int,
        **kwargs,
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/models-1
        """
        opts = {
            "name": name,
            "model_number": model_number,
            "category_id": category_id,
            "manufacturer_id": manufacturer_id,
            "eol": "",
            "fieldset_id": "",
        }
        query = self._opt_sorter(opts, **kwargs)

        return self._requester("POST", "models", params=query)

    def get_model(self, model_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/modelsid
        """
        return self._requester("GET", f"models/{model_id}")

    def delete_model(self, model_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/modelsid-3
        """
        return self._requester("DELETE", f"models/{model_id}")

    def update_model(
        self,
        method: str,
        model_id: str,
        name: str,
        category_id: int,
        manufacturer_id: int,
        **kwargs,
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/modelsid-1
        https://snipe-it.readme.io/reference/modelsid-2
        """
        opts = {
            "name": name,
            "model_number": "",
            "category_id": category_id,
            "manufacturer_id": manufacturer_id,
            "fieldset_id": "",
            "eol": "",
            "depreciation_id": "",
            "notes": "",
            "requestable": False,
        }
        payload = self._opt_sorter(opts, **kwargs)

        if method not in ["PATCH", "PUT"]:
            raise self.InvalidArg(f"{method} must be either PATCH or PUT")

        return self._requester(method, f"models/{model_id}", json=payload)
