provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_image" "abtests_all_20161129220131_nginx_dev" {
  description = "abtests-all-20161129220131-nginx-dev"
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "abtests-all-20161129220131-nginx-dev"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/packer-583dfabd-f197-55a3-a5e2-e73a4bc505d0"
}

resource "google_compute_image" "artifactory_image" {
  description = "An image of artifactory"
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-9-stretch"]
  name        = "artifactory-image"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/us-east1-c/disks/artifactory"
}

resource "google_compute_image" "cassandra1_data_20180508" {
  family      = "cassandra1-data"
  labels      = {}
  name        = "cassandra1-data-20180508"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/stagingmastercassandra-stg-b-v001-gt2c-1"
}

resource "google_compute_image" "cassandra1_data_20181018" {
  family      = "cassandra1-data"
  labels      = {}
  name        = "cassandra1-data-20181018"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/cassandra1-data-20181018"
}

resource "google_compute_image" "cassandra1_data_20181021" {
  family      = "cassandra1-data"
  labels      = {}
  name        = "cassandra1-data-20181021"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/cassandra-master-sod-data"
}

resource "google_compute_image" "cassandra1_data_20181021_new" {
  family      = "cassandra1-data"
  labels      = {}
  name        = "cassandra1-data-20181021-new"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/cassandra-master-sod-data"
}

resource "google_compute_image" "cassandra1_data_20181021_new2" {
  family      = "cassandra1-data"
  labels      = {}
  name        = "cassandra1-data-20181021-new2"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/cassandra-master-sod-data"
}

resource "google_compute_image" "cassandra1_data_20181022" {
  family      = "cassandra1-data"
  labels      = {}
  name        = "cassandra1-data-20181022"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/cassandra-master-sod-data"
}

resource "google_compute_image" "cassandra1_data_20181024" {
  family      = "cassandra1-data"
  labels      = {}
  name        = "cassandra1-data-20181024"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/cassandra-master-sod-data"
}

resource "google_compute_image" "cassandra1_data_20181130" {
  family      = "cassandra1-data"
  labels      = {}
  name        = "cassandra1-data-20181130"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/cassandra-master-sod-data"
}

resource "google_compute_image" "cassandra2_data_20180509" {
  family      = "cassandra2-data"
  labels      = {}
  name        = "cassandra2-data-20180509"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/stagingmastercassandra-stg-c-v001-gk74-1"
}

resource "google_compute_image" "cassandra2_data_20181024" {
  family      = "cassandra2-data"
  labels      = {}
  name        = "cassandra2-data-20181024"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/staging-master-cassandra2-data"
}

resource "google_compute_image" "cassandra2_data_20181130" {
  family      = "cassandra2-data"
  labels      = {}
  name        = "cassandra2-data-20181130"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/staging-master-cassandra2-data"
}

resource "google_compute_image" "cassandra3_data_20180509" {
  family      = "cassandra3-data"
  labels      = {}
  name        = "cassandra3-data-20180509"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/stagingmastercassandra-stg-d-v001-fnn3-1"
}

resource "google_compute_image" "cassandra3_data_20181024" {
  family      = "cassandra3-data"
  labels      = {}
  name        = "cassandra3-data-20181024"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/staging-master-cassandra3-data"
}

resource "google_compute_image" "cassandra3_data_20181130" {
  family      = "cassandra3-data"
  labels      = {}
  name        = "cassandra3-data-20181130"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/staging-master-cassandra3-data"
}

resource "google_compute_image" "empty_disk_image" {
  name    = "empty-disk-image"
  project = "waze-development"
}

resource "google_compute_image" "empty_disk_image_500gb" {
  labels      = {}
  name        = "empty-disk-image-500gb"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/us-east1-b/disks/emptydisk"
}

resource "google_compute_image" "engagement_all_20161129214252_nginx_dev" {
  description = "engagement-all-20161129214252-nginx-dev"
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "engagement-all-20161129214252-nginx-dev"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/packer-583df65e-ea41-d978-d6e8-417a0c915120"
}

resource "google_compute_image" "engagement_all_20161129214930_nginx_dev" {
  description = "engagement-all-20161129214930-nginx-dev"
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "engagement-all-20161129214930-nginx-dev"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/packer-583df7ec-ac46-6082-801e-1ff6c2fa34e4"
}

resource "google_compute_image" "engagement_all_20161129215543_nginx_dev" {
  description = "engagement-all-20161129215543-nginx-dev"
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "engagement-all-20161129215543-nginx-dev"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/packer-583df961-8d1e-fb41-29a5-39fc8dd9cec7"
}

resource "google_compute_image" "hermetic_routing_na_v1" {
  description = "An hermetic routing image for GCE p3rf to run"
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "hermetic-routing-na-v1"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/routing-standalone-na-nirta"
}

resource "google_compute_image" "hermetic_routing_na_v2" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "hermetic-routing-na-v2"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/routing-standalone-na-nirta"
}

resource "google_compute_image" "isolated_routing_test" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "isolated-routing-test"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-staging-routing-regression-201709151312"
}

resource "google_compute_image" "loghost_image_v1" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "loghost-image-v1"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/loghost-deleteme1"
}

resource "google_compute_image" "managed_memcache_v1" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "managed-memcache-v1"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/managed-memcache-base-image"
}

resource "google_compute_image" "managed_memcache_v2" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "managed-memcache-v2"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/managed-memcache-base-image-stg"
}

resource "google_compute_image" "managed_memcache_v3" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "managed-memcache-v3"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-managed-memcache-base-image"
}

resource "google_compute_image" "na_benchmark_dev_v02" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "na-benchmark-dev-v02"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/us-east1-b/disks/na-benchmark-dev-v01"
}

resource "google_compute_image" "na_benchmark_dev_v03" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "na-benchmark-dev-v03"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/us-central1-f/disks/na-benchmark-dev-v02-template-3"
}

resource "google_compute_image" "pg_master_15_root" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1204-precise"]
  name        = "pg-master-15-root"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-staging-db-live1-recover-boot"
}

resource "google_compute_image" "realtime_20171026_1208" {
  description = "realtime-20171026-1208"

  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "realtime-20171026-1208"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/packer-59f1d037-58b0-10e1-68b4-44db9a7ca767"
}

resource "google_compute_image" "realtime_20171030_1516" {
  description = "realtime-20171030-1516"

  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "realtime-20171030-1516"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/packer-59f74250-466e-6252-521c-f37486554d90"
}

resource "google_compute_image" "routing_standalone_na_nirta_v1" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "routing-standalone-na-nirta-v1"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/routing-standalone-na-nirta"
}

resource "google_compute_image" "row_autoscale_1404_base_image_1404_20180729_0144" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-1404-base-image-1404-20180729-0144"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-1404-base-image-1404-d3d0e3a0"
}

resource "google_compute_image" "row_autoscale_1404_base_image_1404_20180805_0144" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-1404-base-image-1404-20180805-0144"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-1404-base-image-1404-f3fa729f"
}

resource "google_compute_image" "row_autoscale_1604_base_image_1604_20181230_0121" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-1604-base-image-1604-20181230-0121"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-1604-base-image-1604-a37cf8bc"
}

resource "google_compute_image" "row_autoscale_1604_base_image_1604_20190106_0121" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-1604-base-image-1604-20190106-0121"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-1604-base-image-1604-4df70ad9"
}

resource "google_compute_image" "row_autoscale_base_image_20160922_1757" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-base-image-20160922-1757"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-base-image-310f0b27"
}

resource "google_compute_image" "row_autoscale_base_image_20160922_2210" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-base-image-20160922-2210"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-base-image-2e88bc81"
}

resource "google_compute_image" "row_autoscale_cassandra_2_1_19_base_image_1604_20171227_0633" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-cassandra-2-1-19-base-image-1604-20171227-0633"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-cassandra-2-1-19-base-image-1604-571956fa"
}

resource "google_compute_image" "row_autoscale_cassandra_2_1_base_image_1604_20180226_0743" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-cassandra-2-1-base-image-1604-20180226-0743"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-cassandra-2-1-base-image-1604-015e6d7b"
}

resource "google_compute_image" "row_autoscale_cassandra_2_1_base_image_1604_20180226_1553" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-cassandra-2-1-base-image-1604-20180226-1553"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-cassandra-2-1-base-image-1604-0acff1f9"
}

resource "google_compute_image" "row_autoscale_cassandra_2_1_base_image_1604_20180226_1702" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-cassandra-2-1-base-image-1604-20180226-1702"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-cassandra-2-1-base-image-1604-76b1399d"
}

resource "google_compute_image" "row_autoscale_cassandra_2_1_base_image_1604_20180903_1332" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-cassandra-2-1-base-image-1604-20180903-1332"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-cassandra-2-1-base-image-1604-e7e0cb9d"
}

resource "google_compute_image" "row_autoscale_cassandra_2_1_base_image_1604_20180903_1356" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-cassandra-2-1-base-image-1604-20180903-1356"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-cassandra-2-1-base-image-1604-6c618200"
}

resource "google_compute_image" "row_autoscale_cassandra_2_1_base_image_1604_20180903_1458" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-cassandra-2-1-base-image-1604-20180903-1458"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-cassandra-2-1-base-image-1604-f067cb82"
}

resource "google_compute_image" "row_autoscale_clean_base_image_1404_20170209_0809" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-clean-base-image-1404-20170209-0809"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-clean-base-image-1404-cdb0e641"
}

resource "google_compute_image" "row_autoscale_clean_base_image_1404_20180729_0211" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-clean-base-image-1404-20180729-0211"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-clean-base-image-1404-39d7c7b2"
}

resource "google_compute_image" "row_autoscale_clean_base_image_1404_20180805_0211" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-clean-base-image-1404-20180805-0211"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-clean-base-image-1404-01f7a004"
}

resource "google_compute_image" "row_autoscale_clean_base_image_1604_20171112_0019" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-clean-base-image-1604-20171112-0019"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-clean-base-image-1604-977171c6"
}

resource "google_compute_image" "row_autoscale_clean_base_image_1604_20180304_0152" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-clean-base-image-1604-20180304-0152"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-clean-base-image-1604-b1b30dc7"
}

resource "google_compute_image" "row_autoscale_clean_base_image_1604_20181021_0136" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-clean-base-image-1604-20181021-0136"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-clean-base-image-1604-0d9b1dac"
}

resource "google_compute_image" "row_autoscale_clean_base_image_1604_20181230_0138" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-clean-base-image-1604-20181230-0138"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-clean-base-image-1604-77f8da6e"
}

resource "google_compute_image" "row_autoscale_clean_base_image_1604_20181231_1451" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-clean-base-image-1604-20181231-1451"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-clean-base-image-1604-bc1dc135"
}

resource "google_compute_image" "row_autoscale_clean_base_image_1604_20190106_0137" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-clean-base-image-1604-20190106-0137"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-clean-base-image-1604-d9e7b536"
}

resource "google_compute_image" "row_autoscale_docker_base_image_1604_20180621_0814" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-docker-base-image-1604-20180621-0814"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-docker-base-image-1604-08652b72"
}

resource "google_compute_image" "row_autoscale_docker_base_image_1604_20180621_0837" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-docker-base-image-1604-20180621-0837"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-docker-base-image-1604-0bd07c9b"
}

resource "google_compute_image" "row_autoscale_fluentd_base_image_1604_20180304_1119" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-fluentd-base-image-1604-20180304-1119"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-fluentd-base-image-1604-c2d5640c"
}

resource "google_compute_image" "row_autoscale_fluentd_base_image_1604_20180304_1524" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-fluentd-base-image-1604-20180304-1524"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-fluentd-base-image-1604-06a30947"
}

resource "google_compute_image" "row_autoscale_fluentd_base_image_1604_20180425_1832" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-fluentd-base-image-1604-20180425-1832"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-fluentd-base-image-1604-83efe062"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1404_20180729_0211" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-nginx-base-image-1404-20180729-0211"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1404-9461d1ea"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1404_20180805_0212" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-nginx-base-image-1404-20180805-0212"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1404-45aa86f2"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170305_0037" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170305-0037"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-8bed90aa"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170423_1401" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170423-1401"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-14b7c2e7"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170625_0028" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170625-0028"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-4c54fccc"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170629_0715" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170629-0715"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-5fb578ec"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170702_0031" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170702-0031"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-c957dc24"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170709_1318" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170709-1318"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-329b86c7"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170723_0037" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170723-0037"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-6897e86e"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170730_0029" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170730-0029"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-b8b36fda"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170807_1134" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170807-1134"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-fab40dd3"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170813_0715" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170813-0715"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-12fc030c"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170817_1348" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170817-1348"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-86a39df6"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170820_0031" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170820-0031"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-f64c5a7b"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170820_1101" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170820-1101"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-2dec3a23"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170824_0727" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170824-0727"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-d6cce184"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170827_0030" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170827-0030"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-1571aab8"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170903_0031" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170903-0031"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-575b640a"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20170924_0038" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20170924-0038"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-c1a1409d"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20171015_0030" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20171015-0030"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-e74254dd"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20171024_2018" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20171024-2018"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-92bd35c7"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20171210_0030" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20171210-0030"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-1e66f343"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20171217_0036" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20171217-0036"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-52212b0a"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180107_0032" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180107-0032"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-21d8679c"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180114_0032" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180114-0032"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-d50731f4"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180128_0038" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180128-0038"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-3146cb17"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180214_1229" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180214-1229"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-d3600c37"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180219_1252" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180219-1252"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-1b681f5c"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180306_1101" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180306-1101"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-84564a5b"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180313_2103" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180313-2103"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-cd0695a2"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180318_0248" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180318-0248"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-5496b8f9"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180424_2347" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180424-2347"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-ec038177"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180429_0206" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180429-0206"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-5488efad"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180520_0208" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180520-0208"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-a361e12c"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180527_0208" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180527-0208"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-c529d5be"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180603_0517" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180603-0517"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-8d8471ce"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180610_0201" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180610-0201"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-251aae2c"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180617_1155" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180617-1155"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-b02e98e0"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180624_0224" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180624-0224"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-297648c1"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180722_0848" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180722-0848"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-7b7c476a"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180805_0211" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180805-0211"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-87ae9476"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180812_0116" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180812-0116"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-35b35420"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180819_0120" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180819-0120"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-1c513afa"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180826_0122" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180826-0122"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-cc9c16bd"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180902_0117" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180902-0117"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-7f6d35f1"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180909_0118" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180909-0118"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-7caad7dc"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20180916_0118" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20180916-0118"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-600b1f5c"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181014_0141" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181014-0141"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-004f089e"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181015_1720" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181015-1720"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-016ace50"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181018_1948" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181018-1948"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-ac6bc3e8"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181021_1115" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181021-1115"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-845df081"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181028_1357" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181028-1357"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-cdf8fa14"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181104_0139" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181104-0139"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-b43c08e0"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181111_0139" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181111-0139"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-72e2945c"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181113_0726" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181113-0726"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-976f17d1"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181118_0139" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181118-0139"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-3257e998"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181209_0139" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181209-0139"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-273f8456"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181212_0852" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181212-0852"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-ffcfe495"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181216_1131" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181216-1131"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-47a1dff8"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181223_1333" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181223-1333"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-d1ae3838"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181226_1347" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181226-1347"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-0887676a"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181230_0138" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181230-0138"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-facb5ebc"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181231_1436" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181231-1436"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-1604-6a04d331"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20181231_1451" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20181231-1451"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-3901a51e"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20190101_1135" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20190101-1135"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-01fa81f0"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20190103_0837" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20190103-0837"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-nginx-base-image-1604-6eba888f"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_1604_20190106_0137" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-1604-20190106-0137"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-1604-978f87d4"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_20160420_2146" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-nginx-base-image-20160420-2146"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-8c3c481a"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_20160616_1011" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-nginx-base-image-20160616-1011"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-3adf305c"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_20160811_1336" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-nginx-base-image-20160811-1336"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-6c32daa8"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_20161202_1831" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-nginx-base-image-20161202-1831"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-nginx-base-image-b9ff8c3f"
}

resource "google_compute_image" "row_autoscale_nginx_base_image_20170116_0824" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-nginx-base-image-20170116-0824"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-nginx-base-image-8686107f"
}

resource "google_compute_image" "row_autoscale_realtime_base_image_1404_20180729_0222" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-realtime-base-image-1404-20180729-0222"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-realtime-base-image-1404-7ea1c482"
}

resource "google_compute_image" "row_autoscale_realtime_base_image_1404_20180805_0218" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-realtime-base-image-1404-20180805-0218"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-realtime-base-image-1404-8ddc4f85"
}

resource "google_compute_image" "row_autoscale_realtime_base_image_1604_20180708_0242" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-realtime-base-image-1604-20180708-0242"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-realtime-base-image-1604-cc12c41a"
}

resource "google_compute_image" "row_autoscale_realtime_base_image_1604_20181216_0138" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-realtime-base-image-1604-20181216-0138"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-realtime-base-image-1604-16f3c93b"
}

resource "google_compute_image" "row_autoscale_realtime_base_image_1604_20181230_0138" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-realtime-base-image-1604-20181230-0138"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-realtime-base-image-1604-459b3f32"
}

resource "google_compute_image" "row_autoscale_realtime_base_image_1604_20181231_1451" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-realtime-base-image-1604-20181231-1451"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-realtime-base-image-1604-81ffc492"
}

resource "google_compute_image" "row_autoscale_realtime_base_image_1604_20181231_1537" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-realtime-base-image-1604-20181231-1537"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-realtime-base-image-1604-63978339"
}

resource "google_compute_image" "row_autoscale_realtime_base_image_1604_20190106_0138" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-realtime-base-image-1604-20190106-0138"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-realtime-base-image-1604-8d8ec6dc"
}

resource "google_compute_image" "row_autoscale_rtproxy_base_image_20151029_1418" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-rtproxy-base-image-20151029-1418"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-rtproxy-base-image-1b89dafc"
}

resource "google_compute_image" "row_autoscale_rtproxy_base_image_20151029_1612" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-rtproxy-base-image-20151029-1612"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-rtproxy-base-image-37474823"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1404_20170207_1145" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-tomcat-base-image-1404-20170207-1145"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1404-48973a30"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1404_20180729_0218" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-tomcat-base-image-1404-20180729-0218"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-1404-44091bba"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1404_20180805_0216" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-tomcat-base-image-1404-20180805-0216"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-1404-f0467a9c"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20170305_0037" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20170305-0037"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1604-f66c1e74"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20170709_1256" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20170709-1256"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-1604-6a2a0d60"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20170806_1229" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20170806-1229"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-1604-cd150e22"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20170814_0735" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20170814-0735"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-1604-9e174ff6"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20170822_0754" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20170822-0754"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-1604-13a83cde"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20170827_0031" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20170827-0031"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-1604-38f4a192"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20170928_0616" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20170928-0616"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-1604-f95a6152"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20171018_0929" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20171018-0929"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-1604-92ad9eb0"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20171023_1355" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20171023-1355"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-1604-47635e65"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20171123_0954" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20171123-0954"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1604-bc9f182d"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20171205_1300" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20171205-1300"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-1604-c679b520"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20171224_0041" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20171224-0041"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-1604-26c5ada5"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20171231_0032" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20171231-0032"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-1604-dad95f2c"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20180114_0033" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20180114-0033"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1604-cfd9f621"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20180128_0036" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20180128-0036"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1604-1701e356"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20180214_0709" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20180214-0709"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-1604-297bcf25"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20180215_0806" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20180215-0806"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1604-12120247"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20180506_0211" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20180506-0211"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-1604-21711eac"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20180516_0717" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20180516-0717"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1604-7ae302be"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20180520_0209" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20180520-0209"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1604-f49a2c1b"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20180624_0229" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20180624-0229"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-1604-d4e3bd3b"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20180708_0241" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20180708-0241"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-1604-81da2e60"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20180729_0219" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20180729-0219"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1604-d7864ff5"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20180819_0120" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20180819-0120"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-1604-cc98b162"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20180826_0123" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20180826-0123"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1604-c7ae0cf0"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20180902_0118" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20180902-0118"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1604-79744440"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20180930_0118" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20180930-0118"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-1604-81c98898"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20181014_0142" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20181014-0142"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1604-a575e8d5"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20181021_0136" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20181021-0136"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1604-b28da17d"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20181118_0139" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20181118-0139"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-1604-6e912ee6"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20181209_0138" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20181209-0138"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-1604-5e1e2e98"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20181216_0139" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20181216-0139"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-1604-cc13f59c"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20181223_0138" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20181223-0138"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1604-593eb96c"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20181230_0138" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20181230-0138"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-1604-8ab732e0"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20181231_1451" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20181231-1451"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-1604-0ead4566"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_1604_20190106_0138" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-autoscale-tomcat-base-image-1604-20190106-0138"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-autoscale-tomcat-base-image-1604-f710d2cc"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_20160530_1456" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-tomcat-base-image-20160530-1456"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-490a43d4"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_20160615_1045" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-tomcat-base-image-20160615-1045"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-7d760efd"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_20160929_1047" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-tomcat-base-image-20160929-1047"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-6c3a7775"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_20161124_0949" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-tomcat-base-image-20161124-0949"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-4010ff66"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_20161211_0029" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-tomcat-base-image-20161211-0029"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-autoscale-tomcat-base-image-ff6b98f5"
}

resource "google_compute_image" "row_autoscale_tomcat_base_image_20170103_0735" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-autoscale-tomcat-base-image-20170103-0735"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-autoscale-tomcat-base-image-b4aa0aba"
}

resource "google_compute_image" "row_cassandra_1_2_19_clean_base_image_1604_20171108_0802" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-cassandra-1-2-19-clean-base-image-1604-20171108-0802"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-cassandra-1-2-19-clean-base-image-1604-18294f7f"
}

resource "google_compute_image" "row_cassandra_2_1_15_clean_base_image_1604_20170801_0906" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-cassandra-2-1-15-clean-base-image-1604-20170801-0906"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-cassandra-2-1-15-clean-base-image-1604-f666a2b5"
}

resource "google_compute_image" "row_cassandra_2_1_15_clean_base_image_1604_20170802_0928" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-cassandra-2-1-15-clean-base-image-1604-20170802-0928"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-cassandra-2-1-15-clean-base-image-1604-559bf496"
}

resource "google_compute_image" "row_cassandra_2_1_18_clean_base_image_1604_20170807_0740" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-cassandra-2-1-18-clean-base-image-1604-20170807-0740"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-cassandra-2-1-18-clean-base-image-1604-44354c30"
}

resource "google_compute_image" "row_dev_editor1_28_8_17" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1204-precise"]
  name        = "row-dev-editor1-28-8-17"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-dev-editor1-new"
}

resource "google_compute_image" "row_emailserver_20161117_1049" {
  description = "row-emailserver-20161117-1049"

  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-emailserver-20161117-1049"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/packer-582d8b50-a950-5270-8a89-a0f069649a55"
}

resource "google_compute_image" "row_fluentd_forwarder_base_image_1604_20170904_0939" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-fluentd-forwarder-base-image-1604-20170904-0939"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-fluentd-forwarder-base-image-1604-0afdf8d3"
}

resource "google_compute_image" "row_fluentd_forwarder_base_image_1604_20180304_0843" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-fluentd-forwarder-base-image-1604-20180304-0843"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-fluentd-forwarder-base-image-1604-690e7a4e"
}

resource "google_compute_image" "row_realtime_base_image_1404_20170117_1116" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-realtime-base-image-1404-20170117-1116"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-realtime-base-image-1404-f7739076"
}

resource "google_compute_image" "row_realtime_base_image_1604_20170709_1258" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-realtime-base-image-1604-20170709-1258"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-realtime-base-image-1604-77c85099"
}

resource "google_compute_image" "row_realtime_base_image_1604_20170725_0802" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-realtime-base-image-1604-20170725-0802"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-realtime-base-image-1604-793c3b5f"
}

resource "google_compute_image" "row_realtime_base_image_20170112_1457" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-realtime-base-image-20170112-1457"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-realtime-base-image-4fad9bee"
}

resource "google_compute_image" "row_realtime_base_image_20170115_1132" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-realtime-base-image-20170115-1132"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-realtime-base-image-3ed199fd"
}

resource "google_compute_image" "row_realtime_proxy_base_image_1604_20170718_1507" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-realtime-proxy-base-image-1604-20170718-1507"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-realtime-proxy-base-image-1604-6d5349b9"
}

resource "google_compute_image" "row_realtime_proxy_base_image_1604_20170719_0717" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-realtime-proxy-base-image-1604-20170719-0717"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-realtime-proxy-base-image-1604-dadf49d9"
}

resource "google_compute_image" "row_realtime_proxy_base_image_1604_20170719_0731" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-realtime-proxy-base-image-1604-20170719-0731"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-realtime-proxy-base-image-1604-8e84bc5b"
}

resource "google_compute_image" "row_realtime_proxy_base_image_1604_20180723_1127" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "row-realtime-proxy-base-image-1604-20180723-1127"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-realtime-proxy-base-image-1604-27644e72"
}

resource "google_compute_image" "row_realtime_proxy_base_image_20160809_0806" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-realtime-proxy-base-image-20160809-0806"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-realtime-proxy-base-image-9e7ab936"
}

resource "google_compute_image" "row_realtime_proxy_base_image_20160809_0917" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-realtime-proxy-base-image-20160809-0917"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-realtime-proxy-base-image-ca4a4602"
}

resource "google_compute_image" "row_routingserver_beta_20161102_1308" {
  description = "row-routingserver-beta-20161102-1308"

  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-routingserver-beta-20161102-1308"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/packer-5819e54a-3d82-7a31-1db6-d1f2d93d7ace"
}

resource "google_compute_image" "row_routingserver_beta_20161109_1712" {
  description = "row-routingserver-beta-20161109-1712"

  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-routingserver-beta-20161109-1712"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/packer-5823590e-15b0-a50c-134c-2f94fa3aae2d"
}

resource "google_compute_image" "row_staging_base_image_20151102_1318" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-staging-base-image-20151102-1318"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-staging-base-image-c844a410"
}

resource "google_compute_image" "row_staging_base_image_20151104_1748" {
  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-staging-base-image-20151104-1748"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-staging-base-image-b889fb49"
}

resource "google_compute_image" "row_staging_cassandra_dse_template_test_1" {
  family      = "cassandra-template-dse"
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-staging-cassandra-dse-template-test-1"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-staging-cassandra-dse-template-test-1"
}

resource "google_compute_image" "row_staging_cassandra_dse_template_test_provision" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-staging-cassandra-dse-template-test-provision"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-staging-cassandra-dse-test-sessionprocessor-9oog"
}

resource "google_compute_image" "row_staging_cassandra_dse_template_test_provision2" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-staging-cassandra-dse-template-test-provision2"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/row-staging-cassandra-sessionprocessor-test2-v001-q9u8"
}

resource "google_compute_image" "row_staging_cassandra_dse_template_test_provision3" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-staging-cassandra-dse-template-test-provision3"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/tom-test-template-cassandra1"
}

resource "google_compute_image" "row_staging_cassandra_dse_template_test_provision7" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "row-staging-cassandra-dse-template-test-provision7"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-staging-cassandra-sessionprocessor-test2-v001-y4z1"
}

resource "google_compute_image" "rt_stg_boot_disk_v1" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-boot-disk-v1"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/gabi-rt-disk-delme"
}

resource "google_compute_image" "rt_stg_boot_disk_v10" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-boot-disk-v10"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/gabi-rt-disk-delme"
}

resource "google_compute_image" "rt_stg_boot_disk_v11" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-boot-disk-v11"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/gabi-rt-disk-delme"
}

resource "google_compute_image" "rt_stg_boot_disk_v12" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-boot-disk-v12"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/gabi-rt-disk-delme"
}

resource "google_compute_image" "rt_stg_boot_disk_v13" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-boot-disk-v13"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/gabi-rt-disk-delme"
}

resource "google_compute_image" "rt_stg_boot_disk_v14" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-boot-disk-v14"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/gabi-rt-disk-delme"
}

resource "google_compute_image" "rt_stg_boot_disk_v15" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-boot-disk-v15"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/gabi-rt-disk-delme"
}

resource "google_compute_image" "rt_stg_boot_disk_v16" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-boot-disk-v16"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/gabi-rt-disk-delme"
}

resource "google_compute_image" "rt_stg_boot_disk_v17" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-boot-disk-v17"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/gabi-rt-disk-delme"
}

resource "google_compute_image" "rt_stg_boot_disk_v18" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-boot-disk-v18"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/gabi-rt-disk-delme"
}

resource "google_compute_image" "rt_stg_boot_disk_v19" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-boot-disk-v19"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/gabi-rt-disk-delme"
}

resource "google_compute_image" "rt_stg_boot_disk_v20" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-boot-disk-v20"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/gabi-rt-disk-delme"
}

resource "google_compute_image" "rt_stg_boot_disk_v21" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-boot-disk-v21"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/gabi-rt-disk-delme"
}

resource "google_compute_image" "rt_stg_fe_2_boot_disk_v2" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-fe-2-boot-disk-v2"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/rt-fe-2-gabi-delme"
}

resource "google_compute_image" "rt_stg_fe_2_boot_disk_v3" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-fe-2-boot-disk-v3"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-staging-rt-fe-2-gvqv"
}

resource "google_compute_image" "rt_stg_fe_2_boot_disk_v4" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-fe-2-boot-disk-v4"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-staging-rt-fe-2-139y"
}

resource "google_compute_image" "rt_stg_fe_2_boot_disk_v5" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-fe-2-boot-disk-v5"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-staging-rt-fe-2-xfdo"
}

resource "google_compute_image" "rt_stg_fe_2_boot_disk_v6" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-fe-2-boot-disk-v6"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/row-staging-rt-fe-2-ivvf"
}

resource "google_compute_image" "rt_stg_fe_2_boot_disk_v8" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "rt-stg-fe-2-boot-disk-v8"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/gabi-rt-disk-delme"
}

resource "google_compute_image" "staging_cassandra1_data_20180501" {
  family      = "staging-cassandra1-data"
  labels      = {}
  name        = "staging-cassandra1-data-20180501"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/cassandra-data"
}

resource "google_compute_image" "staging_cassandra2_data_20180501" {
  family      = "staging-cassandra2-data"
  labels      = {}
  name        = "staging-cassandra2-data-20180501"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/cassandra-data"
}

resource "google_compute_image" "staging_cassandra3_data_20180501" {
  family      = "staging-cassandra3-data"
  labels      = {}
  name        = "staging-cassandra3-data-20180501"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/cassandra-data"
}

resource "google_compute_image" "staging_cassandra_image_20180426" {
  family      = "staging-cassandra-image"
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "staging-cassandra-image-20180426"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/cassandra-base-image"
}

resource "google_compute_image" "staging_db_image_20180618" {
  family      = "staging-db-image"
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "staging-db-image-20180618"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/postgres-stg"
}

resource "google_compute_image" "staging_db_image_20180618_1" {
  family      = "staging-db-image"
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "staging-db-image-20180618-1"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/postgres-stg"
}

resource "google_compute_image" "staging_db_image_20180619" {
  family      = "staging-db-image"
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "staging-db-image-20180619"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/postgres-stg"
}

resource "google_compute_image" "staging_db_image_2018_04_10" {
  family      = "staging-db-image"
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "staging-db-image-2018-04-10"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/staging-db-template"
}

resource "google_compute_image" "staging_db_image_2018_04_16" {
  family      = "staging-db-image"
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1604-xenial"]
  name        = "staging-db-image-2018-04-16"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/postgres-stg"
}

resource "google_compute_image" "stg1_emailserver_20161117_1216" {
  description = "stg1-emailserver-20161117-1216"

  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "stg1-emailserver-20161117-1216"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/packer-582da81a-5021-a044-a9d4-3edfc08e8f02"
}

resource "google_compute_image" "stg1_routingserver_20161123_1459" {
  description = "stg1-routingserver-20161123-1459"

  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "stg1-routingserver-20161123-1459"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/packer-5835aee9-8c33-ed30-553e-4cb2f77b8478"
}

resource "google_compute_image" "stg1_routingserver_20161124_0845" {
  description = "stg1-routingserver-20161124-0845"

  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "stg1-routingserver-20161124-0845"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/packer-5836a8f2-e17a-ca2c-b3f3-381bd3300c84"
}

resource "google_compute_image" "stg1_routingserver_beta_20161114_1042" {
  description = "stg1-routingserver-beta-20161114-1042"

  labels = {
    prune = true
  }

  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "stg1-routingserver-beta-20161114-1042"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-d/disks/packer-58299543-2e81-fe48-18f2-4df6ba055484"
}

resource "google_compute_image" "tts_engine_v01" {
  labels      = {}
  licenses    = ["https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-1404-trusty"]
  name        = "tts-engine-v01"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-b/disks/row-staging-tts-engine1"
}

resource "google_compute_image" "tts_engine_v02" {
  labels      = {}
  name        = "tts-engine-v02"
  project     = "waze-development"
  source_disk = "https://www.googleapis.com/compute/v1/projects/waze-development/zones/europe-west1-c/disks/tts-engine-deleteme"
}
