### Use with Mackerel

- https://registry.terraform.io/providers/mackerelio-labs/mackerel

### Getting started

- terraformer doesn't support `source` at provider directive, you should replace by-hand

```
# initialize for terraformer
vi init.tf
====
terraform {
  required_providers {
    mackerel = {
      source  = "mackerelio-labs/mackerel"
      version = "~> 0.0.1"
    }
  }
}
====

terraform init

# run terraformer
terraformer import mackerel --api-key=XXXXXXXXXX --resources=service,role,monitor,downtime,notification_group,channel,alert_group_setting,service_metadata,role_metadata --service-name=my-test-service --path-pattern {output}/{provider}/my-test-service/

# edit provider and replace provider source
vi provider.tf
====
+ source  = "mackerelio-labs/mackerel"
version = "~> 0.0.1"
====

rm init.tf
terraform state replace-provider registry.terraform.io/-/mackerel registry.terraform.io/mackerelio-labs/mackerel
terraform fmt
terraform plan
```

Example:

```bash
 terraformer import mackerel --resources=service --api-key=YOUR_MACKEREL_API_KEY // or MACKEREL_API_KEY in env --app-key=YOUR_MACKEREL_API_KEY
 terraformer import mackerel --resources=service --filter=service=name1:name2:name4 --api-key=YOUR_MACKEREL_API_KEY // or MACKEREL_API_KEY in env --app-key=YOUR_MACKEREL_API_KEY
 terraformer import mackerel --resources=aws_integration --filter=aws_integration=id1:id2:id4 --api-key=YOUR_MACKEREL_API_KEY // or MACKEREL_API_KEY in env --app-key=YOUR_MACKEREL_API_KEY

terraformer import mackerel --api-key=XXXXXXXXXX --resources=service,role,monitor,downtime,notification_group,channel,alert_group_setting,service_metadata,role_metadata --service-name=my-test-service --path-pattern {output}/{provider}/my-test-service/
MACKEREL_API_KEY=XXXXXXXXXX terraformer import mackerel --resources=service,role,monitor,downtime,notification_group,channel,alert_group_setting,service_metadata,role_metadata --service-name=my-test-service --path-pattern {output}/{provider}/my-test-service/
```

- all of resources are specified by mackerel service name
  - https://mackerel.io/api-docs/entry/services

List of supported Mackerel services:

- `alert_group_setting`
  - `mackerel_alert_group_setting`
- `aws_integration`
  - `mackerel_aws_integration`
    - Sensitive field `secret_key` is not generated and needs to be manually set
    - Sensitive field `external_id` is not generated and needs to be manually set
- `channel`
  - `mackerel_channel`
- `downtime`
  - `mackerel_downtime`
- `monitor`
  - `mackerel_monitor`
- `notification_group`
  - `mackerel_notification_group`
- `role`
  - `mackerel_role`
- `role_metadata`
  - `mackerel_role_metadata`
- `service`
  - `mackerel_service`
- `service_metadata`
  - `mackerel_service_metadata`

[1]: https://github.com/GoogleCloudPlatform/terraformer/blob/master/README.md#filtering
