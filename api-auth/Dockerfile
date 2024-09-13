FROM alpine:3.18 as root-certs
RUN apk update
RUN apk add -U --no-cache ca-certificates
RUN addgroup -g 1001 app
RUN adduser app -u 1001 -D -G app /home/app

FROM golang:1.23.1 as builder
# Instalar gcc y otras dependencias necesarias
RUN apt-get update && apt-get install -y gcc musl-dev

WORKDIR /gpt-playground-api-files
COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./server cmd/api/main.go

FROM scratch as final
COPY --from=root-certs /etc/passwd /etc/passwd
COPY --from=root-certs /etc/group /etc/group
COPY --chown=1001:1001 --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=1001:1001 --from=builder /gpt-playground-api-files/server /server
USER app
ENTRYPOINT ["/server"]
