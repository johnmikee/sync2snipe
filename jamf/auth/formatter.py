import xml.sax.saxutils

from collections import defaultdict
from xml.etree import cElementTree as ElementTree


class Formatter:
    def _dict_to_xml(self, data):
        """
        Convert python dict to xml string
        :returns:  xml string
        """
        if isinstance(data, dict):
            xml_str = ""
            for key, value in data.items():
                if value is None:
                    xml_str += f"<{key}/>"
                elif isinstance(value, list):
                    # if the value is a list, wrap each entry with the key
                    for i in value:
                        result = self._dict_to_xml(i)
                        xml_str += f"<{key}>{result}</{key}>"
                else:
                    # otherwise, wrap the entire result
                    result = self._dict_to_xml(value)
                    xml_str += f"<{key}>{result}</{key}>"
        elif isinstance(data, list):
            raise Exception("unable to properly tag nested lists")
        else:
            # string, boolean, integers, floats, etc
            xml_str = xml.sax.saxutils.escape(f"{data}")

        return xml_str

    def etree_to_dict(self, elem):
        """
        converts xml.cElementTree to python dict
        adapted from: https://stackoverflow.com/a/10077069/12020818
        removed attribute support
        """
        result = {elem.tag: None}
        children = list(elem)
        if children:
            defd = defaultdict(list)
            for dct in map(self.etree_to_dict, children):
                for key, val in dct.items():
                    defd[key].append(val)
            result = {
                elem.tag: {
                    key: val[0] if len(val) == 1 else val for key, val in defd.items()
                }
            }
        elif elem.text:
            result[elem.tag] = elem.text.strip()
        return result

    def xml_to_dict(self, xml_string):
        """
        Convert xml string to python dict
        :returns:  dict
        """
        root = ElementTree.XML(xml_string)
        return self.etree_to_dict(root)
