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

`kubearchinspect` will query the Kubernetes cluster in the current context.

**_NOTE:_** Kubernetes contexts can be shown using `kubectl config get-contexts` and set with `kubectl config set-context`.

Using the pre-built binary:
```sh
kubearchinspect images
```

or clone the repo and run:

```sh
go run . images
```

## Example Output

Output from a small cluster in EKS:

```
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

## Usage

```md
completion  : Generate the autocompletion script for the specified shell
help        : Help about any command
images      : Check which images in your cluster support arm64
```

## Private Registry Authentication

`kubearchinspect` uses the credential helper defined in `~/.docker/config.json` for authenticating with private registries.

## Releases

For release notes and a history of changes of all **production** releases, please see the following:

- [Changelog](CHANGELOG.md)

## Project Structure

The follow described the major aspects of the project structure:

- `cmd/` - Application command logic.
- `internal/` - Go project source files.
- `changes/` - Collection of news files for unreleased changes.
- `assets/` - Project images.

## Getting Help

- For a list of known issues and possible workarounds, please see [Known Issues](KNOWN_ISSUES.md).
- To raise a defect or enhancement please use [GitHub Issues](https://github.com/ArmDeveloperEcosystem/kubearchinspect/issues).

## Contributing

- We are committed to fostering a welcoming community, please see our
  [Code of Conduct](CODE_OF_CONDUCT.md) for more information.
- For ways to contribute to the project, please see the [Contributions Guidelines](CONTRIBUTING.md)
- For a technical introduction into developing this package, please see the [Development Guide](DEVELOPMENT.md)
