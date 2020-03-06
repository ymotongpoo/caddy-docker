# caddy reverse proxy in docker sample

This is sample container image that achieve:

* Run [caddy](https://caddyserver.com/) as revserse proxy to the app
* The app is the normal HTTP server implmeneted in Go
* caddy handles all TLS connection at its edge
* TLS should be handled by caddy's auto certification renewal feature (TODO)

## Prerequisites

In order to run this proof of concept sample, you need:

* docker
* [mkcert](https://github.com/FiloSottile/mkcert) [^cert]

[^cert]: Though you can generate self signed certificate with any tools, mkcert let you have your own CA that Chrome/Chromium can trust.

## How to run

0. `$ mkcert -install` (In the case you run `mkcert` for the first time)
1. `$ mkcert localhost`
2. `$ docker build -t caddy-test .`
3. `$ docker run -p 9000:9000 --rm -it caddy-test:latest`
4. From another terminal: `$ curl -X GET https://localhost:9000/`

On executing #4, you'll see the following message:

> hello test server
