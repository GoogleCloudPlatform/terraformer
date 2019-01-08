provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_target_http_proxy" "carpoolmatchingdispatcher_stg_stg1_target_http_proxy" {
  name    = "carpoolmatchingdispatcher-stg-stg1-target-http-proxy"
  project = "waze-development"
  url_map = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/carpoolmatchingdispatcher-stg"
}

resource "google_compute_target_http_proxy" "stg_carpoolrouting_lb_target_proxy" {
  name    = "stg-carpoolrouting-lb-target-proxy"
  project = "waze-development"
  url_map = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-carpoolrouting-lb"
}

resource "google_compute_target_http_proxy" "stg_mobile_web_target_2" {
  name    = "stg-mobile-web-target-2"
  project = "waze-development"
  url_map = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-mobile-web"
}

resource "google_compute_target_http_proxy" "stg_search_lb" {
  name    = "stg-search-lb"
  project = "waze-development"
  url_map = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-search-lb-urlmap"
}

resource "google_compute_target_http_proxy" "stg_sms_lb" {
  name    = "stg-sms-lb"
  project = "waze-development"
}

resource "google_compute_target_http_proxy" "stg_usersprofile_lb" {
  name    = "stg-usersprofile-lb"
  project = "waze-development"
  url_map = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/stg-usersprofile-lb"
}

resource "google_compute_target_http_proxy" "website_stg_target_proxy_2" {
  name    = "website-stg-target-proxy-2"
  project = "waze-development"
  url_map = "https://www.googleapis.com/compute/v1/projects/waze-development/global/urlMaps/website-stg"
}
