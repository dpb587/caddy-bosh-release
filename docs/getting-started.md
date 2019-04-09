---
title: Getting Started
---

# Getting Started


To start, you will need to include this release in your deployment's `releases` section. The [`caddy`](jobs/caddy) job manages the Caddy server and, at a minimum, you will need to configure the `caddyfile` property ([learn more](https://caddyserver.com/tutorial/caddyfile)) and attach a persistent disk to store the certificates. Assuming the instance is publicly accessible (for verification), the following will get a certificate for `caddy.example.com` and forward all requests to the server running at `localhost:8080`.

```yaml
caddyfile: |
  caddy.example.com
  proxy / localhost:8080 {
    transparent
  }
```

If you are using DNS to verify your domain name, you will also need to configure the `env` property with credentials for your DNS provider ([learn more](https://caddyserver.com/docs/automatic-https#dns-challenge)) and configure the `dns` setting within `caddyfile`.

If you are looking for a sample deployment to experiment with, start with the [`caddy.yml`](https://github.com/dpb587/caddy-bosh-release/blob/master/manifests/caddy.yml) manifest in a [test environment](https://bosh.io/docs/quick-start/).

    bosh -d caddy deploy manifests/caddy.yml


## Firewall

The following firewall changes may be necessary for ACME challenges:

 * egress udp/53 -- required for DNS verification
 * egress tcp/443 -- required for ACME certificate requests
 * ingress tcp/80 -- required for HTTP verification
