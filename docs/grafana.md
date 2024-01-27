### Use with [Grafana](https://grafana.com)

This provider uses the [terraform-provider-grafana](https://registry.terraform.io/providers/grafana/grafana/latest).

#### Example

```
GRAFANA_AUTH=api_token GRAFANA_URL=https://stack.grafana.net ./terraformer import grafana -r=grafana_dashboard // Import with Grafana API token
GRAFANA_AUTH=username:password GRAFANA_URL=https://stack.grafana.net ./terraformer import grafana -r=grafana_dashboard // Import with HTTP basic auth
```

#### Configuration

| Env variable               | Description | Required | Default |
| -------------------------- | -------------------------------------------------------------------- | --- | - |
| GRAFANA_AUTH               | API token or HTTP basic auth (if pattern is `username:password`)     | yes | - |
| GRAFANA_URL                | URL to the Grafana instance, e.g. https://stack.grafana.net          | yes | - |
| GRAFANA_ORG_ID             | Grafana organisation ID                                              | no  | 1 |
| HTTPS_TLS_KEY              | Path to TLS key file                                                 | no  | - |
| HTTPS_TLS_CERT             | Path to TLS cert file                                                | no  | - |
| HTTPS_CA_CERT              | Path to CA cert file                                                 | no  | - |
| HTTPS_INSECURE_SKIP_VERIFY | Whether to skip TLS certificate validation (1 for true, 0 for false) | no  | 0 |

List of supported [Grafana](https://grafana.com) resources:

* `dashboard`
  * `grafana_dashboard`
* `folder`
  * `grafana_folder`
