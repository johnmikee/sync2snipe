from .errors import InvalidAPIArg


def arg_checker(value: str, key: str, info: dict) -> any:
    if value not in api_values[key]:
        raise InvalidAPIArg(
            f"""{value} not acceptable search paramater.
            Valid options are: {api_values[key]}"""
        )
    return info[key][value]


api_values = {
    "general": [
        "id",
        "name",
        "network_adapter_type",
        "mac_address",
        "alt_network_adapter_type",
        "alt_mac_address",
        "ip_address",
        "last_reported_ip",
        "serial_number",
        "udid",
        "jamf_version",
        "platform",
        "barcode_1",
        "barcode_2",
        "asset_tag",
        "remote_management",
        "supervised",
        "mdm_capable",
        "mdm_capable_users",
        "management_status",
        "report_date",
        "report_date_epoch",
        "report_date_utc",
        "last_contact_time",
        "last_contact_time_epoch",
        "last_contact_time_utc",
        "initial_entry_date",
        "initial_entry_date_epoch",
        "initial_entry_date_utc",
        "last_cloud_backup_date_epoch",
        "last_cloud_backup_date_utc",
        "last_enrolled_date_epoch",
        "last_enrolled_date_utc",
        "mdm_profile_expiration_epoch",
        "mdm_profile_expiration_utc",
        "distribution_point",
        "sus",
        "netboot_server",
        "site",
        "itunes_store_account_is_active",
    ],
    "location": [
        "username",
        "realname",
        "real_name",
        "email_address",
        "position",
        "phone",
        "phone_number",
        "department",
        "building",
        "room",
    ],
    "purchasing": [
        "is_purchased",
        "is_leased",
        "po_number",
        "vendor",
        "applecare_id",
        "purchase_price",
        "purchasing_account",
        "po_date",
        "po_date_epoch",
        "po_date_utc",
        "warranty_expires",
        "warranty_expires_epoch",
        "warranty_expires_utc",
        "lease_expires",
        "lease_expires_epoch",
        "lease_expires_utc",
        "life_expectancy",
        "purchasing_contact",
        "os_applecare_id",
        "os_maintenance_expires",
        "attachments",
    ],
    "peripherals": "",
    "hardware": [
        "make",
        "model",
        "model_identifier",
        "os_name",
        "os_version",
        "os_build",
        "active_directory_status",
        "service_pack",
        "processor_type",
        "processor_architecture",
        "processor_speed",
        "processor_speed_mhz",
        "number_processors",
        "number_cores",
        "total_ram",
        "total_ram_mb",
        "boot_rom",
        "bus_speed",
        "bus_speed_mhz",
        "battery_capacity",
        "cache_size",
        "cache_size_kb",
        "available_ram_slots",
        "optical_drive",
        "nic_speed",
        "smc_version",
        "ble_capable",
        "supports_ios_app_installs",
        "sip_status",
        "gatekeeper_status",
        "xprotect_version",
        "institutional_recovery_key",
        "disk_encryption_configuration",
        "filevault2_users",
        "storage",
        "mapped_printers",
    ],
    "certificates": "",
    "security": ["activation_lock", "secure_boot_level", "external_boot_level"],
    "software": [
        "unix_executables",
        "licensed_software",
        "installed_by_casper",
        "installed_by_installer_swu",
        "cached_by_casper",
        "available_software_updates",
        "available_updates",
        "running_services",
        "applications",
        "fonts",
        "plugins",
    ],
    "extension_attributes": ["id", "name", "type", "multi_value", "value"],
    "groups_accounts": [
        "computer_group_memberships",
        "local_accounts",
        "user_inventories",
    ],
    "iphones": [],
    "configuration_profiles": ["id", "name", "uuid", "is_removable"],
}
