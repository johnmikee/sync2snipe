class Utils:
    def _replace_items(self, obj: any, key: str, replace_value: str):
        if isinstance(obj, dict):
            for k, v in obj.items():
                if isinstance(v, list):
                    for item in v:
                        if isinstance(item, dict):
                            if key in item:
                                item[key] = replace_value

                if isinstance(v, dict):
                    for val in v.values():
                        if isinstance(val, list):
                            for i in val:
                                if isinstance(i, dict):
                                    if key in i:
                                        i[key] = replace_value
                                if key in i:
                                    i[key] = replace_value

                    obj[k] = self._replace_items(v, key, replace_value)

        elif isinstance(obj, list):
            for index, item in enumerate(obj):
                if isinstance(item, dict):
                    if key in item:
                        item[key] = replace_value

                if item == key:
                    obj[index] == replace_value

        if key in obj:
            obj[key] = replace_value

        return obj

    def _validate_args(self, opts: any, instance: (dict | None) = None, **kwargs):
        if instance:
            for k, v in instance.items():
                opts = self._replace_items(opts, k, v)

            return opts

        for k, v in kwargs.items():
            opts = self._replace_items(opts, k, v)

        # convert to list to avoid RunTimeError when
        # removing an item in the list while iterating over it
        for item in list(opts.keys()):
            if not kwargs.get(item):
                opts.pop(item)

        return opts
