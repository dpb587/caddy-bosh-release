---
title: Microsoft Azure
---

# Microsoft Azure


## Using DNS Verification

If you are using [Azure DNS](https://azure.microsoft.com/en-us/services/dns/), configure your service account key through your `AZURE_*` environment variable:

```yaml
caddyfile: |
  ...
  tls {
    dns azure
  }
env:
  AZURE_CLIENT_ID: "d37241d0-..."
  AZURE_CLIENT_SECRET: "e0ad9eec-..."
  AZURE_SUBSCRIPTION_ID: "20c70a7b-..."
  AZURE_TENANT_ID: "a6154ed3-..."
```
