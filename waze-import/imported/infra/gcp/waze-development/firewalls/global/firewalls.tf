provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_firewall" "aef_realtime__login__prober__prod__il_20181231t154600_hcfw" {
  allow = {
    ports    = ["8443"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "aef-realtime--login--prober--prod--il-20181231t154600-hcfw"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["130.211.0.0/22", "35.191.0.0/16"]
  target_tags    = ["aef-realtime--login--prober--prod--il-20181231t154600"]
}

resource "google_compute_firewall" "aef_realtime__login__prober__stg_20181231t100523_hcfw" {
  allow = {
    ports    = ["8443"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "aef-realtime--login--prober--stg-20181231t100523-hcfw"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["130.211.0.0/22", "35.191.0.0/16"]
  target_tags    = ["aef-realtime--login--prober--stg-20181231t100523"]
}

resource "google_compute_firewall" "aef_routing__regression_20180124t102927_hcfw" {
  allow = {
    ports    = ["8443"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "aef-routing--regression-20180124t102927-hcfw"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["130.211.0.0/22", "35.191.0.0/16"]
  target_tags    = ["aef-routing--regression-20180124t102927"]
}

resource "google_compute_firewall" "allow_health_check" {
  allow = {
    ports    = ["8088"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "allow-health-check"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["130.211.0.0/22", "35.191.0.0/16"]
  target_tags    = ["health-check-tag"]
}

resource "google_compute_firewall" "allow_https_from_logserver" {
  allow = {
    ports    = ["443"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "allow-https-from-logserver"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["54.204.41.28/32"]
  target_tags    = ["row-staging-topic-events"]
}

resource "google_compute_firewall" "allow_ssh_from_logserver" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  description    = "Allows SSH access to stg from waze-logserver1 for Ron Desta"
  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "allow-ssh-from-logserver"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["176.231.65.5/32", "54.204.41.28/32"]
}

resource "google_compute_firewall" "bastion_incoming" {
  allow = {
    ports    = ["443"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "bastion-incoming"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["54.197.192.34/32", "54.195.80.116/32", "54.197.192.35/32", "54.195.80.117/32"]
  target_tags    = ["bastion-https"]
}

resource "google_compute_firewall" "cassandra_on_k8s_allow_icmp" {
  allow = {
    protocol = "icmp"
  }

  description    = "Allows ICMP connections from any source to any instance on the network."
  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "cassandra-on-k8s-allow-icmp"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority       = "65534"
  project        = "waze-development"
  source_ranges  = ["0.0.0.0/0"]
}

resource "google_compute_firewall" "cassandra_on_k8s_allow_ssh" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  description    = "Allows TCP connections from any source to any instance on the network using port 22."
  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "cassandra-on-k8s-allow-ssh"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority       = "65534"
  project        = "waze-development"
  source_ranges  = ["0.0.0.0/0"]
}

resource "google_compute_firewall" "default_ads_memcached_stg1_service" {
  allow = {
    ports    = ["11211"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "default-ads-memcached-stg1-service"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["stg-poi", "ads-stg-stg1-client", "staging-stg-stg1-client"]
}

resource "google_compute_firewall" "default_allow_http" {
  allow = {
    ports    = ["80"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "default-allow-http"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["0.0.0.0/0"]
  target_tags    = ["http-server"]
}

resource "google_compute_firewall" "default_allow_https" {
  allow = {
    ports    = ["443"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "default-allow-https"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["0.0.0.0/0"]
  target_tags    = ["https-server"]
}

resource "google_compute_firewall" "default_allow_internal" {
  allow = {
    ports    = ["0-65535"]
    protocol = "udp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    ports    = ["0-65535"]
    protocol = "tcp"
  }

  description    = "Internal traffic from default allowed"
  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "default-allow-internal"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "65534"
  project        = "waze-development"
  source_ranges  = ["10.0.0.0/8"]
}

resource "google_compute_firewall" "default_carpool_memcache_stg1_service" {
  allow = {
    ports    = ["11211"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "default-carpool-memcache-stg1-service"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["carpoolpayments-stg-stg1-client"]
  target_tags    = ["carpool-memcached-stg1-service"]
}

resource "google_compute_firewall" "default_general_memcached_stg1_service" {
  allow = {
    ports    = ["11211"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "default-general-memcached-stg1-service"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["general-stg-stg1-client"]
}

resource "google_compute_firewall" "default_load_balanced_service" {
  allow = {
    ports    = ["443"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["24224-24228"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["30000-32767"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["8088"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["18080"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["8001"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["11211"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["20888"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["80-82"]
    protocol = "tcp"
  }

  description    = "Allow access from Google's GFE network range for Load Balanced services"
  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "default-load-balanced-service"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["130.211.0.0/22", "35.191.0.0/16", "10.240.0.0/16"]
  target_tags    = ["load-balanced-service", "gke-adman2-stg-1ce64671-node"]
}

resource "google_compute_firewall" "default_ssh_from_deployment" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  description    = "SSH From Deployment Console"
  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "default-ssh-from-deployment"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["104.155.113.61/32", "104.134.21.0/24", "35.205.205.222/32", "54.197.192.19/32"]
  target_tags    = ["default-ssh-from-deployment"]
}

resource "google_compute_firewall" "default_staging_memcached_stg1_service" {
  allow = {
    ports    = ["11211"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "default-staging-memcached-stg1-service"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["staging-stg-stg1-client"]
}

resource "google_compute_firewall" "from_qa_22" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "from-qa-22"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["104.155.3.206/32", "10.240.51.112/32"]
  target_tags    = ["qa-ssh"]
}

resource "google_compute_firewall" "gke_adman2_stg_1ce64671_all" {
  allow = {
    protocol = "esp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    protocol = "udp"
  }

  allow = {
    protocol = "ah"
  }

  allow = {
    protocol = "sctp"
  }

  allow = {
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-adman2-stg-1ce64671-all"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.64.0.0/14"]
}

resource "google_compute_firewall" "gke_adman2_stg_1ce64671_ssh" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-adman2-stg-1ce64671-ssh"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["130.211.65.152/32"]
  target_tags    = ["gke-adman2-stg-1ce64671-node"]
}

resource "google_compute_firewall" "gke_adman2_stg_1ce64671_vms" {
  allow = {
    ports    = ["1-65535"]
    protocol = "tcp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    ports    = ["1-65535"]
    protocol = "udp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-adman2-stg-1ce64671-vms"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.240.0.0/16"]
  target_tags    = ["gke-adman2-stg-1ce64671-node"]
}

resource "google_compute_firewall" "gke_adman2_stg_ng_d67c5900_all" {
  allow = {
    protocol = "esp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    protocol = "udp"
  }

  allow = {
    protocol = "ah"
  }

  allow = {
    protocol = "sctp"
  }

  allow = {
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-adman2-stg-ng-d67c5900-all"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.40.0.0/14"]
  target_tags    = ["gke-adman2-stg-ng-d67c5900-node"]
}

resource "google_compute_firewall" "gke_adman2_stg_ng_d67c5900_ssh" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-adman2-stg-ng-d67c5900-ssh"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["35.233.19.217/32", "35.240.96.69/32", "130.211.105.27/32"]
  target_tags    = ["gke-adman2-stg-ng-d67c5900-node"]
}

resource "google_compute_firewall" "gke_adman2_stg_ng_d67c5900_vms" {
  allow = {
    ports    = ["1-65535"]
    protocol = "tcp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    ports    = ["1-65535"]
    protocol = "udp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-adman2-stg-ng-d67c5900-vms"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.240.0.0/16"]
  target_tags    = ["gke-adman2-stg-ng-d67c5900-node"]
}

resource "google_compute_firewall" "gke_help_b3fef0b0_all" {
  allow = {
    protocol = "esp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    protocol = "udp"
  }

  allow = {
    protocol = "ah"
  }

  allow = {
    protocol = "sctp"
  }

  allow = {
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-help-b3fef0b0-all"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.44.0.0/14"]
}

resource "google_compute_firewall" "gke_help_b3fef0b0_ssh" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-help-b3fef0b0-ssh"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["35.233.104.16/32"]
  target_tags    = ["gke-help-b3fef0b0-node"]
}

resource "google_compute_firewall" "gke_help_b3fef0b0_vms" {
  allow = {
    ports    = ["1-65535"]
    protocol = "tcp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    ports    = ["1-65535"]
    protocol = "udp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-help-b3fef0b0-vms"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.240.0.0/16"]
  target_tags    = ["gke-help-b3fef0b0-node"]
}

resource "google_compute_firewall" "gke_maint_jobs_stg_5325aa72_all" {
  allow = {
    protocol = "esp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    protocol = "udp"
  }

  allow = {
    protocol = "ah"
  }

  allow = {
    protocol = "sctp"
  }

  allow = {
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-maint-jobs-stg-5325aa72-all"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.56.0.0/14"]
}

resource "google_compute_firewall" "gke_maint_jobs_stg_5325aa72_ssh" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-maint-jobs-stg-5325aa72-ssh"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["104.199.30.155/32"]
  target_tags    = ["gke-maint-jobs-stg-5325aa72-node"]
}

resource "google_compute_firewall" "gke_maint_jobs_stg_5325aa72_vms" {
  allow = {
    ports    = ["1-65535"]
    protocol = "tcp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    ports    = ["1-65535"]
    protocol = "udp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-maint-jobs-stg-5325aa72-vms"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.240.0.0/16"]
  target_tags    = ["gke-maint-jobs-stg-5325aa72-node"]
}

resource "google_compute_firewall" "gke_stg_tools_9fc715a0_all" {
  allow = {
    protocol = "esp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    protocol = "udp"
  }

  allow = {
    protocol = "ah"
  }

  allow = {
    protocol = "sctp"
  }

  allow = {
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-stg-tools-9fc715a0-all"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.16.0.0/14"]
  target_tags    = ["gke-stg-tools-9fc715a0-node"]
}

resource "google_compute_firewall" "gke_stg_tools_9fc715a0_ssh" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-stg-tools-9fc715a0-ssh"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["35.233.120.71/32"]
  target_tags    = ["gke-stg-tools-9fc715a0-node"]
}

resource "google_compute_firewall" "gke_stg_tools_9fc715a0_vms" {
  allow = {
    ports    = ["1-65535"]
    protocol = "tcp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    ports    = ["1-65535"]
    protocol = "udp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-stg-tools-9fc715a0-vms"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.240.0.0/16"]
  target_tags    = ["gke-stg-tools-9fc715a0-node"]
}

resource "google_compute_firewall" "gke_us_central1_tiles_builder_1be8ed5a_gke_6444a243_all" {
  allow = {
    protocol = "esp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    protocol = "udp"
  }

  allow = {
    protocol = "ah"
  }

  allow = {
    protocol = "sctp"
  }

  allow = {
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-us-central1-tiles-builder-1be8ed5a-gke-6444a243-all"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.4.0.0/14"]
  target_tags    = ["gke-us-central1-tiles-builder-1be8ed5a-gke-6444a243-node"]
}

resource "google_compute_firewall" "gke_us_central1_tiles_builder_1be8ed5a_gke_6444a243_ssh" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-us-central1-tiles-builder-1be8ed5a-gke-6444a243-ssh"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["35.225.183.94/32"]
  target_tags    = ["gke-us-central1-tiles-builder-1be8ed5a-gke-6444a243-node"]
}

resource "google_compute_firewall" "gke_us_central1_tiles_builder_1be8ed5a_gke_6444a243_vms" {
  allow = {
    ports    = ["1-65535"]
    protocol = "tcp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    ports    = ["1-65535"]
    protocol = "udp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-us-central1-tiles-builder-1be8ed5a-gke-6444a243-vms"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.240.0.0/16"]
  target_tags    = ["gke-us-central1-tiles-builder-1be8ed5a-gke-6444a243-node"]
}

resource "google_compute_firewall" "gke_waze_tools_stg_cb448316_all" {
  allow = {
    protocol = "esp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    protocol = "udp"
  }

  allow = {
    protocol = "ah"
  }

  allow = {
    protocol = "sctp"
  }

  allow = {
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-waze-tools-stg-cb448316-all"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.20.0.0/14"]
  target_tags    = ["gke-waze-tools-stg-cb448316-node"]
}

resource "google_compute_firewall" "gke_waze_tools_stg_cb448316_ssh" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-waze-tools-stg-cb448316-ssh"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["35.195.24.76/32"]
  target_tags    = ["gke-waze-tools-stg-cb448316-node"]
}

resource "google_compute_firewall" "gke_waze_tools_stg_cb448316_vms" {
  allow = {
    ports    = ["1-65535"]
    protocol = "tcp"
  }

  allow = {
    protocol = "icmp"
  }

  allow = {
    ports    = ["1-65535"]
    protocol = "udp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "gke-waze-tools-stg-cb448316-vms"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.240.0.0/16"]
  target_tags    = ["gke-waze-tools-stg-cb448316-node"]
}

resource "google_compute_firewall" "googlea_https" {
  allow = {
    ports    = ["443"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "googlea-https"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["31.154.8.71/32", "104.132.52.25/32", "104.134.21.0/24", "79.179.70.120/32", "31.154.8.68/32", "31.154.8.70/32", "216.239.32.0/19", "74.125.0.0/16"]
  target_tags    = ["feed", "googlea-https"]
}

resource "google_compute_firewall" "googlea_topictest_flume_debug" {
  allow = {
    ports    = ["41414"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "googlea-topictest-flume-debug"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["31.154.8.71/32", "104.132.52.25/32", "79.179.70.120/32", "31.154.8.68/32", "31.154.8.70/32", "216.239.32.0/19", "74.125.0.0/16"]
  target_tags    = ["topictest-flume-debug"]
}

resource "google_compute_firewall" "http" {
  allow = {
    ports    = ["80"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "http"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["0.0.0.0/0"]
  target_tags    = ["http-access"]
}

resource "google_compute_firewall" "https" {
  allow = {
    ports    = ["443"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "https"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["0.0.0.0/0"]
  target_tags    = ["https-access"]
}

resource "google_compute_firewall" "icmp_test_from_google" {
  allow = {
    protocol = "icmp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "icmp-test-from-google"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["31.154.8.71/32", "79.179.70.120/32", "31.154.8.68/32", "31.154.8.70/32", "216.239.32.0/19", "74.125.0.0/16"]
  target_tags    = ["test-icmp"]
}

resource "google_compute_firewall" "k8s_865dbaeaed1d6cc4_node_http_hc" {
  allow = {
    ports    = ["10256"]
    protocol = "tcp"
  }

  description    = "{\"kubernetes.io/cluster-id\":\"865dbaeaed1d6cc4\"}"
  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "k8s-865dbaeaed1d6cc4-node-http-hc"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["209.85.204.0/22", "209.85.152.0/22", "130.211.0.0/22", "35.191.0.0/16"]
  target_tags    = ["gke-maint-jobs-stg-5325aa72-node"]
}

resource "google_compute_firewall" "k8s_fw_a7d56307ce78011e8bd1242010a84003" {
  allow = {
    ports    = ["14141"]
    protocol = "tcp"
  }

  description    = "{\"kubernetes.io/service-name\":\"default/esp-grpc-wam-stg-p2p-test-service\", \"kubernetes.io/service-ip\":\"35.240.79.195\"}"
  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "k8s-fw-a7d56307ce78011e8bd1242010a84003"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["0.0.0.0/0"]
  target_tags    = ["gke-maint-jobs-stg-5325aa72-node"]
}

resource "google_compute_firewall" "logserver" {
  allow = {
    ports    = ["514"]
    protocol = "udp"
  }

  allow = {
    ports    = ["514"]
    protocol = "tcp"
  }

  description    = "Logserver to rampart"
  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "logserver"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.240.0.0/16"]
  target_tags    = ["logserver"]
}

resource "google_compute_firewall" "monitoring_server" {
  allow = {
    ports    = ["6557"]
    protocol = "tcp"
  }

  description    = "port 6557 from AWS eu monitoring"
  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "monitoring-server"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["54.195.80.50/32"]
  target_tags    = ["monitor-server"]
}

resource "google_compute_firewall" "prod_realtime_ddb_cassandra_stg_cluster" {
  allow = {
    ports    = ["7001", "7000", "7199", "23456"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["7000", "7001", "7199"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-realtime-ddb-cassandra-stg-cluster"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["realtime-ddb-cassandra-stg-cluster"]
  target_tags    = ["realtime-ddb-cassandra-stg-cluster"]
}

resource "google_compute_firewall" "prod_realtime_ddb_cassandra_stg_service" {
  allow = {
    ports    = ["9042", "9160"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-realtime-ddb-cassandra-stg-service"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["realtime-ddb-cassandra-stg-service"]
  target_tags    = ["realtime-ddb-cassandra-stg-service"]
}

resource "google_compute_firewall" "prod_staging_cassandra_stg_cluster" {
  allow = {
    ports    = ["7001", "7000", "7199", "23456"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["7000", "7001", "7199"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-staging-cassandra-stg-cluster"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["staging-cassandra-stg-cluster"]
  target_tags    = ["staging-cassandra-stg-cluster"]
}

resource "google_compute_firewall" "prod_staging_cassandra_stg_service" {
  allow = {
    ports    = ["9042", "9160"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-staging-cassandra-stg-service"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["staging-cassandra-stg-service"]
  target_tags    = ["staging-cassandra-stg-service"]
}

resource "google_compute_firewall" "prod_staging_general21_cassandra_stg_cluster" {
  allow = {
    ports    = ["7001", "7000", "7199", "23456"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["7000", "7001", "7199"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-staging-general21-cassandra-stg-cluster"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["staging-general21-cassandra-stg-cluster"]
  target_tags    = ["staging-general21-cassandra-stg-cluster"]
}

resource "google_compute_firewall" "prod_staging_general21_cassandra_stg_service" {
  allow = {
    ports    = ["9042", "9160"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-staging-general21-cassandra-stg-service"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["staging-general21-cassandra-stg-service"]
  target_tags    = ["staging-general21-cassandra-stg-service"]
}

resource "google_compute_firewall" "prod_staging_master_cassandra_stg_cluster" {
  allow = {
    ports    = ["7001", "7000", "7199", "23456"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["7000", "7001", "7199"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-staging-master-cassandra-stg-cluster"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["staging-master-cassandra-stg-cluster"]
  target_tags    = ["staging-master-cassandra-stg-cluster"]
}

resource "google_compute_firewall" "prod_staging_master_cassandra_stg_service" {
  allow = {
    ports    = ["9042", "9160"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-staging-master-cassandra-stg-service"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["staging-master-cassandra-stg-service"]
  target_tags    = ["staging-master-cassandra-stg-service"]
}

resource "google_compute_firewall" "prod_staging_youying_test_cassandra_stg_cluster" {
  allow = {
    ports    = ["7001", "7000", "7199", "23456"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["7000", "7001", "7199"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-staging-youying-test-cassandra-stg-cluster"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["staging-youying-test-cassandra-stg-cluster"]
  target_tags    = ["staging-youying-test-cassandra-stg-cluster"]
}

resource "google_compute_firewall" "prod_staging_youying_test_cassandra_stg_service" {
  allow = {
    ports    = ["9042", "9160"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-staging-youying-test-cassandra-stg-service"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["staging-youying-test-cassandra-stg-service"]
  target_tags    = ["staging-youying-test-cassandra-stg-service"]
}

resource "google_compute_firewall" "prod_stgregulator_cassandra_stg_cluster" {
  allow = {
    ports    = ["7000", "7001", "7199", "23456"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-stgregulator-cassandra-stg-cluster"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["stgregulator-cassandra-stg-cluster"]
  target_tags    = ["stgregulator-cassandra-stg-cluster"]
}

resource "google_compute_firewall" "prod_stgregulator_cassandra_stg_service" {
  allow = {
    ports    = ["9042", "9160"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-stgregulator-cassandra-stg-service"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["stgregulator-cassandra-stg-service"]
  target_tags    = ["stgregulator-cassandra-stg-service"]
}

resource "google_compute_firewall" "prod_stgtopics_events_cassandra_stg_cluster" {
  allow = {
    ports    = ["7000", "7001", "7199", "23456"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-stgtopics-events-cassandra-stg-cluster"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["stgtopics-events-cassandra-stg-cluster"]
  target_tags    = ["stgtopics-events-cassandra-stg-cluster"]
}

resource "google_compute_firewall" "prod_stgtopics_events_cassandra_stg_service" {
  allow = {
    ports    = ["9042", "9160"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-stgtopics-events-cassandra-stg-service"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["stgtopics-events-cassandra-stg-service"]
  target_tags    = ["stgtopics-events-cassandra-stg-service"]
}

resource "google_compute_firewall" "prod_test_cassandra_stg_cluster" {
  allow = {
    ports    = ["7000", "7001", "7199"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-test-cassandra-stg-cluster"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["test-cassandra-stg-cluster"]
  target_tags    = ["test-cassandra-stg-cluster"]
}

resource "google_compute_firewall" "prod_test_cassandra_stg_service" {
  allow = {
    ports    = ["9042", "9160"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-test-cassandra-stg-service"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["test-cassandra-stg-service"]
  target_tags    = ["test-cassandra-stg-service"]
}

resource "google_compute_firewall" "prod_topics_events_cassandra_stg_cluster" {
  allow = {
    ports    = ["7001", "7000", "7199", "23456"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["7000", "7001", "7199"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-topics-events-cassandra-stg-cluster"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["topics-events-cassandra-stg-cluster"]
  target_tags    = ["topics-events-cassandra-stg-cluster"]
}

resource "google_compute_firewall" "prod_topics_events_cassandra_stg_service" {
  allow = {
    ports    = ["9042", "9160"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "prod-topics-events-cassandra-stg-service"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["topics-events-cassandra-stg-service"]
  target_tags    = ["topics-events-cassandra-stg-service"]
}

resource "google_compute_firewall" "puppet_for_staging_qa" {
  allow = {
    ports    = ["8140"]
    protocol = "tcp"
  }

  description    = "Allow Puppet access from staging QA hosts"
  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "puppet-for-staging-qa"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["104.155.3.206/32"]
  target_tags    = ["staging-puppet-server"]
}

resource "google_compute_firewall" "rampart_ssh" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  description    = "used for rampart access"
  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "rampart-ssh"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["0.0.0.0/0"]
  target_tags    = ["rampart"]
}

resource "google_compute_firewall" "realtime_proxy_allow_83" {
  allow = {
    ports    = ["83"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "realtime-proxy-allow-83"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["realtime-proxy-83-client"]
  target_tags    = ["realtime-proxy-83"]
}

resource "google_compute_firewall" "sergey_test_redis" {
  allow = {
    protocol = "all"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "sergey-test-redis"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["10.154.0.0/20", "10.164.0.0/20", "10.128.0.0/20", "10.138.0.0/20", "10.158.0.0/20", "10.146.0.0/20", "10.156.0.0/20", "10.148.0.0/20", "10.166.0.0/20", "10.168.0.0/20", "10.140.0.0/20", "10.150.0.0/20", "10.160.0.0/20", "10.132.0.0/20", "10.170.0.0/20", "10.142.0.0/20", "10.152.0.0/20", "10.162.0.0/20", "10.130.0.0/20"]
}

resource "google_compute_firewall" "ssh" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "ssh"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["74.125.59.22/32", "31.154.2.234/32", "216.239.55.84/31", "54.197.192.34/32", "66.102.14.16/30", "72.14.224.0/21", "216.239.55.85/32", "130.211.71.75/32", "216.239.45.0/24", "172.20.0.0/16", "104.134.21.0/24", "104.132.78.0/24", "74.125.116.0/22", "212.179.82.64/28", "104.132.36.64/27", "216.239.55.188/31", "74.125.59.17/32", "54.195.80.116/32", "104.132.52.0/24", "216.239.55.189/32", "212.179.82.66/31", "31.154.8.66/32", "31.154.8.68/30", "212.179.82.74/32", "74.125.56.128/29", "31.154.8.65/32", "54.197.192.19/32", "31.154.8.70/32", "74.125.56.129/32", "74.125.120.0/22", "54.197.192.35/32", "176.106.227.66/32", "104.132.154.0/24", "104.132.34.64/27", "54.195.80.117/32", "66.102.14.24/30", "31.154.8.64/28", "74.125.56.132/31", "216.239.55.42/31", "74.125.61.227/32", "216.239.33.60/30", "216.239.35.0/24", "54.247.177.108/32", "104.132.51.0/24"]
}

resource "google_compute_firewall" "ssh_22" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  description    = "port 22 for SSH"
  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "ssh-22"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/routing-qa-network"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["31.154.2.234/32", "216.239.55.84/31", "54.197.192.34/32", "66.102.14.16/30", "54.84.123.13/32", "72.14.224.0/21", "104.155.3.206/32", "216.239.55.85/32", "130.211.71.75/32", "216.239.45.0/24", "172.20.0.0/16", "74.125.116.0/22", "212.179.82.64/28", "216.239.55.188/31", "54.195.80.116/32", "104.132.52.0/24", "216.239.55.189/32", "212.179.82.66/31", "31.154.8.66/32", "31.154.8.68/30", "212.179.82.74/32", "74.125.56.128/29", "31.154.8.65/32", "74.125.56.129/32", "74.125.120.0/22", "54.197.192.35/32", "54.164.36.196/32", "104.132.34.64/27", "54.195.80.117/32", "66.102.14.24/30", "31.154.8.64/28", "74.125.56.132/31", "54.197.192.80/32", "216.239.55.42/31", "74.125.61.227/32", "216.239.33.60/30", "216.239.35.0/24", "54.247.177.108/32"]
}

resource "google_compute_firewall" "ssh_from_il_waze_aws" {
  allow = {
    ports    = ["22"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "ssh-from-il-waze-aws"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["54.84.123.13/32", "52.16.237.68/32", "54.197.221.118/32", "52.49.194.217/32", "54.195.80.7/32", "79.125.17.183/32"]
  target_tags    = ["ssh-from-il-waze-aws"]
}

resource "google_compute_firewall" "stackstorm_google_a" {
  allow = {
    ports    = ["444"]
    protocol = "tcp"
  }

  description    = "Access to StackStorm POC from Google-A only"
  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "stackstorm-google-a"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_tags    = ["row-stackstorm"]
  target_tags    = ["row-stackstorm"]
}

resource "google_compute_firewall" "stunnel" {
  allow = {
    ports    = ["9161"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["30001"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["30000"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["30002"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["9162"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["9163"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["30003"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["30006"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["30007"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["30005"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["30004"]
    protocol = "tcp"
  }

  allow = {
    ports    = ["30008"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "stunnel"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["104.134.21.13/32", "74.125.122.33", "104.134.21.26/32", "104.134.21.19/32", "104.134.21.14/32", "31.154.8.70"]
  target_tags    = ["stunnel"]
}

resource "google_compute_firewall" "suportool_testing" {
  allow = {
    ports    = ["443"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "suportool-testing"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["104.134.21.0/24", "74.125.122.33", "54.195.80.116", "54.195.80.117"]
  target_tags    = ["supportool-testing"]
}

resource "google_compute_firewall" "test" {
  allow = {
    protocol = "all"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "test"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["0.0.0.0/0"]
  target_tags    = ["test"]
}

resource "google_compute_firewall" "ttsgw" {
  allow = {
    ports    = ["80"]
    protocol = "tcp"
  }

  direction      = "INGRESS"
  disabled       = false
  enable_logging = false
  name           = "ttsgw"
  network        = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority       = "1000"
  project        = "waze-development"
  source_ranges  = ["130.211.0.0/22", "35.191.0.0/16"]
  target_tags    = ["ttsgw"]
}
