class Fields:
    def __init__(
        self,
    ):
        self.elements = ["text", "textarea", "checkbox", "radio", "listbox"]

    def get_custom_fields(self) -> dict:
        """
        https://snipe-it.readme.io/reference/fields-1
        """
        return self._requester("GET", "fields")

    def create_custom_field(self, name: str, element: str, **kwargs) -> dict:
        """
        https://snipe-it.readme.io/reference/fields-2
        """

        if element not in self.elements:
            raise self.InvalidArg(f"{element} not in valid args {self.elements}")

        opts = {
            "show_in_email": False,
            "field_encrypted": False,
            "name": name,
            "element": element,
            "field_values": "",
            "format": "",
            "help_text": "",
        }
        payload = self._opt_sorter(opts, **kwargs)

        return self._requester("POST", "fields", json=payload)

    def get_field_by_id(self, id) -> dict:
        """
        https://snipe-it.readme.io/reference/fieldsid
        """
        return self._requester("GET", f"fields/{id}")

    def update_custom_field(self, field_id: str, name: str, element: str) -> dict:
        """
        https://snipe-it.readme.io/reference/update-fields
        """
        if element not in self.elements:
            raise self.InvalidArg(f"{element} not in valid args {self.elements}")

        payload = {"name": name, "element": element}

        return self._requester("PUT", f"fields/{field_id}", json=payload)

    def delete_custom_field(self, field_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/fieldsid-1
        """
        return self._requester("DELETE", f"fields/{field_id}")

    def associate_custom_field(self, field_id: str, fieldset_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/fieldsidassociate
        """
        payload = {"fieldset_id": fieldset_id}

        return self._requester("POST", f"fields/{field_id}/associate", json=payload)

    def remove_custom_field_association(self, field_id: str, fieldset_id: int) -> dict:
        """
        https://snipe-it.readme.io/reference/fieldsidassociate
        """
        payload = {"fieldset_id": fieldset_id}

        return self._requester("POST", f"fields/{field_id}/disassociate", json=payload)


class FieldSets:
    def list_custom_fieldsets(self) -> dict:
        """
        https://snipe-it.readme.io/reference/fields
        """
        return self._requester("GET", "fieldsets")

    def create_fieldset(self, name: str) -> dict:
        """
        https://snipe-it.readme.io/reference/fieldsets
        """
        payload = {"name": name}
        return self._requester("POST", "fieldsets", json=payload)

    def get_fieldset(self, fieldset_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/fieldsetsid
        """
        return self._requester("GET", f"fieldsets/{fieldset_id}")

    def update_fieldset(self, fieldset_id: str, name: str) -> dict:
        """
        https://snipe-it.readme.io/reference/fieldsetsid-1
        """
        payload = {"name": name}
        return self._requester("GET", f"fieldsets/{fieldset_id}", json=payload)

    def delete_fieldset(self, fieldset_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/fieldsetsid-2
        """
        return self._requester("DELETE", f"fieldsets/{fieldset_id}")

    def get_associated_fieldsets(self, fieldset_id: str) -> dict:
        """
        https://snipe-it.readme.io/reference/fieldsetsidfields
        """
        return self._requester("GET", f"fieldsets/{fieldset_id}/fields")
