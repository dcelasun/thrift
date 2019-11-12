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

Coding Standards
================

    Anyone can write code that a computer can understand. Good programmers write code that humans can understand. -- Martin Fowler, 1999

The purpose of this document is to make everyone's life easier. It's easier when you read good, well formatted code with a clearly defined purpose. The only way to read clean code is to *write* clean code.

This document can help you achieve that, but keep in mind that nothing in this list is a silver bullet. Simply, focus on readability and imagine reading your own code ten years from now.

General Standards
-----------------

Thrift has a long history and not all existing code follows these rules, but we still want to improve the codebase over time. When making small changes do not refactor whole files as it makes the commit unreadable. If you are working on something larger, follow these rules as best you can.

When in doubt - reach out to other developers. Code review is the best way to improve readability.

- Use spaces not tabs
- Use only ASCII characters in file and directory names
- Use Unix-style line endings (LF) (on Windows use ``git config core.autocrlf true``)
- Stick to maximum 100 characters per line
- If not specified otherwise in a language specific standard, use 2 spaces for indentation
- Each file must start with a comment containing the `Apache License <https://www.apache.org/licenses/LICENSE-2.0>`__
- Public API of libraries should be documented, preferably using a format native to that language (Javadoc, Doxygen etc.)


Language Specific Standards
---------------------------

Some languages might have their own standards, see ``lib/LANG/coding_standards.md``.


