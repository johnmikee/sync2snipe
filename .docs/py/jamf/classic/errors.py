class InvalidAPIArg(IndexError):
    """raise when invalid arg passed to search on"""


class InvalidJamfCreds(Exception):
    """raise when no creds were passed"""


class InvalidComputerGroupID(Exception):
    """raise when invalid group_id was passed"""


class InvalidMobileGroupID(Exception):
    """raise when invalid group_id was passed"""


class InvalidParamaters(Exception):
    """raise when invalid options or lack of options passed"""


class AttachmentError(Exception):
    """raise when attachment cannot be found"""
