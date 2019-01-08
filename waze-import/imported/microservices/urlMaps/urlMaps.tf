provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_url_map" "admanage_stg_stg1" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/admanage-stg-stg1-backend-service"
  name            = "admanage-stg-stg1"
  project         = "waze-development"
}

resource "google_compute_url_map" "adsmanagementapi_stg_stg1_proxy" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/adsmanagementapi-stg-stg1-proxy-be"
  name            = "adsmanagementapi-stg-stg1-proxy"
  project         = "waze-development"
}

resource "google_compute_url_map" "carpoolmatchingdispatcher_stg" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/carpoolmatchingdispatcher-stg-stg1"
  name            = "carpoolmatchingdispatcher-stg"
  project         = "waze-development"
}

resource "google_compute_url_map" "carpoolpayments_stg_stg1_map" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/carpoolpayments-stg-stg1-bes"
  name            = "carpoolpayments-stg-stg1-map"
  project         = "waze-development"
}

resource "google_compute_url_map" "carpooltesting_proxy_stg_be" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/carpooltesting-realtime-proxy-stg-be"
  name            = "carpooltesting-proxy-stg-be"
  project         = "waze-development"
}

resource "google_compute_url_map" "clienttiles_stg_stg1" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/clienttiles-stg-stg1-backend-service"
  name            = "clienttiles-stg-stg1"
  project         = "waze-development"
}

resource "google_compute_url_map" "inbox_stg_stg1" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/inbox-stg-stg1"
  name            = "inbox-stg-stg1"
  project         = "waze-development"
}

resource "google_compute_url_map" "incidents_stg_stg1_map" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/incidents-stg-stg1-ebes"
  name            = "incidents-stg-stg1-map"
  project         = "waze-development"
}

resource "google_compute_url_map" "incidentsserver_stg_stg1_map" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/incidentsserver-stg-stg1-ebes"
  name            = "incidentsserver-stg-stg1-map"
  project         = "waze-development"
}

resource "google_compute_url_map" "mapnikserver_livemap_stg1" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/mapnikserver-livemap-stg1"
  name            = "mapnikserver-livemap-stg1"
  project         = "waze-development"
}

resource "google_compute_url_map" "parkingrouting_stg_stg1_map" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/parkingrouting-stg-stg1-ebes"
  name            = "parkingrouting-stg-stg1-map"
  project         = "waze-development"
}

resource "google_compute_url_map" "parkingserver_stg_stg1_map" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/parkingserver-stg-stg1-ebes"
  name            = "parkingserver-stg-stg1-map"
  project         = "waze-development"
}

resource "google_compute_url_map" "prompto_stg_stg1_ext" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/stg-prompto-backend"
  name            = "prompto-stg-stg1-ext"
  project         = "waze-development"
}

resource "google_compute_url_map" "realtime_georss_stg1" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/realtime-georss-stg1-backend-service"
  name            = "realtime-georss-stg1"
  project         = "waze-development"
}

resource "google_compute_url_map" "repository_stg_external" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/repositoy-server-stg-stg1-bes"
  name            = "repository-stg-external"
  project         = "waze-development"
}

resource "google_compute_url_map" "reverseproxy_stg_stg1_map" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/reverseproxy-stg-stg1-ebes"
  name            = "reverseproxy-stg-stg1-map"
  project         = "waze-development"
}

resource "google_compute_url_map" "searchserver_stg_stg1_external_map" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/searchserver-stg-stg1-external-bes"
  name            = "searchserver-stg-stg1-external-map"
  project         = "waze-development"
}

resource "google_compute_url_map" "servermanager_external_stg" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/servermanager-external"
  name            = "servermanager-external-stg"
  project         = "waze-development"
}

resource "google_compute_url_map" "socialmediaserver_stg_external" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/socialmediaserver-stg-ext-be"
  name            = "socialmediaserver-stg-external"
  project         = "waze-development"
}

resource "google_compute_url_map" "staging_marx" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/staging-marx"
  name            = "staging-marx"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_adman2_stg" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/stg-adman2-stg-backend-service"
  name            = "stg-adman2-stg"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_ads" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/stg-ads-backend-service"
  name            = "stg-ads"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_adsassets" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendBuckets/waze-ads-resources-test"
  name            = "stg-adsassets"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_carpoolrouting_lb" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/stg-carpoolrouting-future-p6-bes"
  name            = "stg-carpoolrouting-lb"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_dev_editor" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/stg-dev-editor-backend"
  name            = "stg-dev-editor"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_elton" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/stg-elton"
  name            = "stg-elton"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_mapnik_editor" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/stg-mapnik-editor-backend-service"
  name            = "stg-mapnik-editor"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_mobile_web" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/stg-mobile-web-backend-service"
  name            = "stg-mobile-web"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_routing_lb" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/stg-routing-2-backend-service"
  name            = "stg-routing-lb"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_rtproxy" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/stg-rtproxy-backend-service"
  name            = "stg-rtproxy"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_search_lb_urlmap" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/stg-search-lb-service"
  name            = "stg-search-lb-urlmap"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_sms_lb_urlmap" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/stg-sms-lb-service"
  name            = "stg-sms-lb-urlmap"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_ttsgw" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/stg-ttsgw-backend-service"
  name            = "stg-ttsgw"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_usersprofile_lb" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/stg-usersprofile-lb-backend-service"
  name            = "stg-usersprofile-lb"
  project         = "waze-development"
}

resource "google_compute_url_map" "stg_voice_prompts" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendBuckets/stg-voice-prompts-bucket"
  name            = "stg-voice-prompts"
  project         = "waze-development"
}

resource "google_compute_url_map" "storagegateway_stg_ext" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/storagegateway-stg-stg1-bs"
  name            = "storagegateway-stg-ext"
  project         = "waze-development"
}

resource "google_compute_url_map" "usersprofile_stg_stg1_map" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/usersprofile-stg-stg1-ebes"
  name            = "usersprofile-stg-stg1-map"
  project         = "waze-development"
}

resource "google_compute_url_map" "was_stg_stg1" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/was-stg-stg1-be"
  name            = "was-stg-stg1"
  project         = "waze-development"
}

resource "google_compute_url_map" "waze_carpool_groups_images_stg" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendBuckets/waze-carpool-groups-images-stg"
  name            = "waze-carpool-groups-images-stg"
  project         = "waze-development"
}

resource "google_compute_url_map" "waze_reverseproxy_stg" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/reverseproxy-stg-stg1-ebes"
  name            = "waze-reverseproxy-stg"
  project         = "waze-development"
}

resource "google_compute_url_map" "website_stg" {
  default_service = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/website-stg-be"
  name            = "website-stg"
  project         = "waze-development"
}
