resource "octopusdeploy_tag_set" "DevOps" {
  name = "DevOps"
  tag {
    name  = "Monitoring"
    color = "#3692F2"
  }
  tag {
    name  = "CI"
    color = "#FF5733"
  }
}
