provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_http_health_check" "admanage_stg" {
  check_interval_sec  = "15"
  healthy_threshold   = "2"
  name                = "admanage-stg"
  port                = "81"
  project             = "waze-development"
  request_path        = "/AdManage/jax/sanity"
  timeout_sec         = "10"
  unhealthy_threshold = "5"
}

resource "google_compute_http_health_check" "clienttiles_stg_stg1_hc" {
  check_interval_sec  = "15"
  healthy_threshold   = "2"
  name                = "clienttiles-stg-stg1-hc"
  port                = "80"
  project             = "waze-development"
  request_path        = "/IL1.1/77001_17/77001_170c/77001_170c1e/77001_170c1e87.wdf?sessionid=269435532\u0026cookie=9kYSva7T0RDIDVtY"
  timeout_sec         = "5"
  unhealthy_threshold = "5"
}

resource "google_compute_http_health_check" "codelab_test_hc_1525801785983" {
  check_interval_sec  = "1"
  healthy_threshold   = "1"
  name                = "codelab-test-hc-1525801785983"
  port                = "80"
  project             = "waze-development"
  request_path        = "/"
  timeout_sec         = "1"
  unhealthy_threshold = "1"
}

resource "google_compute_http_health_check" "codelab_test_hc_1525802299888" {
  check_interval_sec  = "1"
  healthy_threshold   = "1"
  name                = "codelab-test-hc-1525802299888"
  port                = "8080"
  project             = "waze-development"
  request_path        = "/hello"
  timeout_sec         = "1"
  unhealthy_threshold = "1"
}

resource "google_compute_http_health_check" "default_health_check" {
  check_interval_sec  = "5"
  healthy_threshold   = "2"
  name                = "default-health-check"
  port                = "80"
  project             = "waze-development"
  request_path        = "/"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_http_health_check" "http_healthcheck" {
  check_interval_sec  = "120"
  healthy_threshold   = "2"
  name                = "http-healthcheck"
  port                = "80"
  project             = "waze-development"
  request_path        = "/healthcheck"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_http_health_check" "k8s_865dbaeaed1d6cc4_node" {
  check_interval_sec  = "2"
  description         = "{\"kubernetes.io/service-name\":\"k8s-865dbaeaed1d6cc4-node\"}"
  healthy_threshold   = "1"
  name                = "k8s-865dbaeaed1d6cc4-node"
  port                = "10256"
  project             = "waze-development"
  request_path        = "/healthz"
  timeout_sec         = "1"
  unhealthy_threshold = "5"
}

resource "google_compute_http_health_check" "lb_check" {
  check_interval_sec  = "5"
  healthy_threshold   = "2"
  name                = "lb-check"
  port                = "443"
  project             = "waze-development"
  request_path        = "/"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_http_health_check" "mapnik_livemap_hc" {
  check_interval_sec  = "10"
  healthy_threshold   = "2"
  name                = "mapnik-livemap-hc"
  port                = "80"
  project             = "waze-development"
  request_path        = "/tiles/ping"
  timeout_sec         = "5"
  unhealthy_threshold = "5"
}

resource "google_compute_http_health_check" "microservice_healthcheck" {
  check_interval_sec  = "10"
  description         = "Common MicroService health check"
  healthy_threshold   = "2"
  name                = "microservice-healthcheck"
  port                = "80"
  project             = "waze-development"
  request_path        = "/healthcheck"
  timeout_sec         = "10"
  unhealthy_threshold = "10"
}

resource "google_compute_http_health_check" "microservice_port_82" {
  check_interval_sec  = "120"
  healthy_threshold   = "1"
  name                = "microservice-port-82"
  port                = "82"
  project             = "waze-development"
  request_path        = "/healthcheck"
  timeout_sec         = "62"
  unhealthy_threshold = "5"
}

resource "google_compute_http_health_check" "realtime_stg_stg1_hc" {
  check_interval_sec  = "15"
  description         = "Realtime staging internal healthcheck"
  healthy_threshold   = "2"
  name                = "realtime-stg-stg1-hc"
  port                = "80"
  project             = "waze-development"
  request_path        = "/elbtest"
  timeout_sec         = "10"
  unhealthy_threshold = "3"
}

resource "google_compute_http_health_check" "stg_adman2_exp_hc" {
  check_interval_sec  = "15"
  healthy_threshold   = "2"
  name                = "stg-adman2-exp-hc"
  port                = "30081"
  project             = "waze-development"
  request_path        = "/health"
  timeout_sec         = "10"
  unhealthy_threshold = "5"
}

resource "google_compute_http_health_check" "stg_adman2_hc" {
  check_interval_sec  = "15"
  healthy_threshold   = "2"
  name                = "stg-adman2-hc"
  port                = "30080"
  project             = "waze-development"
  request_path        = "/health"
  timeout_sec         = "10"
  unhealthy_threshold = "5"
}

resource "google_compute_http_health_check" "stg_http_elbtest_generic" {
  check_interval_sec  = "60"
  healthy_threshold   = "2"
  name                = "stg-http-elbtest-generic"
  port                = "80"
  project             = "waze-development"
  request_path        = "/elbtest"
  timeout_sec         = "50"
  unhealthy_threshold = "10"
}

resource "google_compute_http_health_check" "stg_poi_healthcheck" {
  check_interval_sec  = "30"
  healthy_threshold   = "2"
  name                = "stg-poi-healthcheck"
  port                = "80"
  project             = "waze-development"
  request_path        = "/ServerManager/ExecuteCommand?command=HealthCheck\u0026server=PoiServer\u0026id=local"
  timeout_sec         = "30"
  unhealthy_threshold = "10"
}

resource "google_compute_http_health_check" "stg_routing_http" {
  check_interval_sec  = "60"
  healthy_threshold   = "2"
  name                = "stg-routing-http"
  port                = "80"
  project             = "waze-development"
  request_path        = "/elbtest"
  timeout_sec         = "59"
  unhealthy_threshold = "10"
}

resource "google_compute_http_health_check" "stg_routing_http_healthcheck" {
  check_interval_sec  = "60"
  healthy_threshold   = "2"
  name                = "stg-routing-http-healthcheck"
  port                = "80"
  project             = "waze-development"
  request_path        = "/elbtest"
  timeout_sec         = "59"
  unhealthy_threshold = "10"
}

resource "google_compute_http_health_check" "stg_search_check" {
  check_interval_sec  = "5"
  healthy_threshold   = "2"
  name                = "stg-search-check"
  port                = "80"
  project             = "waze-development"
  request_path        = "/"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_http_health_check" "stg_search_lb_http_check" {
  check_interval_sec  = "5"
  healthy_threshold   = "2"
  name                = "stg-search-lb-http-check"
  port                = "80"
  project             = "waze-development"
  request_path        = "/"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_http_health_check" "stg_sms_lb_http_check" {
  check_interval_sec  = "5"
  healthy_threshold   = "2"
  name                = "stg-sms-lb-http-check"
  port                = "80"
  project             = "waze-development"
  request_path        = "/"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_http_health_check" "stg_sms_lb_https_check" {
  check_interval_sec  = "5"
  healthy_threshold   = "2"
  name                = "stg-sms-lb-https-check"
  port                = "443"
  project             = "waze-development"
  request_path        = "/"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_http_health_check" "stg_usersprofile_lb_health_check" {
  check_interval_sec  = "5"
  healthy_threshold   = "2"
  name                = "stg-usersprofile-lb-health-check"
  port                = "80"
  project             = "waze-development"
  request_path        = "/"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_http_health_check" "topicserver_online_stg1_hc_1466065312029" {
  check_interval_sec  = "10"
  healthy_threshold   = "2"
  name                = "topicserver-online-stg1-hc-1466065312029"
  port                = "80"
  project             = "waze-development"
  request_path        = "/elbtest"
  timeout_sec         = "5"
  unhealthy_threshold = "10"
}

resource "google_compute_http_health_check" "usage_stg_stg1_hc_1466946190545" {
  check_interval_sec  = "5"
  healthy_threshold   = "2"
  name                = "usage-stg-stg1-hc-1466946190545"
  port                = "80"
  project             = "waze-development"
  request_path        = "/"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_http_health_check" "website_hc" {
  check_interval_sec  = "10"
  healthy_threshold   = "2"
  name                = "website-hc"
  port                = "80"
  project             = "waze-development"
  request_path        = "/monitoring/nginx"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}
