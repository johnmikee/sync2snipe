import os
import shutil
import subprocess
import sys

TAPE = sys.argv[1]
OUTPUT_GIF = sys.argv[2]


def shell_out(cmd: list) -> str:
    return subprocess.check_output(cmd).decode("utf-8").strip()


TOP_LEVEL = shell_out(["git", "rev-parse", "--show-toplevel"])
RESOURCE_DIR = f"{TOP_LEVEL}/resources"

os.chdir(TOP_LEVEL)

shutil.copy(src=f"{RESOURCE_DIR}/tapes/{TAPE}", dst=TOP_LEVEL)

shell_out(["vhs", sys.argv[1]])

os.remove(TAPE)

shutil.move(src=OUTPUT_GIF, dst=f"{RESOURCE_DIR}/gifs/{OUTPUT_GIF}")
