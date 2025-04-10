### Use with New Relic

Example:

```
./terraformer import newrelic -r alert,infra,synthetics --api-key=NRAK-XXXXXXXX --account-id=XXXXX
```

List of supported New Relic resources:

*   `alert`
    * `newrelic_alert_channel`
    * `newrelic_alert_condition`
    * `newrelic_alert_policy`
    * `newrelic_nrql_alert_condition`
*   `alertchannel`
    * `newrelic_alert_channel`
*   `alertcondition`
    * `newrelic_alert_condition`
    * `newrelic_nrql_alert_condition`
*   `alertpolicy`
    * `newrelic_alert_policy`
*   `infra`
    * `newrelic_infra_alert_condition`
*   `synthetics`
    * `newrelic_synthetics_monitor`
*   `tags`
    * `newrelic_entity_tags`
