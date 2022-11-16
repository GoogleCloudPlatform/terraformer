# Use with Squadcast

Syntax:

`export SQUADCAST_REFRESH_TOKEN=<YOUR_SQUADCAST_REFRESH_TOKEN>`

OR 

Add `--refresh-token` flag in cmd


```
terraformer import squadcast --resources=<SERVICE_NAMES> --region=SQUADCAST_REGION
```

### Examples:

- `Import Resource by providing refresh-token as a flag`

```
terraformer import squadcast --resources=team --region=us --team-name="Default Team" --refresh-token=YOUR_REFRESH_TOKEN
```

- `Import User Resource`

```
terraformer import squadcast --resources=user --region=us --team-name="Defualt Team"
```

- `Import Squad Resource`

```
terraformer import squadcast --resources=team --region=us --team-name="Default Team"
```

- `Import Deduplication Rules Resource` (without `--service-name` flag)
  - Deduplication Rules for all the services under Default Team will be generated.
```
terraformer import squadcast --resources=deduplication_rules --region=us --team-name="Default Team"
```

- `Import Deduplication Rules Resource`
  - Deduplication Rules only for Example Service will be generated.

```
terraformer import squadcast --resources=deduplication_rules --region=us --team-name="Default Team" --service-name="Example Service"
```

### In order to use terraform files:

- Update version and add source in  `provider.tf`
  - Go to `/generated/squadcast/<RESOURCE>/<REGION>/provider.tf`
  - Add `source = "SquadcastHub/squadcast"` to squadcast inside `required_providers`
  - Update `version` in `required_providers` by removing `.exe` (Windows users only)
- Update `terraform_version`
  - `cd /generated/squadcast/<RESOURCE>/<REGION>`
  - `terraform state replace-provider -auto-approve "registry.terraform.io/-/squadcast" "SquadcastHub/squadcast"`
  
### Example:
```
terraform {
    required_providers {
        squadcast = {
            version = "~> 1.0.5"
            source = "SquadcastHub/squadcast"
        }
    }
}
```

### Flags:

- `--team-name`
  - Required for the following resources:
    - deduplication_rules
    - escalation_policy
    - routing_rules
    - runbook
    - service
    - slo
    - squad
    - suppression_rules
    - tagging_rules
    - team_member
    - team_roles
    - user

- `--region`
  - Supported Values:
    - `us`
    - `eu`

- `--service-name` (optional)
  - Supported for the following resources:
    - deduplication_rules
    - routing_rules
    - suppression_rules
    - tagging_rules
  - If service name is not provided, resources for specified automation rule for all the service within the specified team will be generated. However it will only generate for a specific service when this flag is used. [see examples](squadcast.md:36)

- `--refresh-token` (optional)
  - Supported Values:
    - <YOUR_SQUADCAST_REFRESH_TOKEN>

### Supported resources:

- [`deduplication_rules`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/deduplication_rules)
- [`escalation_policy`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/escalation_policy)
- [`routing_rules`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/routing_rules)
- [`runbook`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/runbook)
- [`service`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/service)
- [`slo`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/slo)
- [`squad`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/squad)
- [`suppression_rules`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/suppression_rules)
- [`tagging_rules`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/tagging_rules)
- [`team`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/team)
- [`team_member`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/team_member)
- [`team_roles`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/team_role)
- [`user`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/user)
