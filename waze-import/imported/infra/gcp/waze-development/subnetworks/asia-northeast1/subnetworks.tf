provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_subnetwork" "cassandra_on_k8s" {
  enable_flow_logs         = false
  ip_cidr_range            = "10.146.0.0/20"
  name                     = "cassandra-on-k8s"
  network                  = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  private_ip_google_access = false
  project                  = "waze-development"
  region                   = "asia-northeast1"
}
