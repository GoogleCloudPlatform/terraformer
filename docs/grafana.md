### Use with [Grafana](https://grafana.com)

This provider uses the [terraform-provider-commercetools](https://registry.terraform.io/providers/grafana/grafana/latest).

Example:

```
GRAFANA_AUTH=api_token GRAFANA_URL=https://mygrafana.grafana.net ./terraformer import grafana -r=dashboards // Import with Grafana API token
GRAFANA_AUTH=username:password GRAFANA_URL=https://mygrafana.grafana.net ./terraformer import grafana -r=dashboards // Import with HTTP basic auth
```

List of supported [Grafana](https://grafana.com) resources:

* `dashboard`
  * `grafana_dashboard`
* `folder`
  * `grafana_folder`
