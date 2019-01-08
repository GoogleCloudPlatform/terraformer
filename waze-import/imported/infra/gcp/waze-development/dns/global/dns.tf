provider "google" {
  project = ""
  region  = ""
}

resource "google_dns_managed_zone" "funl" {
  description = "IFS Domain"
  dns_name    = "fu.nl."
  labels      = {}
  name        = "funl"
  project     = "waze-development"
}

resource "google_dns_managed_zone" "mapathon" {
  description = "Mapathon Domain"
  dns_name    = "waze-mapping.com."
  labels      = {}
  name        = "mapathon"
  project     = "waze-development"
}

resource "google_dns_managed_zone" "wazestg" {
  description = "Waze Staging Domain"
  dns_name    = "wazestg.com."
  labels      = {}
  name        = "wazestg"
  project     = "waze-development"
}

resource "google_dns_record_set" "6ci2gcpdmstz-gcp-wazestg-com-_CNAME" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "6ci2gcpdmstz.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["gv-3bt62pv7lzt4cj.dv.googlehosted.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "_validate_domain-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "_validate_domain.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["_validate_domain.pki.goog."]
  ttl          = "300"
  type         = "CNAME"
}

resource "google_dns_record_set" "_validate_domain-wazestg-com-_CNAME" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "_validate_domain.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["_validate_domain.pki.goog."]
  ttl          = "300"
  type         = "CNAME"
}

resource "google_dns_record_set" "adjust-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "adjust.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["146.148.24.225"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "admanage-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "admanage.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.216.95"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "admanage_int-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "admanage-int.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.81"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "admanage_proxy-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "admanage-proxy.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.235.74"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "admanage_stg12-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "admanage-stg12.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.235.115"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "admanage_stg14-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "admanage-stg14.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.138.224"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "admanage_stg19-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "admanage-stg19.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.138.224"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "admanage_stg30-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "admanage-stg30.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.214.219"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "admanage_stg5-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "admanage-stg5.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.198.1"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "admanage_stg7-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "admanage-stg7.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.198.1"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "admiral-gcp-wazestg-com-_CNAME" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "admiral.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["ghs.googlehosted.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "adorn-gcp-wazestg-com-_CNAME" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "adorn.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["ghs.googlehosted.com."]
  ttl          = "300"
  type         = "CNAME"
}

resource "google_dns_record_set" "ads-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "ads.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["146.148.24.225"]
  ttl          = "900"
  type         = "A"
}

resource "google_dns_record_set" "ads_memcached_stg1-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "ads-memcached-stg1.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.1.14"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "adsapi_proxy-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "adsapi-proxy.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.227.250.129"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "adsassets-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "adsassets.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.54.247"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "adslb-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "adslb.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.246.185"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "advil-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "advil.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["146.148.24.225"]
  ttl          = "900"
  type         = "A"
}

resource "google_dns_record_set" "advil-wazestg-com-_CNAME" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "advil.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["ghs.googlehosted.com."]
  ttl          = "300"
  type         = "CNAME"
}

resource "google_dns_record_set" "advil_gae-gcp-wazestg-com-_CNAME" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "advil-gae.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["ghs.googlehosted.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "alerts-apigateway-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "alerts.apigateway.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["q1j5c7tg9j.execute-api.us-east-1.amazonaws.com."]
  ttl          = "300"
  type         = "CNAME"
}

resource "google_dns_record_set" "api-cassandra_on_k8s-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "api.cassandra-on-k8s.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.195.139.117"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "api-internal-cassandra_on_k8s-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "api.internal.cassandra-on-k8s.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.132.0.2"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "apt_ci-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "apt-ci.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.132.0.14"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "apt_ci_gcp-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "apt-ci-gcp.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.132.0.14"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "apt_na_aws-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "apt-na-aws.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["172.30.55.205"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "apt_prd_gcp-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "apt-prd-gcp.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.183.123"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "apt_prod_ext_gcp-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "apt-prod-ext-gcp.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["104.196.24.234"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "apt_prod_gcp-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "apt-prod-gcp.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["172.20.0.5"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "apt_row_aws-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "apt-row-aws.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["172.31.24.215"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "apt_row_ext_aws-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "apt-row-ext-aws.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["54.247.177.108"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "apt_stg_gcp-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "apt-stg-gcp.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.65.186"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "ateam-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "ateam.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.89"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "biz-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "biz.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["130.211.10.102"]
  ttl          = "900"
  type         = "A"
}

resource "google_dns_record_set" "biz_dev_eladm-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "biz-dev-eladm.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["146.148.24.225"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "biz_dev_elanh-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "biz-dev-elanh.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["146.148.24.225"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "biz_dev_sagie-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "biz-dev-sagie.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["146.148.24.225"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "biz_exp-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "biz-exp.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["130.211.10.102"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "bticons-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "bticons.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["54.195.80.50"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "bugger-apigateway-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "bugger.apigateway.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["mhpddqp2zg.execute-api.us-east-1.amazonaws.com."]
  ttl          = "300"
  type         = "CNAME"
}

resource "google_dns_record_set" "carpool_groups_images-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "carpool-groups-images.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.92.68"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "carpool_memcached_stg1-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "carpool-memcached-stg1.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.11.64", "10.240.11.65", "10.240.11.61", "10.240.11.62", "10.240.11.60", "10.240.11.33"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "carpool_test_ipc-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "carpool-test-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "carpoolindex_memcached_stg_stg1-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "carpoolindex-memcached-stg-stg1.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.1.106", "10.240.1.108"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "carpoolpayments-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "carpoolpayments.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["130.211.45.79"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "carpooltesting_ipc-gcp-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "carpooltesting-ipc.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "60"
  type         = "TXT"
}

resource "google_dns_record_set" "carpooltesting_rt-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "carpooltesting-rt.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.250.88"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "client_assets-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "client-assets.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.248.105"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "client_tiles_stg14-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "client-tiles-stg14.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.241.3.158"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "clienttiles_stg12-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "clienttiles-stg12.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.214.219"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "clienttiles_stg18-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "clienttiles-stg18.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.211.77"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "clienttiles_stg19-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "clienttiles-stg19.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.235.115"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "clienttiles_stg30-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "clienttiles-stg30.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.147.220"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "clienttiles_stg5-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "clienttiles-stg5.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.232.86"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "clienttiles_stg7-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "clienttiles-stg7.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.232.86"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "community-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "community.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.248.105"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "cost-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "cost.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.20.152.210"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "darth_www-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "darth-www.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.248.105"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "descartes-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "descartes.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.185"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "descartes_dev-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "descartes-dev.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.183"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "descartes_dev1-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "descartes-dev1.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["104.155.14.44"]
  ttl          = "3600"
  type         = "A"
}

resource "google_dns_record_set" "dev-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "dev.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.248.105"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "dev_editor1-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "dev-editor1.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.246.163"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "ecs-fu-nl-_NS" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "ecs.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["ns-1061.awsdns-04.org.", "ns-907.awsdns-49.net.", "ns-1796.awsdns-32.co.uk.", "ns-238.awsdns-29.com."]
  ttl          = "300"
  type         = "NS"
}

resource "google_dns_record_set" "editor_tiles-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "editor-tiles.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.248.132"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "elton-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "elton.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.61.146"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "embed-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "embed.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.248.105"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "etcd_b-internal-cassandra_on_k8s-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "etcd-b.internal.cassandra-on-k8s.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.132.0.2"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "etcd_events_b-internal-cassandra_on_k8s-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "etcd-events-b.internal.cassandra-on-k8s.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.132.0.2"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "eu_ext_feed-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "eu-ext-feed.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["internal-vpc-eu-ext-feed-1683973286.eu-west-1.elb.amazonaws.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "eudynmon2-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "eudynmon2.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["54.195.80.8"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_b-il_ipc-stg21-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-b.il-ipc.stg21.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.41.0.81"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_b-il_ipc-stg30-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-b.il-ipc.stg30.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.50.0.91"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_b-il_ipc-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-b.il-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.98"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_b-il_ipc_offline-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-b.il-ipc-offline.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.1.55"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_c-il_ipc-stg15-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-c.il-ipc.stg15.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.35.0.14"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_c-il_ipc-stg21-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-c.il-ipc.stg21.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.41.0.80"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_c-il_ipc-stg25-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-c.il-ipc.stg25.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.45.0.53"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_c-il_ipc-stg30-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-c.il-ipc.stg30.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.50.0.96", "10.50.0.84"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_c-il_ipc-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-c.il-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.76"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_c-il_ipc_offline-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-c.il-ipc-offline.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.1.3"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_c-offline_ipc-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-c.offline-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.109.44"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_d-carpool_test_ipc-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-d.carpool-test-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.47"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_d-il_ipc-stg21-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-d.il-ipc.stg21.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.41.0.86"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_d-il_ipc-stg30-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-d.il-ipc.stg30.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.50.0.86", "10.50.0.83"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_d-il_ipc-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-d.il-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.1.16"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_d-il_ipc_offline-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-d.il-ipc-offline.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.1.31"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_d-il_test_ipc-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-d.il-test-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.220.42"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_d-topics_test_ipc-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-d.topics-test-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.11"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "europe_west1_d-users_test_eu-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "europe-west1-d.users-test-eu.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.149.8"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "ext_feed-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "ext-feed.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["internal-vpc-na-ext-feed-1360841090.us-east-1.elb.amazonaws.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "ext_feed-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "ext-feed.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["130.211.104.242"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "ext_feed_int-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "ext-feed-int.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["172.30.144.226"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "feed-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "feed.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["130.211.104.242"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "fluentd_forwarder_ng-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "fluentd-forwarder-ng.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.129"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "fu-nl-_NS" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["ns-cloud-c1.googledomains.com.", "ns-cloud-c2.googledomains.com.", "ns-cloud-c3.googledomains.com.", "ns-cloud-c4.googledomains.com."]
  ttl          = "21600"
  type         = "NS"
}

resource "google_dns_record_set" "fu-nl-_SOA" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["ns-cloud-c1.googledomains.com. dns-admin.google.com. 7 21600 3600 1209600 300"]
  ttl          = "21600"
  type         = "SOA"
}

resource "google_dns_record_set" "gcedev-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "gcedev.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["104.155.91.136"]
  ttl          = "30"
  type         = "A"
}

resource "google_dns_record_set" "gcp-wazestg-com-_CNAME" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "*.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["www.gcp.wazestg.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "gcp-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"google-site-verification=I3p5FeM_q-iSVPIM8uEGIGycDv0X1Igudf1ON_3g2UI\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "general_memcached_stg1-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "general-memcached-stg1.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.1.40", "10.240.1.46"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "georegistry_tcp_lb-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "georegistry-tcp-lb.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.97"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "georss-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "georss.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.211.26"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "grr_aws_na-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "grr-aws-na.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["internal-gagentproxy-prod-na-1173360320.us-east-1.elb.amazonaws.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "grr_aws_row-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "grr-aws-row.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["internal-gagentproxy-prod-row-728869596.eu-west-1.elb.amazonaws.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "grr_dev-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "grr-dev.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.240"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "grr_prod-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "grr-prod.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["172.21.1.214"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "il-waze_mapping-com-_NS" {
  managed_zone = "${google_dns_managed_zone.mapathon.name}"
  name         = "il.${google_dns_managed_zone.mapathon.dns_name}"
  project      = "waze-development"
  rrdatas      = ["ns-1754.awsdns-27.co.uk.", "ns-782.awsdns-33.net.", "ns-442.awsdns-55.com.", "ns-1087.awsdns-07.org."]
  ttl          = "300"
  type         = "NS"
}

resource "google_dns_record_set" "il_dev2-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-dev2.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["23.236.49.102"]
  ttl          = "900"
  type         = "A"
}

resource "google_dns_record_set" "il_ipc-stg10-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg10.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg11-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg11.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg12-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg12.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg13-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg13.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg14-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg14.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg15-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg15.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg16-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg16.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg17-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg17.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg18-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg18.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg19-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg19.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg2-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg2.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg20-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg20.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg21-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg21.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg22-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg22.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg23-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg23.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg24-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg24.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg25-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg25.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg26-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg26.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg27-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg27.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg28-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg28.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg29-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg29.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg3-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg3.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg30-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg30.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg4-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg4.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg5-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg5.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg6-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg6.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg7-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg7.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg8-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg8.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-stg9-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.stg9.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "60"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg10-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg10.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg11-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg11.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg12-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg12.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg13-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg13.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg14-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg14.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg15-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg15.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg16-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg16.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg17-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg17.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg18-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg18.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg19-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg19.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg2-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg2.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg20-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg20.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg21-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg21.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg22-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg22.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg23-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg23.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg24-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg24.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg25-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg25.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg26-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg26.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg27-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg27.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg28-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg28.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg29-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg29.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg3-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg3.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg30-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg30.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg4-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg4.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg5-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg5.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg6-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg6.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg7-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg7.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg8-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg8.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-stg9-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.stg9.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_ipc_offline-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-ipc-offline.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-b,europe-west1-c,europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "il_test_ipc-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "il-test-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "importer-aws-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "importer.aws.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["176.34.107.182"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "importer-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "importer.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["176.34.107.182"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "inbox-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "inbox.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["130.211.15.36"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "incidents-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "incidents.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["130.211.10.69"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "it-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "it.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["146.148.15.90"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "it_internal-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "it-internal.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["172.21.0.9"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "it_internal_proxy-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "it-internal-proxy.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["172.20.0.22"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "it_testing-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "it-testing.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["104.199.15.138"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "login-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "login.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.248.219"]
  ttl          = "900"
  type         = "A"
}

resource "google_dns_record_set" "login2-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "login2.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["104.155.20.24"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "luke_www-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "luke-www.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.248.105"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "m-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "m.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["internal-securitymonkey-internal-2003649720.eu-west-1.elb.amazonaws.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "ma-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "ma.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["54.197.192.120"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "map-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "map.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["row-monitoring-644266387.eu-west-1.elb.amazonaws.com."]
  ttl          = "300"
  type         = "CNAME"
}

resource "google_dns_record_set" "mapnik_livemap_stg12-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "mapnik-livemap-stg12.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.138.224"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "mapnik_livemap_stg19-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "mapnik-livemap-stg19.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.147.220"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "mapnik_livemap_stg5-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "mapnik-livemap-stg5.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.214.219"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "mapnik_livemap_stg7-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "mapnik-livemap-stg7.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.190.159"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "marx-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "marx.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["34.227.163.153"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "memcache_test-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcache-test.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.47", "10.240.0.45", "10.240.0.46"]
  ttl          = "0"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg10-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg10.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.30.0.24"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg13-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg13.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.33.0.32"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg15-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg15.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.35.0.32"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg16-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg16.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.36.0.64"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg17-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg17.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.37.0.29"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg18-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg18.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.38.0.64"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg19-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg19.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.39.0.42"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg2-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg2.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.22.0.24"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg23-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg23.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.43.0.26"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg25-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg25.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.45.0.66"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg28-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg28.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.48.0.27"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg30-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg30.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.50.0.100"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg4-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg4.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.24.0.22"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg5-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg5.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.25.0.44"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg6-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg6.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.26.0.18"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached-stg7-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached.stg7.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.27.0.22"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached_stg14-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached-stg14.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.34.0.102"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached_stg15-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached-stg15.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.35.0.78"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached_stg18-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached-stg18.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.38.0.67"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached_stg21-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached-stg21.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.41.0.106"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached_stg3-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached-stg3.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.23.0.68"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached_stg30-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached-stg30.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.50.0.105"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached_stg5-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached-stg5.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.25.0.110"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcached_stg7-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcached-stg7.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.27.0.110"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "memcachetest_memcached_il-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "memcachetest-memcached-il.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.214", "10.240.0.218", "10.240.0.217"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "metrics-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "metrics.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.130"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "metrics_stg12-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "metrics-stg12.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.241.42.52"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "metrics_stg14-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "metrics-stg14.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.207.240"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "metrics_stg15-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "metrics-stg15.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.198.1"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "metrics_stg18-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "metrics-stg18.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.138.224"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "metrics_stg19-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "metrics-stg19.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.202.182"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "metrics_stg21-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "metrics-stg21.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.241.3.158"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "metrics_stg3-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "metrics-stg3.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.251.142"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "metrics_stg30-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "metrics-stg30.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.198.1"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "metrics_stg5-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "metrics-stg5.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.235.115"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "metrics_stg7-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "metrics-stg7.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.211.77"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "mobile_web-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "mobile-web.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.247.127"]
  ttl          = "900"
  type         = "A"
}

resource "google_dns_record_set" "mobile_web_proxy-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "mobile-web-proxy.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.204.44"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "mon-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "mon.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["row-monitoring-644266387.eu-west-1.elb.amazonaws.com."]
  ttl          = "300"
  type         = "CNAME"
}

resource "google_dns_record_set" "mpa-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "mpa.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["104.196.28.175"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "ms-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "ms.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["54.195.80.10"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "na_routing_map_beta-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "na-routing-map-beta.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["vpc-na-routing-map-beta-459670838.us-east-1.elb.amazonaws.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "nsscache_aws_na-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "nsscache-aws-na.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["internal-gagentproxy-prod-na-1173360320.us-east-1.elb.amazonaws.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "nsscache_aws_row-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "nsscache-aws-row.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["internal-gagentproxy-prod-row-728869596.eu-west-1.elb.amazonaws.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "nsscache_dev-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "nsscache-dev.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.240"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "nsscache_jumpserver_row-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "nsscache-jumpserver-row.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["internal-gagentproxy-prod-row-728869596.eu-west-1.elb.amazonaws.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "nsscache_prod-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "nsscache-prod.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["172.21.1.214", "172.20.1.1"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "nsscache_prod_na-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "nsscache-prod-na.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["172.20.1.1"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "offline-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "offline.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["130.211.101.94"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "offline_ipc-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "offline-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-c\""]
  ttl          = "60"
  type         = "TXT"
}

resource "google_dns_record_set" "omertest-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "omertest.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["internal-test-me-1646034577.eu-central-1.elb.amazonaws.com."]
  ttl          = "5"
  type         = "CNAME"
}

resource "google_dns_record_set" "omertest-wazestg-com-_CNAME" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "omertest.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["internal-test-me-1646034577.eu-central-1.elb.amazonaws.com."]
  ttl          = "5"
  type         = "CNAME"
}

resource "google_dns_record_set" "packages_aws_row-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "packages-aws-row.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["internal-packages-prod-row-916221865.eu-west-1.elb.amazonaws.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "parking-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "parking.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["23.251.135.239"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "parkingmanager-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "parkingmanager.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.1.223"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "profiles-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "profiles.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.66.51"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "profiles_int-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "profiles-int.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.181"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "prompto-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "prompto.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.46.127"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "pu-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "pu.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.20.152.210"]
  ttl          = "900"
  type         = "A"
}

resource "google_dns_record_set" "r2d2_www-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "r2d2-www.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.248.105"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "realtime_frontend_2-stg4-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "realtime-frontend-2.stg4.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.12.54"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "realtime_frontend_3-stg4-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "realtime-frontend-3.stg4.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.221.123"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "realtime_frontend_4-stg4-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "realtime-frontend-4.stg4.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.61.146"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "realtime_proxy_stg12-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "realtime-proxy-stg12.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.158.203"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "realtime_proxy_stg19-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "realtime-proxy-stg19.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.207.240"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "realtime_proxy_stg3-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "realtime-proxy-stg3.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.241.42.52"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "regulator_cassandra-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "regulator-cassandra.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.94"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "repository_server-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "repository-server.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.253.253"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "repository_server-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "repository-server.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.146"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "routing_livemap-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "routing-livemap.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["130.211.28.144"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "row_staging_ads_db1-wazestg-com-_CNAME" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "row-staging-ads-db1.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.144."]
  ttl          = "300"
  type         = "CNAME"
}

resource "google_dns_record_set" "row_staging_general12_cassandra-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "row-staging-general12-cassandra.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.212"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "row_staging_general21_cassandra-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "row-staging-general21-cassandra.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.221"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "rr-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "rr.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["104.155.111.242"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "rt-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "rt.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["130.211.10.205"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "rtproxy-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "rtproxy.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.240.186"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "rtproxy_stg12-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "rtproxy-stg12.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.224.29"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "rtproxy_stg3-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "rtproxy-stg3.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.3.12"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "sashatest_memcached_stg1-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "sashatest-memcached-stg1.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.54"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "saw-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "saw.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.112"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "sd_metrics-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "sd-metrics.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.254.86"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "sd_policies-fu-nl-_A" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "sd-policies.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.254.86"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "search-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "search.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.201.127.199"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg10-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg10.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.241.52.1"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg13-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg13.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.235.74"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg15-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg15.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.235.74"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg16-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg16.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.241.3.158"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg17-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg17.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.62.154"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg18-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg18.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.202.246"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg19-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg19.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.241.52.1"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg2-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg2.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.235.74"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg23-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg23.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.241.28.10"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg25-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg25.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.241.3.158"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg28-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg28.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.235.74"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg30-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg30.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.241.3.158"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg4-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg4.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.68.18"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg5-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg5.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.62.248"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg6-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg6.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.201.112.211"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver-stg7-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver.stg7.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.217.166"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver_stg12-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver-stg12.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.3.12"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver_stg14-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver-stg14.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.168.238"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver_stg15-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver-stg15.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.232.86"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver_stg18-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver-stg18.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.204.44"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver_stg19-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver-stg19.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.214.190"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver_stg21-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver-stg21.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.147.220"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver_stg3-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver-stg3.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.197.123"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver_stg30-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver-stg30.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.207.240"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver_stg5-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver-stg5.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.138.224"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "searchserver_stg7-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "searchserver-stg7.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.244.204.44"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "servermanager-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "servermanager.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.120"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "servermanager_proxy-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "servermanager-proxy.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.47.68"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "servermanager_proxy_new-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "servermanager-proxy-new.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.47.68"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "sms-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "sms.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.77.126"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "staging_memcached_stg1-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "staging-memcached-stg1.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.78"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "staging_mmc-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "staging-mmc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.57", "10.240.0.55", "10.240.0.62", "10.240.0.56"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "stats-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "stats.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["vpc-row-stats-362171766.eu-west-1.elb.amazonaws.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "stg11_ipc-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "stg11-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-c\""]
  ttl          = "60"
  type         = "TXT"
}

resource "google_dns_record_set" "stg1_ipc-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "stg1-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-c\""]
  ttl          = "60"
  type         = "TXT"
}

resource "google_dns_record_set" "stg2_ipc-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "stg2-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-c\""]
  ttl          = "60"
  type         = "TXT"
}

resource "google_dns_record_set" "storagegateway_internal-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "storagegateway-internal.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.123"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "storagegateway_stg-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "storagegateway-stg.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.201.81.162"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "test_iap-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "test-iap.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.236.77"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "test_memcache-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "test-memcache.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.44", "10.240.0.49", "10.240.0.45", "10.240.0.53"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "test_memcache_ext-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "test-memcache-ext.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["130.211.55.182", "192.158.28.80", "146.148.21.66"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "testing_http2-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "testing-http2.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["testing-http2-removeme-2058491143.eu-west-1.elb.amazonaws.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "testrecord-stg4-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "testrecord.stg4.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.0.0.101"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "tf-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "tf.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["104.196.18.150"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "tiles-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "tiles.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.201.122.10"]
  ttl          = "900"
  type         = "A"
}

resource "google_dns_record_set" "topic-fu-nl-_CNAME" {
  managed_zone = "${google_dns_managed_zone.funl.name}"
  name         = "topic.${google_dns_managed_zone.funl.dns_name}"
  project      = "waze-development"
  rrdatas      = ["internal-topic-monitoring-web-elb-692155018.eu-west-1.elb.amazonaws.com."]
  ttl          = "60"
  type         = "CNAME"
}

resource "google_dns_record_set" "topicagent-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "topicagent.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.205.77.44"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "topics_test_ipc-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "topics-test-ipc.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "ttsgw-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "ttsgw.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.254.116"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "ultron-gcp-wazestg-com-_CNAME" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "ultron.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["ghs.googlehosted.com."]
  ttl          = "300"
  type         = "CNAME"
}

resource "google_dns_record_set" "us_central1_f-users_test-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "us-central1-f.users-test.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.66.3"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "users_test-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "users-test.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"us-central1-f\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "users_test_eu-wazestg-com-_TXT" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "users-test-eu.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["\"europe-west1-d\""]
  ttl          = "300"
  type         = "TXT"
}

resource "google_dns_record_set" "venues_live_memcached_stg1-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "venues-live-memcached-stg1.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.0.145"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "venues_snapshot_memcached_stg1-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "venues-snapshot-memcached-stg1.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["10.240.1.51"]
  ttl          = "60"
  type         = "A"
}

resource "google_dns_record_set" "voice_prompts-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "voice-prompts.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.186.231.245"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "was-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "was.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["35.190.78.117"]
  ttl          = "900"
  type         = "A"
}

resource "google_dns_record_set" "waze_mapping-com-_NS" {
  managed_zone = "${google_dns_managed_zone.mapathon.name}"
  name         = "${google_dns_managed_zone.mapathon.dns_name}"
  project      = "waze-development"
  rrdatas      = ["ns-cloud-b1.googledomains.com.", "ns-cloud-b2.googledomains.com.", "ns-cloud-b3.googledomains.com.", "ns-cloud-b4.googledomains.com."]
  ttl          = "3600"
  type         = "NS"
}

resource "google_dns_record_set" "waze_mapping-com-_SOA" {
  managed_zone = "${google_dns_managed_zone.mapathon.name}"
  name         = "${google_dns_managed_zone.mapathon.dns_name}"
  project      = "waze-development"
  rrdatas      = ["ns-cloud-b1.googledomains.com. dns-admin.google.com. 0 21600 3600 1209600 300"]
  ttl          = "21600"
  type         = "SOA"
}

resource "google_dns_record_set" "wazestg-com-_NS" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["ns-cloud-c1.googledomains.com.", "ns-cloud-c2.googledomains.com.", "ns-cloud-c3.googledomains.com.", "ns-cloud-c4.googledomains.com."]
  ttl          = "21600"
  type         = "NS"
}

resource "google_dns_record_set" "wazestg-com-_SOA" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["ns-cloud-c1.googledomains.com. dns-admin.google.com. 41 21600 3600 1209600 300"]
  ttl          = "21600"
  type         = "SOA"
}

resource "google_dns_record_set" "www-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "www.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.248.105"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "yoda_www-gcp-wazestg-com-_A" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "yoda-www.gcp.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["107.178.248.105"]
  ttl          = "300"
  type         = "A"
}

resource "google_dns_record_set" "zcxtcjy77ojy-wazestg-com-_CNAME" {
  managed_zone = "${google_dns_managed_zone.wazestg.name}"
  name         = "zcxtcjy77ojy.${google_dns_managed_zone.wazestg.dns_name}"
  project      = "waze-development"
  rrdatas      = ["gv-tbfgu36ltxtzme.dv.googlehosted.com."]
  ttl          = "300"
  type         = "CNAME"
}
