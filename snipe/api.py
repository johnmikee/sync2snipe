from ratelimit import limits, sleep_and_retry

# https://snipe-it.readme.io/reference/api-throttling
CALLS = 120
RATE_LIMIT = 60


class API:
    def _opt_sorter(self, opts: dict, **kwargs) -> dict:
        # add acceptable terms
        for k, v in kwargs.items():
            if k in opts:
                opts[k] = v

        body = {}
        # remove empty values
        for k, v in opts.items():
            if isinstance(v, bool):
                body[k] = v
            elif v:
                body[k] = v

        return body

    @sleep_and_retry
    @limits(calls=CALLS, period=RATE_LIMIT)
    def _requester(
        self,
        method: str,
        url: str,
        raw: bool = False,
        override_url: bool = False,
        **kwargs,
    ):
        return self.requester(
            method=method, url=url, raw=raw, override_url=override_url, **kwargs
        )
