<!--
Copyright (C) 2024 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# Changes directory

This directory comprises information about all the changes that happened since the last release.

A news file should be added to this directory for each PR.

On release of the action, the content of the file becomes part of the [change log](../CHANGELOG.md) and this directory is reset.

### News Files

News files serve a different purpose to commit messages, which are generally written to inform developers of the
project. News files will form part of the release notes so should be written to target the consumer of the package or
tool.

- At least, one news file should be added for each Merge/Pull request to the directory `/changes`.
- The text of the file should be a single line describing the change and/or impact to the user.
- The filename of the news file should take the form `<number>.<extension>`, e.g, `20191231.feature` where:
  - The number is either the issue number or, if no issue exists, the date in the form `YYYYMMDDHHMM`.
  - The extension should indicate the type of change as described in the following table:

| Change Type                                                                                                             | Extension  | Version Impact  |
|-------------------------------------------------------------------------------------------------------------------------|------------|-----------------|
| Backwards compatibility breakages or significant changes denoting a shift direction.                                    | `.major`   | Major increment |
| New features and enhancements (non breaking).                                                                           | `.feature` | Minor increment |
| Bug fixes or corrections (non breaking).                                                                                | `.bugfix`  | Patch increment |
| Documentation impacting the consumer of the package (not repo documentation, such as this file, for this use `.misc`).  | `.doc`     | N/A             |
| Deprecation of functionality or interfaces (not actual removal, for this use `.major`).                                 | `.removal` | None            |
| Changes to the repository that do not impact functionality e.g. build scripts change.                                   | `.misc`    | None            |
