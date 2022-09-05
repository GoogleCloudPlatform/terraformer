# Use with Squadcast

Syntax:

`export SQUADCAST_REFRESH_TOKEN=<YOUR_SQUADCAST_REFRESH_TOKEN>`

OR 

Add `--refresh-token` flag in cmd


```
terraformer import squadcast --resources=<SERVICE_NAMES> --region=SQUADCAST_REGION
```

Examples:

- `Import Resource by providing refresh-token as a flag`

```
terraformer import squadcast --resources=team --region=us --team-name="Default Team" --refresh-token=YOUR_REFRESH_TOKEN
```

- `Import User Resource`

```
terraformer import squadcast --resources=user --region=us
```

- `Import Squad Resource`

```
terraformer import squadcast --resources=team --region=us --team-name="Default Team"
```



### Flags:

- `--team-name`
  - Required for the following resources:
    - deduplication_rules
    - escalation_policy
    - routing_rules
    - runbook
    - schedules
    - service
    - slo
    - squad
    - suppression_rules
    - tagging_rules
    - team_member
    - team_roles

- `--service-name`
  - Required for the following resources:
    - deduplication_rules
    - routing_rules
    - suppression_rules
    - tagging_rules

- `--schedule-name`
  - Required for the following resources:
    - schedule

- `--region`
  - Supported Values:
    - `us`
    - `eu`

- `--refresh-token` (optional)
  - Supported Values:
    - <YOUR_SQUADCAST_REFRESH_TOKEN>

### Supported resources:

- [`deduplication_rules`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/deduplication_rules)
- [`escalation_policy`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/escalation_policy)
- [`routing_rules`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/routing_rules)
- [`runbook`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/runbook)
- [`schedule`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/schedule)
- [`service`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/service)
- [`slo`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/slo)
- [`squad`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/squad)
- [`suppression_rules`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/suppression_rules)
- [`tagging_rules`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/tagging_rules)
- [`team`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/team)
- [`team_member`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/team_member)
- [`team_roles`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/team_role)
- [`user`](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/user)
