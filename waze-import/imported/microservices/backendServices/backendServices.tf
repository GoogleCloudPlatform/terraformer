provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_backend_service" "admanage_stg_stg1_backend_service" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-d/instanceGroups/admanage-stg-stg1-v032"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "1000"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/admanage-stg"]
  name                            = "admanage-stg-stg1-backend-service"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "adsmanagementapi_stg_stg1_proxy_be" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/adsmanagementapi-stg-stg1-v004"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/adsmanagementapi-stg-stg1-v005"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/healthChecks/adsmanagementapi-hc"]
  name                            = "adsmanagementapi-stg-stg1-proxy-be"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP2"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "aef_realtime__login__prober__prod__il_20181231t154600_bs" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/us-central1/instanceGroups/aef-realtime--login--prober--prod--il-20181231t154600-00"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpsHealthChecks/aef-realtime--login--prober--prod--il-20181231t154600-hcs"]
  name                            = "aef-realtime--login--prober--prod--il-20181231t154600-bs"
  port_name                       = "https"
  project                         = "waze-development"
  protocol                        = "HTTPS"
  session_affinity                = "NONE"
  timeout_sec                     = "3610"
}

resource "google_compute_backend_service" "aef_realtime__login__prober__stg_20181231t100523_bs" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/us-central1/instanceGroups/aef-realtime--login--prober--stg-20181231t100523-00"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpsHealthChecks/aef-realtime--login--prober--stg-20181231t100523-hcs"]
  name                            = "aef-realtime--login--prober--stg-20181231t100523-bs"
  port_name                       = "https"
  project                         = "waze-development"
  protocol                        = "HTTPS"
  session_affinity                = "NONE"
  timeout_sec                     = "3610"
}

resource "google_compute_backend_service" "aef_routing__regression_20180124t102927_bs" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/us-central1/instanceGroups/aef-routing--regression-20180124t102927-00"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpsHealthChecks/aef-routing--regression-20180124t102927-hcs"]
  name                            = "aef-routing--regression-20180124t102927-bs"
  port_name                       = "https"
  project                         = "waze-development"
  protocol                        = "HTTPS"
  session_affinity                = "NONE"
  timeout_sec                     = "3610"
}

resource "google_compute_backend_service" "carpoolmatchingdispatcher_stg_stg1" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatchingdispatcher-bundle-p4-stg1-v016"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "80"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatchingdispatcher-bundle-p3-stg1-v014"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "80"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatchingdispatcher-bundle-p3-stg1-v013"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "80"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatchingdispatcher-bundle-p0-stg1-v018"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "80"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatchingdispatcher-bundle-p7-stg1-v015"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "80"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatchingdispatcher-bundle-p1-stg1-v018"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "80"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatchingdispatcher-bundle-p2-stg1-v013"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "80"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatchingdispatcher-bundle-p2-stg1-v014"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "80"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatchingdispatcher-bundle-p0-stg1-v017"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "80"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatchingdispatcher-bundle-p5-stg1-v016"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "80"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatchingdispatcher-bundle-p6-stg1-v013"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "80"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatchingdispatcher-bundle-p6-stg1-v014"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "80"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatchingdispatcher-bundle-p5-stg1-v015"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "80"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatchingdispatcher-stg-stg1-v036"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "100"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "carpoolmatchingdispatcher-stg-stg1"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "carpoolpayments_stg_stg1_bes" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolpayments-stg-stg1-v548"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "carpoolpayments-stg-stg1-bes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "carpooltesting_realtime_proxy_stg_be" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "0.5"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpooltesting-stg-realtime-v005"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.5"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "carpooltesting-realtime-proxy-stg-be"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "clienttiles_stg_stg1_backend_service" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/clienttiles-stg-stg1-v008"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/clienttiles-stg-stg1-v007"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/clienttiles-stg-stg1-hc"]
  name                            = "clienttiles-stg-stg1-backend-service"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "inbox_stg_stg1" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-d/instanceGroups/inbox-stg-stg1-v011"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.6"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = true
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "inbox-stg-stg1"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "incidents_stg_stg1_ebes" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/incidents-stg-stg1-v002"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "10000"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "incidents-stg-stg1-ebes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "incidentsserver_stg_stg1_ebes" {
  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "incidentsserver-stg-stg1-ebes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "mapnikserver_livemap_stg1" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/mapnik-livemap-stg1-v004"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/healthChecks/mapnik-livemap-hc"]
  name                            = "mapnikserver-livemap-stg1"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "parkingrouting_stg_stg1_ebes" {
  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "parkingrouting-stg-stg1-ebes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "parkingserver_stg_stg1_ebes" {
  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "parkingserver-stg-stg1-ebes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "realtime_georss_stg1_backend_service" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-d/instanceGroups/realtime-georss-stg-v019"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpsHealthChecks/realtime-georss-stg1-ssl-hc"]
  name                            = "realtime-georss-stg1-backend-service"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTPS"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "realtime_proxy_stg_be" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-d/instanceGroups/realtime-proxy-stg-v019"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.6"
  }

  cdn_policy = {
    cache_key_policy = {
      include_host         = true
      include_protocol     = true
      include_query_string = true
    }
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/healthChecks/realtime-proxy-stg-hc"]
  name                            = "realtime-proxy-stg-be"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "TCP"
  session_affinity                = "NONE"
  timeout_sec                     = "300"
}

resource "google_compute_backend_service" "repositoy_server_stg_stg1_bes" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/repository-server-stg-stg1-v006"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.4"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "repositoy-server-stg-stg1-bes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "reverseproxy_stg_stg1_ebes" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/reverseproxy-stg-stg1-v001"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "reverseproxy-stg-stg1-ebes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "searchserver_stg_stg1_external_bes" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/searchserver-heavyduty-stg1-v006"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.4"
  }

  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/searchserver-stg-stg1-v012"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.4"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "searchserver-stg-stg1-external-bes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "servermanager_external" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "0.7"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/servermanager-stg-stg1-v005"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "1000"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "servermanager-external"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "180"
}

resource "google_compute_backend_service" "socialmediaserver_stg_ext_be" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/socialmediaserver-stg-stg1-v010"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.5"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = true
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "socialmediaserver-stg-ext-be"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "staging_marx" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-b/instanceGroups/staging-marx"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/website-hc"]
  name                            = "staging-marx"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "300"
}

resource "google_compute_backend_service" "stg_adman2_stg_backend_service" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-d/instanceGroups/gke-adman2-stg-pool3-f0c2f23b-grp"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/stg-adman2-hc"]
  name                            = "stg-adman2-stg-backend-service"
  port_name                       = "http-30080"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_adman2_stg_exp_backend_service" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-d/instanceGroups/gke-adman2-stg-pool3-f0c2f23b-grp"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/stg-adman2-exp-hc"]
  name                            = "stg-adman2-stg-exp-backend-service"
  port_name                       = "http-30081"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_ads_backend_service" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-b/instanceGroups/stg-ads-ig"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/default-health-check"]
  name                            = "stg-ads-backend-service"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_carpoolrouting_bes" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolrouting-stg-stg1-v270"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "1000"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "stg-carpoolrouting-bes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_carpoolrouting_future_p1_bes" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolrouting-future-p1-stg1-v214"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "1000"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "stg-carpoolrouting-future-p1-bes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_carpoolrouting_future_p2_bes" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolrouting-future-p2-stg1-v209"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "1000"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "stg-carpoolrouting-future-p2-bes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_carpoolrouting_future_p3_bes" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolrouting-future-p3-stg1-v209"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "1000"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "stg-carpoolrouting-future-p3-bes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_carpoolrouting_future_p4_bes" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolrouting-future-p4-stg1-v209"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "1000"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "stg-carpoolrouting-future-p4-bes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_carpoolrouting_future_p5_bes" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolrouting-future-p5-stg1-v209"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "1000"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "stg-carpoolrouting-future-p5-bes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_carpoolrouting_future_p6_bes" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "0.8"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolrouting-future-p7-stg1-v114"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "1000"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "0.8"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolrouting-future-p8-stg1-v114"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "1000"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "0.8"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/carpoolrouting-future-p6-stg1-v207"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "1000"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "stg-carpoolrouting-future-p6-bes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_dev_editor_backend" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-b/instanceGroups/dev-editor"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/website-hc"]
  name                            = "stg-dev-editor-backend"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_elton" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/elton-stg-stg1-v015"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "1000"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "stg-elton"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_mapnik_editor_backend_service" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-b/instanceGroups/mapnik-editor-stg1-v015"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/default-health-check"]
  name                            = "stg-mapnik-editor-backend-service"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_mobile_web_backend_service" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-b/instanceGroups/stg-mobile-web-lb"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/default-health-check"]
  name                            = "stg-mobile-web-backend-service"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_prompto_backend" {
  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "stg-prompto-backend"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "180"
}

resource "google_compute_backend_service" "stg_routing_2_backend_service" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/routingserver-stg-2-stg1-v486"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "1000"
    max_utilization              = "0"
  }

  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/routingserver-stg-tiles-v001"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "1000"
    max_utilization              = "0"
  }

  cdn_policy = {
    cache_key_policy = {
      include_host         = true
      include_protocol     = true
      include_query_string = true
    }
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/stg-routing-http"]
  name                            = "stg-routing-2-backend-service"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_rtproxy_backend_service" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-d/instanceGroups/realtime-proxy-stg-v019"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.6"
  }

  cdn_policy = {
    cache_key_policy = {
      include_host         = true
      include_protocol     = true
      include_query_string = true
    }
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/healthChecks/realtime-proxy-stg-hc"]
  name                            = "stg-rtproxy-backend-service"
  port_name                       = "http-lb"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_search_backend_service" {
  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/default-health-check"]
  name                            = "stg-search-backend-service"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_search_lb_service" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-c/instanceGroups/stg-search-ig"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/stg-search-lb-http-check"]
  name                            = "stg-search-lb-service"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_sms_lb_service" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-b/instanceGroups/stg-sms-ig"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/stg-sms-lb-http-check"]
  name                            = "stg-sms-lb-service"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_ttsgw_backend_service" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-b/instanceGroups/stg-ttsgw"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/default-health-check"]
  name                            = "stg-ttsgw-backend-service"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "stg_usersprofile_lb_backend_service" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-b/instanceGroups/stg-usersprofile"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "10000"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/stg-usersprofile-lb-health-check"]
  name                            = "stg-usersprofile-lb-backend-service"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "storagegateway_stg_stg1_bs" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/storagegateway-stg-stg1-v007"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "10000"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/http-healthcheck"]
  name                            = "storagegateway-stg-stg1-bs"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "usersprofile_stg_stg1_ebes" {
  backend = {
    balancing_mode               = "RATE"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/usersprofile-stg-stg1-v124"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "10000"
    max_utilization              = "0"
  }

  connection_draining_timeout_sec = "0"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/microservice-healthcheck"]
  name                            = "usersprofile-stg-stg1-ebes"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "was_stg_stg1_be" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/regions/europe-west1/instanceGroups/was-stg-stg1-v002"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "300"
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/healthChecks/was-stg-hc"]
  name                            = "was-stg-stg1-be"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_backend_service" "website_stg_be" {
  backend = {
    balancing_mode               = "UTILIZATION"
    capacity_scaler              = "1"
    group                        = "https://www.googleapis.com/compute/beta/projects/waze-development/zones/europe-west1-c/instanceGroups/website-stg"
    max_connections              = "0"
    max_connections_per_instance = "0"
    max_rate                     = "0"
    max_rate_per_instance        = "0"
    max_utilization              = "0.8"
  }

  connection_draining_timeout_sec = "300"
  custom_request_headers          = ["X-Client-City-Lat-Long: {client_city_lat_long}", "X-Client-Region: {client_region}", "X-Client-City: {client_city}", "X-Rtt-Msec: {client_rtt_msec}"]
  enable_cdn                      = false
  health_checks                   = ["https://www.googleapis.com/compute/beta/projects/waze-development/global/httpHealthChecks/website-hc"]
  name                            = "website-stg-be"
  port_name                       = "http"
  project                         = "waze-development"
  protocol                        = "HTTP"
  session_affinity                = "NONE"
  timeout_sec                     = "300"
}
