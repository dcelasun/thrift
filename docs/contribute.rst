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

How to Contribute
=================

Thank you for your interest in contributing to the Apache Thrift project! Information on why and how to contribute is available on the Apache Software Foundation (ASF) web site. In particular, we recommend the following to become acquainted with Apache Contributions:

- `Contributors Tech Guide <https://www.apache.org/dev/contributors>`__
- `Get Involved <https://www.apache.org/foundation/getinvolved.html>`__
- `Legal aspects on submission of contributions <https://www.apache.org/licenses/LICENSE-2.0.html#contributions>`__

Github Pull Requests
--------------------

This is the preferred method of submitting changes. When you submit a pull request through github, it activates the continuous integration (CI) build systems at Appveyor and Travis to build your changesxi on a variety of Linux and Windows configurations and run all the test suites. Follow these requirements for a successful pull request:

- All significant changes require a `Jira Ticket <https://issues.apache.org/jira/browse/THRIFT>`__. Trivial changes such as fixing a typo or a compiler warning do not.

- All pull requests should contain a single commit per issue, or we will ask you to squash it.

- The pull request title must begin with the Jira ticket identifier if it has an associated ticket, for example:

.. code::

    THRIFT-9999: an example pull request title

- Commit messages must follow this pattern for code changes (deviations will not be merged):

.. code::

    THRIFT-9999: [summary of fix, one line if possible]
    Client: [language(s) affected, comma separated, for example: "cpp,erl,perl"]

Instructions:

1. Fork `the repository <https://github.com/apache/thrift>`__.
2. Clone the fork to your development system.
3. Create a branch for your changes (best practice is issue as branch name, e.g. THRIFT-9999).
4. Modify the source to include the improvement/bugfix, and:

    1. Remember to provide *tests* for all submitted changes!
    2. Use test-driven development (TDD): add a test that will isolate the bug *before* applying the change that fixes it.
    3. Verify that you follow :ref:`Thrift Coding Standards <Coding Standards>` (you can run ``make style``, which ensures proper format for some languages).
    4. *[optional]* Verify that your change works on other platforms by adding a GitHub service hook to Travis CI and AppVeyor. You can use this technique to run the Thrift CI jobs in your account to check your changes before they are made public. Every GitHub pull request into Thrift will run the full CI build and test suite on your changes.
5. Squash your changes to a single commit. This maintains clean change history.
6. Commit and push changes to your branch (please use issue name and description as commit title, e.g. "THRIFT-9999: make it perfect"), with the affected languages on the next line of the description.
7. Use GitHub to create a pull request going from your branch to apache:master. Ensure that the Jira ticket number is at the beginning of the title of your pull request, same as the commit title.
8. Wait for other contributors or committers to review your new addition, and for a CI build to complete.
9. Wait for a committer to commit your patch. You can nudge the committers if necessary by sending a message to the :ref:`developer mailing list <Developer Mailing List>`.

Building Locally
----------------

For Windows systems, see our detailed instructions on the `CMake README <https://github.com/apache/thrift/blob/master/build/cmake/README.md>`__.

For Windows Native C++ builds, see our `Appveyor batch files <https://github.com/apache/thrift/tree/master/build/appveyor>`__.

For Docker, see our detailed instructions on the `Docker README <https://github.com/apache/thrift/blob/master/build/docker/README.md>`__.

Building from Source
^^^^^^^^^^^^^^^^^^^^

If you are building from the first time out of the source repository, you will need to generate the configure scripts. (This is not necessary if you downloaded a released tarball.) From the top directory, do:

.. code::

    ./bootstrap.sh

Once the configure scripts are generated, thrift can be configured. From the top directory, do:

.. code::

    ./configure

Disable languages for which you don't want to generate a library. This does *not* include the code generator:

.. code::

    ./configure --without-java

You may need to specify the location of the boost files explicitly. If you installed boost in /usr/local, you would run configure as follows:

.. code::

    ./configure --with-boost=/usr/local

If you want to override the automatic Java SDK detection, use the JAVAC environment variable:

.. code::

    ./configure JAVAC=/usb/bin/javac

Note that by default the Thrift C++ library is built with debugging symbols included. If you want to customize these options you should use the CXXFLAGS option in configure:

.. code::

    ./configure CXXFLAGS='-g -O2'
    ./configure CFLAGS='-g -O2'
    ./configure CPPFLAGS='-DDEBUG_MY_FEATURE'

To see other configuration options run:

.. code::

    ./configure --help

Once configured, you can build Thrift with make:

.. code::

    make

and run the test suite:

.. code::

    make check

and the cross language test suite:

.. code::

    sh test/test.sh

Finally, to install the generated binary become root and do:

.. code::

    make install

Note that some language packages must be installed manually using build tools better suited to those languages (Java, Ruby, PHP etc).

Look for the README file in the ``lib/<language>/`` folder for more details on the installation of each language library package.


Reviewing Open Issues
---------------------

If you'd like to help by doing reviews, you have two options:

- Review the `Pull Request backlog <https://github.com/apache/thrift/pulls>`__ on Github. Code reviews are open to all.
- Review the `Jira <https://issues.apache.org/jira/browse/THRIFT>`__ issue tracker. You can search for tickets relating to languages you are interested in or currently using with thrift, for example a Jira search (Issues -> Search For Issues) query of ``project = THRIFT AND component in ("Erlang - Library") and status not in (resolved, closed)`` will locate all open Erlang Library issues.

Reporting Bugs
--------------

- Check if the issue is already in the `Jira <https://issues.apache.org/jira/browse/THRIFT>`__ issue tracker.
- If not, create a ticket describing the problem.


Adding a New Language
---------------------

When considering new language bindings, there are certain points to think about. First, you should find out if you are about to implement completely new language bindings that are not yet supported with Thrift, or if you just want to add support for a specific "flavour" of an already implemented language.

If bindings for a language exist, but are incomplete / missing support for a feature,  it is recommended to add the new feature as an option to the existing language. For a good model on how to do this take look at the js/nodejs implementations, or the various options that already exist for Python. ``thrift --help`` is a good point to start. Depending on the amount of changes necessary, you will still find the rest of the document useful.

Preparation
^^^^^^^^^^^

The good news is, although there is some work required, the process is not as hard as it looks. First, make sure you have a fully functional build environment and are able to :ref:`build the compiler from source <Building from Source>`.

Next, search `Jira <https://issues.apache.org/jira/browse/THRIFT>`__ and maybe the mailing list archives. If you do not find anything similar, create a new ticket, shortly describing what you are planning. If you are not quite sure, ask on the :ref:`developers mailing list <Developer Mailing List>`.

Now fork and clone the repository from Github and switch to a new branch for development.

Minimal Feature Set
^^^^^^^^^^^^^^^^^^^

- Implement the `code generator <https://github.com/apache/thrift/tree/master/compiler/cpp/src/generate>`__, typically by picking one from the existing pool that is close to what you need. There are already plenty of languages supported, and you'll find OOP, procedural and functional styles. If you are not sure which one to choose, head to the next point.

- Implement the Thrift library for that particular language, again by picking one of the `existing libraries <https://github.com/apache/thrift/tree/master/lib>`__ as a starting point. Because the libraries differ largely with regard to the depth/featureset of their implementations it is recommended to have a closer look on what is implemented, and what is not.

- Implement the standardized `Thrift cross platform test <https://github.com/apache/thrift/tree/master/test>`__ and make sure all tests succeed when run against at least one other language. This ensures interoperability and makes sure that the code does not only work when talking to itself (= same language). You may also add other tests, but these should be put into ``lib/yourlang/test`` rather than ``test/yourlang`` - the latter is intended to host solely the cross language tests.

- Implement the `tutorial code <https://github.com/apache/thrift/tree/master/tutorial>`__ and test it against another language. If you did everything well, this last step is relatively easy.

The minimum required feature set should cover at least:

- Transports: Sockets, Buffered, Framed required, HTTP client recommended
- Protocols: Binary and Multiplex required, JSON recommended
- Server types: SimpleServer required

Recommended Features
^^^^^^^^^^^^^^^^^^^^

These are not strictly required in the first run, but are commonly used. Depending on the language, some things may be easier to implement than others. Alternatively, consider adding the features listed below later as additional contributions instead of trying to press them into the initial contribution. **If in doubt, focus on quality rather than quantity.**

- Transports: HTTP server, Pipes, NamedPipes (where it makes sense)
- Protocols: Compact
- Server types: Nonblocking, Threaded and/or Threadpool server implementation

Final Steps
^^^^^^^^^^^

- Add a ``README.md`` file to your library's folder, describing requirements, dependencies and whatever else might be important. Look at the existing readme files if you aren't sure.

- Make sure the generator, library, tests and tutorial have proper ``makefile.am`` files. Include everything into the build/test scripts. If you need help with these steps, don't hesitate to ask on the mailing lists.

- All done! Open a pull request with your changes.

Notes
^^^^^

These steps are not a linear process, but an iterative one. Even if the code that comes out of the generator compiles, doesn't spit out any warnings and seemingly runs fine, you'll still run into problems as you work on implementing the library and tests. This is normal and expected.

It's a good idea to create the JIRA ticket and post your work early. This serves not only as announcement, but also provides a good starting point for other people who might be looking for that feature. You'll be able to get valuable early feedback and get support with implementation and testing. Having someone else look over your code is a good thing and will ensure high code quality.

Happy coding!




