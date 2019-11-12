 .. Licensed to the Apache Software Foundation (ASF) under one
    or more contributor license agreements.  See the NOTICE file
    distributed with this work for additional information
    regarding copyright ownership.  The ASF licenses this file
    to you under the Apache License, Version 2.0 (the
    "License"); you may not use this file except in compliance
    with the License.  You may obtain a copy of the License at

 ..   http://www.apache.org/licenses/LICENSE-2.0

 .. Unless required by applicable law or agreed to in writing,
    software distributed under the License is distributed on an
    "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, either express or implied.  See the License for the
    specific language governing permissions and limitations
    under the License.

Tutorial
========

**Install Apache Thrift**

Many Linux distributions have Apache Thrift packaged in their repositories and these packages should be preferred.

MacOS users can use package managers like `Homebrew <https://formulae.brew.sh/formula/thrift>`__ to install Thrift.

Windows users can :ref:`download <Download>` the compiler from our download page.

**Manual Installation**

Alternatively, you can build Thrift from source. Follow our :ref:`building from source <Building From Source>` guide.

**Writing a .thrift file**

After the compiler is installed you will need to create a ``.thrift`` file. This file is an interface definition made up of Thrift types and Services. The services you define in this file are implemented by the server and are called by any clients.

In this tutorial, we'll be using `tutorial.thrift <https://github.com/apache/thrift/blob/master/tutorial/tutorial.thrift>`__ and `shared.thrift <https://github.com/apache/thrift/blob/master/tutorial/shared.thrift>`__.

**Generate Source Code**

The Thrift compiler is used to compile your Thrift file into source code used by your server code and client libraries.

.. code::

    thrift --gen <language> tutorial.thrift

To recursively compile all files imported by a Thrift file, run:

.. code::

    thrift -r --gen <language> tutorial.thrift

The sample `tutorial.thrift <https://github.com/apache/thrift/blob/master/tutorial/tutorial.thrift>`__ file defines a basic calculator service, which includes another file called `shared.thrift <https://github.com/apache/thrift/blob/master/tutorial/shared.thrift>`__. Both files will be used to demonstrate how to build a Thrift client and server pair.


Language specific implementations of this tutorial can be found on `Github <https://github.com/apache/thrift/tree/master/tutorial>`__.



