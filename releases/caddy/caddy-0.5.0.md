Breaking Changes

 * When using Route 53 DNS Verification, the IAM Action of `route53:ListResourceRecordSets` is now also required (caused by dependency [changes](https://github.com/go-acme/lego/commit/20d50a559f077fe8d392957aa1e362041ab64f3c)).

Dependencies

 * go 1.12 ([changelog](https://golang.org/doc/go1.12); was 1.11.5)
 * caddy 0.11.4 ([changelog](https://github.com/mholt/caddy/releases/tag/v0.11.4); was 0.11.0)
 * lego 2.2.0 ([changelog](https://github.com/xenolf/lego/releases); was 0.5.0)
 * dnsproviders 0.1.3 (was 0.1.2)
