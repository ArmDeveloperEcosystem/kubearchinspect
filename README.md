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
Legends:
✅ - Supports arm64, ❌ - Does not support arm64, ⬆ - Upgrade for arm64 support, ❗ - Some error occurred
------------------------------------------------------------------------------------------------

602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/coredns:v1.9.3-eksbuild.10 ❗
602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/csi-snapshotter:v6.3.2-eks-1-28-11 ❌
quay.io/kiwigrid/k8s-sidecar:1.21.0 ✅
grafana/grafana:9.3.1 ✅
redis:6.2.4-alpine ✅
602401143452.dkr.ecr.eu-west-1.amazonaws.com/amazon/aws-network-policy-agent:v1.0.6-eksbuild.1 ❗
registry.k8s.io/autoscaling/cluster-autoscaler:v1.25.3 ✅
602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/csi-node-driver-registrar:v2.9.2-eks-1-28-11 ❌
docker.io/bitnami/metrics-server:0.6.2-debian-11-r20 ⬆
amazon/aws-for-fluent-bit:2.10.0 ✅
quay.io/argoproj/argocd:v2.0.5 ⬆
quay.io/prometheus/node-exporter:v1.5.0 ✅
registry.k8s.io/ingress-nginx/controller:v1.9.4@sha256:5b161f051d017e55d358435f295f5e9a297e66158f136321d9b04520ec6c48a3 ❗
quay.io/prometheus-operator/prometheus-operator:v0.63.0 ✅
registry.k8s.io/kube-state-metrics/kube-state-metrics:v2.8.1 ✅
mirrors--ghcr-io.eu-west-1.artifactory.aws.arm.com/banzaicloud/vault-secrets-webhook:1.18.0 ✅
quay.io/prometheus-operator/prometheus-config-reloader:v0.63.0 ✅
mirrors--dockerhub.eu-west-1.artifactory.aws.arm.com/grafana/grafana:9.3.8 ✅
curlimages/curl:7.85.0 ✅
602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/csi-attacher:v4.4.2-eks-1-28-11 ❗
602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/livenessprobe:v2.11.0-eks-1-28-11 ❗
busybox:1.31.1 ✅
quay.io/prometheus/prometheus:v2.42.0 ✅
docker.io/bitnami/external-dns:0.14.0-debian-11-r2 ✅
dsgcore--docker.eu-west-1.artifactory.aws.arm.com/jcaap:3.7 ❗
602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/csi-provisioner:v3.6.2-eks-1-28-11 ❗
602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/csi-resizer:v1.9.2-eks-1-28-11 ❗
602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/kube-proxy:v1.25.16-minimal-eksbuild.1 ❗
quay.io/kiwigrid/k8s-sidecar:1.22.0 ✅
quay.io/prometheus/blackbox-exporter:v0.24.0 ✅
amazon/cloudwatch-agent:1.247350.0b251780 ✅
602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/aws-ebs-csi-driver:v1.26.0 ❗
sergrua/kube-tagger:release-0.1.1 ❌
docker.io/alpine:3.13 ✅
quay.io/prometheus/alertmanager:v0.25.0 ✅
602401143452.dkr.ecr.eu-west-1.amazonaws.com/amazon-k8s-cni-init:v1.15.4-eksbuild.1 ❗
602401143452.dkr.ecr.eu-west-1.amazonaws.com/amazon-k8s-cni:v1.15.4-eksbuild.1 ❗
```

## Usage

```md
completion  : Generate the autocompletion script for the specified shell
help        : Help about any command
images      : Check which images in your cluster support arm64
```

## Private Registry Authentication

`kubearchinspect` uses the credential helper defined in `~/.docker/config.json` for authenticating with private registries.

## Drawbacks

- It can be slow in large clusters
- Does not handle rate-limiting
