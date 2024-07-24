# Armer

Check if container images in a Kubernetes cluster have arm architecture support.

## Running

```shell
go run . images
```

## Drawbacks

- Does not support registries which require authentication (E.g., AWS ECR)
- It can be slow in large clusters
- Does not handle rate-limiting
