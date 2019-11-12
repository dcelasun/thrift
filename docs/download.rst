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

Download
========

Release
-------

The latest stable release of Thrift is 0.13.0 (released on 2019-OCT-16).

- `thrift-0.13.0.tar.gz <http://www.apache.org/dyn/closer.cgi?path=/thrift/0.13.0/thrift-0.13.0.tar.gz>`__ [`PGP <https://www.apache.org/dist/thrift/0.13.0/thrift-0.13.0.tar.gz.asc>`__] | [`MD5 <https://www.apache.org/dist/thrift/0.13.0/thrift-0.13.0.tar.gz.md5>`__]
- `Thrift compiler for Windows (thrift-0.13.0.exe) <http://www.apache.org/dyn/closer.cgi?path=/thrift/0.13.0/thrift-0.13.0.exe>`__ [`PGP <https://www.apache.org/dist/thrift/0.13.0/thrift-0.13.0.exe.asc>`__] | [`MD5 <https://www.apache.org/dist/thrift/0.13.0/thrift-0.13.0.exe.md5>`__]

When downloading from a mirror, please be sure to `verify <https://www.apache.org/info/verification.html>`__ the checksums and signature (see the MD5 and PGP links above). The `KEYS <https://www.apache.org/dist/thrift/KEYS>`__ file contains the public key(s) used for signing releases.

Git Checkout
------------
For those who would like to participate in Thrift development, you may checkout Thrift from the GitHub repository.
.. code::

    $ git clone https://github.com/apache/thrift.git
    $ cd thrift

We recommend you use our `Docker development environment <https://github.com/apache/thrift/blob/master/build/docker/README.md>`__ - the same environment the CI builds use.

Package Managers
----------------

Thrift is available in various package managers.

- **All Languages**

  - Package Manager: Docker
  - Direct Link: https://hub.docker.com/_/thrift/
  - Control File: Dockerfile
  - Remarks: Unofficial package. Compiler in /usr/local/bin/thrift

- **ActionScript**

  - Package Manager: Maven
  - Direct Link: https://repository.apache.org/#nexus-search;quick~libthrift-as3
  - Control File: `build.gradle <https://github.com/apache/thrift/blob/master/lib/as3/build.gradle>`__
  - Maintainers: jking



- **C++**

  - Package Manager: Conan
  - Remarks: Not supported yet, see `THRIFT-4800 <https://issues.apache.org/jira/browse/THRIFT-4800>`__



- **C#** and **.NET Standard**

    - Package Manager: NuGet
    - Direct Link: https://www.nuget.org/packages/ApacheThrift
    - Control File: `ApacheThrift.nuspec <https://github.com/apache/thrift/blob/master/ApacheThrift.nuspec>`__
    - Maintainers: jfarrell, codesf, jking

- **D**

    - Package Manager: dub
    - Direct Link: https://code.dlang.org/packages/apache-thrift
    - Control File: `dub.json <https://github.com/apache/thrift/blob/master/dub.json>`__
    - Maintainers: jking

- **Dart**

    - Package Manager: Pub
    - Direct Link: https://pub.dartlang.org/packages/thrift
    - Control File: `pubspec.yaml <https://github.com/apache/thrift/blob/master/lib/dart/pubspec.yaml>`__
    - Maintainers: jking

- **Erland**

    - Package Manager: Hex PM
    - Direct Link: https://hex.pm/packages?search=thrift&sort=downloads
    - Remarks: No official ASF package available

- **Haskell**

    - Package Manager: Hackage
    - Direct Link: https://hackage.haskell.org/package/thrift
    - Control File: `thrift.cabal <https://github.com/apache/thrift/blob/master/lib/hs/thrift.cabal>`__
    - Maintainers: jfarrell, clavoie, jking

- **Haxe**

    - Package Manager: Haxelib
    - Control File: `haxelib.json <https://github.com/apache/thrift/blob/master/lib/haxe/haxelib.json>`__
    - Maintainers: jensg
    - Remarks: No official ASF package yet, see `THRIFT-3036 <https://issues.apache.org/jira/browse/THRIFT-3036>`__

- **Go** (golang)

    - Package Manager: Go modules
    - Control File: `go.mod <https://github.com/apache/thrift/blob/master/lib/go/thrift/go.mod>`__
    - Maintainers: dcelasun

- **Java**

    - Package Manager: Maven
    - Direct Link: https://repository.apache.org/#nexus-search;quick~org.apache.thrift
    - Control File: `gradle.properties <https://github.com/apache/thrift/blob/master/lib/java/gradle.properties>`__
    - Maintainers: jking

- **JavaScript**

    - Package Manager: Bower
    - Direct Link: https://libraries.io/bower/thrift
    - Control File: `bower.json <https://github.com/apache/thrift/blob/master/bower.json>`__

- **Lua**

    - Package Manager: LuaRocks
    - Direct Link: https://luarocks.org/modules/drauschenbach/thrift
    - Remarks: No official ASF package yet, see `THRIFT-4708 <https://issues.apache.org/jira/browse/THRIFT-4708>`__

- **Node.js**

    - Package Manager: NPM
    - Direct Link: https://www.npmjs.com/package/thrift
    - Control File: `package.json <https://github.com/apache/thrift/blob/master/package.json>`__
    - Maintainers: wadey, jking

- **OCaml**

    - Package Manager: opam
    - Direct Link: https://opam.ocaml.org/packages/thrift/
    - Control File: `opam <https://github.com/apache/thrift/blob/master/lib/ocaml/opam>`__
    - Remarks: No official ASF package yet, see `THRIFT-4706 <https://issues.apache.org/jira/browse/THRIFT-4706>`__

- **Perl**

    - Package Manager: CPAN
    - Direct Link: https://metacpan.org/release/Thrift
    - Control File: `Makefile.PL <https://github.com/apache/thrift/blob/master/lib/perl/Makefile.PL>`__
    - Maintainers: jking

- **PHP**

    - Package Manager: Composer
    - Direct Link: https://packagist.org/packages/apache/thrift
    - Control File: `composer.json <https://github.com/apache/thrift/blob/master/composer.json>`__
    - Maintainers: jfarrell, bufferoverflow, jking

- **Python**

    - Package Manager: PyPi
    - Direct Link: https://pypi.python.org/pypi/thrift
    - Control File: `setup.py <https://github.com/apache/thrift/blob/master/lib/py/setup.py>`__
    - Maintainers: jfarrell
    - Remarks: Outdated, see `THRIFT-4687 <https://issues.apache.org/jira/browse/THRIFT-4687>`__

- **Ruby**

    - Package Manager: Ruby Gem
    - Direct Link: https://rubygems.org/gems/thrift
    - Control File: `thrift.gemspec <https://github.com/apache/thrift/blob/master/lib/rb/thrift.gemspec>`__
    - Maintainers: jfarrell
    - Remarks: Outdated, see `THRIFT-4707 <https://issues.apache.org/jira/browse/THRIFT-4707>`__

- **Rust**

    - Package Manager: Cargo
    - Direct Link: https://crates.io/crates/thrift
    - Control File: `cargo.toml <https://github.com/apache/thrift/blob/master/lib/rs/Cargo.toml>`__
    - Maintainers: allengeorge, jfarrell
    
- **Common LISP**, **Smalltalk** and **Swift**

    - Remarks: No official ASF package available for these languages.
