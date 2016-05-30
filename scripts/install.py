# Go Compile and install script for the Blue Lite API.

# Go version: 1.6.1
# Python 3.4

import os
import shlex
import subprocess

work_dir = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

bin_dir = "{}".format(os.path.join(work_dir, "cmd", "blulite"))
cmd_install = "go install ."

def execute_command():
    print('Installing...')
    os.chdir(bin_dir)
    build = subprocess.Popen(cmd_install,
                        shell=True,
                        stdout=subprocess.PIPE,
                        stderr=subprocess.PIPE)
    print('Done')

execute_command()
