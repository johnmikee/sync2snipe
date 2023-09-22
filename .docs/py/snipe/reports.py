class Reports:
    def get_reports(self, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/reportsactivity
        """
        query_opts = {
            "search": "",
            "target_type": "",
            "target_id": "",
            "item_type": "",
            "item_id": "",
            "action_type": "",
        }
        query = self._opt_sorter(query_opts, **kwargs)

        return self._requester("GET", "reports/activity", params=query)
