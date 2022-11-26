class ChromeDevices:
    """
    Requires one of the following scopes:
        https://www.googleapis.com/auth/admin.directory.device.chromeos
        https://www.googleapis.com/auth/admin.directory.device.chromeos.readonly

    https://developers.google.com/admin-sdk/directory/reference/rest/v1/chromeosdevices
    """

    def checker(self, device: dict, query: str):
        """
        return key value if exists else None
        """
        if device.get(query):
            return device[query]

        return None

    def _chromeos_opt_validator(self, opts: list) -> list:
        valid_opts = [
            "kind",
            "etag",
            "deviceId",
            "serialNumber",
            "status",
            "lastSync",
            "annotatedUser",
            "model",
            "osVersion",
            "platformVersion",
            "firmwareVersion",
            "macAddress",
            "bootMode",
            "lastEnrollmentTime",
            "firstEnrollmentTime",
            "orgUnitPath",
            "orgUnitId",
            "recentUsers",
            "ethernetMacAddress",
            "activeTimeRanges",
            "tpmVersionInfo",
            "cpuStatusReports",
            "systemRamTotal",
            "systemRamFreeReports",
            "diskVolumeReports",
            "autoUpdateExpiration",
            "cpuInfo",
        ]

        return [i for i in opts if i in valid_opts]

    def list_chromeos_devices(self) -> list:
        """
        https://developers.google.com/admin-sdk/directory/reference/rest/v1/chromeosdevices/get
        """
        devices = []
        token = None

        while True:
            results = (
                self.admin_service.chromeosdevices()
                .list(customerId="my_customer", pageToken=token)
                .execute()
            )

            if results.get("chromeosdevices"):
                devices.extend(results["chromeosdevices"])

            token = results.get("nextPageToken")
            if not token:
                break

        return devices

    def filter_chromeos_devices(self, opts: list, devices: list = None) -> dict:
        """
        if devices are not passed this will pull the list.
        returns a dictionary with the device id as key and opts passed sorted in dict as value.

        used to parse list_chromeos_devices which returns
        a list of chrome devices with following keys

        ['kind', 'etag', 'chromeosdevices', 'nextPageToken']

        within chromeosdevices are the fields of interest
            kind
            etag
            deviceId
            serialNumber
            status
            lastSync
            annotatedUser
            model
            osVersion
            platformVersion
            firmwareVersion
            macAddress
            bootMode
            lastEnrollmentTime
            firstEnrollmentTime
            orgUnitPath
            orgUnitId
            recentUsers
            ethernetMacAddress
            activeTimeRanges
            tpmVersionInfo
            cpuStatusReports
            systemRamTotal
            systemRamFreeReports
            diskVolumeReports
            autoUpdateExpiration
            cpuInfo
        """

        if not devices:
            devices = self.list_chromeos_devices()

        sorted = {}

        for d in devices:
            sorted[d["deviceId"]] = {
                o: self.checker(d, o) for o in self._chromeos_opt_validator(opts=opts)
            }

        return sorted

    def get_chromeos_device(self, device_id: str) -> dict:
        """
        https://developers.google.com/admin-sdk/directory/reference/rest/v1/chromeosdevices/get
        """
        return (
            self.admin_service.chromeosdevices()
            .get(customerId="my_customer", deviceId=device_id)
            .execute()
        )

    def patch_chromeos_device(self, device_id: str, **kwargs) -> dict:
        """
        https://developers.google.com/admin-sdk/directory/reference/rest/v1/chromeosdevices/patch
        """
        valid_opts = [
            "annotatedUser",
            "annotatedLocation",
            "notes",
            "orgUnitPath",
            "annotatedAssetId",
        ]
        for k in kwargs:
            if k not in valid_opts:
                self.log.info(f"removing invalid patch arg: {k}")
                kwargs.pop(k)

        return (
            self.admin_service.chromeosdevices()
            .patch(customerId="my_customer", deviceId=device_id, body=kwargs)
            .execute()
        )
