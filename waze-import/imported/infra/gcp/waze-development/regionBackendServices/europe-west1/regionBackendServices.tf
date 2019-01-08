provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_region_backend_service" "admanage_stg_stg1_internal" {
  connection_draining_timeout_sec = "0"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/admanage-stg-int"]
  name                            = "admanage-stg-stg1-internal"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "adreports_pubsub_to_bq_server_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/adreports-pubsub-to-bq-server-stg-stg1-v002"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "adreports-pubsub-to-bq-server-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "advendors_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/advendors-stg-stg1-v004"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "advendors-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "alerts_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/alerts-stg-stg1-v106"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "alerts-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "ateam_stg_stg1_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/ateam-stg-stg1-v004"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "ateam-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "carpool_dev_nyc_stg_stg1_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpool-dev-nyc-stg-stg1"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "carpool-dev-nyc-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "carpool_memcache_stg_stg1" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/ads-memcached-stg1-v001"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpool-memcached-stg1-v005"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/staging-memcached-stg1-v001"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/generic-memcached-hc"]
  name                            = "carpool-memcache-stg-stg1"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "carpoolcommutedaily_modelsbundle_stg1" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpoolcommutedaily-modelsbundle-stg1-v016"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "carpoolcommutedaily-modelsbundle-stg1"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "carpoolgroups_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpoolgroups-stg-stg1-v079"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "carpoolgroups-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "carpoolindex_stg_stg1_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpoolindex-stg-stg1-v034"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "carpoolindex-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "carpoolmanager_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmanager-stg-stg1-v290"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmanager-test-test-shrink-v003"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "carpoolmanager-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "carpoolranking_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpoolranking-stg-stg1-v349"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "carpoolranking-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "carpoolreviews_stg_stg1_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpoolreviews-stg-stg1-v122"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "carpoolreviews-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "cars_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/cars-stg-stg1-v080"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/cars-nird-stg1-v002"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/cars-nird-stg1-v001"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "cars-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "cartool_stg_stg1_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/cartool-stg-stg1-v005"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "cartool-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "dbeventwriter_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/dbeventwriter-stg-stg1-v002"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "dbeventwriter-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "descartes_dev_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/descartes-dev-stg1-v000"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/descartes-dev-stg1-v001"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "descartes-dev-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "descartes_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/descartes-stg-stg1-v011"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "descartes-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "driveplanner_stg_stg1_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/driveplanner-stg-stg1-v004"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "driveplanner-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "engagement_stg_stg1_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/engagement-stg-stg1-v013"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "engagement-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "fluentd_forwarder_ng" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/fluentd-forwarder-ng-stg1-v013"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/fluentd-hc"]
  name                            = "fluentd-forwarder-ng"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "fluentd_forwarder_stg" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/fluentd-forwarder-stg1-v002"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/fluentd-hc"]
  name                            = "fluentd-forwarder-stg"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "gagentproxy_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/gagentproxy-stg-stg1-v015"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "gagentproxy-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "gamingserver_stg_stg1_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/gamingserver-stg-stg1-v001"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "gamingserver-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "geocodingserver_incremental_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/geocodingserver-incremental-stg-stg1-v018"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "geocodingserver-incremental-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "geocodingserver_offline_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/geocodingserver-offline-stg-stg1-v024"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "geocodingserver-offline-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "georegistry_internal_stg" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/georegistry-stg-stg1-v157"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/georegistry-healthcheck"]
  name                            = "georegistry-internal-stg"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "georegistry_offline_stg_stg1" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/georegistry-offline-stg-stg1-v004"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/georegistry-healthcheck"]
  name                            = "georegistry-offline-stg-stg1"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "georegistry_redis_stg_stg1_ibes" {
  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "georegistry-redis-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "internal_merger_jupyter_staging_debug" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/merger-jupyter-staging"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/jupyter-stg-healthcheck"]
  name                            = "internal-merger-jupyter-staging-debug"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "mapnikserver_editor_stg1" {
  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/mapnik-livemap-hc"]
  name                            = "mapnikserver-editor-stg1"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "metrics_stg_stg1" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/metrics-stg-stg1-v061"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "metrics-stg-stg1"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "microservice_stg_internal" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/geoindex-stg-stg1-v010"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpoolmatching-stg-stg1-v579"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/mte-stg-stg1-v003"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpoolcommute-stg-stg1-v293"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/closures-stg-stg1-v002"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/closures-stg-stg1-v003"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/adreports-session-stg-stg1-v034"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpooladapter-stg-stg1-v642"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/pickups-stg-v004"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpoolpricing-stg-stg1-v520"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/datasetmanager-stg-stg1-v003"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/usertracking-stg-stg1-v001"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/brandsserver-stg-stg1-v001"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/brandsserver-stg-stg1-v000"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/configuration-stg-stg1-v006"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/adreports-placed-stg-stg1-v006"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpooldispatcher-stg-stg1-v120"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpooltesting-stg-infra-stg1-v010"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/dockermanager-stg-stg1-v010"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/adpacing-stg-stg1-v404"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/trip-server-stg1-v016"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/mapcomments-stg-stg1-v011"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/mapcomments-stg-stg1-v013"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpoolactivity-stg-stg1-v187"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/regressionchecker-stg-stg1-v014"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/voiceprompts-stg-stg1-v004"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/adreports-advise-stg-stg1-v005"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpooltesting-stg-stg1-v023"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/elton-stg-stg1-v015"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpoolserver-stg-stg1-v001"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpooltesting-stg-trip-stg1-v003"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/pointsserver-staging-stg1-v008"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/topiceventsdistributor-stg-stg1-v032"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/carpooltesting-stg-scheduler-stg1-v003"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "microservice-stg-internal"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "permits_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/permits-stg-stg1-v011"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "permits-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "piitool_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/piitool-stg-stg1-v029"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "piitool-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "preferences_stg_stg1_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/preferences-stg-stg1-v004"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "preferences-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "prompto_stg_stg1_ibes" {
  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "prompto-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "realtime_proxy_stg_management_stg1" {
  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/realtime-proxy-stg-hc"]
  name                            = "realtime-proxy-stg-management-stg1"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "realtime_stg_stg1_internal" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/realtime-traffic-east-stg-v018"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/realtime-frontend-3-stg-v023"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/realtime-frontend-2-stg-v039"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/realtime-traffic-west-stg-v014"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/realtime-frontend-6-stg-v008"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/realtime-frontend-5-stg-v003"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/realtime-centraldb-stg-v015"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/realtime-frontend-4-stg-v023"
  }

  connection_draining_timeout_sec = "0"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/realtime-stg-internal-hc"]
  name                            = "realtime-stg-stg1-internal"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "redit_test_stg_stg1_ibes" {
  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "redit-test-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "repository_server_stg_stg1_bes" {
  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "repository-server-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "repositoryprocessor_stg_stg1_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/repositoryprocessor-stg-stg1-v002"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/repositoryprocessor-stg-stg1-v001"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "repositoryprocessor-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "row_staging_general12_cassandra_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/instanceGroups/stagingcassandra-stg-c-v001"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/stagingcassandra-stg-d-v001"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/instanceGroups/stagingcassandra-stg-b-v001"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/cassandra-healthcheck"]
  name                            = "row-staging-general12-cassandra-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "row_staging_general21_cassandra_bes" {
  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/cassandra-healthcheck"]
  name                            = "row-staging-general21-cassandra-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "saw_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/saw-stg-stg1-v001"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "saw-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "scheduler_stg_stg1_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/scheduler-stg-stg1-v005"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "scheduler-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "searchserver_stg_stg1_bes" {
  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "searchserver-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "servermanager_stg_stg1_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/servermanager-stg-stg1-v002"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/servermanager-stg-stg1-v005"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "servermanager-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "socialmediaserver_stg_stg1_bes" {
  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "socialmediaserver-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "stg_realtime_ddb_cassandra_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/instanceGroups/realtimeddbcassandra-stg-c-v000"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/realtimeddbcassandra-stg-d-v000"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/instanceGroups/realtimeddbcassandra-stg-b-v000"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/cassandra-healthcheck"]
  name                            = "stg-realtime-ddb-cassandra-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "stg_staging_general21_cassandra_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/staginggeneral21cassandra-stg-d-v001"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/instanceGroups/staginggeneral21cassandra-stg-b-v001"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/instanceGroups/staginggeneral21cassandra-stg-c-v001"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/cassandra-healthcheck"]
  name                            = "stg-staging-general21-cassandra-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "stg_staging_master_cassandra_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/stagingmastercassandra-stg-d-v001"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/instanceGroups/stagingmastercassandra-stg-b-v001"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/instanceGroups/stagingmastercassandra-stg-c-v001"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/cassandra-healthcheck"]
  name                            = "stg-staging-master-cassandra-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "stg_staging_youying_test_cassandra_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/instanceGroups/stagingyouyingtestcassandra-stg-b-v000"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/instanceGroups/stagingyouyingtestcassandra-stg-c-v001"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/stagingyouyingtestcassandra-stg-d-v000"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/cassandra-healthcheck"]
  name                            = "stg-staging-youying-test-cassandra-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "stg_stgregulator_cassandra_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/instanceGroups/stgregulatorcassandra-stg-c-v000"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/instanceGroups/stgregulatorcassandra-stg-b-v000"
  }

  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/instanceGroups/stgregulatorcassandra-stg-d-v000"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/cassandra-healthcheck"]
  name                            = "stg-stgregulator-cassandra-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "stg_stgtopics_events_cassandra_bes" {
  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/cassandra-healthcheck"]
  name                            = "stg-stgtopics-events-cassandra-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "stg_test_cassandra_bes" {
  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/cassandra-healthcheck"]
  name                            = "stg-test-cassandra-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "stg_topics_events_cassandra_bes" {
  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/cassandra-healthcheck"]
  name                            = "stg-topics-events-cassandra-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "storagegateway_stg_stg1_bes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/storagegateway-stg-stg1-v007"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "storagegateway-stg-stg1-bes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "traceutility_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/traceutility-stg-stg1-v004"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "traceutility-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "useractions_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/useractions-stg-stg1-v002"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "useractions-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "userjourney_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/userjourney-stg-stg1-v014"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/microservice-healthcheck"]
  name                            = "userjourney-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "userrouting_stg_stg1_ibes" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/userrouting-stg-stg1-v003"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/http-healthcheck"]
  name                            = "userrouting-stg-stg1-ibes"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}

resource "google_compute_region_backend_service" "usersprofile_stg_internal" {
  backend = {
    group = "https://www.googleapis.com/compute/v1/projects/waze-development/regions/europe-west1/instanceGroups/usersprofile-stg-stg1-v124"
  }

  connection_draining_timeout_sec = "300"
  health_checks                   = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/healthChecks/usersprofile-stg-hc"]
  name                            = "usersprofile-stg-internal"
  project                         = "waze-development"
  protocol                        = "TCP"
  region                          = "europe-west1"
  session_affinity                = "NONE"
  timeout_sec                     = "30"
}
