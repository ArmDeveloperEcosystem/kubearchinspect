# Armer

Check if container images in a Kubernetes cluster have arm architecture support.

https://github.com/Arm-Debug/armer/assets/32394735/d89f6692-1e42-4790-8908-746a5cf9c7b6


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
