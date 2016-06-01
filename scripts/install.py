# Go Compile and install script for the Blue Lite API.

# Go version: 1.6.1
# Python 3.4

import os, sys
import subprocess
import platform
import shutil

arg = sys.argv[1]

work_dir = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
bin_dir = "{}".format(os.path.join(work_dir, 'cmd', 'blulite'))
arch = '{}_{}'.format(platform.system(), platform.machine())
output_dir = os.path.join(work_dir, '_output', arch, 'bin')
cmd_build = 'go build -o {}/blulite'.format(output_dir)

def handle_output(just_clean=False):
    shutil.rmtree(os.path.dirname(os.path.dirname(output_dir)), ignore_errors=True)
    if just_clean:
        return

    if not os.path.exists(output_dir):
        os.makedirs(output_dir)
        return

def install():
    handle_output()
    os.chdir(bin_dir)
    print('Installed in {}'.format(output_dir))
    build = subprocess.Popen(cmd_build,
                        shell=True,
                        stdout=subprocess.PIPE,
                        stderr=subprocess.PIPE)
    stdout, stderr = build.communicate()
    if stdout != b'':
        print(stdout)
    if stderr != b'':
        print(stderr)

if arg == 'install':
    install()

if arg == 'clean':
    handle_output(True)
