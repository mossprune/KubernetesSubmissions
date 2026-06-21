# log_output

Generates a random UUID on startup and prints it with a timestamp every 5 seconds.

## Running

```bash
docker run ghcr.io/mossprune/log_output:1.1
```

## Running in k3d

```
kubectl create deployment log-output  --image=ghcr.io/mossprune/log_output:1.1
kubectl get po
kubectl logs log-output-<STRING>
```