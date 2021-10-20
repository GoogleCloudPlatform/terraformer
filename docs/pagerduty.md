### Use with PagerDuty

Example:

```
./terraformer import pagerduty -r team,schedule,user  -t YOUR_PAGERDUTY_TOKEN // or PAGERDUTY_TOKEN in env
```
Instructions to obtain a Auth Token: https://developer.pagerduty.com/docs/rest-api-v2/authentication/

List of supported PagerDuty resources:

* `addon`
  * `pagerduty_addon`
* `business_service`
  * `pagerduty_business_service`
* `escalation_policy`
  * `pagerduty_escalation_policy`
* `extension`
  * `pagerduty_extension`
* `maintenance_window`
  * `pagerduty_maintenance_window`
* `response_play`
  * `pagerduty_response_play`
* `ruleset`
  * `pagerduty_ruleset`
* `ruleset_rule`
  * `pagerduty_ruleset_rule`
* `schedule`
  * `pagerduty_schedule`
* `service`
  * `pagerduty_service`
  * `pagerduty_service_event_rule`
* `service_integration`
  * `pagerduty_service_integration`
* `team`
  * `pagerduty_team`
  * `pagerduty_team_membership`
* `user`
  * `pagerduty_user`

### ToDos

 - [ ] Fix https://github.com/GoogleCloudPlatform/terraformer/issues/1051 and implement `pagerduty_service_dependency` generator
