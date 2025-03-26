### Use with Logz.io

Example:

```
 LOGZIO_API_TOKEN=foobar LOGZIO_BASE_URL=https://api-eu.logz.io ./terraformer import logzio -r=alerts,alert_notification_endpoints // Import Logz.io alerts and alert notification endpoints
```

List of supported Logz.io resources:

*   `alert_notification_endpoints`
    * `logzio_endpoint`
*   `alerts`
    * `logzio_alert`
