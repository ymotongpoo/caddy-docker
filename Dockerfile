FROM busybox as install
WORKDIR /install
RUN wget -O /install/caddy "https://github.com/caddyserver/caddy/releases/download/v2.0.0-beta.15/caddy2_beta15_linux_amd64"
RUN chmod +x /install/caddy

FROM gcr.io/distroless/base-debian10
COPY --from=install /install/caddy /caddy
COPY ./Caddyfile /etc/Caddyfile
CMD ["/caddy", "run", "--config", "/etc/Caddyfile", "--adapter", "caddyfile"]
