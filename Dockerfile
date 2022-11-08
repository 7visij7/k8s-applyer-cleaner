FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY smib-applyers-cleaner/ .
RUN env GOOS=linux GOARCH=amd64 go build -o smib-applyers-cleaner
CMD /app/smib-applyers-cleaner

############################
FROM scratch
WORKDIR /app
COPY config.yaml /app/config.yaml
COPY --from=builder /app/smib-applyers-cleaner /app/smib-applyers-cleaner
ENTRYPOINT ["/app/smib-applyers-cleaner"]
