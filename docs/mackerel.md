### Use with Mackerel

Example:

```bash
 ./terraformer import mackerel --resources=service --api-key=YOUR_MACKEREL_API_KEY // or MACKEREL_API_KEY in env --app-key=YOUR_MACKEREL_API_KEY
 ./terraformer import mackerel --resources=service --filter=service=name1:name2:name4 --api-key=YOUR_MACKEREL_API_KEY // or MACKEREL_API_KEY in env --app-key=YOUR_MACKEREL_API_KEY
 ./terraformer import mackerel --resources=aws_integration --filter=aws_integration=id1:id2:id4 --api-key=YOUR_MACKEREL_API_KEY // or MACKEREL_API_KEY in env --app-key=YOUR_MACKEREL_API_KEY
```

List of supported Mackerel services:

*   `alert_group_setting`
    * `mackerel_alert_group_setting`
*   `aws_integration`
    * `mackerel_aws_integration`
        * Sensitive field `secret_key` is not generated and needs to be manually set
        * Sensitive field `external_id` is not generated and needs to be manually set
*   `channel`
    * `mackerel_channel`
*   `downtime`
    * `mackerel_downtime`
*   `monitor`
    * `mackerel_monitor`
*   `notification_group`
    * `mackerel_notification_group`
*   `role`
    * `mackerel_role`
*   `service`
    * `mackerel_service`