resource "octopusdeploy_channel" "helm" {
  name        = "Helm"
  description = "The Helm channel"
  project_id  = octopusdeploy_project.deploymark_api.id
}
