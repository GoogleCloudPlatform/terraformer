provider "google" {
  project = ""
  region  = ""
}

resource "google_storage_bucket" "artifacts-waze_development-appspot-com" {
  labels        = {}
  location      = "US"
  name          = "artifacts.waze-development.appspot.com"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "baraktest" {
  labels        = {}
  location      = "EU"
  name          = "baraktest"
  project       = "waze-development"
  storage_class = "COLDLINE"
}

resource "google_storage_bucket" "carpoolrouting_dev" {
  labels = {}

  lifecycle_rule = {
    action = {
      type = "Delete"
    }

    condition = {
      age                = "15"
      created_before     = ""
      is_live            = false
      num_newer_versions = "0"
    }
  }

  location      = "US"
  name          = "carpoolrouting-dev"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "cassandra_backup" {
  labels        = {}
  location      = "US"
  name          = "cassandra-backup"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "cassandra_staging" {
  labels        = {}
  location      = "EUROPE-WEST1"
  name          = "cassandra-staging"
  project       = "waze-development"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket" "cdbg_agent_waze_development" {
  labels        = {}
  location      = "US"
  name          = "cdbg-agent_waze-development"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "cdbg_logs_waze_development" {
  labels = {}

  lifecycle_rule = {
    action = {
      type = "Delete"
    }

    condition = {
      age                = "1"
      created_before     = ""
      is_live            = false
      num_newer_versions = "0"
    }
  }

  location      = "US"
  name          = "cdbg-logs_waze-development"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "countrymap" {
  labels        = {}
  location      = "EU"
  name          = "countrymap"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "dataflow_staging_europe_west1_205061584173" {
  labels        = {}
  location      = "EUROPE-WEST1"
  name          = "dataflow-staging-europe-west1-205061584173"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "dataflow_staging_us_central1_205061584173" {
  labels        = {}
  location      = "US-CENTRAL1"
  name          = "dataflow-staging-us-central1-205061584173"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "dataflow_staging_us_east1_205061584173" {
  labels        = {}
  location      = "US-EAST1"
  name          = "dataflow-staging-us-east1-205061584173"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "dataproc_02dfa3c4_6240_402c_b226_c7b4a4670b1d_us" {
  labels        = {}
  location      = "US"
  name          = "dataproc-02dfa3c4-6240-402c-b226-c7b4a4670b1d-us"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "dataproc_1398de85_a65d_406c_bc3a_d9800739cb66_europe_west1" {
  labels        = {}
  location      = "EUROPE-WEST1"
  name          = "dataproc-1398de85-a65d-406c-bc3a-d9800739cb66-europe-west1"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "dataproc_5f430212_e3ab_4900_9d30_23ac5c12858e_europe_west2" {
  labels        = {}
  location      = "EUROPE-WEST2"
  name          = "dataproc-5f430212-e3ab-4900-9d30-23ac5c12858e-europe-west2"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "dataproc_ab851108_ca3d_44cf_95af_175a851569bd_us_east1" {
  labels        = {}
  location      = "US-EAST1"
  name          = "dataproc-ab851108-ca3d-44cf-95af-175a851569bd-us-east1"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "dataproc_ec593c78_412a_4475_8ceb_a4ae4eb547ad_eu" {
  labels        = {}
  location      = "EU"
  name          = "dataproc-ec593c78-412a-4475-8ceb-a4ae4eb547ad-eu"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "deep_learning" {
  labels        = {}
  location      = "EUROPE-WEST1"
  name          = "deep-learning"
  project       = "waze-development"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket" "gcres-waze-com" {
  labels        = {}
  location      = "US"
  name          = "gcres.waze.com"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "genady" {
  labels        = {}
  location      = "US"
  name          = "genady"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "getdate_lb_4" {
  labels        = {}
  location      = "US"
  name          = "getdate-lb-4"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "gsutil_test_test_cp_unwritable_tracker_file_download_b_0701e591" {
  labels        = {}
  location      = "US-CENTRAL1"
  name          = "gsutil-test-test_cp_unwritable_tracker_file_download-b-0701e591"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "gsutil_test_test_cp_unwritable_tracker_file_download_b_d50a8870" {
  labels        = {}
  location      = "US-CENTRAL1"
  name          = "gsutil-test-test_cp_unwritable_tracker_file_download-b-d50a8870"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "gsutil_test_test_minus_d_resumable_upload_bucket_394137ab" {
  labels        = {}
  location      = "US-CENTRAL1"
  name          = "gsutil-test-test_minus_d_resumable_upload-bucket-394137ab"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "hermetic_routing" {
  labels        = {}
  location      = "US-WEST1"
  name          = "hermetic-routing"
  project       = "waze-development"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket" "kubernetes_clusters_sergey_test_cassandra" {
  labels        = {}
  location      = "US"
  name          = "kubernetes-clusters-sergey-test-cassandra"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "merger_data" {
  labels        = {}
  location      = "US"
  name          = "merger-data"
  project       = "waze-development"
  storage_class = "COLDLINE"
}

resource "google_storage_bucket" "merger_data_jupyterhub" {
  labels        = {}
  location      = "US"
  name          = "merger-data-jupyterhub"
  project       = "waze-development"
  storage_class = "COLDLINE"
}

resource "google_storage_bucket" "pg_staging_master" {
  labels        = {}
  location      = "EUROPE-WEST1"
  name          = "pg-staging-master"
  project       = "waze-development"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket" "routing_regression" {
  labels        = {}
  location      = "EU"
  name          = "routing_regression"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "routing_regression_automation" {
  labels        = {}
  location      = "US"
  name          = "routing_regression_automation"
  project       = "waze-development"
  storage_class = "COLDLINE"
}

resource "google_storage_bucket" "row_app_logs" {
  labels        = {}
  location      = "EUROPE-WEST1"
  name          = "row-app-logs"
  project       = "waze-development"
  storage_class = "DURABLE_REDUCED_AVAILABILITY"
}

resource "google_storage_bucket" "row_pg_backup" {
  labels = {}

  lifecycle_rule = {
    action = {
      type = "Delete"
    }

    condition = {
      age                = "5"
      created_before     = ""
      is_live            = false
      num_newer_versions = "0"
    }
  }

  location      = "EU"
  name          = "row-pg-backup"
  project       = "waze-development"
  storage_class = "NEARLINE"
}

resource "google_storage_bucket" "rr-fu-nl" {
  labels        = {}
  location      = "US"
  name          = "rr.fu.nl"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "sagie_phantom_test" {
  labels        = {}
  location      = "US"
  name          = "sagie-phantom-test"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "sergeylanz" {
  labels        = {}
  location      = "EU"
  name          = "sergeylanz"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "spin_ef0fc51f_8165_4d25_87a8_4bb3bc1d13d6" {
  labels        = {}
  location      = "US"
  name          = "spin-ef0fc51f-8165-4d25-87a8-4bb3bc1d13d6"
  project       = "waze-development"
  storage_class = "STANDARD"

  versioning = {
    enabled = true
  }
}

resource "google_storage_bucket" "spin_f2b48897_afed_4cc1_8507_53f7424a80d4" {
  labels        = {}
  location      = "US"
  name          = "spin-f2b48897-afed-4cc1-8507-53f7424a80d4"
  project       = "waze-development"
  storage_class = "STANDARD"

  versioning = {
    enabled = true
  }
}

resource "google_storage_bucket" "spinnaker_waze_development" {
  labels        = {}
  location      = "US"
  name          = "spinnaker-waze-development"
  project       = "waze-development"
  storage_class = "STANDARD"

  versioning = {
    enabled = true
  }
}

resource "google_storage_bucket" "staging-waze_development-appspot-com" {
  labels = {}

  lifecycle_rule = {
    action = {
      type = "Delete"
    }

    condition = {
      age                = "15"
      created_before     = ""
      is_live            = false
      num_newer_versions = "0"
    }
  }

  location      = "US"
  name          = "staging.waze-development.appspot.com"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "stg_am_app_logs" {
  labels        = {}
  location      = "US"
  name          = "stg-am-app-logs"
  project       = "waze-development"
  storage_class = "DURABLE_REDUCED_AVAILABILITY"
}

resource "google_storage_bucket" "stg_cassandra_backup" {
  labels        = {}
  location      = "US"
  name          = "stg-cassandra-backup"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "stg_row_app_logs" {
  labels        = {}
  location      = "EU"
  name          = "stg-row-app-logs"
  project       = "waze-development"
  storage_class = "DURABLE_REDUCED_AVAILABILITY"
}

resource "google_storage_bucket" "stg_scripts_download" {
  labels        = {}
  location      = "US-EAST1"
  name          = "stg-scripts-download"
  project       = "waze-development"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket" "storage_gateway" {
  labels        = {}
  location      = "US"
  name          = "storage_gateway"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "tiles_generator_il" {
  labels        = {}
  location      = "US"
  name          = "tiles-generator-il"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "tiles_generator_stg" {
  labels        = {}
  location      = "US"
  name          = "tiles-generator-stg"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "tilesbuilder_configs" {
  labels        = {}
  location      = "US"
  name          = "tilesbuilder-configs"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "tilesbuilder_stg" {
  labels        = {}
  location      = "US"
  name          = "tilesbuilder-stg"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "topic_test_libs" {
  labels        = {}
  location      = "EU"
  name          = "topic-test-libs"
  project       = "waze-development"
  storage_class = "NEARLINE"
}

resource "google_storage_bucket" "topic_test_results" {
  labels = {}

  lifecycle_rule = {
    action = {
      type = "Delete"
    }

    condition = {
      age                = "90"
      created_before     = ""
      is_live            = false
      num_newer_versions = "0"
    }
  }

  location      = "EU"
  name          = "topic-test-results"
  project       = "waze-development"
  storage_class = "NEARLINE"
}

resource "google_storage_bucket" "us-artifacts-waze_development-appspot-com" {
  labels        = {}
  location      = "US"
  name          = "us.artifacts.waze-development.appspot.com"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "us_central1_tiles_builder_1be8ed5a_bucket" {
  labels = {
    goog-composer-environment = "tiles-builder"
    goog-composer-location    = "us-central1"
    goog-composer-version     = "composer-0-5-3-airflow-1-9-0"
  }

  location      = "US"
  name          = "us-central1-tiles-builder-1be8ed5a-bucket"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "vm_config-waze_development-appspot-com" {
  labels        = {}
  location      = "US"
  name          = "vm-config.waze-development.appspot.com"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "vm_containers-waze_development-appspot-com" {
  labels        = {}
  location      = "US"
  name          = "vm-containers.waze-development.appspot.com"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "waze01_backup" {
  labels        = {}
  location      = "EUROPE-WEST1"
  name          = "waze01_backup"
  project       = "waze-development"
  storage_class = "NEARLINE"
}

resource "google_storage_bucket" "waze_ads_adapt_stg" {
  labels = {}

  lifecycle_rule = {
    action = {
      type = "Delete"
    }

    condition = {
      age                = "2"
      created_before     = ""
      is_live            = false
      num_newer_versions = "0"
    }
  }

  location      = "EUROPE-WEST1"
  name          = "waze-ads-adapt-stg"
  project       = "waze-development"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket" "waze_ads_bi_benchmarks_stg" {
  labels        = {}
  location      = "EUROPE-WEST1"
  name          = "waze-ads-bi-benchmarks-stg"
  project       = "waze-development"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket" "waze_ads_bi_countries_stg" {
  labels        = {}
  location      = "EU"
  name          = "waze-ads-bi-countries-stg"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_ads_elon_test" {
  labels        = {}
  location      = "US"
  name          = "waze-ads-elon-test"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_ads_resources_test" {
  cors = {
    max_age_seconds = "3600"
    method          = ["GET", "OPTIONS"]
    origin          = ["https://waze.com", "https://*.waze.com", "https://waze.co.il", "https://*.waze.co.il", "https://*.gcp.wazestg.com", "http://localhost", "http://*.localhost", "http://localhost:*", "https://*.witools.foo"]
    response_header = ["Content-Type"]
  }

  labels        = {}
  location      = "EU"
  name          = "waze-ads-resources-test"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_ads_rotem_test" {
  labels = {}

  lifecycle_rule = {
    action = {
      type = "Delete"
    }

    condition = {
      age                = "1"
      created_before     = ""
      is_live            = false
      num_newer_versions = "0"
    }
  }

  lifecycle_rule = {
    action = {
      type = "Delete"
    }

    condition = {
      age                = "2"
      created_before     = ""
      is_live            = false
      num_newer_versions = "0"
    }
  }

  lifecycle_rule = {
    action = {
      type = "Delete"
    }

    condition = {
      age                = "0"
      created_before     = "2018-09-30"
      is_live            = false
      num_newer_versions = "0"
    }
  }

  location      = "US"
  name          = "waze-ads-rotem-test"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_artifactory_backup" {
  labels        = {}
  location      = "EU"
  name          = "waze-artifactory-backup"
  project       = "waze-development"
  storage_class = "NEARLINE"
}

resource "google_storage_bucket" "waze_backups" {
  labels        = {}
  location      = "US"
  name          = "waze_backups"
  project       = "waze-development"
  storage_class = "NEARLINE"
}

resource "google_storage_bucket" "waze_carpool_groups_images_stg" {
  labels        = {}
  location      = "EUROPE-WEST1"
  name          = "waze-carpool-groups-images-stg"
  project       = "waze-development"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket" "waze_client_resources_staging" {
  labels        = {}
  location      = "EUROPE-WEST1"
  name          = "waze-client-resources-staging"
  project       = "waze-development"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket" "waze_cloud_functions_stg" {
  labels        = {}
  location      = "EUROPE-WEST1"
  name          = "waze-cloud-functions-stg"
  project       = "waze-development"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket" "waze_dataproc_dev" {
  labels        = {}
  location      = "US"
  name          = "waze-dataproc-dev"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_deployment_staging" {
  labels        = {}
  location      = "EU"
  name          = "waze-deployment-staging"
  project       = "waze-development"
  storage_class = "STANDARD"

  versioning = {
    enabled = true
  }
}

resource "google_storage_bucket" "waze_deployment_staging_sod" {
  labels        = {}
  location      = "US"
  name          = "waze-deployment-staging-sod"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "waze_development-appspot-com" {
  labels        = {}
  location      = "US"
  name          = "waze-development.appspot.com"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "waze_development_cloudbuild" {
  labels        = {}
  location      = "US"
  name          = "waze-development_cloudbuild"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "waze_development_daisy_bkt" {
  labels        = {}
  location      = "US"
  name          = "waze-development-daisy-bkt"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "waze_dexter_resources" {
  labels        = {}
  location      = "US"
  name          = "waze-dexter-resources"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "waze_dirt_dategiver" {
  labels        = {}
  location      = "EU"
  name          = "waze-dirt-dategiver"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_dirt_dategiver1" {
  labels        = {}
  location      = "EU"
  name          = "waze-dirt-dategiver1"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_dirt_gabik" {
  labels        = {}
  location      = "US"
  name          = "waze-dirt-gabik"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_feedback_screenshots_stg" {
  labels        = {}
  location      = "US"
  name          = "waze-feedback-screenshots-stg"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_fluentd_buffer" {
  labels = {}

  lifecycle_rule = {
    action = {
      type = "Delete"
    }

    condition = {
      age                = "1"
      created_before     = ""
      is_live            = false
      num_newer_versions = "0"
    }
  }

  location      = "EUROPE-WEST1"
  name          = "waze-fluentd-buffer"
  project       = "waze-development"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket" "waze_fluentd_buffer_stg" {
  labels = {}

  lifecycle_rule = {
    action = {
      type = "Delete"
    }

    condition = {
      age                = "1"
      created_before     = ""
      is_live            = false
      num_newer_versions = "0"
    }
  }

  location      = "EUROPE-WEST1"
  name          = "waze-fluentd-buffer-stg"
  project       = "waze-development"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket" "waze_images" {
  labels        = {}
  location      = "US"
  name          = "waze_images"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "waze_maint_jobs_stg" {
  labels        = {}
  location      = "EUROPE-WEST1"
  name          = "waze-maint-jobs-stg"
  project       = "waze-development"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket" "waze_merger" {
  labels        = {}
  location      = "EU"
  name          = "waze-merger"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_merger_jupyterhub" {
  labels        = {}
  location      = "EU"
  name          = "waze-merger-jupyterhub"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_nsscache_stg" {
  labels        = {}
  location      = "EU"
  name          = "waze-nsscache-stg"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_piitool_stg" {
  labels        = {}
  location      = "EU"
  name          = "waze-piitool-stg"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_routing_regression" {
  labels        = {}
  location      = "EU"
  name          = "waze-routing-regression"
  project       = "waze-development"
  storage_class = "DURABLE_REDUCED_AVAILABILITY"
}

resource "google_storage_bucket" "waze_routing_tiles" {
  labels        = {}
  location      = "US"
  name          = "waze-routing-tiles"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_tile_baseline" {
  labels        = {}
  location      = "US"
  name          = "waze-tile-baseline"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "waze_tts" {
  labels        = {}
  location      = "EU"
  name          = "waze_tts"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "waze_tts_heb" {
  labels        = {}
  location      = "EU"
  name          = "waze_tts_heb"
  project       = "waze-development"
  storage_class = "STANDARD"
}

resource "google_storage_bucket" "waze_voice" {
  labels        = {}
  location      = "US"
  name          = "waze-voice"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket" "wme_staging_frontend_assets" {
  cors = {
    max_age_seconds = "86400"
    method          = ["GET", "HEAD"]
    origin          = ["https://*.wazestg.com"]
  }

  labels        = {}
  location      = "EUROPE-WEST1"
  name          = "wme-staging-frontend-assets"
  project       = "waze-development"
  storage_class = "REGIONAL"
}

resource "google_storage_bucket" "www-gcp-wazestg-com" {
  labels        = {}
  location      = "EU"
  name          = "www.gcp.wazestg.com"
  project       = "waze-development"
  storage_class = "MULTI_REGIONAL"
}

resource "google_storage_bucket_acl" "artifacts-waze_development-appspot-com" {
  bucket = "artifacts.waze-development.appspot.com"
}

resource "google_storage_bucket_acl" "baraktest" {
  bucket = "baraktest"
}

resource "google_storage_bucket_acl" "carpoolrouting_dev" {
  bucket = "carpoolrouting-dev"
}

resource "google_storage_bucket_acl" "cassandra_backup" {
  bucket = "cassandra-backup"
}

resource "google_storage_bucket_acl" "cassandra_staging" {
  bucket = "cassandra-staging"
}

resource "google_storage_bucket_acl" "cdbg_agent_waze_development" {
  bucket = "cdbg-agent_waze-development"
}

resource "google_storage_bucket_acl" "cdbg_logs_waze_development" {
  bucket = "cdbg-logs_waze-development"
}

resource "google_storage_bucket_acl" "countrymap" {
  bucket = "countrymap"
}

resource "google_storage_bucket_acl" "dataflow_staging_europe_west1_205061584173" {
  bucket = "dataflow-staging-europe-west1-205061584173"
}

resource "google_storage_bucket_acl" "dataflow_staging_us_central1_205061584173" {
  bucket = "dataflow-staging-us-central1-205061584173"
}

resource "google_storage_bucket_acl" "dataflow_staging_us_east1_205061584173" {
  bucket = "dataflow-staging-us-east1-205061584173"
}

resource "google_storage_bucket_acl" "dataproc_02dfa3c4_6240_402c_b226_c7b4a4670b1d_us" {
  bucket = "dataproc-02dfa3c4-6240-402c-b226-c7b4a4670b1d-us"
}

resource "google_storage_bucket_acl" "dataproc_1398de85_a65d_406c_bc3a_d9800739cb66_europe_west1" {
  bucket = "dataproc-1398de85-a65d-406c-bc3a-d9800739cb66-europe-west1"
}

resource "google_storage_bucket_acl" "dataproc_5f430212_e3ab_4900_9d30_23ac5c12858e_europe_west2" {
  bucket = "dataproc-5f430212-e3ab-4900-9d30-23ac5c12858e-europe-west2"
}

resource "google_storage_bucket_acl" "dataproc_ab851108_ca3d_44cf_95af_175a851569bd_us_east1" {
  bucket = "dataproc-ab851108-ca3d-44cf-95af-175a851569bd-us-east1"
}

resource "google_storage_bucket_acl" "dataproc_ec593c78_412a_4475_8ceb_a4ae4eb547ad_eu" {
  bucket = "dataproc-ec593c78-412a-4475-8ceb-a4ae4eb547ad-eu"
}

resource "google_storage_bucket_acl" "deep_learning" {
  bucket = "deep-learning"
}

resource "google_storage_bucket_acl" "gcres-waze-com" {
  bucket = "gcres.waze.com"
}

resource "google_storage_bucket_acl" "genady" {
  bucket = "genady"
}

resource "google_storage_bucket_acl" "getdate_lb_4" {
  bucket = "getdate-lb-4"
}

resource "google_storage_bucket_acl" "gsutil_test_test_cp_unwritable_tracker_file_download_b_0701e591" {
  bucket = "gsutil-test-test_cp_unwritable_tracker_file_download-b-0701e591"
}

resource "google_storage_bucket_acl" "gsutil_test_test_cp_unwritable_tracker_file_download_b_d50a8870" {
  bucket = "gsutil-test-test_cp_unwritable_tracker_file_download-b-d50a8870"
}

resource "google_storage_bucket_acl" "gsutil_test_test_minus_d_resumable_upload_bucket_394137ab" {
  bucket = "gsutil-test-test_minus_d_resumable_upload-bucket-394137ab"
}

resource "google_storage_bucket_acl" "hermetic_routing" {
  bucket = "hermetic-routing"
}

resource "google_storage_bucket_acl" "kubernetes_clusters_sergey_test_cassandra" {
  bucket = "kubernetes-clusters-sergey-test-cassandra"
}

resource "google_storage_bucket_acl" "merger_data" {
  bucket = "merger-data"
}

resource "google_storage_bucket_acl" "merger_data_jupyterhub" {
  bucket = "merger-data-jupyterhub"
}

resource "google_storage_bucket_acl" "pg_staging_master" {
  bucket = "pg-staging-master"
}

resource "google_storage_bucket_acl" "routing_regression" {
  bucket = "routing_regression"
}

resource "google_storage_bucket_acl" "routing_regression_automation" {
  bucket = "routing_regression_automation"
}

resource "google_storage_bucket_acl" "row_app_logs" {
  bucket = "row-app-logs"
}

resource "google_storage_bucket_acl" "row_pg_backup" {
  bucket = "row-pg-backup"
}

resource "google_storage_bucket_acl" "rr-fu-nl" {
  bucket = "rr.fu.nl"
}

resource "google_storage_bucket_acl" "sagie_phantom_test" {
  bucket = "sagie-phantom-test"
}

resource "google_storage_bucket_acl" "sergeylanz" {
  bucket = "sergeylanz"
}

resource "google_storage_bucket_acl" "spin_ef0fc51f_8165_4d25_87a8_4bb3bc1d13d6" {
  bucket = "spin-ef0fc51f-8165-4d25-87a8-4bb3bc1d13d6"
}

resource "google_storage_bucket_acl" "spin_f2b48897_afed_4cc1_8507_53f7424a80d4" {
  bucket = "spin-f2b48897-afed-4cc1-8507-53f7424a80d4"
}

resource "google_storage_bucket_acl" "spinnaker_waze_development" {
  bucket = "spinnaker-waze-development"
}

resource "google_storage_bucket_acl" "staging-waze_development-appspot-com" {
  bucket = "staging.waze-development.appspot.com"
}

resource "google_storage_bucket_acl" "stg_am_app_logs" {
  bucket = "stg-am-app-logs"
}

resource "google_storage_bucket_acl" "stg_cassandra_backup" {
  bucket = "stg-cassandra-backup"
}

resource "google_storage_bucket_acl" "stg_row_app_logs" {
  bucket = "stg-row-app-logs"
}

resource "google_storage_bucket_acl" "stg_scripts_download" {
  bucket = "stg-scripts-download"
}

resource "google_storage_bucket_acl" "storage_gateway" {
  bucket = "storage_gateway"
}

resource "google_storage_bucket_acl" "tiles_generator_il" {
  bucket = "tiles-generator-il"
}

resource "google_storage_bucket_acl" "tiles_generator_stg" {
  bucket = "tiles-generator-stg"
}

resource "google_storage_bucket_acl" "tilesbuilder_configs" {
  bucket = "tilesbuilder-configs"
}

resource "google_storage_bucket_acl" "tilesbuilder_stg" {
  bucket = "tilesbuilder-stg"
}

resource "google_storage_bucket_acl" "topic_test_libs" {
  bucket = "topic-test-libs"
}

resource "google_storage_bucket_acl" "topic_test_results" {
  bucket = "topic-test-results"
}

resource "google_storage_bucket_acl" "us-artifacts-waze_development-appspot-com" {
  bucket = "us.artifacts.waze-development.appspot.com"
}

resource "google_storage_bucket_acl" "us_central1_tiles_builder_1be8ed5a_bucket" {
  bucket = "us-central1-tiles-builder-1be8ed5a-bucket"
}

resource "google_storage_bucket_acl" "vm_config-waze_development-appspot-com" {
  bucket = "vm-config.waze-development.appspot.com"
}

resource "google_storage_bucket_acl" "vm_containers-waze_development-appspot-com" {
  bucket = "vm-containers.waze-development.appspot.com"
}

resource "google_storage_bucket_acl" "waze01_backup" {
  bucket = "waze01_backup"
}

resource "google_storage_bucket_acl" "waze_ads_adapt_stg" {
  bucket = "waze-ads-adapt-stg"
}

resource "google_storage_bucket_acl" "waze_ads_bi_benchmarks_stg" {
  bucket = "waze-ads-bi-benchmarks-stg"
}

resource "google_storage_bucket_acl" "waze_ads_bi_countries_stg" {
  bucket = "waze-ads-bi-countries-stg"
}

resource "google_storage_bucket_acl" "waze_ads_elon_test" {
  bucket = "waze-ads-elon-test"
}

resource "google_storage_bucket_acl" "waze_ads_resources_test" {
  bucket = "waze-ads-resources-test"
}

resource "google_storage_bucket_acl" "waze_ads_rotem_test" {
  bucket = "waze-ads-rotem-test"
}

resource "google_storage_bucket_acl" "waze_artifactory_backup" {
  bucket = "waze-artifactory-backup"
}

resource "google_storage_bucket_acl" "waze_backups" {
  bucket = "waze_backups"
}

resource "google_storage_bucket_acl" "waze_carpool_groups_images_stg" {
  bucket = "waze-carpool-groups-images-stg"
}

resource "google_storage_bucket_acl" "waze_client_resources_staging" {
  bucket = "waze-client-resources-staging"
}

resource "google_storage_bucket_acl" "waze_cloud_functions_stg" {
  bucket = "waze-cloud-functions-stg"
}

resource "google_storage_bucket_acl" "waze_dataproc_dev" {
  bucket = "waze-dataproc-dev"
}

resource "google_storage_bucket_acl" "waze_deployment_staging" {
  bucket = "waze-deployment-staging"
}

resource "google_storage_bucket_acl" "waze_deployment_staging_sod" {
  bucket = "waze-deployment-staging-sod"
}

resource "google_storage_bucket_acl" "waze_development-appspot-com" {
  bucket = "waze-development.appspot.com"
}

resource "google_storage_bucket_acl" "waze_development_cloudbuild" {
  bucket = "waze-development_cloudbuild"
}

resource "google_storage_bucket_acl" "waze_development_daisy_bkt" {
  bucket = "waze-development-daisy-bkt"
}

resource "google_storage_bucket_acl" "waze_dexter_resources" {
  bucket = "waze-dexter-resources"
}

resource "google_storage_bucket_acl" "waze_dirt_dategiver" {
  bucket = "waze-dirt-dategiver"
}

resource "google_storage_bucket_acl" "waze_dirt_dategiver1" {
  bucket = "waze-dirt-dategiver1"
}

resource "google_storage_bucket_acl" "waze_dirt_gabik" {
  bucket = "waze-dirt-gabik"
}

resource "google_storage_bucket_acl" "waze_feedback_screenshots_stg" {
  bucket = "waze-feedback-screenshots-stg"
}

resource "google_storage_bucket_acl" "waze_fluentd_buffer" {
  bucket = "waze-fluentd-buffer"
}

resource "google_storage_bucket_acl" "waze_fluentd_buffer_stg" {
  bucket = "waze-fluentd-buffer-stg"
}

resource "google_storage_bucket_acl" "waze_images" {
  bucket = "waze_images"
}

resource "google_storage_bucket_acl" "waze_maint_jobs_stg" {
  bucket = "waze-maint-jobs-stg"
}

resource "google_storage_bucket_acl" "waze_merger" {
  bucket = "waze-merger"
}

resource "google_storage_bucket_acl" "waze_merger_jupyterhub" {
  bucket = "waze-merger-jupyterhub"
}

resource "google_storage_bucket_acl" "waze_nsscache_stg" {
  bucket = "waze-nsscache-stg"
}

resource "google_storage_bucket_acl" "waze_piitool_stg" {
  bucket = "waze-piitool-stg"
}

resource "google_storage_bucket_acl" "waze_routing_regression" {
  bucket = "waze-routing-regression"
}

resource "google_storage_bucket_acl" "waze_routing_tiles" {
  bucket = "waze-routing-tiles"
}

resource "google_storage_bucket_acl" "waze_tile_baseline" {
  bucket = "waze-tile-baseline"
}

resource "google_storage_bucket_acl" "waze_tts" {
  bucket = "waze_tts"
}

resource "google_storage_bucket_acl" "waze_tts_heb" {
  bucket = "waze_tts_heb"
}

resource "google_storage_bucket_acl" "waze_voice" {
  bucket = "waze-voice"
}

resource "google_storage_bucket_acl" "wme_staging_frontend_assets" {
  bucket = "wme-staging-frontend-assets"
}

resource "google_storage_bucket_acl" "www-gcp-wazestg-com" {
  bucket = "www.gcp.wazestg.com"
}

resource "google_storage_bucket_iam_policy" "artifacts-waze_development-appspot-com" {
  bucket = "artifacts.waze-development.appspot.com"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "serviceAccount:205061584173@cloudbuild.gserviceaccount.com",
        "user:stephenhu@google.com"
      ],
      "role": "roles/storage.admin"
    },
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development",
        "serviceAccount:1014978548522-compute@developer.gserviceaccount.com",
        "serviceAccount:am-maint-job@waze-prod.iam.gserviceaccount.com",
        "serviceAccount:row-maint-job@waze-prod.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "group:waze-web@google.com",
        "serviceAccount:docker-builder-am@waze-development.iam.gserviceaccount.com",
        "serviceAccount:waze-web-ci@waze-development.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    },
    {
      "members": [
        "group:waze-maint-jobs-users@google.com",
        "user:afuqua@google.com"
      ],
      "role": "roles/storage.objectCreator"
    },
    {
      "members": [
        "serviceAccount:1010575561383-compute@developer.gserviceaccount.com",
        "serviceAccount:adman2-stg-ng@waze-development.iam.gserviceaccount.com",
        "serviceAccount:spinnaker-dev@waze-development.iam.gserviceaccount.com",
        "serviceAccount:spinnaker-prod-gke@waze-ci.iam.gserviceaccount.com",
        "serviceAccount:spinnaker-waze-dev@waze-development.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.objectViewer"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "baraktest" {
  bucket = "baraktest"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "carpoolrouting_dev" {
  bucket = "carpoolrouting-dev"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:205061584173@project.gserviceaccount.com",
        "user:asafd@google.com",
        "user:galm@google.com",
        "user:uriagassi@google.com",
        "user:zondi@google.com"
      ],
      "role": "roles/storage.objectAdmin"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "cassandra_backup" {
  bucket = "cassandra-backup"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "cassandra_staging" {
  bucket = "cassandra-staging"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "cdbg_agent_waze_development" {
  bucket = "cdbg-agent_waze-development"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "cdbg_logs_waze_development" {
  bucket = "cdbg-logs_waze-development"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "countrymap" {
  bucket = "countrymap"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.admin"
    },
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development",
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development",
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    },
    {
      "members": [
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyObjectOwner"
    },
    {
      "members": [
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyObjectReader"
    },
    {
      "members": [
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.objectAdmin"
    },
    {
      "members": [
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.objectCreator"
    },
    {
      "members": [
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.objectViewer"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "dataflow_staging_europe_west1_205061584173" {
  bucket = "dataflow-staging-europe-west1-205061584173"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "dataflow_staging_us_central1_205061584173" {
  bucket = "dataflow-staging-us-central1-205061584173"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "dataflow_staging_us_east1_205061584173" {
  bucket = "dataflow-staging-us-east1-205061584173"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "dataproc_02dfa3c4_6240_402c_b226_c7b4a4670b1d_us" {
  bucket = "dataproc-02dfa3c4-6240-402c-b226-c7b4a4670b1d-us"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "dataproc_1398de85_a65d_406c_bc3a_d9800739cb66_europe_west1" {
  bucket = "dataproc-1398de85-a65d-406c-bc3a-d9800739cb66-europe-west1"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "dataproc_5f430212_e3ab_4900_9d30_23ac5c12858e_europe_west2" {
  bucket = "dataproc-5f430212-e3ab-4900-9d30-23ac5c12858e-europe-west2"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "dataproc_ab851108_ca3d_44cf_95af_175a851569bd_us_east1" {
  bucket = "dataproc-ab851108-ca3d-44cf-95af-175a851569bd-us-east1"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "dataproc_ec593c78_412a_4475_8ceb_a4ae4eb547ad_eu" {
  bucket = "dataproc-ec593c78-412a-4475-8ceb-a4ae4eb547ad-eu"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "deep_learning" {
  bucket = "deep-learning"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:storage-transfer-8318072285121504523@partnercontent.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    },
    {
      "members": [
        "user:spapini@google.com"
      ],
      "role": "roles/storage.objectAdmin"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "gcres-waze-com" {
  bucket = "gcres.waze.com"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "genady" {
  bucket = "genady"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development",
        "user:jlewi@google.com"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "user:mariand@google.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "getdate_lb_4" {
  bucket = "getdate-lb-4"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "gsutil_test_test_cp_unwritable_tracker_file_download_b_0701e591" {
  bucket = "gsutil-test-test_cp_unwritable_tracker_file_download-b-0701e591"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "gsutil_test_test_cp_unwritable_tracker_file_download_b_d50a8870" {
  bucket = "gsutil-test-test_cp_unwritable_tracker_file_download-b-d50a8870"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "gsutil_test_test_minus_d_resumable_upload_bucket_394137ab" {
  bucket = "gsutil-test-test_minus_d_resumable_upload-bucket-394137ab"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "hermetic_routing" {
  bucket = "hermetic-routing"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "kubernetes_clusters_sergey_test_cassandra" {
  bucket = "kubernetes-clusters-sergey-test-cassandra"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "merger_data" {
  bucket = "merger-data"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:storage-transfer-8318072285121504523@partnercontent.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    },
    {
      "members": [
        "user:spapini@google.com"
      ],
      "role": "roles/storage.objectAdmin"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "merger_data_jupyterhub" {
  bucket = "merger-data-jupyterhub"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "pg_staging_master" {
  bucket = "pg-staging-master"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development",
        "serviceAccount:docker-manager@waze-ta.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:docker-manager@waze-ta.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    },
    {
      "members": [
        "serviceAccount:docker-manager@waze-ta.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyObjectReader"
    },
    {
      "members": [
        "serviceAccount:default@waze-development.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.objectCreator"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "routing_regression" {
  bucket = "routing_regression"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development",
        "user:eladb@google.com"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "routing_regression_automation" {
  bucket = "routing_regression_automation"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "user:omriz@google.com"
      ],
      "role": "roles/storage.objectAdmin"
    },
    {
      "members": [
        "user:omriz@google.com"
      ],
      "role": "roles/storage.objectCreator"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "row_app_logs" {
  bucket = "row-app-logs"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "row_pg_backup" {
  bucket = "row-pg-backup"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "rr-fu-nl" {
  bucket = "rr.fu.nl"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "allUsers",
        "projectViewer:waze-development",
        "user:gabik@google.com"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "sagie_phantom_test" {
  bucket = "sagie-phantom-test"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "sergeylanz" {
  bucket = "sergeylanz"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development",
        "serviceAccount:ahtztakpkfgcdh5ahla7udwsoy@speckle-umbrella-pg-1.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "spin_ef0fc51f_8165_4d25_87a8_4bb3bc1d13d6" {
  bucket = "spin-ef0fc51f-8165-4d25-87a8-4bb3bc1d13d6"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "spin_f2b48897_afed_4cc1_8507_53f7424a80d4" {
  bucket = "spin-f2b48897-afed-4cc1-8507-53f7424a80d4"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "spinnaker_waze_development" {
  bucket = "spinnaker-waze-development"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "staging-waze_development-appspot-com" {
  bucket = "staging.waze-development.appspot.com"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "group:wme-eng@google.com"
      ],
      "role": "roles/storage.admin"
    },
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "user:amirn@google.com"
      ],
      "role": "roles/storage.objectCreator"
    },
    {
      "members": [
        "user:amirn@google.com"
      ],
      "role": "roles/storage.objectViewer"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "stg_am_app_logs" {
  bucket = "stg-am-app-logs"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "stg_cassandra_backup" {
  bucket = "stg-cassandra-backup"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:row-staging-general21-cassandr@waze-development.iam.gserviceaccount.com",
        "serviceAccount:staging-geocoding-tiles-cass@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-realtime-ddb-cassandra@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-staging-general21-cassandr@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-staging-master-cassandra@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-staging-youying-test-cassa@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-stgregulator-cassandra@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-stgtopics-events-cassandra@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-stgvenues-cassandra@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-test-cassandra@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-topics-events-cassandra@waze-development.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.objectAdmin"
    },
    {
      "members": [
        "serviceAccount:row-staging-general21-cassandr@waze-development.iam.gserviceaccount.com",
        "serviceAccount:staging-geocoding-tiles-cass@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-realtime-ddb-cassandra@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-staging-general21-cassandr@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-staging-master-cassandra@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-staging-youying-test-cassa@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-stgregulator-cassandra@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-stgtopics-events-cassandra@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-stgvenues-cassandra@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-test-cassandra@waze-development.iam.gserviceaccount.com",
        "serviceAccount:stg-topics-events-cassandra@waze-development.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.objectCreator"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "stg_row_app_logs" {
  bucket = "stg-row-app-logs"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development",
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.objectCreator"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "stg_scripts_download" {
  bucket = "stg-scripts-download"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:a7kfr757mreq7funju5w3mq5yi@speckle-umbrella-pg-1.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    },
    {
      "members": [
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.objectViewer"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "storage_gateway" {
  bucket = "storage_gateway"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development",
        "serviceAccount:storage-gateway-service-accoun@waze-development.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "tiles_generator_il" {
  bucket = "tiles-generator-il"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "tiles_generator_stg" {
  bucket = "tiles-generator-stg"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "tilesbuilder_configs" {
  bucket = "tilesbuilder-configs"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "tilesbuilder_stg" {
  bucket = "tilesbuilder-stg"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development",
        "serviceAccount:waze-development@appspot.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development",
        "serviceAccount:service-205061584173@gcf-admin-robot.iam.gserviceaccount.com",
        "serviceAccount:waze-development@appspot.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:service-205061584173@gcf-admin-robot.iam.gserviceaccount.com",
        "serviceAccount:waze-development@appspot.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyObjectReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "topic_test_libs" {
  bucket = "topic-test-libs"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "topic_test_results" {
  bucket = "topic-test-results"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "group:cloud-logs@google.com",
        "projectEditor:waze-development",
        "projectOwner:waze-development",
        "serviceAccount:topic-testing@waze-development.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "us-artifacts-waze_development-appspot-com" {
  bucket = "us.artifacts.waze-development.appspot.com"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "us_central1_tiles_builder_1be8ed5a_bucket" {
  bucket = "us-central1-tiles-builder-1be8ed5a-bucket"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectEditor:ma7bd4d68baf7c6cb-tp",
        "serviceAccount:150521821488@cloudbuild.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "vm_config-waze_development-appspot-com" {
  bucket = "vm-config.waze-development.appspot.com"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "serviceAccount:waze-development@appspot.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:admin-console-hr@appspot.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "vm_containers-waze_development-appspot-com" {
  bucket = "vm-containers.waze-development.appspot.com"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development",
        "serviceAccount:waze-development@appspot.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development",
        "serviceAccount:admin-console-hr@appspot.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze01_backup" {
  bucket = "waze01_backup"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_ads_adapt_stg" {
  bucket = "waze-ads-adapt-stg"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_ads_bi_benchmarks_stg" {
  bucket = "waze-ads-bi-benchmarks-stg"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:storage-transfer-8318072285121504523@partnercontent.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    },
    {
      "members": [
        "serviceAccount:slides-editor@waze-development.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.objectViewer"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_ads_bi_countries_stg" {
  bucket = "waze-ads-bi-countries-stg"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:storage-transfer-8318072285121504523@partnercontent.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    },
    {
      "members": [
        "serviceAccount:slides-editor@waze-development.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.objectViewer"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_ads_elon_test" {
  bucket = "waze-ads-elon-test"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_ads_resources_test" {
  bucket = "waze-ads-resources-test"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:waze-web-stg@appspot.gserviceaccount.com",
        "user:ashram@google.com",
        "user:micdavis@google.com",
        "user:tsion@google.com"
      ],
      "role": "roles/storage.objectAdmin"
    },
    {
      "members": [
        "serviceAccount:1014978548522-compute@developer.gserviceaccount.com",
        "serviceAccount:waze-web-stg@appspot.gserviceaccount.com",
        "user:ashram@google.com",
        "user:levos@google.com",
        "user:micdavis@google.com",
        "user:saarc@google.com",
        "user:shatsinna@google.com",
        "user:tsion@google.com"
      ],
      "role": "roles/storage.objectCreator"
    },
    {
      "members": [
        "serviceAccount:waze-web-stg@appspot.gserviceaccount.com",
        "user:ashram@google.com",
        "user:micdavis@google.com"
      ],
      "role": "roles/storage.objectViewer"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_ads_rotem_test" {
  bucket = "waze-ads-rotem-test"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_artifactory_backup" {
  bucket = "waze-artifactory-backup"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_backups" {
  bucket = "waze_backups"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_carpool_groups_images_stg" {
  bucket = "waze-carpool-groups-images-stg"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_client_resources_staging" {
  bucket = "waze-client-resources-staging"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_cloud_functions_stg" {
  bucket = "waze-cloud-functions-stg"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_dataproc_dev" {
  bucket = "waze-dataproc-dev"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_deployment_staging" {
  bucket = "waze-deployment-staging"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development",
        "serviceAccount:334348044266@cloudbuild.gserviceaccount.com",
        "serviceAccount:docker-manager@waze-ta.iam.gserviceaccount.com",
        "serviceAccount:storage-transfer-8318072285121504523@partnercontent.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:1014978548522-k8tp1fjng8v367utgbg73d8m6g14a2km@developer.gserviceaccount.com",
        "serviceAccount:deployment-console@waze-development.iam.gserviceaccount.com",
        "serviceAccount:elton-261@waze-prod.iam.gserviceaccount.com",
        "serviceAccount:tts-abb-storage@waze-prod.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    },
    {
      "members": [
        "serviceAccount:334348044266@cloudbuild.gserviceaccount.com",
        "serviceAccount:docker-manager@waze-ta.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyObjectReader"
    },
    {
      "members": [
        "serviceAccount:elton-261@waze-prod.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.objectCreator"
    },
    {
      "members": [
        "serviceAccount:elton-261@waze-prod.iam.gserviceaccount.com",
        "serviceAccount:storage-transfer-8318072285121504523@partnercontent.gserviceaccount.com"
      ],
      "role": "roles/storage.objectViewer"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_deployment_staging_sod" {
  bucket = "waze-deployment-staging-sod"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development",
        "serviceAccount:storage-transfer-8318072285121504523@partnercontent.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:storage-transfer-8318072285121504523@partnercontent.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    },
    {
      "members": [
        "serviceAccount:storage-transfer-8318072285121504523@partnercontent.gserviceaccount.com"
      ],
      "role": "roles/storage.objectAdmin"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_development-appspot-com" {
  bucket = "waze-development.appspot.com"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_development_cloudbuild" {
  bucket = "waze-development_cloudbuild"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "serviceAccount:205061584173@cloudbuild.gserviceaccount.com"
      ],
      "role": "roles/storage.admin"
    },
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "group:waze-maint-jobs-users@google.com"
      ],
      "role": "roles/storage.objectCreator"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_development_daisy_bkt" {
  bucket = "waze-development-daisy-bkt"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "user:chriswilkes@google.com"
      ],
      "role": "roles/storage.objectViewer"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_dexter_resources" {
  bucket = "waze-dexter-resources"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:1014978548522-compute@developer.gserviceaccount.com"
      ],
      "role": "roles/storage.objectViewer"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_dirt_dategiver" {
  bucket = "waze-dirt-dategiver"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_dirt_dategiver1" {
  bucket = "waze-dirt-dategiver1"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_dirt_gabik" {
  bucket = "waze-dirt-gabik"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_feedback_screenshots_stg" {
  bucket = "waze-feedback-screenshots-stg"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "group:waze-feedback-screenshots@google.com",
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_fluentd_buffer" {
  bucket = "waze-fluentd-buffer"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "serviceAccount:gcf-dc75b37e37be1876@appspot.gserviceaccount.com"
      ],
      "role": "roles/storage.admin"
    },
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:waze-development@appspot.gserviceaccount.com"
      ],
      "role": "roles/storage.objectViewer"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_fluentd_buffer_stg" {
  bucket = "waze-fluentd-buffer-stg"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_images" {
  bucket = "waze_images"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_maint_jobs_stg" {
  bucket = "waze-maint-jobs-stg"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_merger" {
  bucket = "waze-merger"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_merger_jupyterhub" {
  bucket = "waze-merger-jupyterhub"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_nsscache_stg" {
  bucket = "waze-nsscache-stg"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_piitool_stg" {
  bucket = "waze-piitool-stg"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "user:boazg@google.com",
        "user:levydaniel@google.com"
      ],
      "role": "roles/storage.admin"
    },
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:pii-tool-test@waze-development.iam.gserviceaccount.com",
        "serviceAccount:waze-web-stg@appspot.gserviceaccount.com"
      ],
      "role": "roles/storage.objectAdmin"
    },
    {
      "members": [
        "serviceAccount:pii-tool-test@waze-development.iam.gserviceaccount.com"
      ],
      "role": "roles/storage.objectCreator"
    },
    {
      "members": [
        "serviceAccount:waze-web-stg@appspot.gserviceaccount.com",
        "user:boazg@google.com",
        "user:levydaniel@google.com",
        "user:waze-web-eng@prod.google.com"
      ],
      "role": "roles/storage.objectViewer"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_routing_regression" {
  bucket = "waze-routing-regression"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "user:eladb@google.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_routing_tiles" {
  bucket = "waze-routing-tiles"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_tile_baseline" {
  bucket = "waze-tile-baseline"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_tts" {
  bucket = "waze_tts"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.legacyBucketWriter"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_tts_heb" {
  bucket = "waze_tts_heb"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "waze_voice" {
  bucket = "waze-voice"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "serviceAccount:205061584173@project.gserviceaccount.com"
      ],
      "role": "roles/storage.objectAdmin"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "wme_staging_frontend_assets" {
  bucket = "wme-staging-frontend-assets"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "group:wme-eng@google.com",
        "serviceAccount:waze-web-ci@waze-development.iam.gserviceaccount.com",
        "user:amirn@google.com",
        "user:avia@google.com"
      ],
      "role": "roles/storage.admin"
    },
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "www-gcp-wazestg-com" {
  bucket = "www.gcp.wazestg.com"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:waze-development",
        "projectOwner:waze-development"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:waze-development"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}
