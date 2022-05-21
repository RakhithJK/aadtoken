FROM alpine:latest as builder

RUN apk add -U --no-cache ca-certificates

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY aadtoken /

ENTRYPOINT [ "/aadtoken" ]
