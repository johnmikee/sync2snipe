class Users:
    def get_all_users(self) -> dict:
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
                }

            token = results.get("nextPageToken")
            if not token:
                break

        return sanitized

    def get_user(self, user: str) -> dict:
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

    def get_user_id(self, user: str) -> str:
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

    def get_user_token(self, user: str) -> list:
        tokens = []

        results = self.admin_service.tokens().list(userKey=user).execute()

        try:
            for r in results["items"]:
                tokens.append(r["clientId"])

            self.log.debug(f"{user} has the following access tokens: {tokens}")

        except KeyError:
            self.log.debug(f"{user} does not have any access tokens")

        return tokens

    def get_user_ou(self, user: str, user_info=False) -> str:
        if user_info:
            user_ou = user_info["users"][0]["orgUnitPath"]
        else:
            res = self.get_user(user)
            user_ou = res["users"][0]["orgUnitPath"]

        return user_ou
