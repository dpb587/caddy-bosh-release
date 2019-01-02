# caddy-bosh-release

A [BOSH release](https://bosh.io/) to run [Caddy](https://caddyserver.com/), a basic web and proxy server with support for dynamically provisioning certificates. You may find this useful if you are looking for an easy method to front another BOSH job with a certificate from [Let's Encrypt](https://letsencrypt.org/).

*Note:* key pairs are managed on each VM, so you should only use this for instance groups with a single instance. Using this job across multiple instances may result in inefficient TLS connections and exceeding Certificate Authority rate limits.


## Getting Started

To start, you will need to include this release in your deployment's `releases` section. The [`caddy`](jobs/caddy) job manages the Caddy server and, at a minimum, you will need to configure the `caddyfile` property ([learn more](https://caddyserver.com/tutorial/caddyfile)) and attach a persistent disk to store the certificates. Assuming the instance is publicly accessible (for verification), the following will get a certificate for `caddy.example.com` and forward all requests to the server running at `localhost:8080`.

    caddyfile: |
      caddy.example.com
      proxy / localhost:8080 {
        transparent
      }

If you are using DNS to verify your domain name, you will also need to configure the `env` property with credentials for your DNS provider ([learn more](https://caddyserver.com/docs/automatic-https#dns-challenge)) and configure the `dns` setting within `caddyfile`.

    caddyfile: |
      caddy.example.com
      proxy / localhost:8080 {
        transparent
      }
      tls {
        dns route53
      }
    env:
      AWS_ACCESS_KEY_ID: AKIA...
      AWS_SECRET_ACCESS_KEY: ...
      AWS_HOSTED_ZONE_ID: Z... # optional

If you are using Google Cloud DNS, you will configure your service account key as a json string.

    env:
      GCE_PROJECT: sample-project-001
      GCE_SERVICE_ACCOUNT_FILE: |
        {
          "type": "service_account",
          "project_id": "sample-project-001",
          "private_key_id": "sample4c02016ac2d8abf5b1577993ded31626a8",
          "private_key": "-----BEGIN PRIVATE KEY-----\nMIIE...k8LA==\n-----END PRIVATE KEY-----\n",
          "client_email": "caddy@sample-project-001.iam.gserviceaccount.com",
          "client_id": "sample012345678901234",
          "auth_uri": "https://accounts.google.com/o/oauth2/auth",
          "token_uri": "https://accounts.google.com/o/oauth2/token",
          "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
          "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/caddy%40sample-project-001.iam.gserviceaccount.com"
        }

If you are looking for a sample deployment to experiment with, start with the [`caddy.yml`](manifests/caddy.yml) manifest in a [test environment](https://bosh.io/docs/quick-start/).

    bosh -d caddy deploy manifests/caddy.yml


### Examples

Some examples of putting Caddy in front of other BOSH releases and their relevant configuration properties.

 * [cloudfoundry/bosh](https://github.com/cloudfoundry/bosh) - [`manifests/examples/bosh-ops.yml`](manifests/examples/bosh-ops.yml)
 * [cloudfoundry/uaa-release](https://github.com/cloudfoundry/uaa-release) - [`manifests/examples/uaa-ops.yml`](manifests/examples/uaa-ops.yml)
 * [concourse/concourse](https://github.com/concourse/concourse) - [`manifests/examples/concourse-ops.yml`](manifests/examples/concourse-ops.yml)
 * [dpb587/ssoca-bosh-release](https://github.com/dpb587/ssoca-bosh-release) - [`manifests/examples/ssoca-ops.yml`](manifests/examples/ssoca-ops.yml)
 * [minio/minio-boshrelease](https://github.com/minio/minio-boshrelease) - [`manifests/examples/minio-ops.yml`](manifests/examples/minio-ops.yml)
 * [vito/grafana-boshrelease](https://github.com/vito/grafana-boshrelease) - [`manifests/examples/grafana-ops.yml`](manifests/examples/grafana-ops.yml)


## Technical Notes


### Caddy

 * HTTP and DNS ([azure](https://caddyserver.com/docs/tls.dns.azure), [googlecloud](https://caddyserver.com/docs/tls.dns.googlecloud), [route53](https://caddyserver.com/docs/tls.dns.route53)) verifications are currently supported.
 * It is assumed that you have already read and agree to the Subscriber Agreement for your Certificate Authorities.
 * Generated keys and certificates will be stored in `/var/vcap/store/caddy`. You should use a persistent disk.
 * This release is focused on being a simple proxy server with certificate generation features, and it intentionally does not include most optional plugins of Caddy. For more advanced behavior, you should probably continue using your existing server, although you may still want to introduce this job for the [Automatic HTTPS](https://caddyserver.com/docs/automatic-https) behavior.


### Let's Encrypt

 * [Let's Encrypt](https://letsencrypt.org/) is the default Certificate Authority.
 * You can find their Subscriber Agreement at [letsencrypt.org/repository](https://letsencrypt.org/repository/).
 * It is strongly recommended to use their [staging environment](https://letsencrypt.org/docs/staging-environment/) for testing before using in production. To change, set `acme.ca.url` to `https://acme-staging-v02.api.letsencrypt.org/directory`.


### Security Reminders

 * [Certificate Transparency](https://www.certificate-transparency.org/) logs are published by most Certificate Authorities, including [Let's Encrypt](https://letsencrypt.org/). This may contain the domain name or IP address from your certificate request.
 * DNS Verification - consider using credentials from an account which has limited access to update the `TXT` record for `_acme-challenge.*` (note `_acme-challenge` can be delegated via `NS` records).
    * AWS - a sample IAM policy for Route 53 can be found in [`aws-iam-policy.json`](src/aws-iam-policy.json)
 * The following firewall changes may be necessary for ACME challenges:
    * egress udp/53 - required for DNS verification
    * egress tcp/443 - required for ACME certificate requests
    * ingress tcp/80 - required for HTTP verification


## License

[Apache License 2.0](LICENSE)
