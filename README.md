# Armer

Check if container images in a Kubernetes cluster have arm architecture support.

## Running

```shell
go run . images
```

It uses credential helper defined in `~/.docker/config.json` for authenticating with private registries.

## Drawbacks

- It can be slow in large clusters
- Does not handle rate-limiting
