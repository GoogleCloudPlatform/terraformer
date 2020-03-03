resource "octopusdeploy_lifecycle" "stage_prod" {
  name        = "Stage-Production"
  description = "Stage-Production Lifecycle"

  release_retention_policy {
    unit             = "Days"
    quantity_to_keep = 3
  }

  tentacle_retention_policy {
    unit             = "Items"
    quantity_to_keep = 3
  }

  phase {
    name                                  = "Stage"
    minimum_environments_before_promotion = 0
    is_optional_phase                     = false
    optional_deployment_targets           = ["${octopusdeploy_environment.stage.id}"]
  }

  phase {
    name                                  = "Production"
    minimum_environments_before_promotion = 0
    is_optional_phase                     = false
    optional_deployment_targets           = ["${octopusdeploy_environment.production.id}"]
  }
}
