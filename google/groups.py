class Groups:
    def get_google_groups(self, user):
        """
        https://developers.google.com/admin-sdk/directory/v1/reference/groups/list
        https://developers.google.com/admin-sdk/directory/v1/reference/members
        """
        results = (
            self.admin_service.groups()
            .list(customer="my_customer", query=f"memberKey={user}")
            .execute()
        )

        groups = {}
        try:
            for g in results["groups"]:
                groups[g["id"]] = g["email"]

            return groups

        except KeyError:
            self.log.debug(f"{user} was not a member of any groups")
            return False
