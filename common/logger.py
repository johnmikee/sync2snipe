import sys
import logging


def get_logger(
    level: bool = None,
    log_file: bool = False,
    log_location: bool = False,
    disable_requests_logging: bool = True,
) -> logging.Logger:
    stdout_handler = logging.StreamHandler(sys.stdout)
    handlers = [stdout_handler]

    if log_file:
        if log_location:
            file_handler = logging.FileHandler(log_location)
        else:
            import time

            file_handler = logging.FileHandler(
                f"sync2snipe-{time.strftime('%m-%d-%Y')}.log"
            )

        handlers.append(file_handler)

    log_level = logging.INFO
    if level:
        level = level.upper()
        if level == "DEBUG":
            log_level = logging.DEBUG
        elif level == "ERROR":
            log_level = logging.ERROR
        elif level == "CRITICAL":
            log_level = logging.CRITICAL
        elif level == "WARNING":
            log_level = logging.WARNING

    logging.basicConfig(
        level=log_level,
        format="[%(asctime)s] {%(filename)s:%(lineno)d} %(levelname)s - %(message)s",
        handlers=handlers,
    )

    if disable_requests_logging:
        # requests and urllib3 make a whole lot of noise. supress that.
        logging.getLogger("urllib3").setLevel(logging.WARNING)
        # Disable all child loggers of urllib3, e.g. urllib3.connectionpool
        logging.getLogger("urllib3").propagate = False

    return logging.getLogger(__file__)
