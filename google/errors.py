class ServiceFileNotFound(Exception):
    """Google Service file does not exist - cannot continue"""


class InvalidScopes(Exception):
    """One or more scopes passed is not valid"""


class MissingConfig(Exception):
    """Need to pass all three options to instantiate the class"""


class InvalidAction(Exception):
    """Invalid action performed"""
