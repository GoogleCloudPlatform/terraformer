provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_target_https_proxy" "admanage_stg_stg1_target_proxy" {
  name             = "admanage-stg-stg1-target-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/admanage-stg-stg1"
}

resource "google_compute_target_https_proxy" "adsmanagementapi_stg_stg1_prox_target_proxy" {
  name             = "adsmanagementapi-stg-stg1-prox-target-proxy"
  project          = "waze-development"
  quic_override    = "NONE"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/adsmanagementapi-stg-stg1-proxy"
}

resource "google_compute_target_https_proxy" "carpoolpayments_stg_stg1_tps" {
  name             = "carpoolpayments-stg-stg1-tps"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/carpoolpayments-stg-stg1-map"
}

resource "google_compute_target_https_proxy" "carpooltesting_proxy_stg_be_target_proxy" {
  name             = "carpooltesting-proxy-stg-be-target-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/carpooltesting-proxy-stg-be"
}

resource "google_compute_target_https_proxy" "clienttiles_stg_stg1_target_https_proxy" {
  name             = "clienttiles-stg-stg1-target-https-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/clienttiles-stg-stg1"
}

resource "google_compute_target_https_proxy" "inbox_stg_stg1_target_proxy" {
  name             = "inbox-stg-stg1-target-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/inbox-stg-stg1"
}

resource "google_compute_target_https_proxy" "incidents_stg_stg1_tps" {
  name             = "incidents-stg-stg1-tps"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/incidents-stg-stg1-map"
}

resource "google_compute_target_https_proxy" "incidentsserver_stg_stg1_tps" {
  name             = "incidentsserver-stg-stg1-tps"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/incidentsserver-stg-stg1-map"
}

resource "google_compute_target_https_proxy" "mapnikserver_livemap_stg1_target_proxy" {
  name             = "mapnikserver-livemap-stg1-target-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/mapnikserver-livemap-stg1"
}

resource "google_compute_target_https_proxy" "parkingrouting_stg_stg1_tps" {
  name             = "parkingrouting-stg-stg1-tps"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/parkingrouting-stg-stg1-map"
}

resource "google_compute_target_https_proxy" "parkingserver_stg_stg1_tps" {
  name             = "parkingserver-stg-stg1-tps"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/parkingserver-stg-stg1-map"
}

resource "google_compute_target_https_proxy" "prompto_stg_stg1_ifr_ext_target_https_proxy" {
  name             = "prompto-stg-stg1-ifr-ext-target-https-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/prompto-stg-stg1-ext"
}

resource "google_compute_target_https_proxy" "realtime_georss_stg1_target_proxy" {
  name             = "realtime-georss-stg1-target-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/realtime-georss-stg1"
}

resource "google_compute_target_https_proxy" "repository_server_stg_stg1_external_target_https_proxy" {
  name             = "repository-server-stg-stg1-external-target-https-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/repository-stg-external"
}

resource "google_compute_target_https_proxy" "reverseproxy_stg_stg1_map_target_proxy" {
  name             = "reverseproxy-stg-stg1-map-target-proxy"
  project          = "waze-development"
  quic_override    = "NONE"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/reverseproxy-stg-stg1-map"
}

resource "google_compute_target_https_proxy" "searchserver_stg_stg1_external_fr_target_https_proxy" {
  name             = "searchserver-stg-stg1-external-fr-target-https-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/searchserver-stg-stg1-external-map"
}

resource "google_compute_target_https_proxy" "servermanager_external_stg_target_proxy" {
  name             = "servermanager-external-stg-target-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/servermanager-external-stg"
}

resource "google_compute_target_https_proxy" "socialmediaserver_stg_ext_fe_target_https_proxy" {
  name             = "socialmediaserver-stg-ext-fe-target-https-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/socialmediaserver-stg-external"
}

resource "google_compute_target_https_proxy" "staging_marx_target_proxy" {
  name             = "staging-marx-target-proxy"
  project          = "waze-development"
  quic_override    = "NONE"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/marx-stg"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/staging-marx"
}

resource "google_compute_target_https_proxy" "stg_adman2_stg_target_1" {
  name             = "stg-adman2-stg-target-1"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-adman2-stg"
}

resource "google_compute_target_https_proxy" "stg_ads_target" {
  name             = "stg-ads-target"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-ads"
}

resource "google_compute_target_https_proxy" "stg_adsassets_target_proxy" {
  name             = "stg-adsassets-target-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-adsassets"
}

resource "google_compute_target_https_proxy" "stg_dev_editor_target_proxy" {
  name             = "stg-dev-editor-target-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-dev-editor"
}

resource "google_compute_target_https_proxy" "stg_elton_target_proxy" {
  name             = "stg-elton-target-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-elton"
}

resource "google_compute_target_https_proxy" "stg_mapnik_editor_target_1" {
  name             = "stg-mapnik-editor-target-1"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-mapnik-editor"
}

resource "google_compute_target_https_proxy" "stg_mobile_web_target_1" {
  name             = "stg-mobile-web-target-1"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-mobile-web"
}

resource "google_compute_target_https_proxy" "stg_routing_lb_target_proxy" {
  name             = "stg-routing-lb-target-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-routing-lb"
}

resource "google_compute_target_https_proxy" "stg_rtproxy_default_target_https" {
  name             = "stg-rtproxy-default-target-https"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-rtproxy"
}

resource "google_compute_target_https_proxy" "stg_search_lb_https_v2" {
  name             = "stg-search-lb-https-v2"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-search-lb-urlmap"
}

resource "google_compute_target_https_proxy" "stg_sms_lb_https_v2" {
  name             = "stg-sms-lb-https-v2"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-sms-lb-urlmap"
}

resource "google_compute_target_https_proxy" "stg_ttsgw_target_1" {
  name             = "stg-ttsgw-target-1"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-ttsgw"
}

resource "google_compute_target_https_proxy" "stg_usersprofile_lb_https_v1" {
  name             = "stg-usersprofile-lb-https-v1"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-usersprofile-lb"
}

resource "google_compute_target_https_proxy" "stg_voice_prompts_target_proxy" {
  name             = "stg-voice-prompts-target-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-voice-prompts"
}

resource "google_compute_target_https_proxy" "storagegateway_stg_https_target_https_proxy" {
  name             = "storagegateway-stg-https-target-https-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/storagegateway-stg-ext"
}

resource "google_compute_target_https_proxy" "usersprofile_stg_stg1_tps" {
  name             = "usersprofile-stg-stg1-tps"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/usersprofile-stg-stg1-map"
}

resource "google_compute_target_https_proxy" "was_stg_stg1_target_proxy" {
  name             = "was-stg-stg1-target-proxy"
  project          = "waze-development"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/was-stg-stg1"
}

resource "google_compute_target_https_proxy" "waze_carpool_groups_images_stg_target_proxy" {
  name             = "waze-carpool-groups-images-stg-target-proxy"
  project          = "waze-development"
  quic_override    = "NONE"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/waze-carpool-groups-images-stg"
}

resource "google_compute_target_https_proxy" "waze_reverseproxy_stg_target_proxy" {
  name             = "waze-reverseproxy-stg-target-proxy"
  project          = "waze-development"
  quic_override    = "NONE"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/waze-reverseproxy-stg"
}

resource "google_compute_target_https_proxy" "website_stg_target_proxy" {
  name             = "website-stg-target-proxy"
  project          = "waze-development"
  quic_override    = "DISABLE"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
  url_map          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/website-stg"
}
