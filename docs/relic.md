### Use with New Relic

Example:

```
./terraformer import newrelic -r alert,dashboard,infra,synthetics --api-key=NRAK-XXXXXXXX --account-id=XXXXX
```

List of supported New Relic resources:

*   `alert`
    * `newrelic_alert_channel`
    * `newrelic_alert_condition`
    * `newrelic_nrql_alert_condition`
    * `newrelic_alert_policy`
*   `dashboard`
    * `newrelic_dashboard`
*   `infra`
    * `newrelic_infra_alert_condition`
*   `synthetics`
    * `newrelic_synthetics_monitor`
    * `newrelic_synthetics_alert_condition`
