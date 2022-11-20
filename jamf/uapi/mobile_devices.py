from .api_params import mobile_opts, mobile_search_opts
from .errors import InvalidParamaters


class Mobile:
    def get_mobile_devices(self) -> dict:
        """
        Gets Mobile Device objects.
        https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-devices
        """
        return self._requester("GET", "api/v2/mobile-devices")["results"]

    def get_mobile_device(self, mobile_id: int, detail: bool = False) -> dict:
        """
        Get MobileDevice with our without detail

        https://developer.jamf.com/jamf-pro/reference/mobile-devices-1#get_v1-mobile-devices-id
        https://developer.jamf.com/jamf-pro/reference/mobile-devices-1#get_v1-mobile-devices-id-detail
        """
        url = f"api/v1/mobile-devices/{mobile_id}"
        if detail:
            url = f"{url}/detail"

        res = self._requester("GET", url)

        return res["results"]

    def get_mobile_device_compliance(self, mobile_id: int) -> dict:
        """
        Return basic compliance information for the given mobile device

        https://developer.jamf.com/jamf-pro/reference/client-check-in-1
        """
        return self._requester(
            "GET",
            f"api/v1/conditional-access/device-compliance-information/mobile/{mobile_id}",
        )

    def get_mobile_device_groups(self) -> dict:
        """
        Returns the list of all mobile device groups.

        https://developer.jamf.com/jamf-pro/reference/mobile-device-groups-preview-1
        """
        return self._requester("GET", "api/v1/mobile-device-groups")

    def search_mobile_devices(self, instance=None, **kwargs) -> dict:
        """
        Search Mobile Devices

        https://developer.jamf.com/jamf-pro/reference/mobile-devices-1#post_v1-search-mobile-devices
        """
        if not "orderBy" in kwargs:
            raise InvalidParamaters(
                "orderBy is a required paramater to search on mobile device"
            )
        if instance:
            payload = self._validate_args(mobile_search_opts, instance)
        else:
            payload = self._validate_args(mobile_search_opts, **kwargs)

        return self._requester("PATCH", "api/v1/search-mobile-devices", json=payload)

    def get_mobile_device_enrollment_profile(self, mobile_id: int) -> dict:
        """
        Retrieve the MDM Enrollment Profile
        https://developer.jamf.com/jamf-pro/reference/mobile-device-enrollment-profile-1
        """
        return self._requester(
            "GET",
            f"api/v1/mobile-device-enrollment-profile/{mobile_id}/download-profile",
        )

    def update_mobile_device(self, mobile_id: int, instance=None, **kwargs) -> dict:
        """
        Updates fields on a mobile device that are allowed to be modified by users.
        https://developer.jamf.com/jamf-pro/reference/patch_v2-mobile-devices-id
        """
        if instance:
            payload = self._validate_args(mobile_opts, instance)
        else:
            payload = self._validate_args(mobile_opts, **kwargs)

        return self._requester(
            "PATCH", f"api/v1/mobile-devices/{mobile_id}", json=payload
        )
