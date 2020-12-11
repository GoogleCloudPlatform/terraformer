# Testing the Datadog provider 

The CLI script provided is used to test importing Datadog resources with terraformer. The tool will create resources using Terraform CLI and import them using terraformer. The imported resources will be stored in the `generated/` directory.

_Note_: The script will create and destroy real resources. Never run this on a production Datadog organization.

### Requirements 
* terraform version >= 0.12.x

### Script usage

**Terraform 0.12.x**

Run the script from the projects root directory:

```
go run ./tests/datadog/
``` 

- Test should run successfully without exiting

**Terraform 0.13.x**

Run the script from the projects root directory and pass the terraform version using DATADOG_TF_VERSION env var:

```
DATADOG_TF_VERSION=0.13.x LOG_CMD_OUTPUT=true go run ./tests/datadog/
```
- Terraformer currently generates resources using terraform version 0.12.29 and HCLv1 standards. When using terraform version 0.13.x, the script will fail due to `outputs` diffs when running `terraform plan` on the generated resources. This is due to differences in how outputs are references in state files.

- Manually ensure that generated diffs are regarding outputs only. E.g:

```
Plan: 0 to add, 0 to change, 0 to destroy.

Changes to Outputs:
    ~ datadog_dashboard_tfer--dashboard_gwh-002D-a7r-002D-cfs_id = "<resource-ID>" -> "datadog_dashboard.tfer--dashboard_gwh-002D-a7r-002D-cfs.id"
```

##### Available configuration options

| Configuration Options     | Description        |
| -------------             |:-------------      |
| **DD_TEST_CLIENT_API_KEY**    | Datadog api key      |
| **DD_TEST_CLIENT_APP_KEY**    | Datadog APP key      |
| **DATADOG_HOST**              | The API Url. if you're working with "EU" version of Datadog, use `https://api.datadoghq.eu/`. Default: `https://api.datadoghq.com/`      |
| **DATADOG_TF_VERSION**        | Terraform version installed. Pass the terraform version number if using Terraform version >= 0.13.x      |
| **DATADOG_TERRAFORM_TARGET**    | Colon separated list of resource addresses to [target](https://www.terraform.io/docs/commands/plan.html#resource-targeting). Example: `DATADOG_TERRAFORM_TARGET="datadog_dashboard.free_dashboard_example:datadog_monitor.monitor_example"`      |
| **LOG_CMD_OUTPUT**    | Print outputs to stderr and stdout when running `terraform` commands. Default `false`      |

**Frequently Asked Questions**

```
2020/10/26 15:08:34 Message: Error while importing resources. Error: fork/exec : no such file or directory
```
- Above error indicates that Terraformer is unable to locate the datadog provider executable. Manually pass the dir of the plugin's directory using env var `TF_PLUGIN_DIR`. E.g. `TF_PLUGIN_DIR=~/.terraform.d/`
