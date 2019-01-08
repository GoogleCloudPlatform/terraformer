provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_https_health_check" "aef_realtime__login__prober__prod__il_20181231t154600_00ahs" {
  check_interval_sec  = "30"
  healthy_threshold   = "1"
  name                = "aef-realtime--login--prober--prod--il-20181231t154600-00ahs"
  port                = "8443"
  project             = "waze-development"
  request_path        = "/_ah/health"
  timeout_sec         = "1"
  unhealthy_threshold = "10"
}

resource "google_compute_https_health_check" "aef_realtime__login__prober__prod__il_20181231t154600_hcs" {
  check_interval_sec  = "1"
  healthy_threshold   = "1"
  name                = "aef-realtime--login--prober--prod--il-20181231t154600-hcs"
  port                = "8443"
  project             = "waze-development"
  request_path        = "/_ah/health"
  timeout_sec         = "1"
  unhealthy_threshold = "1"
}

resource "google_compute_https_health_check" "aef_realtime__login__prober__stg_20181231t100523_00ahs" {
  check_interval_sec  = "30"
  healthy_threshold   = "1"
  name                = "aef-realtime--login--prober--stg-20181231t100523-00ahs"
  port                = "8443"
  project             = "waze-development"
  request_path        = "/_ah/health"
  timeout_sec         = "1"
  unhealthy_threshold = "10"
}

resource "google_compute_https_health_check" "aef_realtime__login__prober__stg_20181231t100523_hcs" {
  check_interval_sec  = "1"
  healthy_threshold   = "1"
  name                = "aef-realtime--login--prober--stg-20181231t100523-hcs"
  port                = "8443"
  project             = "waze-development"
  request_path        = "/_ah/health"
  timeout_sec         = "1"
  unhealthy_threshold = "1"
}

resource "google_compute_https_health_check" "aef_routing__regression_20180124t102927_00ahs" {
  check_interval_sec  = "30"
  healthy_threshold   = "1"
  name                = "aef-routing--regression-20180124t102927-00ahs"
  port                = "8443"
  project             = "waze-development"
  request_path        = "/_ah/health"
  timeout_sec         = "1"
  unhealthy_threshold = "10"
}

resource "google_compute_https_health_check" "aef_routing__regression_20180124t102927_hcs" {
  check_interval_sec  = "1"
  healthy_threshold   = "1"
  name                = "aef-routing--regression-20180124t102927-hcs"
  port                = "8443"
  project             = "waze-development"
  request_path        = "/_ah/health"
  timeout_sec         = "1"
  unhealthy_threshold = "1"
}

resource "google_compute_https_health_check" "default_health_check" {
  check_interval_sec  = "5"
  healthy_threshold   = "2"
  name                = "default-health-check"
  port                = "443"
  project             = "waze-development"
  request_path        = "/"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_https_health_check" "realtime_georss_stg1_ssl_hc" {
  check_interval_sec  = "10"
  healthy_threshold   = "2"
  name                = "realtime-georss-stg1-ssl-hc"
  port                = "443"
  project             = "waze-development"
  request_path        = "/rtserver/web/TGeoRSS?ma=600\u0026mj=100\u0026mu=100\u0026left=34.092498779296875\u0026right=35.810760498046875\u0026bottom=31.621848297182485\u0026top=32.40064431296675\u0026_=1438583611101\u0026tk=web"
  timeout_sec         = "5"
  unhealthy_threshold = "5"
}

resource "google_compute_https_health_check" "stg_rtproxy_health_check" {
  check_interval_sec  = "60"
  healthy_threshold   = "2"
  name                = "stg-rtproxy-health-check"
  port                = "443"
  project             = "waze-development"
  request_path        = "/rtserver/distrib/login"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_https_health_check" "tcp_test" {
  check_interval_sec  = "5"
  healthy_threshold   = "2"
  name                = "tcp-test"
  port                = "443"
  project             = "waze-development"
  request_path        = "/rtserver/distrib/login"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_https_health_check" "website_hc_https" {
  check_interval_sec  = "5"
  description         = "https://buganizer.corp.google.com/issues/70453472"
  healthy_threshold   = "2"
  name                = "website-hc-https"
  port                = "443"
  project             = "waze-development"
  request_path        = "/health"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}
