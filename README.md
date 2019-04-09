# caddy-bosh-release

A [BOSH release](https://bosh.io/) to run [Caddy](https://caddyserver.com/), a basic web and proxy server with support for dynamically provisioning certificates. You may find this useful if you are looking for an easy method to front another BOSH job with a certificate from [Let's Encrypt](https://letsencrypt.org/).


## Examples

Some examples of putting Caddy in front of other BOSH releases and their relevant configuration properties.

 * [cloudfoundry/bosh](https://github.com/cloudfoundry/bosh) - [`manifests/examples/bosh-ops.yml`](manifests/examples/bosh-ops.yml)
 * [cloudfoundry/uaa-release](https://github.com/cloudfoundry/uaa-release) - [`manifests/examples/uaa-ops.yml`](manifests/examples/uaa-ops.yml)
 * [concourse/concourse](https://github.com/concourse/concourse) - [`manifests/examples/concourse-ops.yml`](manifests/examples/concourse-ops.yml)
 * [dpb587/ssoca-bosh-release](https://github.com/dpb587/ssoca-bosh-release) - [`manifests/examples/ssoca-ops.yml`](manifests/examples/ssoca-ops.yml)
 * [minio/minio-boshrelease](https://github.com/minio/minio-boshrelease) - [`manifests/examples/minio-ops.yml`](manifests/examples/minio-ops.yml)
 * [vito/grafana-boshrelease](https://github.com/vito/grafana-boshrelease) - [`manifests/examples/grafana-ops.yml`](manifests/examples/grafana-ops.yml)


## License

[Apache License 2.0](LICENSE)
