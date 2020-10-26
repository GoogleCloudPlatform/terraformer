# Testing the Datadog provider 

The CLI script in this directory is provided to test importing Datadog resources. The tool will create resources using Terraform CLI and import them using terraformer. The imported resources are 

_Note_: The script will create and destroy real resources. Never run this on a production Datadog organization.

### Requirements 
* terraform version >= 0.12.x

### Script usage

#####Terraform 0.12.x

   Run the script from the projects root directory:
    
    ```
    go run ./tests/datadog/
    ``` 

#####Terraform 0.13.x

   Run the script from the projects root directory and pass the terraform version using DATADOG_TF_VERSION env var:

   ```
   DATADOG_TF_VERSION=0.13.x LOG_CMD_OUTPUT=true go run ./tests/datadog/
   ```
   _Note_ Terraformer currently generates resources using terraform version 0.12.29 and HCLv1 standards. When using terraform version 0.13.x, the script will fail due to `output` diffs when running `terraform plan` on the generated resources. This is due to differences in how outputs are references in state files.
   
   - Manually ensure that generated diffs are regarding outputs only. E.g:
   
   ```
   Changes to Outputs:
     ~ datadog_dashboard_tfer--dashboard_gwh-002D-a7r-002D-cfs_id = "<resource-ID>" -> "datadog_dashboard.tfer--dashboard_gwh-002D-a7r-002D-cfs.id"
   ```

##### Available configuration options

| Configuration Options     | Description        |
| -------------             |:-------------      |
| DD_TEST_CLIENT_API_KEY    | Datadog api key      |
| DD_TEST_CLIENT_APP_KEY    | Datadog APP key      |
| DATADOG_HOST              | The API Url. if you're working with "EU" version of Datadog, use `https://api.datadoghq.eu/`. Default: `https://api.datadoghq.com/`      |
| DATADOG_TF_VERSION        | Terraform version installed. Pass the terraform version number if using Terraform version >= 0.13.x      |
| DATADOG_TERRAFORM_TARGET    | Colon separated list of resource addresses to [target](https://www.terraform.io/docs/commands/plan.html#resource-targeting). Example: `DATADOG_TERRAFORM_TARGET="datadog_dashboard.free_dashboard_example:datadog_monitor.monitor_example"`      |
| LOG_CMD_OUTPUT    | Print outputs of stderr and stdout from `terraform` commands. Default `false`      |

