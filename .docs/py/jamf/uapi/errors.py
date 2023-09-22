class InvalidJamfCreds(Exception):
    """raise when no creds were passed"""


class InvalidParamaters(Exception):
    """raise when invalid options or lack of options passed"""


class AttachmentError(Exception):
    """raise when attachment cannot be found"""
