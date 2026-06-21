FROM golang:1.26 AS builder

LABEL org.opencontainers.image.source=https://github.com/mossprune/KubernetesSubmissions

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o log_output .

FROM scratch
COPY --from=builder /app/log_output /log_output
ENTRYPOINT ["/log_output"]
