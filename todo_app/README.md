# todo_app

A minimal Go HTTP server that responds with "Hello, World!" at the root endpoint.

## Run locally

```bash
go run .
```

Server starts on port `8080` by default. Set the `PORT` environment variable to override.

## Build

```bash
go build -o todo_app .
./todo_app
```

## Docker

```bash
docker build -t todo_app .
docker run -p 8080:8080 todo_app
```

## Running in k3d

```
kubectl apply -f manifests/deployment.yaml
kubectl get po
kubectl logs todo-app-<STRING>
```
