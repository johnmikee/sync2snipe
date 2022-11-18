class Companies:
    def get_company_list(self) -> dict:
        """
        https://snipe-it.readme.io/reference/companies
        """
        return self._requester("GET", "companies")

    def create_company(self, name: str) -> dict:
        """
        https://snipe-it.readme.io/reference/companies-1
        """
        payload = {"name": name}
        return self._requester("POST", "companies", json=payload)

    def get_company_detail(self, company_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/companiesid
        """
        return self._requester("GET", f"companies/{company_id}")

    def delete_company(self, company_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/companiesid-3
        """
        return self._requester("DELETE", f"companies/{company_id}")

    def update_company_detail(self, company_id: str, name: str) -> dict:
        """
        https://snipe-it.readme.io/reference/companiesid-1
        """
        payload = {"name": name}
        return self._requester("GET", f"companies/{company_id}", json=payload)
