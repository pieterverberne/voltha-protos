# -*- coding: utf-8 -*-
# -----------------------------------------------------------------------
# Copyright 2019-2024 Open Networking Foundation Contributors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# -----------------------------------------------------------------------
# SPDX-FileCopyrightText: 2019-2024 Open Networking Foundation Contributors
# SPDX-License-Identifier: Apache-2.0
# -----------------------------------------------------------------------

from setuptools import setup, find_packages


def version():
    with open('VERSION') as f:
        return f.read()


setup(
    name='voltha-protos',
    version=version(),
    description='Protobuf interface definitions',
    author='VOLTHA project',
    author_email='voltha-dev@opencord.org',
    url='https://gerrit.opencord.org/gitweb?p=voltha-protos.git',
    license='Apache Software License',
    classifiers=[
        'Development Status :: 5 - Production/Stable',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: Apache Software License',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6',
        'Programming Language :: Python :: 3.7',
    ],
    keywords='protobuf voltha',
    packages = find_packages(where="python"),
    package_dir = {"": "python"},
    install_requires = [
        "grpcio==1.66.1",
        "protobuf==5.28.1",
        "grpcio-tools==1.66.1",
        "googleapis-common-protos==1.65.0"
    ],
    include_package_data=True,
)
