### Use with PagerDuty

Example:

```
./terraformer import pagerduty -r team,schedule,user  -t YOUR_PAGERDUTY_TOKEN // or PAGERDUTY_TOKEN in env
```
Instructions to obtain a Auth Token: https://developer.pagerduty.com/docs/rest-api-v2/authentication/

List of supported PagerDuty resources:

* `business_service`
  * `pagerduty_business_service`
* `escalation_policy`
  * `pagerduty_escalation_policy`
* `ruleset`
  * `pagerduty_ruleset`
  * `pagerduty_ruleset_rule`
* `schedule`
  * `pagerduty_schedule`
* `service`
  * `pagerduty_service`
  * `pagerduty_service_event_rule`
* `team`
  * `pagerduty_team`
  * `pagerduty_team_membership`
* `user`
  * `pagerduty_user`
