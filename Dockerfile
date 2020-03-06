FROM busybox as install
WORKDIR /install
RUN wget -O /install/caddy "https://github.com/caddyserver/caddy/releases/download/v2.0.0-beta.15/caddy2_beta15_linux_amd64"
RUN chmod +x /install/caddy

FROM golang:1.14.0-buster as app
WORKDIR /build
COPY /server/main.go /build/main.go
RUN CGO_ENABLED=0 go build -o /build/server /build/main.go

FROM golang:1.14.0-buster as exec
WORKDIR /build
COPY /exec/main.go /build/main.go
RUN CGO_ENABLED=0 go build -o /build/exec /build/main.go

FROM gcr.io/distroless/base-debian10
COPY --from=install /install/caddy /caddy
COPY --from=app /build/server /server
COPY --from=exec /build/exec /exec
COPY ./localhost-key.pem key.pem
COPY ./localhost.pem cert.pem
COPY ./Caddyfile /etc/Caddyfile
CMD ["/exec"]