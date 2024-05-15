# Use terraformer with [Opal](https://opal.dev)

##  Usage
### 1. Installation
First you will need to install terraformer with the opal provider. See the [readme](https://github.com/GoogleCloudPlatform/terraformer#installation).

### 2. Set up a template terraform workspace
Before you can use terraformer, you need to create a template workspace so that terraformer
can access the [opalsecurity/opal](https://registry.terraform.io/providers/opalsecurity/opal/latest) provider.

To do this, create a new directory with a basic `provider.tf` file:
```hcl
terraform {
  required_providers {
    opal = {
      source = "opalsecurity/opal"
      version = "0.0.2"
    }
  }
}

provider "opal" {
  # Configuration options
}
```

then run:
```bash
$ terraform init
````

You should see the output: `Terraform has been successfully initialized!`

### 3. Run terraformer:

```bash
export OPAL_AUTH_TOKEN=Your token from https://app.opal.dev/settings#api
# If you are running an on-prem installation, you will need to provide a base url as well:
# export OPAL_BASE_URL=Your token from https://my.opal.com

./terraformer import opal --resources=* --path-pattern {output}/{provider}
```

You can also specify only certain kinds of resources to import as well, i.e. `--resources=owner`.

Note that we currently do not support the terraformer `--filter` flag.

### 4. Inspect the imported terraform files

You should now see a `generated/` subdirectory with generated files. If you are using
terraform version `>= 0.13`, you will need to run a state migration:
```bash
$ cd generated/opal/
$ terraform state replace-provider -auto-approve "registry.terraform.io/-/opal" "opalsecurity/opal"
```

You can now initialize and use your new generated resources:
```bash
$ terraform init
$ terraform plan # No changes. Your infrastructure matches the configuration.
```

## Supported Opal resources:

*   `owner`
    * `opal_owner`
*   `resource`
    * `opal_resource`
*   `group`
    * `opal_group`
*   `on_call_schedules`
    * `opal_on_call_schedules`
*   `message_channels`
    * `opal_message_channels`
