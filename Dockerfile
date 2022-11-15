FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY k8s-applyer-cleaner/ .
RUN env GOOS=linux GOARCH=amd64 go build -o k8s-applyer-cleaner
CMD /app/k8s-applyer-cleaner

############################
FROM scratch
WORKDIR /app
COPY config.yaml /app/config.yaml
COPY --from=builder /app/k8s-applyer-cleaner /app/k8s-applyer-cleaner
ENTRYPOINT ["/app/k8s-applyer-cleaner"]
