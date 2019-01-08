provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_target_ssl_proxy" "realtime_proxy_tp_ssl" {
  backend_service  = "https://www.googleapis.com/compute/v1/projects/waze-development/global/backendServices/realtime-proxy-stg-be"
  name             = "realtime-proxy-tp-ssl"
  project          = "waze-development"
  proxy_header     = "PROXY_V1"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/waze-development/global/sslCertificates/gcpwazestg2019"]
}
