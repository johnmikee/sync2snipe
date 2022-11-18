from .errors import InvalidParamaters


class Category:
    def get_advanced_search_objects(self) -> dict:
        """
        Gets Advanced Search Objects

        https://developer.jamf.com/jamf-pro/reference/advanced-mobile-device-searches-1#get_v1-advanced-mobile-device-searches
        """
        return self._requester("GET", "api/v1/advanced-mobile-device-searches")

    def get_categories(self) -> dict:
        """
        Gets Category objects.

        returns a list of dictionaries
        keys:
            id, name, priority

        https://developer.jamf.com/jamf-pro/reference/categories-2#get_v1-categories
        """
        res = self._requester("GET", "api/v1/categories")

        return res["results"]

    def create_category(self, name: str, priority: int) -> dict:
        """
        Create category record

        https://developer.jamf.com/jamf-pro/reference/post_v1-categories
        """
        payload = {"name": name, "priority": priority}

        return self._requester("POST", "api/v1/categories", json=payload)

    def delete_multiple_category_ids(self, category_ids: list) -> dict:
        """
        Delete multiple Categories by their IDs

        Pass an array of ids

        https://developer.jamf.com/jamf-pro/reference/post_v1-categories-delete-multiple
        """
        if not isinstance(category_ids, list):
            raise InvalidParamaters(f"{category_ids} is not a list of ids to delete.")

        payload = {"ids": category_ids}
        return self._requester(
            "POST", "api/v1/categories/delete-multiple", json=payload
        )

    def get_category(self, category_id: int) -> dict:
        """
        Gets specified Category object

        returns:
        id, name, priority

        https://developer.jamf.com/jamf-pro/reference/get_v1-categories-id
        """
        return self._requester("GET", f"api/v1/categories/{category_id}")

    def update_category(self, category_id: int, name: str, priority: int) -> dict:
        """
        Update specified category object

        https://developer.jamf.com/jamf-pro/reference/categories-2#put_v1-categories-id
        """
        payload = {"name": name, "priority": priority}
        return self._requester("PUT", f"api/v1/categories/{category_id}", json=payload)

    def delete_category(self, category_id: int) -> dict:
        """
        Removes specified category record

        https://developer.jamf.com/jamf-pro/reference/categories-2#delete_v1-categories-id
        """
        return self._requester("DELETE", f"api/v1/categories/{category_id}")
