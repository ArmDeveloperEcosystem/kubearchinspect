# Armer

Check if container images in a Kubernetes cluster have arm architecture support.

## Setup

### Install Skopeo

```shell
brew install skopeo
```

### Setup a demo cluster (optional)

Use k3d to set up a demo cluster to test with.

```shell
k3d cluster create armer
```

### Deploy containers to the demo cluster

```shell
kubectl apply -f init/deployments.yaml
```

## Running the script

```shell
./armer.sh
```
