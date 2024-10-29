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

* `kubectl` - `kubearchinspect` must be executed on a client with `kubectl` installed and configured to connect
  to the target Kubernetes cluster. If multiple clusters are configured, it will query the cluster in the current
  default context.
* `docker` client - The Docker credential store is used to authenticate to private registries, use [`docker login`](https://docs.docker.com/reference/cli/docker/login/) to add credentials.

### Usage

```console
kubearchinspect [OPTIONS]
```

### Options

* `images` : Check which images in your cluster support arm64
* `completion` : Generate the autocompletion script for the specified shell
* `help` : Help about any command

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

🚫 602401143452.dkr.ecr.eu-west-1.amazonaws.com/amazon-k8s-cni-init:v1.15.4-eksbuild.1
🚫 602401143452.dkr.ecr.eu-west-1.amazonaws.com/amazon-k8s-cni:v1.15.4-eksbuild.1
🚫 602401143452.dkr.ecr.eu-west-1.amazonaws.com/amazon/aws-network-policy-agent:v1.0.6-eksbuild.1
🚫 602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/aws-ebs-csi-driver:v1.26.0
🚫 602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/coredns:v1.9.3-eksbuild.10
🚫 602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/csi-attacher:v4.4.2-eks-1-28-11
🚫 602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/csi-node-driver-registrar:v2.9.2-eks-1-28-11
🚫 602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/csi-provisioner:v3.6.2-eks-1-28-11
🚫 602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/csi-resizer:v1.9.2-eks-1-28-11
🚫 602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/csi-snapshotter:v6.3.2-eks-1-28-11
🚫 602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/kube-proxy:v1.25.16-minimal-eksbuild.1
🚫 602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/livenessprobe:v2.11.0-eks-1-28-11
✅ amazon/aws-for-fluent-bit:2.10.0
✅ amazon/cloudwatch-agent:1.247350.0b251780
✅ busybox:1.31.1
✅ curlimages/curl:7.85.0
✅ docker.io/alpine:3.13
✅ docker.io/bitnami/external-dns:0.14.0-debian-11-r2
🆙 docker.io/bitnami/metrics-server:0.6.2-debian-11-r20
🚫 dsgcore--docker.internal.aws.arm.com/jcaap:3.7
✅ mirrors--internal.aws.arm.com/grafana/grafana:9.3.8
✅ mirrors--internal.aws.arm.com/banzaicloud/vault-secrets-webhook:1.18.0
🆙 quay.io/argoproj/argocd:v2.0.5
✅ quay.io/kiwigrid/k8s-sidecar:1.22.0
✅ quay.io/prometheus-operator/prometheus-config-reloader:v0.63.0
✅ quay.io/prometheus-operator/prometheus-operator:v0.63.0
✅ quay.io/prometheus/alertmanager:v0.25.0
✅ quay.io/prometheus/blackbox-exporter:v0.24.0
✅ quay.io/prometheus/node-exporter:v1.5.0
✅ quay.io/prometheus/prometheus:v2.42.0
✅ redis:6.2.4-alpine
✅ registry.k8s.io/autoscaling/cluster-autoscaler:v1.25.3
🚫 registry.k8s.io/ingress-nginx/controller:v1.9.4@sha256:5b161f051d017e55d358435f295f5e9a297e66158f136321d9b04520ec6c48a3
✅ registry.k8s.io/kube-state-metrics/kube-state-metrics:v2.8.1
```

### Errors in Output

At present the tool will show an error if it cannot connect to the repository for an image for any reason.
This could include:
* Authentication failure
* Communication error
* Image no longer available

## Private Registry Authentication

If `kubearchinspect` discovers an image from a registry that requires authentication, it uses the `docker` credential
store located at `~/.docker/config.json` to obtain the required credentials.

## Releases

For release notes and a history of changes of all releases, please see the following:

- [Changelog](CHANGELOG.md)

## Project Structure

The follow described the major aspects of the project structure:

* `cmd/` - Application command logic.
* `internal/` - Go project source files.
* `changes/` - Collection of news files for unreleased changes.

## Getting Help

* For a list of known issues and possible workarounds, please see [Known Issues](KNOWN_ISSUES.md).
* To raise a defect or enhancement please use [GitHub Issues](https://github.com/ArmDeveloperEcosystem/kubearchinspect/issues).

## Contributing

* We are committed to fostering a welcoming community, please see our
  [Code of Conduct](CODE_OF_CONDUCT.md) for more information.
* For ways to contribute to the project, please see the [Contributions Guidelines](CONTRIBUTING.md)
* For a technical introduction into developing this package, please see the [Development Guide](DEVELOPMENT.md)
