<!--
Copyright (C) 2024 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# Development and Testing
## Local Development

### Build manually

Clone the repo and run:

```sh
go build
```

### Running tests

To run all tests:
```bash
go test ./...
```

### Static analysis and linting

Static analysis tools and linters are run as part of CI.
They come from [golangci-lint](https://golangci-lint.run/). To run this locally:
```bash
# Must be in a directory with a go.mod file
cd <directory_with_go_module>
golangci-lint run ./...
```

# Releasing

### Release workflow

1. Navigate to the [GitHub Actions](https://github.com/ArmDeveloperEcosystem/kubearchinspect/actions/workflows/release.yml) page.
2. Select the **Run Workflow** button.

### Version Numbers

The version number will be automatically calculated, based on the news files.
