FROM golang:1.26 AS builder

LABEL org.opencontainers.image.source=https://github.com/mossprune/KubernetesSubmissions

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o todo_app .

FROM scratch
COPY --from=builder /app/todo_app /todo_app
EXPOSE 8080
ENTRYPOINT ["/todo_app"]
