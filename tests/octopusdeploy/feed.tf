resource "octopusdeploy_feed" "feed" {
  name          = "feedme"
  feed_type     = "Helm"
  feed_uri      = "https://charts.helm.sh/stable"
  username      = "foo"
  password      = "bar"
  enhanced_mode = false
}
