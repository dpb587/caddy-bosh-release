---
title: Technical Notes
---

# Technical Notes


## Caddy

 * It is assumed that you have already read and agree to the Subscriber Agreement for your Certificate Authorities.
 * This release is focused on being a simple proxy server with certificate generation features, and it intentionally does not include most optional plugins of Caddy. For more advanced behavior, you should probably continue using your existing server, although you may still want to introduce this job for the [Automatic HTTPS](https://caddyserver.com/docs/automatic-https) behavior.


### Certificate Storage

By default, Caddy will manage keys and certificates on the local filesystem at `/var/vcap/store/caddy`. For simple, single-instance groups this is sufficient, but you must configure a persistent disk to retain the data across VM recreations. *Using this job with the default configuration with multiple instances may result in inefficient TLS connections and exceeding Certificate Authority rate limits.*

For more complex or multi-instance groups, you may want to use [Consul](https://www.consul.io/) as the credential backend by configuring additional environment variables ([learn more](https://caddyserver.com/docs/consul)).


## Let's Encrypt

 * [Let's Encrypt](https://letsencrypt.org/) is the default Certificate Authority.
 * You can find their Subscriber Agreement at [letsencrypt.org/repository](https://letsencrypt.org/repository/).
 * It is strongly recommended to use their [staging environment](https://letsencrypt.org/docs/staging-environment/) for testing before using in production. To change, set `acme.ca.url` to `https://acme-staging-v02.api.letsencrypt.org/directory`.


## Security Reminders

 * [Certificate Transparency](https://www.certificate-transparency.org/) logs are published by most Certificate Authorities, including [Let's Encrypt](https://letsencrypt.org/). This may contain the domain name or IP address from your certificate request.
