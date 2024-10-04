# KubeArchInspect

![kubearchinspect logo](./assets/kubearchinspect_logo-small.webp)

`kubearchinspect` is a utility to check if container images in a Kubernetes cluster have arm architecture support.

[![Main CI/CD](https://github.com/ArmDeveloperEcosystem/kubearchinspect/actions/workflows/main.yml/badge.svg)](https://github.com/ArmDeveloperEcosystem/kubearchinspect/actions/workflows/main.yml)

## Installation

### From binary

You can directly [download the kubearchinspect executable](https://github.com/ArmDeveloperEcosystem/kubearchinspect/releases).

### Build manually

Clone the repo and run:

```sh
go build
```

## Running

```sh
kubearchinspect images
```

or clone the repo and run:

```sh
go run . images
```

## Usage

```md
completion  : Generate the autocompletion script for the specified shell
help        : Help about any command
images      : Check which images in your cluster support arm64.
```

## Authentication

`kubearchinspect` uses the credential helper defined in `~/.docker/config.json` for authenticating with private registries.

## Drawbacks

- It can be slow in large clusters
- Does not handle rate-limiting
