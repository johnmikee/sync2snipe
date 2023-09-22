class MDM:
    def get_smart_group_members(self, group_id: int) -> dict:
        """
        *This endpoint is in preview*

        Recalculates the smart group for the given id and then
        returns the ids for the users in the smart group
        https://developer.jamf.com/jamf-pro/reference/smart-user-groups-preview-1
        """
        return self._requester(
            "POST", f"api/v1/smart-user-groups/{group_id}/recalculate"
        )

    def get_static_groups(self):
        """
        *This endpoint is in preview*

        Returns a list of all static user groups.
        https://developer.jamf.com/jamf-pro/reference/get_v1-static-user-groups
        """
        return self._requester("GET", "api/v1/static-user-groups")

    def get_static_group(self, group_id: int) -> dict:
        """
        *This endpoint is in preview*

        Returns a specific static user group by id.
        https://developer.jamf.com/jamf-pro/reference/get_v1-static-user-groups-id
        """
        return self._requester("GET", f"api/v1/static-user-groups/{group_id}")

    def get_smart_groups_users(self, group_id: int) -> dict:
        """
        *This endpoint is in preview*

        Recalculates a smart group for the given user id and then
        returns the count of smart groups the user falls into
        https://developer.jamf.com/jamf-pro/reference/smart-user-groups-preview-1#post_v1-users-id-recalculate-smart-groups
        """
        return self._requester(
            "POST", f"api/v1/users/{group_id}/recalculate-smart-groups"
        )
