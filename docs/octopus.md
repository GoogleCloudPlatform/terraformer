# Use with OctopusDeploy

Example:

```bash
export OCTOPUS_CLI_SERVER=http://localhost:8081/
export OCTOPUS_CLI_API_KEY=API-CK7DQ8BMJCUUBSHAJCDIATXUO

terraformer import octopusdeploy --resources=tagsets
```

```powershell
$Env:OCTOPUS_CLI_SERVER = 'http://localhost:8081/'
$Env:OCTOPUS_CLI_API_KEY = 'API-CK7DQ8BMJCUUBSHAJCDIATXUO'

terraformer import octopusdeploy --resources=tagsets
```

## Suported Resources

* `accounts`
  * `octopusdeploy_account`
* `certificates`
  * `octopusdeploy_certificate`
* `environments`
  * `octopusdeploy_environment`
* `feeds`
  * `octopusdeploy_feed`
* `libraryvariablesets`
  * `octopusdeploy_library_variable_set`
* `lifecycles`
  * `octopusdeploy_lifecycle`
* `projects`
  * `octopusdeploy_project`
* `projectgroups`
  * `octopusdeploy_project_group`
* `projecttriggers`
  * `octopusdeploy_project_deployment_target_trigger`
* `tagsets`
  * `octopusdeploy_tag_set`
