class Departments:
    def get_departments(self) -> dict:
        """
        https://snipe-it.readme.io/reference/departments
        """
        return self._requester("GET", "departments")

    def create_department(self, name: str) -> dict:
        """
        https://snipe-it.readme.io/reference/departments-1
        """
        payload = {"name": name}
        return self._requester("POST", "departments", json=payload)

    def get_department(self, department_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/departmentsid
        """
        return self._requester("GET", f"departments/{department_id}")

    def delete_department(self, department_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/departmentsid-2
        """
        return self._requester("DELETE", f"departments/{department_id}")

    def update_department(self, method: str, name: str) -> dict:
        """
        https://snipe-it.readme.io/reference/departments-1
        """
        payload = {"name": name}

        if method.upper() not in ["PATCH", "PUT"]:
            raise self.InvalidArg(f"{method} must be either PATCH or PUT")

        return self._requester(method.upper(), "departments", json=payload)
