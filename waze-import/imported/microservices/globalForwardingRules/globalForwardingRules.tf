provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_global_forwarding_rule" "admanage_stg_stg1_forwarding_rule" {
  ip_address  = "35.186.216.95"
  ip_protocol = "TCP"
  labels      = {}
  name        = "admanage-stg-stg1-forwarding-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/admanage-stg-stg1-target-proxy"
}

resource "google_compute_global_forwarding_rule" "adsmanagementapi_stg_stg1_prox_forwarding_rule" {
  ip_address  = "35.227.250.129"
  ip_protocol = "TCP"
  ip_version  = "IPV4"
  labels      = {}
  name        = "adsmanagementapi-stg-stg1-prox-forwarding-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/adsmanagementapi-stg-stg1-prox-target-proxy"
}

resource "google_compute_global_forwarding_rule" "carpool_groups_images_stg" {
  ip_address  = "35.190.92.68"
  ip_protocol = "TCP"
  ip_version  = "IPV4"
  labels      = {}
  name        = "carpool-groups-images-stg"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/waze-carpool-groups-images-stg-target-proxy"
}

resource "google_compute_global_forwarding_rule" "carpoolmatchingdispatcher_stg_stg1" {
  ip_address  = "35.244.202.182"
  ip_protocol = "TCP"
  labels      = {}
  name        = "carpoolmatchingdispatcher-stg-stg1"
  port_range  = "80-80"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpProxies/carpoolmatchingdispatcher-stg-stg1-target-http-proxy"
}

resource "google_compute_global_forwarding_rule" "carpoolpayments_stg_stg1_fr" {
  ip_address  = "130.211.45.79"
  ip_protocol = "TCP"
  labels      = {}
  name        = "carpoolpayments-stg-stg1-fr"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/carpoolpayments-stg-stg1-tps"
}

resource "google_compute_global_forwarding_rule" "carpooltesting_realtime_proxy_gfr" {
  ip_address  = "35.186.250.88"
  ip_protocol = "TCP"
  ip_version  = "IPV4"
  labels      = {}
  name        = "carpooltesting-realtime-proxy-gfr"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/carpooltesting-proxy-stg-be-target-proxy"
}

resource "google_compute_global_forwarding_rule" "clienttiles_stg_stg1" {
  ip_address  = "35.201.122.10"
  ip_protocol = "TCP"
  labels      = {}
  name        = "clienttiles-stg-stg1"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/clienttiles-stg-stg1-target-https-proxy"
}

resource "google_compute_global_forwarding_rule" "inbox_stg_stg1_forwarding_rule" {
  ip_address  = "130.211.15.36"
  ip_protocol = "TCP"
  labels      = {}
  name        = "inbox-stg-stg1-forwarding-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/inbox-stg-stg1-target-proxy"
}

resource "google_compute_global_forwarding_rule" "incidents_stg_stg1_efr" {
  ip_address  = "130.211.10.69"
  ip_protocol = "TCP"
  labels      = {}
  name        = "incidents-stg-stg1-efr"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/incidents-stg-stg1-tps"
}

resource "google_compute_global_forwarding_rule" "incidentsserver_stg_stg1_efr" {
  ip_address  = "35.190.43.211"
  ip_protocol = "TCP"
  labels      = {}
  name        = "incidentsserver-stg-stg1-efr"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/incidentsserver-stg-stg1-tps"
}

resource "google_compute_global_forwarding_rule" "mapnikserver_livemap_stg1_forwarding_rule" {
  ip_address  = "35.190.51.239"
  ip_protocol = "TCP"
  ip_version  = "IPV4"
  labels      = {}
  name        = "mapnikserver-livemap-stg1-forwarding-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/mapnikserver-livemap-stg1-target-proxy"
}

resource "google_compute_global_forwarding_rule" "mobile_web" {
  ip_address  = "107.178.247.127"
  ip_protocol = "TCP"
  labels      = {}
  name        = "mobile-web"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/stg-mobile-web-target-1"
}

resource "google_compute_global_forwarding_rule" "mobile_web_http" {
  ip_address  = "107.178.247.127"
  ip_protocol = "TCP"
  labels      = {}
  name        = "mobile-web-http"
  port_range  = "80-80"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpProxies/stg-mobile-web-target-2"
}

resource "google_compute_global_forwarding_rule" "parkingrouting_stg_stg1_efr" {
  ip_address  = "35.201.91.29"
  ip_protocol = "TCP"
  labels      = {}
  name        = "parkingrouting-stg-stg1-efr"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/parkingrouting-stg-stg1-tps"
}

resource "google_compute_global_forwarding_rule" "parkingserver_stg_stg1_efr" {
  ip_address  = "35.190.1.223"
  ip_protocol = "TCP"
  labels      = {}
  name        = "parkingserver-stg-stg1-efr"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/parkingserver-stg-stg1-tps"
}

resource "google_compute_global_forwarding_rule" "prompto_stg_stg1_ifr_ext" {
  ip_address  = "35.190.46.127"
  ip_protocol = "TCP"
  labels      = {}
  name        = "prompto-stg-stg1-ifr-ext"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/prompto-stg-stg1-ifr-ext-target-https-proxy"
}

resource "google_compute_global_forwarding_rule" "realtime_georss_stg1_forwarding_rule" {
  ip_address  = "35.186.211.26"
  ip_protocol = "TCP"
  labels      = {}
  name        = "realtime-georss-stg1-forwarding-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/realtime-georss-stg1-target-proxy"
}

resource "google_compute_global_forwarding_rule" "realtime_proxy_stg_gfr" {
  ip_address  = "130.211.10.205"
  ip_protocol = "TCP"
  labels      = {}
  name        = "realtime-proxy-stg-gfr"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetSslProxies/realtime-proxy-tp-ssl"
}

resource "google_compute_global_forwarding_rule" "repository_server_stg_stg1_external" {
  ip_address  = "35.186.253.253"
  ip_protocol = "TCP"
  labels      = {}
  name        = "repository-server-stg-stg1-external"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/repository-server-stg-stg1-external-target-https-proxy"
}

resource "google_compute_global_forwarding_rule" "searchserver_stg_stg1_external_fr" {
  ip_address  = "35.201.127.199"
  ip_protocol = "TCP"
  labels      = {}
  name        = "searchserver-stg-stg1-external-fr"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/searchserver-stg-stg1-external-fr-target-https-proxy"
}

resource "google_compute_global_forwarding_rule" "servermanager_external_stg" {
  ip_address  = "35.190.47.68"
  ip_protocol = "TCP"
  labels      = {}
  name        = "servermanager-external-stg"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/servermanager-external-stg-target-proxy"
}

resource "google_compute_global_forwarding_rule" "socialmediaserver_stg_ext_fe" {
  ip_address  = "35.190.77.126"
  ip_protocol = "TCP"
  labels      = {}
  name        = "socialmediaserver-stg-ext-fe"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/socialmediaserver-stg-ext-fe-target-https-proxy"
}

resource "google_compute_global_forwarding_rule" "staging_marx" {
  ip_address  = "35.244.168.238"
  ip_protocol = "TCP"
  labels      = {}
  name        = "staging-marx"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/staging-marx-target-proxy"
}

resource "google_compute_global_forwarding_rule" "stg_adman2_stg_gfr" {
  ip_address  = "130.211.10.102"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-adman2-stg-gfr"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/stg-adman2-stg-target-1"
}

resource "google_compute_global_forwarding_rule" "stg_ads_gfr" {
  ip_address  = "107.178.246.185"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-ads-gfr"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/stg-ads-target"
}

resource "google_compute_global_forwarding_rule" "stg_adsassets" {
  ip_address  = "35.190.54.247"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-adsassets"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/stg-adsassets-target-proxy"
}

resource "google_compute_global_forwarding_rule" "stg_carpoolrouting_grf" {
  ip_address  = "35.201.115.46"
  ip_protocol = "TCP"
  ip_version  = "IPV4"
  labels      = {}
  name        = "stg-carpoolrouting-grf"
  port_range  = "80-80"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpProxies/stg-carpoolrouting-lb-target-proxy"
}

resource "google_compute_global_forwarding_rule" "stg_dev_editor_forwarding_rule" {
  ip_address  = "107.178.246.163"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-dev-editor-forwarding-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/stg-dev-editor-target-proxy"
}

resource "google_compute_global_forwarding_rule" "stg_elton" {
  ip_address  = "35.190.61.146"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-elton"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/stg-elton-target-proxy"
}

resource "google_compute_global_forwarding_rule" "stg_mapnik_editor" {
  ip_address  = "107.178.248.132"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-mapnik-editor"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/stg-mapnik-editor-target-1"
}

resource "google_compute_global_forwarding_rule" "stg_routing_lb_forwarding_rule" {
  ip_address  = "130.211.28.144"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-routing-lb-forwarding-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/stg-routing-lb-target-proxy"
}

resource "google_compute_global_forwarding_rule" "stg_rtproxy_https_rule" {
  ip_address  = "107.178.240.186"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-rtproxy-https-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/stg-rtproxy-default-target-https"
}

resource "google_compute_global_forwarding_rule" "stg_search_lb_http_rule" {
  ip_address  = "107.178.248.9"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-search-lb-http-rule"
  port_range  = "80-80"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpProxies/stg-search-lb"
}

resource "google_compute_global_forwarding_rule" "stg_search_lb_https_rule" {
  ip_address  = "107.178.248.9"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-search-lb-https-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/stg-search-lb-https-v2"
}

resource "google_compute_global_forwarding_rule" "stg_sms_lb_http_rule" {
  ip_address  = "107.178.252.170"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-sms-lb-http-rule"
  port_range  = "80-80"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpProxies/stg-sms-lb"
}

resource "google_compute_global_forwarding_rule" "stg_sms_lb_https_rule" {
  ip_address  = "107.178.252.170"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-sms-lb-https-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/stg-sms-lb-https-v2"
}

resource "google_compute_global_forwarding_rule" "stg_ttsgw" {
  ip_address  = "107.178.254.116"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-ttsgw"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/stg-ttsgw-target-1"
}

resource "google_compute_global_forwarding_rule" "stg_usersprofile_lb_http_rule" {
  ip_address  = "107.178.248.219"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-usersprofile-lb-http-rule"
  port_range  = "80-80"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpProxies/stg-usersprofile-lb"
}

resource "google_compute_global_forwarding_rule" "stg_usersprofile_lb_https_rule" {
  ip_address  = "107.178.248.219"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-usersprofile-lb-https-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/stg-usersprofile-lb-https-v1"
}

resource "google_compute_global_forwarding_rule" "stg_voice_prompts_forwarding_rule" {
  ip_address  = "35.186.231.245"
  ip_protocol = "TCP"
  labels      = {}
  name        = "stg-voice-prompts-forwarding-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/stg-voice-prompts-target-proxy"
}

resource "google_compute_global_forwarding_rule" "storagegateway_stg_https" {
  ip_address  = "35.201.81.162"
  ip_protocol = "TCP"
  labels      = {}
  name        = "storagegateway-stg-https"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/storagegateway-stg-https-target-https-proxy"
}

resource "google_compute_global_forwarding_rule" "usersprofile_stg_stg1_efr" {
  ip_address  = "35.190.66.51"
  ip_protocol = "TCP"
  labels      = {}
  name        = "usersprofile-stg-stg1-efr"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/usersprofile-stg-stg1-tps"
}

resource "google_compute_global_forwarding_rule" "was_stg_stg1_forwarding_rule" {
  ip_address  = "35.190.78.117"
  ip_protocol = "TCP"
  ip_version  = "IPV4"
  labels      = {}
  name        = "was-stg-stg1-forwarding-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/was-stg-stg1-target-proxy"
}

resource "google_compute_global_forwarding_rule" "waze_reverseproxy_stg_forwarding_rule" {
  ip_address  = "35.241.33.157"
  ip_protocol = "TCP"
  labels      = {}
  name        = "waze-reverseproxy-stg-forwarding-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/waze-reverseproxy-stg-target-proxy"
}

resource "google_compute_global_forwarding_rule" "web_reverse_proxy_stg" {
  ip_address  = "35.244.204.44"
  ip_protocol = "TCP"
  labels      = {}
  name        = "web-reverse-proxy-stg"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/reverseproxy-stg-stg1-map-target-proxy"
}

resource "google_compute_global_forwarding_rule" "website_stg_forwarding_rule" {
  ip_address  = "107.178.248.105"
  ip_protocol = "TCP"
  labels      = {}
  name        = "website-stg-forwarding-rule"
  port_range  = "443-443"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpsProxies/website-stg-target-proxy"
}

resource "google_compute_global_forwarding_rule" "website_stg_forwarding_rule_2" {
  ip_address  = "107.178.248.105"
  ip_protocol = "TCP"
  labels      = {}
  name        = "website-stg-forwarding-rule-2"
  port_range  = "80-80"
  project     = "waze-development"
  target      = "https://www.googleapis.com/compute/beta/projects/waze-development/global/targetHttpProxies/website-stg-target-proxy-2"
}
