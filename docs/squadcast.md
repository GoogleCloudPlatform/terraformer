# Use with Squadcast

Syntax:

`export SQUADCAST_REFRESH_TOKEN=<YOUR_SQUADCAST_REFRESH_TOKEN>`

```
terraformer import squadcast --resources=<SERVICE_NAMES> --region=SQUADCAST_REGION
```

Examples:

- `Import User Resource`

```
terraformer import squadcast --resources=user --region=us
```

- `Import Squad Resource`

```
terraformer import squadcast --resources=team --region=us --team-name="Default Team"
```

- `Import Resource by providing refresh-token as a flag`

```
terraformer import squadcast --resources=team --region=us --team-name="Default Team" --refresh-token=YOUR_REFRESH_TOKEN
```

### Flags:

- `--team-name`

  - Required for the following resources:
    - squad
    - service
    - escalation_policy
    - team_member
    - team_roles
    - slo
    - runbook

- `--region`
  - Supported Values:
    - `us`
    - `eu`

### Supported resources:

- `user`
  - [squadcast_user](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/user)
- `team`
  - [squadcast_team](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/team)
- `team_member`
  - [squadcast_team_member](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/team_member)
- `team_roles`
  - [squadcast_team_roles](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/team_roles)
- `squad`
  - [squadcast_squad](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/squad)
- `service`
  - [squadcast_service](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/service)
- `escalation_policy`
  - [squadcast_escalation_policy](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/escalation_policy)
- `runbook`
  - [squadcast_runbook](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/runbook)
- `slo`
  - [squadcast_slo](https://registry.terraform.io/providers/SquadcastHub/squadcast/latest/docs/resources/slo)
