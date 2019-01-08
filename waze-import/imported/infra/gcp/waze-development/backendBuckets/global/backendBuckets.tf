provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_backend_bucket" "stg_voice_prompts_bucket" {
  bucket_name = "waze-voice"
  enable_cdn  = true
  name        = "stg-voice-prompts-bucket"
  project     = "waze-development"
}

resource "google_compute_backend_bucket" "waze_ads_resources_test" {
  bucket_name = "waze-ads-resources-test"
  enable_cdn  = false
  name        = "waze-ads-resources-test"
  project     = "waze-development"
}

resource "google_compute_backend_bucket" "waze_carpool_groups_images_stg" {
  bucket_name = "waze-carpool-groups-images-stg"
  enable_cdn  = true
  name        = "waze-carpool-groups-images-stg"
  project     = "waze-development"
}

resource "google_compute_backend_bucket" "waze_dirt_dategiver" {
  bucket_name = "waze-dirt-dategiver"
  enable_cdn  = true
  name        = "waze-dirt-dategiver"
  project     = "waze-development"
}

resource "google_compute_backend_bucket" "waze_dirt_dategiver1" {
  bucket_name = "waze-dirt-dategiver1"
  enable_cdn  = false
  name        = "waze-dirt-dategiver1"
  project     = "waze-development"
}

resource "google_compute_backend_bucket" "waze_editor_staging_gcs" {
  bucket_name = "editor.gcp.wazestg.com"
  enable_cdn  = true
  name        = "waze-editor-staging-gcs"
  project     = "waze-development"
}

resource "google_compute_backend_bucket" "waze_web_staging_gcs" {
  bucket_name = "www.gcp.wazestg.com"
  enable_cdn  = true
  name        = "waze-web-staging-gcs"
  project     = "waze-development"
}
