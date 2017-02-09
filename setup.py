#!/usr/bin/env python

import os
import subprocess

from setuptools import setup

requirements = None
with open('requirements.txt') as reqsfile:
    requirements = reqsfile.readlines()

with open(os.path.join(os.path.dirname(__file__), 'README.rst')) as fp:
    LONG_DESCRIPTION = fp.read()

setup(
    name='px',
    version=subprocess.check_output(['git', 'describe', '--dirty']).decode('utf-8').strip(),
    description='Cross Functional Process Explorer',
    long_description=LONG_DESCRIPTION,
    author='Johan Walles',
    author_email='walles@gmail.com',
    url='https://github.com/walles/px',
    license='MIT',

    packages=['px'],

    install_requires=requirements,

    # See: http://setuptools.readthedocs.io/en/latest/setuptools.html#setting-the-zip-safe-flag
    zip_safe=True,

    setup_requires=[
        'pytest-runner',
        'pytest-cov',
        ],

    tests_require=[
        'pytest',
        ],

    entry_points={
        'console_scripts': [
            'px = px.px:main',
            'ptop = px.px:main'
        ],
    }
)
