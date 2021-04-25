resource "octopusdeploy_feed" "feed" {
  name          = "feedme"
  feed_type     = "Helm"
  feed_uri      = "https://kubernetes-charts.storage.googleapis.com"
  username      = "foo"
  password      = "bar"
  enhanced_mode = false
}
