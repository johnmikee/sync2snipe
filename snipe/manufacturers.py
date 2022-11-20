class Manufacturers:
    def get_manufacturers(self) -> dict:
        """
        https://snipe-it.readme.io/reference/manufacturers
        """
        return self._requester("GET", "manufacturers")

    def create_manufacturer(self, name: str) -> dict:
        """
        https://snipe-it.readme.io/reference/manufacturers-1
        """
        payload = {"name": name}
        return self._requester("POST", "manufacturers", json=payload)

    def get_manufacturer(self, manufacturer_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/manufacturers
        """
        return self._requester("GET", f"manufacturers/{manufacturer_id}")

    def delete_manufacturer(self, manufacturer_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/manufacturersid-2
        """
        return self._requester("DELETE", f"manufacturers/{manufacturer_id}")

    def update_manufacturer(self, manufacturer_id: int, name: str) -> dict:
        """
        https://snipe-it.readme.io/reference/manufacturersid-3
        """
        payload = {"name": name}
        return self._requester("PUT", f"manufacturers/{manufacturer_id}", json=payload)
