# Testing the OctopusDeploy provider

You will need an OctopusDeploy server:

```bash
git clone git@github.com:MattHodge/VagrantBoxes.git
cd VagrantBoxes/OctopusDeployServer
vagrant up
cd -
```

Download the provider (pick the right one for your platform):

```bash
PLATFORM=$(uname -s | tr '[:upper:]' '[:lower:]')
PLUGIN_DIR=.terraform/plugins/${PLATFORM}_amd64/
PROVIDER=terraform-provider-octopusdeploy_${PLATFORM}_amd64_v0.5.0
mkdir -p "${PLUGIN_DIR}"
cd "${PLUGIN_DIR}"
curl -sLO "https://github.com/OctopusDeploy/terraform-provider-octopusdeploy/releases/download/v0.5.0/${PROVIDER}.zip"
unzip "${PROVIDER}.zip"
mv "${PROVIDER}" terraform-provider-octopusdeploy_v0.5.0
cd -
```

Run `terraform` to create the resources (adjust the values in `provider.tf`):

```bash
terraform init
terraform plan
terraform apply --auto-approve
```

Import them back with `terraformer`:

```bash
terraformer import octopusdeploy \
  --server "http://localhost:8081" \
  --apikey "API-YVLL2ML1XRIBUU8GKJKEMXKPWQ" \
  -r accounts,environments,feeds,libraryvariablesets,lifecycles,projects,projectgroups,projecttriggers,tagsets
```

Compare the output from `generated/octopusdeploy` with the original files.
