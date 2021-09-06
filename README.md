## Intro

inlets&reg;^H^H^H^H^H^H^Hcevka is how you connect services between different networks. You won't have to think about managing firewalls, NAT or VPNs again. Services can be tunnelled securely over a websocket and accessed on a remote network privately, or exposed on the Internet using an exit-server (5-10USD / mo).

Why do we need this project? Similar tools such as [inlets](https://inlets.dev/), [ngrok](https://ngrok.com/) and [Argo Tunnel](https://developers.cloudflare.com/argo-tunnel/) from [Cloudflare](https://www.cloudflare.com/) are closed-source, have limits built-in, can work out expensive, and have limited support for arm/arm64, Docker and Kubernetes. Ngrok's domain is also often banned by corporate firewall policies meaning it can be unusable. Other open-source tunnel tools are designed to only set up a single static tunnel.

With inlets^H^H^H^H^H^Hcevka you can set up your own self-hosted tunnel, copy over the static binary and start tunnelling traffic without any arbitrary limits or restrictions. When used with TLS, inlets^H^H^H^H^H^Hcevka can be used with most corporate HTTP proxies.

![Conceptual diagram](docs/inlets.png)

_Conceptual diagram for inlets^H^H^H^H^H^Hcevka_

## About inlets^H^H^H^H^H^Hcevka OSS

inlets^H^H^H^H^H^Hcevka OSS uses a websocket to create a tunnel between a client and a server. The server is typically a machine with a public IP address, and the client is on a private network with no public address.

### Features

- Tunnel HTTP or websockets
- Expose a site on through the use of a DNS entry and a `Host` header
- Shared authentication token for the client and server
- Automatic reconnects for when the connection drops

Distribution:

- Source code only (for now)
