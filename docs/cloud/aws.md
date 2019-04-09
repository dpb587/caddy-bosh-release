---
title: Amazon Web Services
---

# Amazon Web Services


## Using DNS Verification

If you are using [Route 53](https://aws.amazon.com/route53/), configure your service account key through your `AWS_*` environment variable:

```yaml
caddyfile: |
  ...
  tls {
    dns route53
  }
env:
  AWS_ACCESS_KEY_ID: AKIA...
  AWS_SECRET_ACCESS_KEY: ...
  AWS_HOSTED_ZONE_ID: Z... # optional
```


## Limited IAM Policy

Consider using credentials from an account which has limited access to update the `TXT` record for `_acme-challenge.*` (note `_acme-challenge` can be delegated via `NS` records).

The following policy restricts updates for a single `{{ZONE_ID}}`.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "route53:GetChange"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Effect": "Allow",
      "Action": [
        "route53:ChangeResourceRecordSets",
        "route53:ListResourceRecordSets"
      ],
      "Resource": [
        "arn:aws:route53:::hostedzone/{{ZONE_ID}}"
      ]
    }
  ]
}
```
