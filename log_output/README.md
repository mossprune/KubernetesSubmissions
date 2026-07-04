# log_output

Generates a random UUID on startup and prints it with a timestamp every 5 seconds.

## Running

```bash
docker run ghcr.io/mossprune/log_output:1.1
```

## Running in k3d

```
kubectl apply -f manifests/deployment.yaml
kubectl get po
kubectl logs log-output-<STRING>
```
