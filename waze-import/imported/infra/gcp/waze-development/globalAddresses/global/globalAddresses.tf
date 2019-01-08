provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_global_address" "admanage_stg_stg1_proxy_address" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "admanage-stg-stg1-proxy-address"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "adsmanagementapi_stg_stg1_prox_forwarding_rule_2_ip" {
  description   = "IP Address for Ads Management API Staging proxy. This is used to test multi-plexing of requests to backends from WAM API java grpc proxy."
  ip_version    = "IPV4"
  labels        = {}
  name          = "adsmanagementapi-stg-stg1-prox-forwarding-rule-2-ip"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "carpoolpayments_stg_stg1_address" {
  description   = "carpoolpayments-stg-stg1 global address"
  labels        = {}
  name          = "carpoolpayments-stg-stg1-address"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "cartodb_ip" {
  labels        = {}
  name          = "cartodb-ip"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "clienttiles_stg_stg1_lb" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "clienttiles-stg-stg1-lb"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "inbox_stg_stg1_lb" {
  labels        = {}
  name          = "inbox-stg-stg1-lb"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "incidents_stg_stg1_address" {
  description   = "incidents-stg-stg1 global address"
  labels        = {}
  name          = "incidents-stg-stg1-address"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "incidentsserver_stg_stg1_address" {
  description   = "incidentsserver-stg-stg1 global address"
  labels        = {}
  name          = "incidentsserver-stg-stg1-address"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "mapnikserver_livemap_stg1_lb" {
  labels        = {}
  name          = "mapnikserver-livemap-stg1-lb"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "merger_jupyter_staging__ip" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "merger-jupyter-staging--ip"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "merger_jupyter_staging_debug_ip" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "merger-jupyter-staging-debug-ip"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "merger_jupyter_staging_http_ip" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "merger-jupyter-staging-http-ip"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "merger_jupyter_staging_ip" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "merger-jupyter-staging-ip"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "merger_jupyter_staging_other_ip" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "merger-jupyter-staging-other-ip"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "merger_jupyter_staging_spark_ip" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "merger-jupyter-staging-spark-ip"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "mobile_web_lb" {
  labels        = {}
  name          = "mobile-web-lb"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "multigrok_static_ip" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "multigrok-static-ip"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "parkingrouting_stg_stg1_address" {
  description   = "parkingrouting-stg-stg1 global address"
  labels        = {}
  name          = "parkingrouting-stg-stg1-address"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "parkingserver_stg_stg1_address" {
  description   = "parkingserver-stg-stg1 global address"
  labels        = {}
  name          = "parkingserver-stg-stg1-address"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "realtime_georss_stg1" {
  labels        = {}
  name          = "realtime-georss-stg1"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "realtime_proxy_stg_lb" {
  labels        = {}
  name          = "realtime-proxy-stg-lb"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "reverseproxy_stg_stg1_address" {
  description   = "reverseproxy-stg-stg1 global address"
  labels        = {}
  name          = "reverseproxy-stg-stg1-address"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "servermanager_external_stg" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "servermanager-external-stg"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "staging_marx" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "staging-marx"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "stg_adman2_stg" {
  labels        = {}
  name          = "stg-adman2-stg"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "stg_ads_lb" {
  labels        = {}
  name          = "stg-ads-lb"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "stg_adsassets" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "stg-adsassets"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "stg_dev_editor_lb" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "stg-dev-editor-lb"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "stg_elton_lb" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "stg-elton-lb"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "stg_mapnik_editor_lb" {
  labels        = {}
  name          = "stg-mapnik-editor-lb"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "stg_prompto" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "stg-prompto"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "stg_realtimeproxy_lb" {
  description   = "Staging Realtime Proxy load balancer"
  labels        = {}
  name          = "stg-realtimeproxy-lb"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "stg_search_lb_http_rule" {
  labels        = {}
  name          = "stg-search-lb-http-rule"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "stg_sms_lb" {
  labels        = {}
  name          = "stg-sms-lb"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "stg_ttsgw" {
  labels        = {}
  name          = "stg-ttsgw"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "stg_usersprofile_lb" {
  labels        = {}
  name          = "stg-usersprofile-lb"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "stg_voice_prompts_lb" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "stg-voice-prompts-lb"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "storagegateway_external_ip" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "storagegateway-external-ip"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "usersprofile_stg_stg1_address" {
  description   = "usersprofile-stg-stg1 global address"
  labels        = {}
  name          = "usersprofile-stg-stg1-address"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "was_stg_stg1_lb" {
  labels        = {}
  name          = "was-stg-stg1-lb"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "waze_artifactory_ip" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "waze-artifactory-ip"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "waze_reverseproxy_stg" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "waze-reverseproxy-stg"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "web_reverse_proxy_stg" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "web-reverse-proxy-stg"
  prefix_length = "0"
  project       = "waze-development"
}

resource "google_compute_global_address" "website_stg_lb" {
  ip_version    = "IPV4"
  labels        = {}
  name          = "website-stg-lb"
  prefix_length = "0"
  project       = "waze-development"
}
