provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_health_check" "admanage_stg_int" {
  check_interval_sec = "15"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "81"
    proxy_header = "NONE"
    request_path = "/AdManage/jax/sanity"
  }

  name                = "admanage-stg-int"
  project             = "waze-development"
  timeout_sec         = "10"
  unhealthy_threshold = "5"
}

resource "google_compute_health_check" "adsmanagementapi_hc" {
  check_interval_sec  = "60"
  healthy_threshold   = "2"
  name                = "adsmanagementapi-hc"
  project             = "waze-development"
  timeout_sec         = "15"
  unhealthy_threshold = "2"
}

resource "google_compute_health_check" "cassandra_healthcheck" {
  check_interval_sec = "60"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "80"
    proxy_header = "NONE"
    request_path = "/healthcheck"
  }

  name                = "cassandra-healthcheck"
  project             = "waze-development"
  timeout_sec         = "30"
  unhealthy_threshold = "2"
}

resource "google_compute_health_check" "dategiver_test" {
  check_interval_sec = "5"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "8088"
    proxy_header = "NONE"
    request_path = "/getdate"
  }

  name                = "dategiver-test"
  project             = "waze-development"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_health_check" "fluentd_hc" {
  check_interval_sec = "5"
  healthy_threshold  = "2"
  name               = "fluentd-hc"
  project            = "waze-development"

  tcp_health_check = {
    port         = "24224"
    proxy_header = "NONE"
  }

  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_health_check" "generic_memcached_hc" {
  check_interval_sec = "10"
  description        = "Generic memcached health check"
  healthy_threshold  = "2"
  name               = "generic-memcached-hc"
  project            = "waze-development"

  tcp_health_check = {
    port         = "11211"
    proxy_header = "NONE"
  }

  timeout_sec         = "5"
  unhealthy_threshold = "10"
}

resource "google_compute_health_check" "georegistry_healthcheck" {
  check_interval_sec = "10"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "80"
    proxy_header = "NONE"
    request_path = "/GeoRegistry/registry/describeIPC"
  }

  name                = "georegistry-healthcheck"
  project             = "waze-development"
  timeout_sec         = "10"
  unhealthy_threshold = "2"
}

resource "google_compute_health_check" "http_basic_check" {
  check_interval_sec = "5"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "8088"
    proxy_header = "NONE"
    request_path = "/cluster"
  }

  name                = "http-basic-check"
  project             = "waze-development"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_health_check" "http_healthcheck" {
  check_interval_sec = "120"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "80"
    proxy_header = "NONE"
    request_path = "/healthcheck"
  }

  name                = "http-healthcheck"
  project             = "waze-development"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_health_check" "internal_hc" {
  check_interval_sec = "5"
  healthy_threshold  = "2"
  name               = "internal-hc"
  project            = "waze-development"

  tcp_health_check = {
    port         = "83"
    proxy_header = "NONE"
  }

  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_health_check" "jupyter_http_basic_check_http" {
  check_interval_sec = "5"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "8088"
    proxy_header = "NONE"
    request_path = "/cluster"
  }

  name                = "jupyter-http-basic-check-http"
  project             = "waze-development"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_health_check" "jupyter_http_basic_check_other" {
  check_interval_sec = "5"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "8088"
    proxy_header = "NONE"
    request_path = "/cluster"
  }

  name                = "jupyter-http-basic-check-other"
  project             = "waze-development"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_health_check" "jupyter_http_basic_check_spark" {
  check_interval_sec = "5"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "8088"
    proxy_header = "NONE"
    request_path = "/cluster"
  }

  name                = "jupyter-http-basic-check-spark"
  project             = "waze-development"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_health_check" "jupyter_stg_healthcheck" {
  check_interval_sec = "5"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "8088"
    proxy_header = "NONE"
    request_path = "/cluster"
  }

  name                = "jupyter-stg-healthcheck"
  project             = "waze-development"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_health_check" "mapnik_livemap_hc" {
  check_interval_sec = "10"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "80"
    proxy_header = "NONE"
    request_path = "/tiles/ping"
  }

  name                = "mapnik-livemap-hc"
  project             = "waze-development"
  timeout_sec         = "5"
  unhealthy_threshold = "5"
}

resource "google_compute_health_check" "microservice_healthcheck" {
  check_interval_sec = "10"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "80"
    proxy_header = "NONE"
    request_path = "/healthcheck"
  }

  name                = "microservice-healthcheck"
  project             = "waze-development"
  timeout_sec         = "10"
  unhealthy_threshold = "10"
}

resource "google_compute_health_check" "microservice_port_82" {
  check_interval_sec = "120"
  healthy_threshold  = "1"

  http_health_check = {
    port         = "82"
    proxy_header = "NONE"
    request_path = "/healthcheck"
  }

  name                = "microservice-port-82"
  project             = "waze-development"
  timeout_sec         = "62"
  unhealthy_threshold = "5"
}

resource "google_compute_health_check" "realtime_proxy_stg_hc" {
  check_interval_sec = "12"
  healthy_threshold  = "3"

  http_health_check = {
    port         = "80"
    proxy_header = "PROXY_V1"
    request_path = "/rtserver/distrib/login"
  }

  name                = "realtime-proxy-stg-hc"
  project             = "waze-development"
  timeout_sec         = "5"
  unhealthy_threshold = "2"
}

resource "google_compute_health_check" "realtime_stg_internal_hc" {
  check_interval_sec = "15"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "80"
    proxy_header = "NONE"
    request_path = "/elbtest"
  }

  name                = "realtime-stg-internal-hc"
  project             = "waze-development"
  timeout_sec         = "10"
  unhealthy_threshold = "5"
}

resource "google_compute_health_check" "routingserver_int_hc" {
  check_interval_sec = "20"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "80"
    proxy_header = "NONE"
    request_path = "/elbtest"
  }

  name                = "routingserver-int-hc"
  project             = "waze-development"
  timeout_sec         = "15"
  unhealthy_threshold = "10"
}

resource "google_compute_health_check" "stg_carpoolrouting_lb_hc" {
  check_interval_sec = "120"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "82"
    proxy_header = "NONE"
    request_path = "/healthcheck"
  }

  name                = "stg-carpoolrouting-lb-hc"
  project             = "waze-development"
  timeout_sec         = "62"
  unhealthy_threshold = "2"
}

resource "google_compute_health_check" "usersprofile_stg_hc" {
  check_interval_sec = "30"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "81"
    proxy_header = "NONE"
    request_path = "/healthcheck"
  }

  name                = "usersprofile-stg-hc"
  project             = "waze-development"
  timeout_sec         = "10"
  unhealthy_threshold = "6"
}

resource "google_compute_health_check" "was_stg_hc" {
  check_interval_sec = "15"
  healthy_threshold  = "2"

  http_health_check = {
    port         = "80"
    proxy_header = "NONE"
    request_path = "/WAS/map_info"
  }

  name                = "was-stg-hc"
  project             = "waze-development"
  timeout_sec         = "5"
  unhealthy_threshold = "6"
}
