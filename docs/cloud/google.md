---
title: Google Cloud
---

# Google Cloud


## Using DNS Verification

If you are using [Cloud DNS](https://cloud.google.com/dns/), configure your service account key through the `GCE_SERVICE_ACCOUNT` environment variable:

```yaml
caddyfile: |
  ...
  tls {
    dns googlecloud
  }
env:
  GCE_SERVICE_ACCOUNT: |
    {
      "type": "service_account",
      "project_id": "sample-project-001",
      "private_key_id": "sample4c02016ac2d8abf5b1577993ded31626a8",
      "private_key": "-----BEGIN PRIVATE KEY-----\nMIIE...k8LA==\n-----END PRIVATE KEY-----\n",
      ...
    }
```
