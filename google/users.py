class Users:
    def get_all_users(self):
        sanitized = {}

        token = None
        while True:
            results = (
                self.admin_service.users()
                .list(customer="my_customer", pageToken=token)
                .execute()
            )

            for user in results["users"]:
                sanitized[user["primaryEmail"]] = {
                    "first_name": user["name"]["givenName"],
                    "last_name": user["name"]["familyName"],
                    "last_login": user["lastLoginTime"],
                    "admin": user["isAdmin"],
                    "suspended": user["suspended"],
                }

            token = results.get("nextPageToken")
            if not token:
                break

        return sanitized

    def get_user(self, user):
        results = (
            self.admin_service.users()
            .list(
                customer="my_customer",
                maxResults=10,
                query=user,
                orderBy="email",
            )
            .execute()
        )

        return results

    def get_user_id(self, user):
        self.log.debug(f"getting user id for {user}")
        # pylint: disable=no-member
        results = (
            self.admin_service.users()
            .list(
                customer="my_customer",
                maxResults=10,
                query=user,
                orderBy="email",
            )
            .execute()
        )

        return results["users"][0]["id"]

    def get_user_token(self, user):
        results = self.admin_service.tokens().list(userKey=user).execute()

        tokens = []

        try:
            for r in results["items"]:
                tokens.append(r["clientId"])

            self.log.debug(f"{user} has the following access tokens: {tokens}")

        except KeyError:
            self.log.debug(f"{user} does not have any access tokens")

        return results

    def get_user_ou(self, user, user_info=False):
        if user_info:
            user_ou = user_info["users"][0]["orgUnitPath"]
        else:
            res = self.get_user(user)
            user_ou = res["users"][0]["orgUnitPath"]

        return user_ou
