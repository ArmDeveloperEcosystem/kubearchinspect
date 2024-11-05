<!--
Copyright (C) 2024 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->

# KubeArchInspect

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Main CI/CD](https://github.com/ArmDeveloperEcosystem/kubearchinspect/actions/workflows/main.yml/badge.svg)](https://github.com/ArmDeveloperEcosystem/kubearchinspect/actions/workflows/main.yml)

## Overview

Migrating your websites and services to run on Arm infrastructure can bring benefits in cost savings and performance improvements. The first phase in migrating to Arm is to determine whether the container images in the Kubernetes cluster have support for the Arm architecture. It can be a manual and time consuming task to check compatibility. To make it easier we have developed the `kubearchinspect` tool which automates this process. The tool will check the metadata of your images for Arm architecture support and if the current version of the image lacks support, it will also check newer versions for compatibility so that you can easily upgrade.

This is Open Source Software and we appreciate contributions and feedback, please see the [Contribution Guidelines](CONTRIBUTING.md) for more information.

## Installation

Pre-built binaries are available from the [releases page](https://github.com/ArmDeveloperEcosystem/kubearchinspect/releases), alternatively see the [Development Guide](DEVELOPMENT.md) for building locally.

## Running

### Prerequsites

- `kubectl` - `kubearchinspect` must be executed on a client with `kubectl` installed and configured to connect
  to the target Kubernetes cluster. If multiple clusters are configured, it will query the cluster in the current
  default context.
- `docker` client - The Docker credential store is used to authenticate to private registries, use [`docker login`](https://docs.docker.com/reference/cli/docker/login/) to add credentials.

### Usage

```console
kubearchinspect [OPTIONS]
```

### Options

- `images` : Check which images in your cluster support arm64
- `completion` : Generate the autocompletion script for the specified shell
- `help` : Help about any command

## Example

Output from a small cluster in EKS:

```console
% kubearchinspect images

Legend:
-------
✅ - arm64 supported
🆙 - arm64 supported (with update)
❌ - arm64 not supported
🚫 - error occurred
------------------------------------------------------------------------------------------------

🚫 some.ecr.link.com/eks/csi-snapshotter:v1 || Authentication error.
🚫 some.ecr.link.com/eks/kube-proxy:v1 || Authentication error.
🚫 some.ecr.link.com/eks/livenessprobe:v1 || Authentication error.
✅ amazon/aws-for-fluent-bit:1
✅ amazon/cloudwatch-agent:1
✅ busybox:1
✅ curlimages/curl:1
✅ docker.io/alpine:1
✅ docker.io/bitnami/external-dns:1
🆙 docker.io/bitnami/metrics-server:1
🚫 some.mirror.link.com/jcaap:1 || Image not found.
✅ mirrors--internal.aws.arm.com/grafana/grafana:1
✅ mirrors--internal.aws.arm.com/banzaicloud/vault-secrets-webhook:1
🆙 quay.io/argoproj/argocd:1
✅ quay.io/kiwigrid/k8s-sidecar:1
✅ quay.io/prometheus-operator/prometheus-config-reloader:1
✅ quay.io/prometheus-operator/prometheus-operator:1
✅ quay.io/prometheus/alertmanager:1
✅ quay.io/prometheus/blackbox-exporter:1
✅ quay.io/prometheus/node-exporter:1
✅ quay.io/prometheus/prometheus:1
✅ redis:1
✅ registry.k8s.io/autoscaling/cluster-autoscaler:1
✅ registry.k8s.io/kube-state-metrics/kube-state-metrics:1
```

### Errors in Output

If there is an error checking an image the tool will display the 🚫 symbol and give a short description of the error at the end of the line. The current common errors are:

- Authentication error. This usually means that you are missing your docker credentials or that you are denied access to the registry.
- Communication error. Could mean that the URL host doesn't exist or that its down.
- Image not found.
- Unknown error, run in debug mode for more info -d

## Private Registry Authentication

If `kubearchinspect` discovers an image from a registry that requires authentication, it uses the `docker` credential
store located at `~/.docker/config.json` to obtain the required credentials.

## Releases

For release notes and a history of changes of all releases, please see the following:

- [Changelog](CHANGELOG.md)

## Project Structure

The follow described the major aspects of the project structure:

- `cmd/` - Application command logic.
- `internal/` - Go project source files.
- `changes/` - Collection of news files for unreleased changes.

## Getting Help

- For a list of known issues and possible workarounds, please see [Known Issues](KNOWN_ISSUES.md).
- To raise a defect or enhancement please use [GitHub Issues](https://github.com/ArmDeveloperEcosystem/kubearchinspect/issues).

## Contributing

- We are committed to fostering a welcoming community, please see our
  [Code of Conduct](CODE_OF_CONDUCT.md) for more information.
- For ways to contribute to the project, please see the [Contributions Guidelines](CONTRIBUTING.md)
- For a technical introduction into developing this package, please see the [Development Guide](DEVELOPMENT.md)
