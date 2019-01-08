provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_network" "cassandra_on_k8s" {
  auto_create_subnetworks = true
  name                    = "cassandra-on-k8s"
  project                 = "waze-development"
  routing_mode            = "GLOBAL"
}

resource "google_compute_network" "default" {
  auto_create_subnetworks = false
  description             = "Default network for the project"
  ipv4_range              = "10.240.0.0/16"
  name                    = "default"
  project                 = "waze-development"
  routing_mode            = "REGIONAL"
}

resource "google_compute_network" "restricted" {
  auto_create_subnetworks = false
  ipv4_range              = "10.240.0.0/16"
  name                    = "restricted"
  project                 = "waze-development"
  routing_mode            = "REGIONAL"
}

resource "google_compute_network" "routing_qa_network" {
  auto_create_subnetworks = false
  description             = "Allow port 8080 from IL office"
  ipv4_range              = "10.240.0.0/16"
  name                    = "routing-qa-network"
  project                 = "waze-development"
  routing_mode            = "REGIONAL"
}
