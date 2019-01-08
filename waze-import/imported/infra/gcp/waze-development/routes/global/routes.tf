provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_route" "cassandra_on_k8s_wazestg_c_00e0c6a7_a2dd_11e8_bdac_42010a840002" {
  description            = "k8s-node-route"
  dest_range             = "100.96.2.0/24"
  name                   = "cassandra-on-k8s-wazestg-c-00e0c6a7-a2dd-11e8-bdac-42010a840002"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-b/instances/nodes-77pj"
  next_hop_instance_zone = "europe-west1-b"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "cassandra_on_k8s_wazestg_c_bd85b483_a2dc_11e8_bdac_42010a840002" {
  description            = "k8s-node-route"
  dest_range             = "100.96.0.0/24"
  name                   = "cassandra-on-k8s-wazestg-c-bd85b483-a2dc-11e8-bdac-42010a840002"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-b/instances/master-europe-west1-b-kvb5"
  next_hop_instance_zone = "europe-west1-b"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "cassandra_on_k8s_wazestg_c_fe22f239_a2dc_11e8_bdac_42010a840002" {
  description            = "k8s-node-route"
  dest_range             = "100.96.1.0/24"
  name                   = "cassandra-on-k8s-wazestg-c-fe22f239-a2dc-11e8-bdac-42010a840002"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-b/instances/nodes-94rj"
  next_hop_instance_zone = "europe-west1-b"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "default_route_189da4192ebd3164" {
  description = "Default local route to the subnetwork 10.154.0.0/20."
  dest_range  = "10.154.0.0/20"
  name        = "default-route-189da4192ebd3164"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_27582a7da223781f" {
  description = "Default local route to the subnetwork 10.150.0.0/20."
  dest_range  = "10.150.0.0/20"
  name        = "default-route-27582a7da223781f"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_2c702df28cf900a4" {
  description = "Default local route to the subnetwork 10.156.0.0/20."
  dest_range  = "10.156.0.0/20"
  name        = "default-route-2c702df28cf900a4"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_2ec431d2124831ca" {
  description = "Default local route to the subnetwork 10.128.0.0/20."
  dest_range  = "10.128.0.0/20"
  name        = "default-route-2ec431d2124831ca"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_38820a5f3c879bfa" {
  description      = "Default route to the Internet."
  dest_range       = "0.0.0.0/0"
  name             = "default-route-38820a5f3c879bfa"
  network          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/routing-qa-network"
  next_hop_gateway = "https://www.googleapis.com/compute/v1/projects/waze-development/global/gateways/default-internet-gateway"
  priority         = "1000"
  project          = "waze-development"
}

resource "google_compute_route" "default_route_3e8189e287c400ff" {
  description = "Default local route to the subnetwork 10.160.0.0/20."
  dest_range  = "10.160.0.0/20"
  name        = "default-route-3e8189e287c400ff"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_4082486f8c482793" {
  description = "Default route to the virtual network."
  dest_range  = "10.240.0.0/16"
  name        = "default-route-4082486f8c482793"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/routing-qa-network"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_5197e155b9a28697" {
  description = "Default local route to the subnetwork 10.132.0.0/20."
  dest_range  = "10.132.0.0/20"
  name        = "default-route-5197e155b9a28697"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_5430cc5983aa2fd6" {
  description = "Default local route to the subnetwork 10.162.0.0/20."
  dest_range  = "10.162.0.0/20"
  name        = "default-route-5430cc5983aa2fd6"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_72d1c76cb09f423c" {
  description = "Default local route to the subnetwork 10.164.0.0/20."
  dest_range  = "10.164.0.0/20"
  name        = "default-route-72d1c76cb09f423c"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_7934dab84f38a9a3" {
  description = "Default local route to the subnetwork 10.148.0.0/20."
  dest_range  = "10.148.0.0/20"
  name        = "default-route-7934dab84f38a9a3"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_82f2218f58cb44f2" {
  description      = "Default route to the Internet."
  dest_range       = "0.0.0.0/0"
  name             = "default-route-82f2218f58cb44f2"
  network          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  next_hop_gateway = "https://www.googleapis.com/compute/v1/projects/waze-development/global/gateways/default-internet-gateway"
  priority         = "1000"
  project          = "waze-development"
}

resource "google_compute_route" "default_route_887c5fada992909b" {
  description = "Default route to the virtual network."
  dest_range  = "10.240.0.0/16"
  name        = "default-route-887c5fada992909b"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_90d85b4a4b5374df" {
  description      = "Default route to the Internet."
  dest_range       = "0.0.0.0/0"
  name             = "default-route-90d85b4a4b5374df"
  network          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_gateway = "https://www.googleapis.com/compute/v1/projects/waze-development/global/gateways/default-internet-gateway"
  priority         = "1000"
  project          = "waze-development"
}

resource "google_compute_route" "default_route_9ae63a7ded3e0c5c" {
  description = "Default local route to the subnetwork 10.170.0.0/20."
  dest_range  = "10.170.0.0/20"
  name        = "default-route-9ae63a7ded3e0c5c"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_a3ce467de904700d" {
  description = "Default local route to the subnetwork 10.140.0.0/20."
  dest_range  = "10.140.0.0/20"
  name        = "default-route-a3ce467de904700d"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_bec559788fc1a802" {
  description = "Default local route to the subnetwork 10.152.0.0/20."
  dest_range  = "10.152.0.0/20"
  name        = "default-route-bec559788fc1a802"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_c20621a6deb270bb" {
  description = "Default local route to the subnetwork 10.146.0.0/20."
  dest_range  = "10.146.0.0/20"
  name        = "default-route-c20621a6deb270bb"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_c5050c99834a226e" {
  description = "Default local route to the subnetwork 10.142.0.0/20."
  dest_range  = "10.142.0.0/20"
  name        = "default-route-c5050c99834a226e"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_cb8f7d4bc03cf3a9" {
  description = "Default local route to the subnetwork 10.130.0.0/20."
  dest_range  = "10.130.0.0/20"
  name        = "default-route-cb8f7d4bc03cf3a9"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_d9af6f4ec5dbcb26" {
  description = "Default local route to the subnetwork 10.166.0.0/20."
  dest_range  = "10.166.0.0/20"
  name        = "default-route-d9af6f4ec5dbcb26"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_e261d2f1f1a19aa3" {
  description = "Default local route to the subnetwork 10.138.0.0/20."
  dest_range  = "10.138.0.0/20"
  name        = "default-route-e261d2f1f1a19aa3"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_e877b23714d53dc4" {
  description      = "Default route to the Internet."
  dest_range       = "0.0.0.0/0"
  name             = "default-route-e877b23714d53dc4"
  network          = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/restricted"
  next_hop_gateway = "https://www.googleapis.com/compute/v1/projects/waze-development/global/gateways/default-internet-gateway"
  priority         = "1000"
  project          = "waze-development"
}

resource "google_compute_route" "default_route_ef19dc7b3fd21d2d" {
  description = "Default local route to the subnetwork 10.168.0.0/20."
  dest_range  = "10.168.0.0/20"
  name        = "default-route-ef19dc7b3fd21d2d"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_f8e6809461766946" {
  description = "Default local route to the subnetwork 10.158.0.0/20."
  dest_range  = "10.158.0.0/20"
  name        = "default-route-f8e6809461766946"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/cassandra-on-k8s"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "default_route_fa1f80e967e23987" {
  description = "Default route to the virtual network."
  dest_range  = "10.240.0.0/16"
  name        = "default-route-fa1f80e967e23987"
  network     = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/restricted"
  priority    = "1000"
  project     = "waze-development"
}

resource "google_compute_route" "gke_adman2_stg_1ce64671_367e6c9c_e8cc_11e8_bd2f_42010a840102" {
  description            = "k8s-node-route"
  dest_range             = "10.64.1.0/24"
  name                   = "gke-adman2-stg-1ce64671-367e6c9c-e8cc-11e8-bd2f-42010a840102"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-adman2-stg-pool3-f0c2f23b-09b8"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_adman2_stg_1ce64671_89076974_b51d_11e8_b879_42010a8401af" {
  description            = "k8s-node-route"
  dest_range             = "10.64.2.0/24"
  name                   = "gke-adman2-stg-1ce64671-89076974-b51d-11e8-b879-42010a8401af"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-adman2-stg-pool3-f0c2f23b-33pd"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_adman2_stg_1ce64671_b265f6d9_9cc8_11e8_afcb_42010a84009d" {
  description            = "k8s-node-route"
  dest_range             = "10.64.0.0/24"
  name                   = "gke-adman2-stg-1ce64671-b265f6d9-9cc8-11e8-afcb-42010a84009d"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-adman2-stg-pool3-f0c2f23b-0080"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_adman2_stg_ng_d67c5900_2219c117_0042_11e9_b748_42010a84012d" {
  description            = "k8s-node-route"
  dest_range             = "10.40.1.0/24"
  name                   = "gke-adman2-stg-ng-d67c5900-2219c117-0042-11e9-b748-42010a84012d"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-b/instances/gke-adman2-stg-ng-default-pool-da0543ca-r376"
  next_hop_instance_zone = "europe-west1-b"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_adman2_stg_ng_d67c5900_49185a48_e4d5_11e8_a0f2_42010a8400e1" {
  description            = "k8s-node-route"
  dest_range             = "10.40.2.0/24"
  name                   = "gke-adman2-stg-ng-d67c5900-49185a48-e4d5-11e8-a0f2-42010a8400e1"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-adman2-stg-ng-default-pool-c5692e52-vmj6"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_adman2_stg_ng_d67c5900_c76142de_e919_11e8_9543_42010a8400ca" {
  description            = "k8s-node-route"
  dest_range             = "10.40.0.0/24"
  name                   = "gke-adman2-stg-ng-d67c5900-c76142de-e919-11e8-9543-42010a8400ca"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-c/instances/gke-adman2-stg-ng-default-pool-2b3c1d16-1dnq"
  next_hop_instance_zone = "europe-west1-c"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_help_b3fef0b0_ef74f3c8_56b4_11e8_8654_42010af0029a" {
  description            = "k8s-node-route"
  dest_range             = "10.44.0.0/24"
  name                   = "gke-help-b3fef0b0-ef74f3c8-56b4-11e8-8654-42010af0029a"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-help-default-pool-651aaec2-mg2n"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_help_b3fef0b0_ef7c911f_56b4_11e8_8654_42010af0029a" {
  description            = "k8s-node-route"
  dest_range             = "10.44.1.0/24"
  name                   = "gke-help-b3fef0b0-ef7c911f-56b4-11e8-8654-42010af0029a"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-help-default-pool-651aaec2-4t6t"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_help_b3fef0b0_f078e663_56b4_11e8_8654_42010af0029a" {
  description            = "k8s-node-route"
  dest_range             = "10.44.2.0/24"
  name                   = "gke-help-b3fef0b0-f078e663-56b4-11e8-8654-42010af0029a"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-help-default-pool-651aaec2-b5rz"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_maint_jobs_stg_5325aa7_367f1faf_9a4d_11e8_8c34_42010a840027" {
  description            = "k8s-node-route"
  dest_range             = "10.57.6.0/24"
  name                   = "gke-maint-jobs-stg-5325aa7-367f1faf-9a4d-11e8-8c34-42010a840027"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-maint-jobs-stg-n4-pool-7d89cf24-2q47"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_maint_jobs_stg_5325aa7_aa3fcaaa_9a4d_11e8_8c34_42010a840027" {
  description            = "k8s-node-route"
  dest_range             = "10.57.7.0/24"
  name                   = "gke-maint-jobs-stg-5325aa7-aa3fcaaa-9a4d-11e8-8c34-42010a840027"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-maint-jobs-stg-n4-pool-7d89cf24-frc1"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_maint_jobs_stg_5325aa7_ba25bb03_1362_11e9_9e78_42010a84018e" {
  description            = "k8s-node-route"
  dest_range             = "10.58.13.0/24"
  name                   = "gke-maint-jobs-stg-5325aa7-ba25bb03-1362-11e9-9e78-42010a84018e"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-maint-jobs-stg-n4-pool-7d89cf24-9dck"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_maint_jobs_stg_5325aa7_d4803e5b_9a4c_11e8_8c34_42010a840027" {
  description            = "k8s-node-route"
  dest_range             = "10.57.5.0/24"
  name                   = "gke-maint-jobs-stg-5325aa7-d4803e5b-9a4c-11e8-8c34-42010a840027"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-maint-jobs-stg-n4-pool-7d89cf24-0rmw"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_stg_tools_9fc715a0_5c224ad7_0943_11e9_ba0c_42010a840023" {
  description            = "k8s-node-route"
  dest_range             = "10.16.0.0/24"
  name                   = "gke-stg-tools-9fc715a0-5c224ad7-0943-11e9-ba0c-42010a840023"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-stg-tools-default-pool-f5023440-3mtf"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_stg_tools_9fc715a0_5cb6867d_0943_11e9_ba0c_42010a840023" {
  description            = "k8s-node-route"
  dest_range             = "10.16.1.0/24"
  name                   = "gke-stg-tools-9fc715a0-5cb6867d-0943-11e9-ba0c-42010a840023"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-stg-tools-default-pool-f5023440-lqx1"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_stg_tools_9fc715a0_5cc5ee6d_0943_11e9_ba0c_42010a840023" {
  description            = "k8s-node-route"
  dest_range             = "10.16.2.0/24"
  name                   = "gke-stg-tools-9fc715a0-5cc5ee6d-0943-11e9-ba0c-42010a840023"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-stg-tools-default-pool-f5023440-zlbv"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_us_central1_tiles_buil_3dccac5a_e80e_11e8_96df_42010a80005c" {
  description            = "k8s-node-route"
  dest_range             = "10.4.1.0/24"
  name                   = "gke-us-central1-tiles-buil-3dccac5a-e80e-11e8-96df-42010a80005c"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/us-central1-f/instances/gke-us-central1-tiles-bu-default-pool-1aff38e0-pj8w"
  next_hop_instance_zone = "us-central1-f"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_us_central1_tiles_buil_adcdb8c0_e80d_11e8_96df_42010a80005c" {
  description            = "k8s-node-route"
  dest_range             = "10.4.0.0/24"
  name                   = "gke-us-central1-tiles-buil-adcdb8c0-e80d-11e8-96df-42010a80005c"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/us-central1-f/instances/gke-us-central1-tiles-bu-default-pool-1aff38e0-gll0"
  next_hop_instance_zone = "us-central1-f"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_us_central1_tiles_buil_cd296838_e80e_11e8_96df_42010a80005c" {
  description            = "k8s-node-route"
  dest_range             = "10.4.2.0/24"
  name                   = "gke-us-central1-tiles-buil-cd296838-e80e-11e8-96df-42010a80005c"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/us-central1-f/instances/gke-us-central1-tiles-bu-default-pool-1aff38e0-vq0b"
  next_hop_instance_zone = "us-central1-f"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_waze_tools_stg_cb44831_c40d3324_0c3e_11e9_8398_42010a840019" {
  description            = "k8s-node-route"
  dest_range             = "10.20.0.0/24"
  name                   = "gke-waze-tools-stg-cb44831-c40d3324-0c3e-11e9-8398-42010a840019"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-waze-tools-stg-default-pool-4fb5d594-lt0v"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_waze_tools_stg_cb44831_c41003cc_0c3e_11e9_8398_42010a840019" {
  description            = "k8s-node-route"
  dest_range             = "10.20.1.0/24"
  name                   = "gke-waze-tools-stg-cb44831-c41003cc-0c3e-11e9-8398-42010a840019"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-waze-tools-stg-default-pool-4fb5d594-zg4l"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_waze_tools_stg_cb44831_c415cf10_0c3e_11e9_8398_42010a840019" {
  description            = "k8s-node-route"
  dest_range             = "10.20.2.0/24"
  name                   = "gke-waze-tools-stg-cb44831-c415cf10-0c3e-11e9-8398-42010a840019"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-d/instances/gke-waze-tools-stg-default-pool-4fb5d594-3x9m"
  next_hop_instance_zone = "europe-west1-d"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_waze_tools_stg_cb44831_c417e521_0c3e_11e9_8398_42010a840019" {
  description            = "k8s-node-route"
  dest_range             = "10.20.3.0/24"
  name                   = "gke-waze-tools-stg-cb44831-c417e521-0c3e-11e9-8398-42010a840019"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-b/instances/gke-waze-tools-stg-default-pool-7cbe39a4-44c2"
  next_hop_instance_zone = "europe-west1-b"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_waze_tools_stg_cb44831_c419a71e_0c3e_11e9_8398_42010a840019" {
  description            = "k8s-node-route"
  dest_range             = "10.20.4.0/24"
  name                   = "gke-waze-tools-stg-cb44831-c419a71e-0c3e-11e9-8398-42010a840019"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-b/instances/gke-waze-tools-stg-default-pool-7cbe39a4-4680"
  next_hop_instance_zone = "europe-west1-b"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_waze_tools_stg_cb44831_c419fd0b_0c3e_11e9_8398_42010a840019" {
  description            = "k8s-node-route"
  dest_range             = "10.20.5.0/24"
  name                   = "gke-waze-tools-stg-cb44831-c419fd0b-0c3e-11e9-8398-42010a840019"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-b/instances/gke-waze-tools-stg-default-pool-7cbe39a4-w1dx"
  next_hop_instance_zone = "europe-west1-b"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_waze_tools_stg_cb44831_c48effad_0c3e_11e9_8398_42010a840019" {
  description            = "k8s-node-route"
  dest_range             = "10.20.6.0/24"
  name                   = "gke-waze-tools-stg-cb44831-c48effad-0c3e-11e9-8398-42010a840019"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-c/instances/gke-waze-tools-stg-default-pool-966685b3-js58"
  next_hop_instance_zone = "europe-west1-c"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_waze_tools_stg_cb44831_c4e76aff_0c3e_11e9_8398_42010a840019" {
  description            = "k8s-node-route"
  dest_range             = "10.20.7.0/24"
  name                   = "gke-waze-tools-stg-cb44831-c4e76aff-0c3e-11e9-8398-42010a840019"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-c/instances/gke-waze-tools-stg-default-pool-966685b3-lp5g"
  next_hop_instance_zone = "europe-west1-c"
  priority               = "1000"
  project                = "waze-development"
}

resource "google_compute_route" "gke_waze_tools_stg_cb44831_c59b21c0_0c3e_11e9_8398_42010a840019" {
  description            = "k8s-node-route"
  dest_range             = "10.20.8.0/24"
  name                   = "gke-waze-tools-stg-cb44831-c59b21c0-0c3e-11e9-8398-42010a840019"
  network                = "https://www.googleapis.com/compute/v1/projects/waze-development/global/networks/default"
  next_hop_instance      = "projects/waze-development/zones/europe-west1-c/instances/gke-waze-tools-stg-default-pool-966685b3-k4sc"
  next_hop_instance_zone = "europe-west1-c"
  priority               = "1000"
  project                = "waze-development"
}
