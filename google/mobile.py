from .errors import InvalidAction


class Mobile:
    def get_mobile_devices(self, user, customer_id=False):
        """
        https://developers.google.com/admin-sdk/directory/v1/reference/mobiledevices/get

        if this fails, replace my_customer with the org customer id
        """
        my_customer = "my_customer"
        if customer_id:
            my_customer = customer_id

        results = (
            self.admin_service.mobiledevices()
            .list(customerId=my_customer, maxResults=10, query=user)
            .execute()
        )

        user_devices = {}
        try:
            for r in results["mobiledevices"]:
                user_devices[r["resourceId"]] = r["lastSync"]

            return user_devices

        except KeyError:
            self.log.debug(
                f"{user} has no mobile devices in Google. Nothing to do here."
            )
            return False

    def action_mobile_devices(self, action, device, user):
        """
        https://developers.google.com/admin-sdk/directory/v1/reference/mobiledevices/action
        """

        """
		action	string	The action to be performed on the device.
        device  string  The device id to be acted upon
        user    string  The user the device belongs to.
                        This is only for logging purposes and can be empty.

		Acceptable values are:
			"admin_account_wipe": Remotely wipes only G Suite data from the device. See the administration help center for more information.
			"admin_remote_wipe": Remotely wipes all data on the device. See the administration help center for more information.
			"approve": Approves the device. If you've selected Enable device activation, devices that register after the device activation setting is enabled will need to be approved before they can start syncing with your domain. Enabling device activation forces the device user to install the Device Policy app to sync with G Suite.
			"block": Blocks access to G Suite data (mail, calendar, and contacts) on the device. The user can still access their mail, calendar, and contacts from a desktop computer or mobile browser.
			"cancel_remote_wipe_then_activate": Cancels a remote wipe of the device and then reactivates it.
			"cancel_remote_wipe_then_block": Cancels a remote wipe of the device and then blocks it.
		"""
        acceptable_actions = [
            "admin_account_wipe",
            "admin_remote_wipe",
            "approve",
            "block",
            "cancel_remote_wipe_then_activate",
            "cancel_remote_wipe_then_block",
        ]
        perform = action.lower()
        if perform not in acceptable_actions:
            raise InvalidAction(f"{action} not in {acceptable_actions}")

        action = {
            "action": perform,
        }

        try:
            results = (
                self.admin_service.mobiledevices()
                .action(customerId="my_customer", resourceId=device, body=action)
                .execute()
            )
            self.log.debug(results)
        except Exception as e:
            self.log.debug(
                f"encountered the following error deleting mobile devices for {user}: {e}"
            )
