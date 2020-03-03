resource "octopusdeploy_environment" "stage" {
  name                         = "Stage"
  description                  = "Y (SS1)"
  use_guided_failure           = false
  allow_dynamic_infrastructure = true
}

resource "octopusdeploy_environment" "production" {
  name                         = "Production"
  description                  = "B, C, D, E, F & G and more"
  use_guided_failure           = true
  allow_dynamic_infrastructure = true
}
