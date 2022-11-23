from google import Google

"""
for auth user put your g suite admin
for google user put your standard account

for both of these do not add the @domain 
"""
g = Google(
    google_auth_user="",
    google_user="",
    domain="",
    scopes=[
        "https://www.googleapis.com/auth/admin.directory.user",
        "https://www.googleapis.com/auth/admin.directory.group",
        "https://www.googleapis.com/auth/admin.directory.user.security",
        "https://www.googleapis.com/auth/admin.directory.device.mobile",
        "https://www.googleapis.com/auth/admin.directory.device.chromeos",
    ],
    service_file="google_service.json",
)

# grab all the devices
all_chromeos_devices = g.list_chromeos_devices()
# get some values
chrome_devices = g.filter_chromeos_devices(
    opts=["serialNumber", "status", "annotatedUser", "model", "recentUsers"]
)
print(chrome_devices)
# get more values
chrome_devices_more = g.filter_chromeos_devices(
    opts=[
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
)
print(chrome_devices_more)
# use the list you grabbed the first time to parse the values.
chrome_devices_reuse = g.filter_chromeos_devices(
    ["serialNumber", "status", "annotatedUser", "model", "recentUsers"],
    devices=all_chromeos_devices,
)
print(chrome_devices_reuse)
