provider "google" {
  project = ""
  region  = ""
}

resource "google_sql_database" "ads_main1_stg_eu_israel_ads" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "ads-main1-stg-eu"
  name      = "israel_ads"
  project   = "waze-development"
}

resource "google_sql_database" "ads_main1_stg_eu_postgres" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "ads-main1-stg-eu"
  name      = "postgres"
  project   = "waze-development"
}

resource "google_sql_database" "adsdb_stg_il_events" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "adsdb-stg"
  name      = "il_events"
  project   = "waze-development"
}

resource "google_sql_database" "adsdb_stg_israel_ads" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "adsdb-stg"
  name      = "israel_ads"
  project   = "waze-development"
}

resource "google_sql_database" "adsdb_stg_na_ads" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "adsdb-stg"
  name      = "na_ads"
  project   = "waze-development"
}

resource "google_sql_database" "adsdb_stg_na_events" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "adsdb-stg"
  name      = "na_events"
  project   = "waze-development"
}

resource "google_sql_database" "adsdb_stg_postgres" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "adsdb-stg"
  name      = "postgres"
  project   = "waze-development"
}

resource "google_sql_database" "adsdb_stg_row_ads" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "adsdb-stg"
  name      = "row_ads"
  project   = "waze-development"
}

resource "google_sql_database" "adsdb_stg_row_events" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "adsdb-stg"
  name      = "row_events"
  project   = "waze-development"
}

resource "google_sql_database" "datasetmanager1_stg_datasets" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "datasetmanager1-stg"
  name      = "datasets"
  project   = "waze-development"
}

resource "google_sql_database" "datasetmanager1_stg_postgres" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "datasetmanager1-stg"
  name      = "postgres"
  project   = "waze-development"
}

resource "google_sql_database" "editor_test_israel" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "editor-test"
  name      = "israel"
  project   = "waze-development"
}

resource "google_sql_database" "editor_test_israel_old_stg" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "editor-test"
  name      = "israel_old_stg"
  project   = "waze-development"
}

resource "google_sql_database" "editor_test_postgres" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "editor-test"
  name      = "postgres"
  project   = "waze-development"
}

resource "google_sql_database" "live_db_test1_linqmap" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "live-db-test1"
  name      = "linqmap"
  project   = "waze-development"
}

resource "google_sql_database" "live_db_test1_postgres" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "live-db-test1"
  name      = "postgres"
  project   = "waze-development"
}

resource "google_sql_database" "postgres_stg4_il_events" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "postgres-stg4"
  name      = "il_events"
  project   = "waze-development"
}

resource "google_sql_database" "postgres_stg4_israel" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "postgres-stg4"
  name      = "israel"
  project   = "waze-development"
}

resource "google_sql_database" "postgres_stg4_israel_ads" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "postgres-stg4"
  name      = "israel_ads"
  project   = "waze-development"
}

resource "google_sql_database" "postgres_stg4_merger" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "postgres-stg4"
  name      = "merger"
  project   = "waze-development"
}

resource "google_sql_database" "postgres_stg4_postgres" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "postgres-stg4"
  name      = "postgres"
  project   = "waze-development"
}

resource "google_sql_database" "postgres_stg4_users" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "postgres-stg4"
  name      = "users"
  project   = "waze-development"
}

resource "google_sql_database" "row_dev_countrymap2_postgres" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "row-dev-countrymap2"
  name      = "postgres"
  project   = "waze-development"
}

resource "google_sql_database" "sonarqube_postgres" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "sonarqube"
  name      = "postgres"
  project   = "waze-development"
}

resource "google_sql_database" "sonarqube_sonarqube" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "sonarqube"
  name      = "sonarqube"
  project   = "waze-development"
}

resource "google_sql_database" "stg_users_postgres" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "stg-users"
  name      = "postgres"
  project   = "waze-development"
}

resource "google_sql_database" "stg_users_users" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = "stg-users"
  name      = "users"
  project   = "waze-development"
}

resource "google_sql_database" "topicstats_postgres" {
  charset   = "SQL_ASCII"
  collation = "C"
  instance  = "topicstats"
  name      = "postgres"
  project   = "waze-development"
}

resource "google_sql_database" "topicstats_topics" {
  charset   = "SQL_ASCII"
  collation = "C"
  instance  = "topicstats"
  name      = "topics"
  project   = "waze-development"
}

resource "google_sql_database_instance" "ads_main1_stg_eu" {
  database_version = "POSTGRES_9_6"
  ip_address       = []
  name             = "ads-main1-stg-eu"
  project          = "waze-development"
  region           = "europe-west1"

  settings = {
    activation_policy = "ALWAYS"

    backup_configuration = {
      binary_log_enabled = false
      enabled            = true
      start_time         = "18:00"
    }

    crash_safe_replication = false
    disk_autoresize        = true
    disk_size              = "200"
    disk_type              = "PD_SSD"

    ip_configuration = {
      ipv4_enabled = true
      require_ssl  = false
    }

    maintenance_window = {
      day  = "0"
      hour = "0"
    }

    pricing_plan     = "PER_USE"
    replication_type = "SYNCHRONOUS"
    tier             = "db-custom-4-16384"
    user_labels      = {}
  }
}

resource "google_sql_database_instance" "adsdb_stg" {
  database_version = "POSTGRES_9_6"
  ip_address       = []
  name             = "adsdb-stg"
  project          = "waze-development"
  region           = "europe-west1"

  settings = {
    activation_policy = "ALWAYS"
    availability_type = "ZONAL"

    backup_configuration = {
      binary_log_enabled = false
      enabled            = true
      start_time         = "11:00"
    }

    crash_safe_replication = false
    disk_autoresize        = true
    disk_size              = "500"
    disk_type              = "PD_SSD"

    ip_configuration = {
      authorized_networks = {
        name  = "bastion1"
        value = "54.195.80.116/32"
      }

      authorized_networks = {
        name  = "office"
        value = "104.132.36.76/32"
      }

      authorized_networks = {
        name  = "bastion2"
        value = "104.132.36.64/27"
      }

      ipv4_enabled = true
      require_ssl  = false
    }

    location_preference = []

    maintenance_window = {
      day  = "0"
      hour = "0"
    }

    pricing_plan     = "PER_USE"
    replication_type = "SYNCHRONOUS"
    tier             = "db-custom-1-3840"
    user_labels      = {}
  }
}

resource "google_sql_database_instance" "datasetmanager1_stg" {
  database_version = "POSTGRES_9_6"
  ip_address       = []
  name             = "datasetmanager1-stg"
  project          = "waze-development"
  region           = "europe-west1"

  settings = {
    activation_policy = "ALWAYS"
    availability_type = "REGIONAL"

    backup_configuration = {
      binary_log_enabled = false
      enabled            = true
      start_time         = "16:00"
    }

    crash_safe_replication = false
    disk_autoresize        = true
    disk_size              = "10"
    disk_type              = "PD_SSD"

    ip_configuration = {
      ipv4_enabled = true
      require_ssl  = false
    }

    location_preference = []

    maintenance_window = {
      day  = "0"
      hour = "0"
    }

    pricing_plan     = "PER_USE"
    replication_type = "SYNCHRONOUS"
    tier             = "db-custom-2-7680"
    user_labels      = {}
  }
}

resource "google_sql_database_instance" "editor_test" {
  database_version = "POSTGRES_9_6"
  ip_address       = []
  name             = "editor-test"
  project          = "waze-development"
  region           = "europe-west1"

  settings = {
    activation_policy = "ALWAYS"
    availability_type = "ZONAL"

    backup_configuration = {
      binary_log_enabled = false
      enabled            = true
      start_time         = "04:00"
    }

    crash_safe_replication = false
    disk_autoresize        = true
    disk_size              = "49"
    disk_type              = "PD_SSD"

    ip_configuration = {
      authorized_networks = {
        name  = "office"
        value = "31.154.8.70/32"
      }

      authorized_networks = {
        name  = "mapnik"
        value = "35.195.62.248/32"
      }

      authorized_networks = {
        name  = "row-dev-editor"
        value = "35.205.128.22/32"
      }

      authorized_networks = {
        name  = "live-stg"
        value = "23.251.141.161/32"
      }

      ipv4_enabled = true
      require_ssl  = true
    }

    location_preference = {
      zone = "europe-west1-b"
    }

    maintenance_window = {
      day  = "0"
      hour = "0"
    }

    pricing_plan     = "PER_USE"
    replication_type = "SYNCHRONOUS"
    tier             = "db-custom-1-3840"
    user_labels      = {}
  }
}

resource "google_sql_database_instance" "live_db_test1" {
  database_version = "POSTGRES_9_6"
  ip_address       = []
  name             = "live-db-test1"
  project          = "waze-development"
  region           = "europe-west1"

  settings = {
    activation_policy = "ALWAYS"
    availability_type = "REGIONAL"

    backup_configuration = {
      binary_log_enabled = false
      enabled            = true
      start_time         = "14:00"
    }

    crash_safe_replication = false
    disk_autoresize        = true
    disk_size              = "700"
    disk_type              = "PD_SSD"

    ip_configuration = {
      authorized_networks = {
        name  = "office"
        value = "104.132.36.64/27"
      }

      authorized_networks = {
        name  = "live-db-test-instance"
        value = "146.148.3.30/32"
      }

      authorized_networks = {
        name  = "mapnik"
        value = "35.195.62.248/32"
      }

      authorized_networks = {
        name  = "row-dev-editor"
        value = "35.205.128.22/32"
      }

      authorized_networks = {
        name  = "live-stg"
        value = "23.251.141.161/32"
      }

      ipv4_enabled = true
      require_ssl  = true
    }

    location_preference = []

    maintenance_window = {
      day  = "0"
      hour = "0"
    }

    pricing_plan     = "PER_USE"
    replication_type = "SYNCHRONOUS"
    tier             = "db-custom-16-106496"
    user_labels      = {}
  }
}

resource "google_sql_database_instance" "postgres_stg4" {
  database_version = "POSTGRES_9_6"
  ip_address       = []
  name             = "postgres-stg4"
  project          = "waze-development"
  region           = "europe-west1"

  settings = {
    activation_policy = "ALWAYS"
    availability_type = "ZONAL"

    backup_configuration = {
      binary_log_enabled = false
      enabled            = true
      start_time         = "17:00"
    }

    crash_safe_replication = false
    disk_autoresize        = true
    disk_size              = "200"
    disk_type              = "PD_SSD"

    ip_configuration = {
      authorized_networks = {
        name  = "tile-builder-ext"
        value = "35.187.56.219/32"
      }

      authorized_networks = {
        name  = "office"
        value = "31.154.8.68/30"
      }

      ipv4_enabled = true
      require_ssl  = true
    }

    location_preference = []
    pricing_plan        = "PER_USE"
    replication_type    = "SYNCHRONOUS"
    tier                = "db-custom-4-15360"
    user_labels         = {}
  }
}

resource "google_sql_database_instance" "row_dev_countrymap2" {
  database_version = "POSTGRES_9_6"
  ip_address       = []
  name             = "row-dev-countrymap2"
  project          = "waze-development"
  region           = "europe-west1"

  settings = {
    activation_policy = "ALWAYS"
    availability_type = "ZONAL"

    backup_configuration = {
      binary_log_enabled = false
      enabled            = true
      start_time         = "17:00"
    }

    crash_safe_replication = false

    database_flags = {
      name  = "temp_file_limit"
      value = "2147483647"
    }

    disk_autoresize = true
    disk_size       = "10"
    disk_type       = "PD_SSD"

    ip_configuration = {
      ipv4_enabled = true
      require_ssl  = false
    }

    location_preference = {
      zone = "europe-west1-b"
    }

    maintenance_window = {
      day  = "0"
      hour = "0"
    }

    pricing_plan     = "PER_USE"
    replication_type = "SYNCHRONOUS"
    tier             = "db-custom-8-30720"
    user_labels      = {}
  }
}

resource "google_sql_database_instance" "sonarqube" {
  database_version = "POSTGRES_9_6"
  ip_address       = []
  name             = "sonarqube"
  project          = "waze-development"
  region           = "europe-west4"

  settings = {
    activation_policy = "ALWAYS"
    availability_type = "REGIONAL"

    backup_configuration = {
      binary_log_enabled = false
      enabled            = true
      start_time         = "22:00"
    }

    crash_safe_replication = false
    disk_autoresize        = true
    disk_size              = "20"
    disk_type              = "PD_SSD"

    ip_configuration = {
      ipv4_enabled = true
      require_ssl  = false
    }

    location_preference = {
      zone = "europe-west4-c"
    }

    maintenance_window = {
      day  = "5"
      hour = "21"
    }

    pricing_plan     = "PER_USE"
    replication_type = "SYNCHRONOUS"
    tier             = "db-custom-1-3840"
    user_labels      = {}
  }
}

resource "google_sql_database_instance" "stg_users" {
  database_version = "POSTGRES_9_6"
  ip_address       = []
  name             = "stg-users"
  project          = "waze-development"
  region           = "europe-west1"

  settings = {
    activation_policy = "ALWAYS"
    availability_type = "ZONAL"

    backup_configuration = {
      binary_log_enabled = false
      enabled            = true
      start_time         = "02:00"
    }

    crash_safe_replication = false
    disk_autoresize        = true
    disk_size              = "10"
    disk_type              = "PD_SSD"

    ip_configuration = {
      authorized_networks = {
        name  = "points"
        value = "10.240.0.0/24"
      }

      authorized_networks = {
        name  = "row-staging1"
        value = "23.251.135.0/24"
      }

      authorized_networks = {
        name  = "descartes-stg"
        value = "35.187.179.0/24"
      }

      authorized_networks = {
        name  = "points2"
        value = "35.195.148.0/24"
      }

      ipv4_enabled = true
      require_ssl  = true
    }

    location_preference = []

    maintenance_window = {
      day  = "0"
      hour = "0"
    }

    pricing_plan     = "PER_USE"
    replication_type = "SYNCHRONOUS"
    tier             = "db-custom-1-3840"
    user_labels      = {}
  }
}

resource "google_sql_database_instance" "topicstats" {
  database_version = "POSTGRES_9_6"
  ip_address       = []
  name             = "topicstats"
  project          = "waze-development"
  region           = "europe-west1"

  settings = {
    activation_policy = "ALWAYS"

    backup_configuration = {
      binary_log_enabled = false
      enabled            = false
      start_time         = "03:00"
    }

    crash_safe_replication = false
    disk_autoresize        = false
    disk_size              = "500"
    disk_type              = "PD_SSD"

    ip_configuration = {
      authorized_networks = {
        name  = "waze"
        value = "10.240.0.0/16"
      }

      authorized_networks = {
        name  = "monitoring"
        value = "104.155.91.136/32"
      }

      ipv4_enabled = true
      require_ssl  = false
    }

    maintenance_window = {
      day          = "0"
      hour         = "0"
      update_track = "stable"
    }

    pricing_plan     = "PER_USE"
    replication_type = "SYNCHRONOUS"
    tier             = "db-custom-6-20480"
    user_labels      = {}
  }
}
