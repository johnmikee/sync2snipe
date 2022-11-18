import os

from .errors import AttachmentError
from .api_params import computer_opts


class Computers:
    def delete_computer(self, computer_id: int) -> dict:
        """
        Remove specified Computer record

        https://developer.jamf.com/jamf-pro/reference/computer-inventory-1#delete_v1-computers-inventory-id
        """
        return self._requester("DELETE", f"api/v1/computers-inventory/{computer_id}")

    def get_computer_detail(self, computer_id: int) -> dict:
        """
        Return a Computer details with all sections

        returns dictionary of device info
        keys:

        id,udid,general,diskEncryption,localUserAccounts,purchasing,printers,storage,
        applications,userAndLocation,configurationProfiles,services,plugins,hardware,
        certificates,attachments,packageReceipts,fonts,security,operatingSystem,licensedSoftware,
        softwareUpdates,groupMemberships,extensionAttributes,contentCaching,ibeacons

        https://developer.jamf.com/jamf-pro/reference/computer-inventory-1#get_v1-computers-inventory-detail-id
        """
        return self._requester(
            "GET", f"api/v1/computers-inventory-detail/{computer_id}"
        )

    def get_computer_general_detail(self, computer_id: int) -> dict:
        """
        Return a Computer General details

        returns a dictionary of general info
        keys:
        'name', 'lastIpAddress', 'lastReportedIp', 'jamfBinaryVersion', 'platform', 'barcode1',
        'barcode2', 'assetTag', 'remoteManagement', 'supervised', 'mdmCapable', 'reportDate',
        'lastContactTime', 'lastCloudBackupDate', 'lastEnrolledDate', 'mdmProfileExpiration',
        'initialEntryDate', 'distributionPoint', 'site', 'itunesStoreAccountActive', 'enrolledViaAutomatedDeviceEnrollment',
        'userApprovedMdm', 'enrollmentMethod', 'extensionAttributes'

        https://developer.jamf.com/jamf-pro/reference/computer-inventory-1#get_v1-computers-inventory-id
        """
        res = self._requester("GET", f"api/v1/computers-inventory/{computer_id}")
        return res["general"]

    def get_computer_inventory(self) -> dict:
        """
        Return a Computer Inventory for paginated list of computers

        returns a list of dictionaries
        keys:

        id,udid,general,diskEncryption,localUserAccounts,purchasing,printers,storage,
        applications,userAndLocation,configurationProfiles,services,plugins,hardware,
        certificates,attachments,packageReceipts,fonts,security,operatingSystem,licensedSoftware,
        softwareUpdates,groupMemberships,extensionAttributes,contentCaching,ibeacons

        https://developer.jamf.com/jamf-pro/reference/computer-inventory-1
        """

        res = self._requester("GET", "api/v1/computers-inventory")
        return res["results"]

    def get_computer_prestages(self) -> dict:
        """
        Search for sorted and paged computer prestages

        returns list of dictionaries
        keys:
        'enrollmentSiteId', 'id', 'displayName', 'supportPhoneNumber', 'supportEmailAddress',
        'department', 'authenticationPrompt', 'profileUUID', 'deviceEnrollmentProgramInstanceId',
        'versionLock', 'siteId', 'skipSetupItems', 'locationInformation', 'purchasingInformation',
        'anchorCertificates', 'enrollmentCustomizationId', 'customPackageIds', 'customPackageDistributionPointId',
        'prestageInstalledProfileIds', 'isInstallProfilesDuringSetup', 'isMdmRemovable', 'isKeepExistingSiteMembership',
        'isKeepExistingLocationInformation', 'isMandatory', 'isDefaultPrestage', 'isRequireAuthentication',
        'isPreventActivationLock', 'isEnableDeviceBasedActivationLock'

        https://developer.jamf.com/jamf-pro/reference/computer-prestages-1#get_v1-computer-prestages
        """
        res = self._requester("GET", "api/v1/computer-prestages")
        return res["results"]

    def update_computer(self, computer_id: int, instance=None, **kwargs) -> dict:
        """
        Return a updated computer instance

        https://developer.jamf.com/jamf-pro/reference/computer-inventory-1#patch_v1-computers-inventory-detail-id
        """
        if instance:
            payload = self._validate_args(computer_opts, instance)
        else:
            payload = self._validate_args(computer_opts, **kwargs)

        return self._requester(
            "PATCH", f"v1/computers-inventory-detail/{computer_id}", json=payload
        )
