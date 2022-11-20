class Users:
    def get_users(self, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/users
        """
        query_opts = {
            "search": "",
            "limit": "",
            "offset": "",
            "sort": "",
            "order": "",
            "first_name": "",
            "last_name": "",
            "username": "",
            "email": "",
            "employee_num": "",
            "state": "",
            "zip": "",
            "country": "",
            "group_id": "",
            "department_id": "",
            "company_id": "",
            "location_id": "",
            "deleted": "",
            "all": "",
        }
        query = self._opt_sorter(query_opts, **kwargs)

        return self._requester("GET", "users", params=query)

    def create_user(
        self, first_name: str, last_name: str, username: str, password=None, **kwargs
    ) -> dict:
        """
        https://snipe-it.readme.io/reference/users-2
        """
        """
        this isnt the best idea
        - supply your own method for password generation
        """
        if password is not None:
            import secrets

            password_length = 20
            password = secrets.token_urlsafe(password_length)

        opts = {
            "first_name": first_name,
            "last_name": last_name,
            "username": username,
            "password": password,
            "password_confirmation": password,
            "email": "",
            "permissions": "",
            "activated": False,
            "phone": "",
            "jobtitle": "",
            "manager_id": "",
            "employee_num": "",
            "company_id": "",
            "notes": "",
            "two_factor_enrolled": False,
            "two_factor_optin": False,
            "department_id": "",
            "location_id": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        return self._requester("POST", "users", json=payload)

    def get_user(self, user_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/usersid
        """
        return self._requester("POST", f"users/{user_id: int}")

    def delete_user(self, user_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/usersid-1
        """
        return self._requester("DELETE", f"users/{user_id}")

    def update_user(self, method: str, user_id: int, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/usersid-2
        https://snipe-it.readme.io/reference/users-3
        """
        opts = {
            "first_name": "",
            "last_name": "",
            "username": "",
            "password": "",
            "password_confirmation": "",
            "email": "",
            "permissions": "",
            "activated": False,
            "phone": "",
            "jobtitle": "",
            "manager_id": "",
            "employee_num": "",
            "company_id": "",
            "notes": "",
            "two_factor_enrolled": False,
            "two_factor_optin": False,
            "department_id": "",
            "location_id": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        if method.upper() not in ["PATCH", "PUT"]:
            raise self.InvalidArg(f"{method} must be either PATCH or PUT")

        return self._requester(method.upper(), f"users/{user_id}", json=payload)

    def get_users_items(self, user_id: int, item_type: str) -> dict:
        """
        https://snipe-it.readme.io/reference/usersidassets
        https://snipe-it.readme.io/reference/usersidaccessories
        https://snipe-it.readme.io/reference/usersidlicenses

        item_type can be assets, accessories, licenses
        """
        endpoints = ["assets", "accessories", "licenses"]

        if item_type in endpoints:
            endpoint = f"users/{user_id}/{item_type}"
        else:
            raise self.InvalidArg(
                f"{item_type} must be either assets, accessories, or licenses"
            )

        return self._requester("GET", endpoint)
