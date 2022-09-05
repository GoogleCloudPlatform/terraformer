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

- `deduplication_rules`
- `escalation_policy`
- `routing_rules`
- `runbook`
- `schedule`
- `service`
- `slo`
- `squad`
- `suppression_rules`
- `tagging_rules`
- `team`
- `team_member`
- `team_roles`
- `user`
