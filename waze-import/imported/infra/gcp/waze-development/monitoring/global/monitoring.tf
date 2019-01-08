provider "google" {
  project = ""
  region  = ""
}

resource "google_monitoring_alert_policy" "projects--waze_development--alertPolicies--10624087474126344002" {
  combiner = "OR"

  conditions = {
    condition_threshold = {
      aggregations = {
        alignment_period     = "3600s"
        cross_series_reducer = "REDUCE_MEAN"
        group_by_fields      = ["metric.label.env"]
        per_series_aligner   = "ALIGN_MEAN"
      }

      comparison      = "COMPARISON_GT"
      duration        = "3600s"
      filter          = "metric.type=\"custom.googleapis.com/testing/waze/ads_adapt/Ads_Adapt_missing_entities_in_poi_world\" AND resource.type=\"aws_ec2_instance\""
      threshold_value = "0"

      trigger = {
        count   = "1"
        percent = "0"
      }
    }

    display_name = "STG-Ads-Adapt-Poi-World-Missing-entities"
  }

  display_name          = "Ads-Adapt-Missing-Entities-STG"
  enabled               = true
  notification_channels = ["projects/waze-development/notificationChannels/12740028274668241490"]
  project               = "waze-development"
}

resource "google_monitoring_alert_policy" "projects--waze_development--alertPolicies--18020651424460027856" {
  combiner = "OR"

  conditions = {
    condition_threshold = {
      aggregations = {
        alignment_period     = "3600s"
        cross_series_reducer = "REDUCE_MEAN"
        group_by_fields      = ["metric.label.probe"]
        per_series_aligner   = "ALIGN_MEAN"
      }

      comparison      = "COMPARISON_LT"
      duration        = "3600s"
      filter          = "project=\"waze-development\" AND metric.type=\"custom.googleapis.com/testing/waze/ads_backend_probers/exit_status\" AND resource.type=\"aws_ec2_instance\""
      threshold_value = "1"

      trigger = {
        count   = "1"
        percent = "0"
      }
    }

    display_name = "Dev-Ads-Backend-Prober-Failures"
  }

  display_name          = "Ads-Backend-Prober-Failures-Dev"
  enabled               = true
  notification_channels = ["projects/waze-development/notificationChannels/10651237603377009513"]
  project               = "waze-development"
}

resource "google_monitoring_alert_policy" "projects--waze_development--alertPolicies--267780096609570373" {
  combiner = "OR"

  conditions = {
    condition_threshold = {
      aggregations = {
        alignment_period   = "120s"
        group_by_fields    = ["metric.label.probe"]
        per_series_aligner = "ALIGN_MEAN"
      }

      comparison      = "COMPARISON_LT"
      duration        = "120s"
      filter          = "project=\"waze-development\" AND metric.type=\"custom.googleapis.com/testing/waze/probers/exit_status\" AND resource.type=\"aws_ec2_instance\""
      threshold_value = "0.5"

      trigger = {
        count   = "1"
        percent = "0"
      }
    }

    display_name = "EditLoopProber failure on STG"
  }

  display_name          = "Dev-Prober-Failures"
  enabled               = true
  notification_channels = ["projects/waze-development/notificationChannels/13714067585895474016"]
  project               = "waze-development"
}

resource "google_monitoring_alert_policy" "projects--waze_development--alertPolicies--2903152574431104318" {
  combiner = "OR"

  conditions = {
    condition_threshold = {
      aggregations = {
        alignment_period     = "120s"
        cross_series_reducer = "REDUCE_MAX"
        group_by_fields      = ["metric.label.env"]
        per_series_aligner   = "ALIGN_MEAN"
      }

      comparison      = "COMPARISON_GT"
      duration        = "120s"
      filter          = "metric.type=\"custom.googleapis.com/testing/waze/ads_server/Adapt_imported_model_age_in_minutes\" AND resource.type=\"aws_ec2_instance\""
      threshold_value = "60"

      trigger = {
        count   = "1"
        percent = "0"
      }
    }

    display_name = "STG-Ads-External-Poi-Model-Age"
  }

  display_name          = "Ads-External-Poi-Model-Age-STG"
  enabled               = true
  notification_channels = ["projects/waze-development/notificationChannels/10651237603377009513"]
  project               = "waze-development"
}

resource "google_monitoring_alert_policy" "projects--waze_development--alertPolicies--3958240469237705677" {
  combiner = "OR"

  conditions = {
    condition_threshold = {
      aggregations = {
        alignment_period   = "60s"
        per_series_aligner = "ALIGN_MEAN"
      }

      comparison      = "COMPARISON_GT"
      duration        = "600s"
      filter          = "project=\"waze-development\" AND metric.type=\"agent.googleapis.com/disk/percent_used\" AND resource.type=\"gce_instance\" AND metric.label.state=\"used\""
      threshold_value = "85"

      trigger = {
        count   = "1"
        percent = "0"
      }
    }

    display_name = "Disk-Usage-Development"
  }

  display_name          = "Disk-Usage-Development-API"
  enabled               = true
  notification_channels = ["projects/waze-development/notificationChannels/3447124078559702958"]
  project               = "waze-development"
}

resource "google_monitoring_alert_policy" "projects--waze_development--alertPolicies--5292193251185668978" {
  combiner = "OR"

  conditions = {
    condition_threshold = {
      aggregations = {
        alignment_period     = "600s"
        cross_series_reducer = "REDUCE_MIN"
        group_by_fields      = ["metric.label.env"]
        per_series_aligner   = "ALIGN_MIN"
      }

      comparison      = "COMPARISON_LT"
      duration        = "600s"
      filter          = "metric.type=\"custom.googleapis.com/testing/waze/ads_adapt/Ads_Adapt_poi_world_export_fail\" AND resource.type=\"aws_ec2_instance\""
      threshold_value = "0"

      trigger = {
        count   = "1"
        percent = "0"
      }
    }

    display_name = "STG-Ads-Adapt-Export-Failures"
  }

  display_name          = "Ads-Adapt-Export-Failures-STG"
  enabled               = true
  notification_channels = ["projects/waze-development/notificationChannels/12740028274668241490"]
  project               = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1001295190389753471" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic40-15-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic40-15-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1009846671644319685" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic5-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic5-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1016512" {
  display_name = "users-service"
  filter       = "resource.metadata.name=has_substring(\"am-users\")"
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1020124902375353511" {
  display_name = "mGroup-topic-test-hrsanity-113046-nostalgy-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-113046-nostalgy-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1025685520066682861" {
  display_name = "mGroup-topic-test-clustermanytopics-012932-topic45-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012932-topic45-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1028271408447330752" {
  display_name = "mGroup-topic-test-dailysanity-084734-sanitytopic1-sub"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-084734-sanitytopic1-sub\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--10351" {
  display_name = "RealTime"
  filter       = "resource.metadata.name=has_substring(\"row-staging-rt\") OR resource.metadata.name=has_substring(\"realtime-\")"
  is_cluster   = true
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1070210066333535729" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic25-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic25-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1074814046532227228" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190121-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190121-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1080343915076125089" {
  display_name = "mGroup-topic-test-cgtest-130102-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130102-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1080647628358219733" {
  display_name = "mGroup-topic-test-hrrandpubs-060034-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-060034-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--108957959202467135" {
  display_name = "mGroup-topic-test-cgtest-190129-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1091314146111223633" {
  display_name = "mGroup-topic-test-cgtest-040105-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-040105-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1097999503895417248" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic46-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic46-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1100439628880842801" {
  display_name = "mGroup-topic-test-cnsmrgrptest-094641-cgtopic-testcg1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-094641-cgtopic-testcg1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1107004895156358386" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic35-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic35-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1108031702904771022" {
  display_name = "mGroup-topic-test-cgtest-190140-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190140-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1119046317447967988" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic31-13-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic31-13-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1122508588843117351" {
  display_name = "mGroup-topic-test-dailysanity-190130-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190130-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--112510352168806695" {
  display_name = "mGroup-topic-test-cgtest-190138-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190138-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--114217514723473688" {
  display_name = "mGroup-topic-test-cgtest-190134-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190134-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1148982329035177778" {
  display_name = "mGroup-topic-test-hrrandpubs-020026-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020026-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1180947322502789376" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic48-16-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic48-16-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1184653792414544470" {
  display_name = "mGroup-topic-test-hrrandpubs-020025-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020025-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1189412639122576264" {
  display_name = "mGroup-topic-test-dailysanity-190131-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190131-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--120968779974575911" {
  display_name = "mGroup-topic-test-multi-topics-102522-topic25-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-102522-topic25-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1216854017296824375" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1231651903730177854" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic17-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic17-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1232836982929573465" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic43-15-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic43-15-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1235371028454544937" {
  display_name = "mGroup-topic-test-dailysanity-190139-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190139-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1244538002241683679" {
  display_name = "mGroup-topic-test-cgtest-190109-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190109-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1252437619424901139" {
  display_name = "mGroup-topic-test-dailysanity-190137-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190137-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1265790347625043712" {
  display_name = "mGroup-topic-test-clustermanytopics-012932-topic20-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012932-topic20-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1266013500750296" {
  display_name = "mGroup-topic-test-hrresize-040028-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040028-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1270680573923010753" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic14-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic14-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--127826199979696806" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic15-18-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic15-18-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1284891392206099532" {
  display_name = "mGroup-topic-test-dailysanity-190133-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190133-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1304315445805315994" {
  display_name = "mGroup-topic-test-cgpubsadd-080553-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgpubsadd-080553-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--130493224320034609" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic37-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic37-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1309143368435966327" {
  display_name = "mGroup-topic-test-dailysanity-090105-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-090105-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1309611625795713566" {
  display_name = "mGroup-topic-test-dailysanity-190126-sanitytopic2-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190126-sanitytopic2-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1322029255092868966" {
  display_name = "mGroup-topic-test-dailysanity-190120-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190120-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1341451737308490821" {
  display_name = "mGroup-topic-test-hrsanity-233853-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-233853-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1348233898554250460" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1351833596580603907" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic9-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic9-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1362538086520199223" {
  display_name = "mGroup-topic-test-cgtest-130109-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130109-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1363055845231059773" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic21-19-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic21-19-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1364860169180746535" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic37-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic37-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1376334951319803806" {
  display_name = "mGroup-topic-test-hrresize-040059-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040059-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1384134707290149260" {
  display_name = "mGroup-topic-test-cgtest-190122-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190122-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1389362959531212959" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic46-21-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic46-21-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1391087433039518291" {
  display_name = "mGroup-topic-test-hrrandpubs-020026-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020026-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1407320100670055626" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic26-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic26-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--140893711112967955" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic40-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic40-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1414754130997406173" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic17-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic17-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1423361676889305966" {
  display_name = "mGroup-topic-test-cgtest-090101-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090101-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1425952450450114592" {
  display_name = "mGroup-topic-test-hrsanity-100020-nostalgy-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-100020-nostalgy-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1435577330830429543" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic8-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic8-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1446357266664647710" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic50-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic50-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1491133308794011021" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic40-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic40-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1505396148849068876" {
  display_name = "mGroup-topic-test-clustermanytopics-012932-topic18-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012932-topic18-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1506227982387684638" {
  display_name = "mGroup-topic-test-cgtest-164243-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-164243-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1518004078942038626" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic46-20-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic46-20-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1530958099301381009" {
  display_name = "mGroup-topic-test-cgtest-090101-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090101-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1536836162536368855" {
  display_name = "mGroup-topic-test-dailysanity-190151-sanitytopic1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190151-sanitytopic1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1537563474712251096" {
  display_name = "mGroup-topic-test-hrresize-040027-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040027-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1545921126855207927" {
  display_name = "mGroup-topic-test-cgtest-190129-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1552133151456232602" {
  display_name = "mGroup-topic-test-hrresize-040034-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040034-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1583852573573045385" {
  display_name = "mGroup-topic-test-hrrandpubs-020026-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020026-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1591408154356982606" {
  display_name = "mGroup-topic-test-dailysanity-190119-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190119-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1591673253675030116" {
  display_name = "mGroup-topic-test-shortlongmsgs-190125-long-msg-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-190125-long-msg-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1613477542533664482" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190114-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190114-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1629757926370607764" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic8-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic8-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--163508814557114292" {
  display_name = "mGroup-topic-test-cgtest-130105-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130105-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1635157639235833779" {
  display_name = "xmGroup-cassandra"
  filter       = "resource.metadata.name=has_substring(\"cassandra\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1650373873458421784" {
  display_name = "mGroup-topic-test-dailysanity-190138-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190138-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1650601406731952631" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic43-24-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic43-24-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1653486312471466969" {
  display_name = "mGroup-topic-test-dailysanity-171605-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-171605-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1664309613305682205" {
  display_name = "mGroup-topic-test-dailysanity-090109-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-090109-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1674719110365046477" {
  display_name = "mGroup-topic-test-dailysanity-190137-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190137-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1678868442769010190" {
  display_name = "mGroup-topic-test-dailysanity-190130-sanitytopic2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190130-sanitytopic2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1708433647651619751" {
  display_name = "mGroup-topic-test-cgtest-190136-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190136-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1721030863927379118" {
  display_name = "mGroup-topic-test-dailysanity-190135-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190135-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1724492857301155081" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic15-25-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic15-25-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1729887540818200472" {
  display_name = "mGroup-topic-test-hrresize-040059-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040059-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1734088413229892421" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190129-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190129-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1740353800285621026" {
  display_name = "mGroup-topic-test-dailysanity-190111-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190111-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1745792434128652950" {
  display_name = "mGroup-topic-test-dailysanity-190133-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190133-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1755454969266661739" {
  display_name = "mGroup-topic-test-dailysanity-190151-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190151-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1757175972040202428" {
  display_name = "mGroup-topic-test-highrates-153005-topic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-highrates-153005-topic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1771203181101966054" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic31-29-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic31-29-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1776470507188104427" {
  display_name = "mGroup-topic-test-cgtest-190122-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190122-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1782236921650196321" {
  display_name = "mGroup-topic-test-shortlongmsgs-131745-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-131745-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1807489041321537583" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic39-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic39-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1817747157891370465" {
  display_name = "mGroup-topic-test-dailysanity-190120-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190120-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1832134366823198956" {
  display_name = "mGroup-topic-test-hrresize-040026-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040026-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1838876398104419232" {
  display_name = "mGroup-topic-test-dailysanity-190134-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190134-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--183897950783889259" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic37-25-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic37-25-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--183947896320235991" {
  display_name = "mGroup-topic-test-dailysanity-082300-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-082300-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1843058705487785013" {
  display_name = "mGroup-topic-test-longmsgs-190121-longtopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-190121-longtopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1845108253553446475" {
  display_name = "mGroup-topic-test-hrresize-040029-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040029-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--185786871373015611" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic7-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic7-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1858413929239482568" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic34-33-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic34-33-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1889109218544436684" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic34-26-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic34-26-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1893040940529153913" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic24-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic24-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1897983366983769286" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic6-30-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic6-30-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1903447570730980660" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic22-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic22-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1907804167334210260" {
  display_name = "mGroup-topic-test-cgtest-190135-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190135-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1919484125869408076" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190137-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190137-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1923754217251070963" {
  display_name = "mGroup-topic-test-cgtest-190116-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190116-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1933975316225566403" {
  display_name = "mGroup-topic-test-dailysanity-182312-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-182312-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1937894445239264838" {
  display_name = "mGroup-topic-test-hrsanity-204616-nostalgy-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-204616-nostalgy-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--194545257233186378" {
  display_name = "mGroup-topic-test-cgtest-163042-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-163042-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1955309263911046886" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic6-13-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic6-13-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1956409492501131689" {
  display_name = "mGroup-topic-test-clustermanytopics-150820-topic2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-150820-topic2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1971748499100920819" {
  display_name = "mGroup-topic-test-hrresize-020034-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020034-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--1992733695641721461" {
  display_name = "mGroup-topic-test-cnsmrgrptest-094641-cgtopic-testcg3-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-094641-cgtopic-testcg3-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2008776068488171794" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic42-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic42-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2012078455667937839" {
  display_name = "mGroup-topic-test-cgtest-190132-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190132-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2017907126951508711" {
  display_name = "mGroup-topic-test-cgtest-190109-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190109-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2019857523795503444" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic33-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic33-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2025169679580370264" {
  display_name = "mGroup-topic-test-dailysanity-190105-sanitytopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190105-sanitytopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--203641356565602" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic23-29-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic23-29-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2053577679664959546" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic17-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic17-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2060141899741859895" {
  display_name = "mGroup-topic-test-hrsanity-160725-nostalgy-hrs-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-160725-nostalgy-hrs-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2074781145661486622" {
  display_name = "mGroup-topic-test-hrresize-040025-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040025-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--208420351168373922" {
  display_name = "mGroup-topic-test-hrresize-040027-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040027-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2088151305006022878" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190144-cgtopic-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190144-cgtopic-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2093680707066914074" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic22-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic22-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2104491455652431907" {
  display_name = "mGroup-topic-test-cgtest-190129-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2138226048455987866" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic39-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic39-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2140250920923584703" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2142141713871799848" {
  display_name = "mGroup-topic-test-dailysanity-190121-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190121-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2145042186889948836" {
  display_name = "mGroup-topic-test-hrresize-040028-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040028-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2149338622025880790" {
  display_name = "mGroup-topic-test-hrresize-020039-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020039-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--214990020466620894" {
  display_name = "mGroup-topic-test-dailysanity-190121-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190121-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2164824377885777023" {
  display_name = "mGroup-topic-test-cgtest-163042-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-163042-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2168671959901863444" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic42-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic42-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--217573101355570182" {
  display_name = "mGroup-topic-test-highrates-151311-topic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-highrates-151311-topic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2175902049867102335" {
  display_name = "mGroup-topic-test-cgtest-130108-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130108-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2176248492199889366" {
  display_name = "mGroup-topic-test-dailysanity-190115-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190115-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2181898469047286096" {
  display_name = "mGroup-topic-test-cgtest-181648-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-181648-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2185604229014422325" {
  display_name = "mGroup-topic-test-serverrestart-234520-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-serverrestart-234520-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2193197086267664137" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2193544331133013564" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic12-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic12-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2193606564856422775" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic23-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic23-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2194360637359221922" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic3-23-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic3-23-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2199159666357666339" {
  display_name = "mGroup-topic-test-dailysanity-084456-sanitytopic1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-084456-sanitytopic1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2202280291177791389" {
  display_name = "mGroup-topic-test-clustermanytopics-012932-topic49-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012932-topic49-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2203986806270093693" {
  display_name = "mGroup-topic-test-cgtest-163042-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-163042-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2208771938431581359" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic25-17-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic25-17-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--221143839302357525" {
  display_name = "mGroup-topic-test-dailysanity-070117-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-070117-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2213651739078908122" {
  display_name = "mGroup-topic-test-dailysanity-130106-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-130106-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2224680717469883374" {
  display_name = "mGroup-topic-test-dailysanity-190108-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190108-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--22437322369697481" {
  display_name = "mGroup-topic-test-hrresize-040059-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040059-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2250475403894726899" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic36-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic36-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2275221232276776496" {
  display_name = "mGroup-topic-test-cgtest-071717-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-071717-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2285174958066373996" {
  display_name = "mGroup-topic-test-dailysanity-221642-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-221642-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2286150036414694521" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic38-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic38-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2289148699771149939" {
  display_name = "mGroup-topic-test-dailysanity-130106-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-130106-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--229206233917633561" {
  display_name = "mGroup-topic-test-cgtest-163042-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-163042-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2300531904546823219" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic39-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic39-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2300734082035751565" {
  display_name = "mGroup-topic-test-cgtest-190119-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190119-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2306218904151068514" {
  display_name = "mGroup-topic-test-cgtest-071717-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-071717-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--233141369148401409" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic43-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic43-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2337724871664554889" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190144-cgtopic-testcg2-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190144-cgtopic-testcg2-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2338372050186355216" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic31-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic31-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2349838945816772395" {
  display_name = "mGroup-topic-test-cgtest-131808-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-131808-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2358006121654180059" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic26-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic26-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2378674381525498909" {
  display_name = "mGroup-topic-test-dailysanity-130109-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-130109-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2388032926166033151" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic24-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic24-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2388263905959186309" {
  display_name = "mGroup-topic-test-cgtest-090101-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090101-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2402794386429231908" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic14-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic14-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2417157605600666832" {
  display_name = "mGroup-topic-test-dailysanity-040110-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-040110-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2417319611187125770" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic21-7-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic21-7-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2418756386077731149" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic37-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic37-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2425771700802637248" {
  display_name = "mGroup-topic-test-shortlongmsgs-111151-short-msg-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-111151-short-msg-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2451191911844725581" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic6-7-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic6-7-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2455459760167610111" {
  display_name = "mGroup-topic-test-cgtest-190120-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190120-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2456407010081780799" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic5-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic5-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2463834515254423232" {
  display_name = "mGroup-topic-test-dailysanity-171605-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-171605-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2474235093859991945" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic43-17-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic43-17-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2481766916754006457" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic45-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic45-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2489005793895120885" {
  display_name = "mGroup-topic-test-dailysanity-130109-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-130109-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2496505899113050635" {
  display_name = "mGroup-topic-test-cgtest-090105-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090105-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2500424633929907063" {
  display_name = "mGroup-topic-test-hrrandpubs-020028-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020028-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2511812708981424414" {
  display_name = "mGroup-topic-test-longmsgs-130105-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-130105-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2513536402891946875" {
  display_name = "mGroup-topic-test-dailysanity-130104-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-130104-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2518131938984503366" {
  display_name = "mGroup-topic-test-cgtest-190129-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2524974085241239966" {
  display_name = "xmGroup-cassandra"
  filter       = "resource.metadata.name=has_substring(\"cassandra\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2537642676725106337" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic13-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic13-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2541375446226358224" {
  display_name = "mGroup-topic-test-multi-topics-182311-topic35-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-182311-topic35-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2543793040785424262" {
  display_name = "mGroup-topic-test-hrresize-040026-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040026-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--255261375756292736" {
  display_name = "mGroup-topic-test-dailysanity-190105-sanitytopic1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190105-sanitytopic1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2558855290135208792" {
  display_name = "mGroup-topic-test-clustermanytopics-012932-topic10-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012932-topic10-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--258975462780886738" {
  display_name = "mGroup-topic-test-hrsanity-202131-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-202131-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2609032636745389133" {
  display_name = "mGroup-topic-test-dailysanity-154450-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-154450-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2618411599914519324" {
  display_name = "mGroup-topic-test-cgtest-190109-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190109-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2622056310736775733" {
  display_name = "mGroup-topic-test-cnsmrgrptest-094641-cgtopic-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-094641-cgtopic-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2626618078018375298" {
  display_name = "mGroup-topic-test-dailysanity-190157-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190157-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2636373761224404869" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic31-32-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic31-32-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2643537165208967915" {
  display_name = "mGroup-topic-test-dailysanity-164245-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-164245-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2657370396092917658" {
  display_name = "mGroup-topic-test-cgtest-190126-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190126-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2659234705081248295" {
  display_name = "mGroup-topic-test-clustermanytopics-012932-topic7-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012932-topic7-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2665802594168946230" {
  display_name = "mGroup-topic-test-cgtest-190134-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190134-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2670106425941005862" {
  display_name = "mGroup-topic-test-dailysanity-130109-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-130109-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2684318193451390005" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic47-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic47-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--269520788477772830" {
  display_name = "mGroup-topic-test-hrresize-040025-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040025-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2705408711762385717" {
  display_name = "mGroup-topic-test-cgtest-130108-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130108-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2708855902525389343" {
  display_name = "mGroup-topic-test-cgtest-130108-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130108-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2724446398827945683" {
  display_name = "mGroup-topic-test-dailysanity-190136-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190136-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2728471271541410561" {
  display_name = "mGroup-topic-test-cgtest-090108-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090108-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2734773558285430686" {
  display_name = "mGroup-topic-test-dailysanity-112850-sanitytopic2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-112850-sanitytopic2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2743722890557161838" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic28-32-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic28-32-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2749920006334897010" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic12-31-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic12-31-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2758850308246567301" {
  display_name = "mGroup-topic-test-cgtest-190132-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190132-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--278081289019642572" {
  display_name = "mGroup-topic-test-cgtest-190116-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190116-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2798587581859607270" {
  display_name = "mGroup-topic-test-hrrandpubs-020025-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020025-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2805137173298429642" {
  display_name = "mGroup-topic-test-hrrandpubs-020025-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020025-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2815996701810769507" {
  display_name = "mGroup-topic-test-dailysanity-190130-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190130-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2823248467707634938" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190130-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190130-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2826965118756351512" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic50-6-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic50-6-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2829895623943807501" {
  display_name = "mGroup-topic-test-cgtest-090108-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090108-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2834528709554509801" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic21-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic21-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2853429289754405508" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic20-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic20-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2859042809509175530" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic16-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic16-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2871805986377438051" {
  display_name = "mGroup-topic-test-dailysanity-190121-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190121-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--287187049493681319" {
  display_name = "mGroup-topic-test-longmsgs-190112-longtopic1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-190112-longtopic1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2905991186534724933" {
  display_name = "mGroup-topic-test-hrsanity-124756-nostalgy-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-124756-nostalgy-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2906669594252909031" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic6-15-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic6-15-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2943111292561261088" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic48-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic48-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2958617758652813942" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic8-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic8-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--297520976583762429" {
  display_name = "mGroup-topic-test-dailysanity-122759-sanitytopic2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-122759-sanitytopic2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2978059682440087916" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic13-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic13-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2982128822307774538" {
  display_name = "mGroup-topic-test-cgtest-090105-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090105-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2993274457581911838" {
  display_name = "mGroup-topic-test-dailysanity-163041-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-163041-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2994535130131691999" {
  display_name = "mGroup-topic-test-cgtest-185933-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-185933-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2995066127752256137" {
  display_name = "mGroup-topic-test-dailysanity-190111-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190111-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--2998614201698172678" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic34-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic34-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3006132270434012969" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic15-26-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic15-26-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3010080259077902743" {
  display_name = "mGroup-topic-test-dailysanity-190126-sanitytopic1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190126-sanitytopic1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3015605898410682517" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic8-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic8-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3017897245885914569" {
  display_name = "mGroup-topic-test-longmsgs-190138-longtopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-190138-longtopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3019753645119114954" {
  display_name = "mGroup-topic-test-clustermanytopics-012932-topic25-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012932-topic25-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--303701432106220979" {
  display_name = "mGroup-topic-test-hrsanity-183000-nostalgy-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-183000-nostalgy-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--303744712584193618" {
  display_name = "mGroup-topic-test-cgtest-190132-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190132-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3037806374633561952" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic14-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic14-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3044412243965846987" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic35-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic35-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--305210420775724752" {
  display_name = "mGroup-topic-test-clustermanytopics-012932-topic17-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012932-topic17-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3061255537215274104" {
  display_name = "mGroup-topic-test-hrrandpubs-020025-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020025-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3067291960419210506" {
  display_name = "mGroup-topic-test-hrrandpubs-020033-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020033-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3067890309297308066" {
  display_name = "mGroup-topic-test-dailysanity-040110-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-040110-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3095525874963650700" {
  display_name = "mGroup-topic-test-shortlongmsgs-030059-short-msg-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-030059-short-msg-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3124938773686019744" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic41-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic41-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3125274343095844063" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic40-29-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic40-29-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3126245776848038592" {
  display_name = "mGroup-topic-test-hrresize-040026-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040026-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3130289987455652477" {
  display_name = "mGroup-topic-test-hrsanity-103425-nostalgy-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-103425-nostalgy-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3147403113210526179" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic43-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic43-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3156578195789722370" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic48-21-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic48-21-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--316249153858498727" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic32-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic32-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3166039265939560135" {
  display_name = "aaa"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-084734-sanitytopic1-publisher\")"
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3170100742039690214" {
  display_name = "mGroup-topic-test-hrrandpubs-020025-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020025-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3177979547386921641" {
  display_name = "mGroup-topic-test-dailysanity-190157-sanitytopic2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190157-sanitytopic2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3181948519572329549" {
  display_name = "mGroup-topic-test-cgtest-221639-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-221639-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3186422663541474235" {
  display_name = "mGroup-topic-test-cgtest-071717-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-071717-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3188853151958444259" {
  display_name = "mGroup-topic-test-hrresize-040059-nostalgy-4-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040059-nostalgy-4-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3197497550633126748" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic5-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic5-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3199462737191205580" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic27-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic27-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3203856121343292430" {
  display_name = "mGroup-topic-test-dailysanity-130146-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-130146-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3207885975151359708" {
  display_name = "mGroup-topic-test-cgtest-071717-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-071717-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3225365016855257854" {
  display_name = "mGroup-topic-test-hrresize-020033-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020033-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3225823283753052045" {
  display_name = "mGroup-topic-test-dailysanity-190151-sanitytopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190151-sanitytopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3234178059380592079" {
  display_name = "mGroup-topic-test-cgtest-130105-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130105-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3238431092701677963" {
  display_name = "mGroup-topic-test-hrresize-020032-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020032-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3244531064188834805" {
  display_name = "mGroup-topic-test-highrates-171234-topic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-highrates-171234-topic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3272181958502319131" {
  display_name = "mGroup-topic-test-cgtest-190125-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190125-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--328386771608496050" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic6-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic6-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3294284959119655523" {
  display_name = "mGroup-topic-test-hrresize-020035-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020035-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3295390126205528838" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190144-cgtopic-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190144-cgtopic-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3302058086212517260" {
  display_name = "mGroup-topic-test-hrresize-040026-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040026-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3303921925088235463" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3311102047885102995" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic29-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic29-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--331269609239899607" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic37-20-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic37-20-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3314384600908055994" {
  display_name = "mGroup-topic-test-hrsanity-180645-nostalgy-hrs-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-180645-nostalgy-hrs-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3317003333735361631" {
  display_name = "mGroup-topic-test-cgtest-181648-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-181648-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3325327013595867096" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic25-12-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic25-12-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3351617657263145775" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic4-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic4-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3364415111012034201" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic27-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic27-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--337532191230828801" {
  display_name = "mGroup-topic-test-cgtest-190129-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3399891134728571935" {
  display_name = "mGroup-topic-test-cgtest-090101-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090101-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3414346225988558542" {
  display_name = "mGroup-topic-test-cgtest-190132-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190132-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3418490609578965476" {
  display_name = "mGroup-topic-test-manypubs-204213-manypubs1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-manypubs-204213-manypubs1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3423260731456343441" {
  display_name = "mGroup-topic-test-cnsmrgrptest-151259-cgtopic-testcg1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-151259-cgtopic-testcg1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3426909132648205608" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic50-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic50-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--342867781901013312" {
  display_name = "mGroup-topic-test-serverrestart-215102-2-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-serverrestart-215102-2-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--346109773548182040" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic40-23-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic40-23-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3471296148320021703" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic18-11-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic18-11-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3475280828852783135" {
  display_name = "mGroup-topic-test-cgtest-130102-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130102-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3494470784440076375" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic9-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic9-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3514836580367255086" {
  display_name = "mGroup-topic-test-cgtest-190134-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190134-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3517731024247332202" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic12-15-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic12-15-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3532134201245304761" {
  display_name = "mGroup-topic-test-cnsmrgrptest-094641-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-094641-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3538991777115781099" {
  display_name = "mGroup-topic-test-shortlongmsgs-030103-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-030103-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--354291183148795906" {
  display_name = "mGroup-topic-test-shortlongmsgs-111151-long-msg-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-111151-long-msg-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3561499700827463990" {
  display_name = "mGroup-topic-test-cgtest-190129-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--356351999639615949" {
  display_name = "mGroup-topic-test-manypubs-005819-manypubs1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-manypubs-005819-manypubs1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3563987598534451466" {
  display_name = "mGroup-topic-test-cgtest-190119-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190119-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3571151107635033393" {
  display_name = "mGroup-topic-test-dailysanity-190108-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190108-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3578631463193950031" {
  display_name = "mGroup-topic-test-longmsgs-230037-longtopic1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-230037-longtopic1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3589353887115635922" {
  display_name = "mGroup-topic-test-cgtest-071717-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-071717-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--359360949165533313" {
  display_name = "mGroup-topic-test-hrrandpubs-020026-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020026-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3599235748064362656" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic3-11-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic3-11-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3609389600704573398" {
  display_name = "mGroup-topic-test-cgtest-190127-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190127-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3618038647193298343" {
  display_name = "mGroup-topic-test-cgtest-121442-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-121442-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3630471005131563304" {
  display_name = "mGroup-topic-test-cgtest-131808-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-131808-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3642064329751149201" {
  display_name = "mGroup-topic-test-cgtest-130105-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130105-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3651400558343742750" {
  display_name = "mGroup-topic-test-hrresize-020034-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020034-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--365372284951639727" {
  display_name = "mGroup-topic-test-hrresize-040040-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040040-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3662452967316914862" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic7-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic7-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3672049669195114949" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic38-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic38-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3690567286443563940" {
  display_name = "mGroup-topic-test-dailysanity-190149-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190149-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3703205468650952196" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic37-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic37-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3737074489052170596" {
  display_name = "mGroup-topic-test-dailysanity-131802-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-131802-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3740228710221525228" {
  display_name = "mGroup-topic-test-hrresize-040028-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040028-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3756429098077606687" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic46-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic46-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--381897553127977583" {
  display_name = "mGroup-topic-test-hrresize-040025-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040025-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3840760222699701282" {
  display_name = "mGroup-topic-test-hrresize-040025-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040025-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--38424402627584153" {
  display_name = "mGroup-topic-test-hrrandpubs-050840-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-050840-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3842762345984893423" {
  display_name = "mGroup-topic-test-dailysanity-190134-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190134-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3856918773261977905" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic47-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic47-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3863674216664068091" {
  display_name = "mGroup-topic-test-cgtest-090108-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090108-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3865883286353152003" {
  display_name = "mGroup-topic-test-longmsgs-230030-longtopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-230030-longtopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3881233832427617057" {
  display_name = "mGroup-topic-test-cgtest-190116-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190116-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3887891587680468615" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic40-31-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic40-31-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3897998925753980308" {
  display_name = "mGroup-topic-test-cgtest-181648-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-181648-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3898829984062079682" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic3-8-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic3-8-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3901406894965368220" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic30-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic30-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3931112304520292405" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic49-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic49-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3940514852475026959" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic7-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic7-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3948098288340690673" {
  display_name = "mGroup-topic-test-cgtest-040105-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-040105-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3948267688455392225" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic43-9-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic43-9-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3964379799467600481" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic26-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic26-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3968131101548888588" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic2-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic2-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--3974582913452740922" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic29-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic29-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4003944744345104598" {
  display_name = "mGroup-topic-test-cgtest-154443-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-154443-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4005314683223457709" {
  display_name = "mGroup-topic-test-manysubs-090415-manysubs1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-manysubs-090415-manysubs1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4010041518115464122" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic16-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic16-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4010352924181157581" {
  display_name = "mGroup-topic-test-hrresize-020037-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020037-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4015203940479819788" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic8-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic8-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4026855165972953010" {
  display_name = "mGroup-topic-test-cgtest-190138-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190138-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4041721203568496283" {
  display_name = "mGroup-topic-test-cgtest-190122-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190122-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4042839852239371669" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic25-19-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic25-19-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4050496941789382301" {
  display_name = "mGroup-topic-test-cgtest-185933-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-185933-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4058677774280221941" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic25-23-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic25-23-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4063053770236421923" {
  display_name = "mGroup-topic-test-serverrestart-183446-2-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-serverrestart-183446-2-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4067735131572474674" {
  display_name = "mGroup-topic-test-dailysanity-070117-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-070117-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4073758752736050862" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic30-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic30-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4076216254669484531" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic18-17-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic18-17-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4076589342457223589" {
  display_name = "mGroup-topic-test-cgtest-190202-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190202-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4083364118824636116" {
  display_name = "mGroup-topic-test-dailysanity-122759-sanitytopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-122759-sanitytopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4084219865616295192" {
  display_name = "mGroup-topic-test-serverrestart-215128-2-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-serverrestart-215128-2-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4097330216600824179" {
  display_name = "mGroup-topic-test-cgtest-190138-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190138-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4100499996601204457" {
  display_name = "mGroup-topic-test-shortlongmsgs-190125-short-msg-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-190125-short-msg-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4102570103269685164" {
  display_name = "mGroup-topic-test-dailysanity-190119-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190119-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4104602927062197894" {
  display_name = "mGroup-topic-test-hrrandpubs-020026-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020026-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4120086565393866801" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--412029814747269907" {
  display_name = "mGroup-topic-test-dailysanity-090105-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-090105-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4126220887861952484" {
  display_name = "mGroup-topic-test-hrsanity-140055-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-140055-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4133106665602267089" {
  display_name = "mGroup-topic-test-cgtest-130108-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130108-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4144019851062447057" {
  display_name = "mGroup-topic-test-hrrandpubs-020025-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020025-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4151222036853241323" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic27-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic27-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4194384529234629321" {
  display_name = "mGroup-topic-test-cgtest-190127-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190127-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4207007446288573289" {
  display_name = "mGroup-topic-test-cgtest-190109-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190109-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4208936776947059127" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic35-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic35-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4228070592696042794" {
  display_name = "mGroup-topic-test-dailysanity-181646-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-181646-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4229381741112801527" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic33-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic33-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4234825626715811342" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic38-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic38-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4241695981040157135" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic47-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic47-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4242545496229442107" {
  display_name = "xmGroup-ca1ssandra"
  filter       = "resource.metadata.name=has_substring(\"ca1ssandra\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4249163560493016956" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic22-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic22-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4249508747384418876" {
  display_name = "mGroup-topic-test-hrresize-040026-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040026-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4267377280864827094" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic4-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic4-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4267377280864829226" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic6-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic6-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4270981144786921840" {
  display_name = "mGroup-topic-test-dailysanity-190135-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190135-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4272276037432367550" {
  display_name = "mGroup-topic-test-hrrandpubs-020028-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020028-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4288095671742303199" {
  display_name = "mGroup-topic-test-longmsgs-190121-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-190121-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4307924070837611810" {
  display_name = "mGroup-topic-test-hrrandpubs-020026-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020026-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4309615126803293752" {
  display_name = "mGroup-topic-test-cgtest-040105-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-040105-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4312843732250188544" {
  display_name = "mGroup-topic-test-cgtest-190120-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190120-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4323224975421903463" {
  display_name = "mGroup-topic-test-cgtest-044255-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-044255-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4374143655756290866" {
  display_name = "mGroup-topic-test-cgtest-130109-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130109-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4376570026340171731" {
  display_name = "mGroup-topic-test-cnsmrgrptest-114049-cgtopic-testcg1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-114049-cgtopic-testcg1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4383445139911630382" {
  display_name = "mGroup-topic-test-longmsgs-154448-longtopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-154448-longtopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4411681758859139113" {
  display_name = "mGroup-topic-test-hrresize-020034-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020034-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4420990419559038881" {
  display_name = "mGroup-topic-test-cgtest-070114-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-070114-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4424730553472728303" {
  display_name = "mGroup-topic-test-cgtest-130102-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130102-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--443143823898483146" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic3-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic3-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4433261169218824026" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--444359132365350777" {
  display_name = "mGroup-topic-test-dailysanity-190138-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190138-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4455126680199375430" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4465023574799465332" {
  display_name = "mGroup-topic-test-multi-topics-112852-topic36-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-112852-topic36-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4465779157122992015" {
  display_name = "mGroup-topic-test-hrrandpubs-020026-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020026-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4471885195098090955" {
  display_name = "mGroup-topic-test-cgtest-190140-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190140-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--447938390963890537" {
  display_name = "mGroup-topic-test-cgtest-190136-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190136-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4486398593582216825" {
  display_name = "mGroup-topic-test-hrresize-040040-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040040-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4490789277420248148" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic34-23-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic34-23-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4490886848187748409" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic24-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic24-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4497057158438451279" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic34-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic34-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4501349165027288525" {
  display_name = "mGroup-topic-test-hrresize-040028-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040028-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4505282092472252996" {
  display_name = "mGroup-topic-test-dailysanity-092701-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-092701-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4517424424771557144" {
  display_name = "mGroup-topic-test-dailysanity-190223-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190223-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4519426366766356965" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190129-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190129-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4525879761330996245" {
  display_name = "mGroup-topic-test-cgtest-190136-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190136-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4537567985548376047" {
  display_name = "mGroup-topic-test-dailysanity-190135-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190135-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4551249705062594068" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic23-16-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic23-16-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4554258972211947922" {
  display_name = "mGroup-topic-test-cgtest-190116-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190116-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4563592130009199408" {
  display_name = "mGroup-topic-test-cgtest-190120-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190120-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4567970073053706580" {
  display_name = "mGroup-topic-test-cgtest-190138-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190138-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4568515004010394015" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic26-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic26-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4582244276978480145" {
  display_name = "mGroup-topic-test-hrresize-040026-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040026-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4583689302662825969" {
  display_name = "mGroup-topic-test-manysubs-073859-manysubs1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-manysubs-073859-manysubs1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4599940620292253100" {
  display_name = "mGroup-topic-test-cgtest-190134-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190134-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4604725268543171699" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic23-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic23-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4605131733074145991" {
  display_name = "mGroup-topic-test-dailysanity-070117-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-070117-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4614681097864104952" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic15-33-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic15-33-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4617847820693215749" {
  display_name = "mGroup-topic-test-cgtest-190135-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190135-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4624642741541873749" {
  display_name = "mGroup-topic-test-cgtest-090108-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090108-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4626039099289837765" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic29-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic29-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4628592020789037951" {
  display_name = "mGroup-topic-test-dailysanity-154450-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-154450-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--463873283797455125" {
  display_name = "mGroup-topic-test-dailysanity-190121-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190121-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4663521372195218554" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic49-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic49-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--468162740279345986" {
  display_name = "mGroup-topic-test-cgtest-190119-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190119-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4697241483126381884" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190144-cgtopic-testcg1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190144-cgtopic-testcg1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4697817538549651999" {
  display_name = "mGroup-topic-test-hrsanity-085212-nostalgy-hrs-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-085212-nostalgy-hrs-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4705135000822952989" {
  display_name = "mGroup-topic-test-cgtest-181648-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-181648-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4738030744520277497" {
  display_name = "mGroup-topic-test-hrresize-040027-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040027-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--47386519360787611" {
  display_name = "mGroup-topic-test-hrsanity-151950-nostalgy-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-151950-nostalgy-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4740510618320033466" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic5-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic5-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4755161986694383132" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic31-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic31-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4758081312443610020" {
  display_name = "mGroup-topic-test-cgtest-221639-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-221639-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4772913672412954896" {
  display_name = "mGroup-topic-test-cgtest-190135-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190135-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4773587510482333804" {
  display_name = "mGroup-topic-test-cnsmrgrptest-165059-cgtopic-testcg3-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-165059-cgtopic-testcg3-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4775339725229417576" {
  display_name = "mGroup-topic-test-cgtest-190125-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190125-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4784622186218068316" {
  display_name = "mGroup-topic-test-cgtest-092702-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-092702-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--479322670607692304" {
  display_name = "mGroup-topic-test-cgtest-190202-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190202-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4796585115121212147" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic3-14-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic3-14-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4814021027179800573" {
  display_name = "mGroup-topic-test-cgtest-130105-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130105-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4821195322569202557" {
  display_name = "mGroup-topic-test-dailysanity-130103-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-130103-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4829042812514065251" {
  display_name = "mGroup-topic-test-longmsgs-230055-longtopic1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-230055-longtopic1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4831344751870319280" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic33-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic33-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4837856849066800465" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic9-6-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic9-6-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4839252414273499970" {
  display_name = "mGroup-topic-test-dailysanity-185933-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-185933-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4839252414273500307" {
  display_name = "mGroup-topic-test-cgtest-185933-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-185933-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4843326584646312935" {
  display_name = "mGroup-topic-test-cgtest-190135-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190135-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4843422213311620557" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic31-30-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic31-30-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4850525270523980329" {
  display_name = "mGroup-topic-test-cgtest-130108-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130108-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4887792159603457118" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic45-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic45-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4921908350468552928" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic10-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic10-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4930502069527522894" {
  display_name = "mGroup-topic-test-hrresize-040027-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040027-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4938194035357576856" {
  display_name = "mGroup-topic-test-serverrestart-234520-topic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-serverrestart-234520-topic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4942331119452430123" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic12-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic12-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4942692151143412233" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic36-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic36-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4965076588058428837" {
  display_name = "mGroup-topic-test-dailysanity-190115-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190115-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--4987132891873097500" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic29-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic29-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5010054326935616986" {
  display_name = "mGroup-topic-test-dailysanity-190130-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190130-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5015621519697445900" {
  display_name = "mGroup-topic-test-hrresize-040025-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040025-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5036038992154542370" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic39-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic39-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5037455003296278918" {
  display_name = "mGroup-topic-test-cgtest-190119-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190119-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5057458547328412890" {
  display_name = "mGroup-topic-test-longmsgs-230053-longtopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-230053-longtopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5067727485299799" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic12-33-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic12-33-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--50773018382638593" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic34-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic34-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5078487404448759564" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic20-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic20-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5080179732147655475" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic12-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic12-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5091718405721841355" {
  display_name = "mGroup-topic-test-hrresize-040026-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040026-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5105865114450878525" {
  display_name = "mGroup-topic-test-shortlongmsgs-230122-short-msg-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-230122-short-msg-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5107700959649889731" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic19-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic19-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5109347743304885604" {
  display_name = "mGroup-topic-test-hrresize-040026-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040026-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5130139615529386609" {
  display_name = "mGroup-topic-test-dailysanity-104158-sanitytopic2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-104158-sanitytopic2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5140709551115609797" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic44-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic44-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5146185168683869452" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic18-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic18-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5148673468909323055" {
  display_name = "mGroup-topic-test-cgtest-190127-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190127-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5152874909280207851" {
  display_name = "mGroup-topic-test-cgtest-130109-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130109-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5155911075953019028" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic9-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic9-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5160360066000733750" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic37-30-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic37-30-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5165308031200937793" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic36-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic36-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5170413813943341934" {
  display_name = "mGroup-topic-test-hrrandpubs-020025-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020025-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--517193710352071299" {
  display_name = "mGroup-topic-test-cgtest-190122-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190122-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5174152931136151824" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic50-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic50-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5178062041577454094" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic9-24-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic9-24-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5180956644673171116" {
  display_name = "mGroup-topic-test-multi-topics-200847-topic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-200847-topic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--519943067672923243" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic34-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic34-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5202393481519127114" {
  display_name = "mGroup-topic-test-dailysanity-190157-sanitytopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190157-sanitytopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5202569895311157568" {
  display_name = "mGroup-topic-test-cgtest-190132-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190132-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5203969486513532735" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic6-23-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic6-23-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5206909127184531875" {
  display_name = "mGroup-topic-test-cgtest-154443-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-154443-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5209674765145276940" {
  display_name = "mGroup-topic-test-hrrandpubs-020035-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020035-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5209892184244615461" {
  display_name = "mGroup-topic-test-cgtest-190140-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190140-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5219434595317988629" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic38-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic38-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5250606842343978724" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic31-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic31-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5254746381709319692" {
  display_name = "mGroup-topic-test-intellij-cp-165311-nostalgy-hrs-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-intellij-cp-165311-nostalgy-hrs-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5269979506019190929" {
  display_name = "mGroup-topic-test-serverrestart-234520-topic1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-serverrestart-234520-topic1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5277669941864887936" {
  display_name = "mGroup-topic-test-shortlongmsgs-131745-long-msg-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-131745-long-msg-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5288022667850994535" {
  display_name = "mGroup-topic-test-dailysanity-130103-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-130103-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5294315603648851309" {
  display_name = "mGroup-topic-test-dailysanity-190124-sanitytopic2-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190124-sanitytopic2-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--529495745458201406" {
  display_name = "mGroup-topic-test-longmsgs-190126-longtopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-190126-longtopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5301446027824589898" {
  display_name = "mGroup-topic-test-cgtest-190122-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190122-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--531590885645900661" {
  display_name = "mGroup-topic-test-dailysanity-190117-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190117-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5325064355304419685" {
  display_name = "mGroup-topic-test-cgtest-130102-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130102-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5330254909347509078" {
  display_name = "mGroup-topic-test-cnsmrgrptest-094641-cgtopic-testcg2-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-094641-cgtopic-testcg2-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--533077575760925185" {
  display_name = "mGroup-topic-test-hrresize-040029-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040029-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5339208409058318468" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic21-14-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic21-14-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5342040033728930322" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic50-20-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic50-20-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5348384752992462908" {
  display_name = "mGroup-topic-test-multi-topics-200847-topic38-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-200847-topic38-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5349044061730779047" {
  display_name = "mGroup-topic-test-dailysanity-190136-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190136-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5353649549919665047" {
  display_name = "mGroup-topic-test-dailysanity-190139-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190139-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5363886610635860081" {
  display_name = "mGroup-topic-test-dailysanity-190135-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190135-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5373460889803754094" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic7-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic7-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5376172670017992865" {
  display_name = "mGroup-topic-test-shortlongmsgs-230115-long-msg-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-230115-long-msg-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5380744169124543032" {
  display_name = "mGroup-topic-test-dailysanity-131741-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-131741-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5397089094061109112" {
  display_name = "mGroup-topic-test-hrresize-040026-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040026-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5412279304071763971" {
  display_name = "mGroup-topic-test-hrresize-040025-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040025-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5418491331952400467" {
  display_name = "mGroup-topic-test-cgtest-190109-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190109-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--542445760322119466" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic46-26-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic46-26-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--542528469704329148" {
  display_name = "mGroup-topic-test-cgtest-190129-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5448847330846035810" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic40-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic40-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5491066199870396778" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic15-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic15-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5501598839420142632" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic34-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic34-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5510833576319892026" {
  display_name = "mGroup-topic-test-hrresize-020033-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020033-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5511259749892312005" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic8-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic8-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5526844007686535470" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic3-5-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic3-5-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5544338906624417641" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic20-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic20-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5557255177307310703" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190137-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190137-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5569276124827139873" {
  display_name = "mGroup-topic-test-cgtest-190127-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190127-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5569739259706342917" {
  display_name = "mGroup-topic-test-dailysanity-040118-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-040118-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5573681147542041835" {
  display_name = "mGroup-topic-test-cgtest-221639-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-221639-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5576631106679976227" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic45-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic45-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5585311613742914557" {
  display_name = "mGroup-topic-test-dailysanity-221642-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-221642-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5601406125069942809" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic15-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic15-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5602957619633014198" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic50-30-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic50-30-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5605717138003448233" {
  display_name = "mGroup-topic-test-dailysanity-190158-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190158-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--560798158389033221" {
  display_name = "mGroup-topic-test-highrates-221204-topic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-highrates-221204-topic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5609128214542367821" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic11-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic11-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5615440552581712431" {
  display_name = "mGroup-topic-test-cgtest-190125-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190125-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5634048455557038498" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic28-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic28-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5640092969668238670" {
  display_name = "mGroup-topic-test-hrrandpubs-020026-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020026-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5640616776070923460" {
  display_name = "mGroup-topic-test-hrresize-040026-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040026-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5642332825518381758" {
  display_name = "mGroup-topic-test-dailysanity-040118-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-040118-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5645320692444766719" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic42-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic42-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5651383151826289245" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic15-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic15-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5652316815060149701" {
  display_name = "mGroup-topic-test-dailysanity-130103-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-130103-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5659871142724044191" {
  display_name = "mGroup-topic-test-longmsgs-104209-longtopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-104209-longtopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5691548280221424810" {
  display_name = "topics-test-cassandra"
  filter       = "resource.metadata.name=has_substring(\"row-staging-topics-test-cassandra\")"
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5702321891824010899" {
  display_name = "mGroup-topic-test-cgtest-190116-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190116-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5703453288171070323" {
  display_name = "mGroup-topic-test-clustermanytopics-012932-topic12-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012932-topic12-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5704277780270690523" {
  display_name = "mGroup-topic-test-hrresize-040028-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040028-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5732540008431651656" {
  display_name = "mGroup-topic-test-dailysanity-130104-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-130104-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5737583725225130411" {
  display_name = "mGroup-topic-test-hrresize-020032-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020032-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5739376348990878111" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic35-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic35-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5741200149990540666" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic16-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic16-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5767797775654057750" {
  display_name = "mGroup-topic-test-dailysanity-111154-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-111154-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5771646425212810828" {
  display_name = "mGroup-topic-test-dailysanity-190130-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190130-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5780152910285736806" {
  display_name = "mGroup-topic-test-dailysanity-145959-sanitytopic2-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-145959-sanitytopic2-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--580118917335357508" {
  display_name = "mGroup-topic-test-cgtest-130109-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130109-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5824963282086613195" {
  display_name = "mGroup-topic-test-dailysanity-190157-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190157-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5837449927084533344" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic37-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic37-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5849483295984963884" {
  display_name = "mGroup-topic-test-cgtest-090108-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090108-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5866515508934720491" {
  display_name = "mGroup-topic-test-dailysanity-145959-sanitytopic2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-145959-sanitytopic2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5874959344315408591" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic41-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic41-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5884381101108429187" {
  display_name = "mGroup-topic-test-dailysanity-144645-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-144645-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5894920251058306103" {
  display_name = "mGroup-topic-test-hrresize-040158-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040158-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5904076536939113946" {
  display_name = "mGroup-topic-test-hrresize-020037-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020037-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5904090610015219772" {
  display_name = "mGroup-topic-test-hrsanity-093938-nostalgy-hrs-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-093938-nostalgy-hrs-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5905661494359737543" {
  display_name = "mGroup-topic-test-hrresize-020039-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020039-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5906320381756431860" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic13-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic13-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5913316075393735548" {
  display_name = "mGroup-topic-test-dailysanity-130106-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-130106-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5915395917723062092" {
  display_name = "mGroup-topic-test-cgtest-190129-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--592106586269976472" {
  display_name = "mGroup-topic-test-hrrandpubs-020031-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020031-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5928924428605351436" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic44-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic44-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5936781798447887846" {
  display_name = "mGroup-topic-test-hrrandpubs-020027-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020027-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5963266846832418032" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic4-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic4-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5970322286156400210" {
  display_name = "mGroup-topic-test-cgtest-090108-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090108-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5978181075234593293" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic2-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic2-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5986570414350887411" {
  display_name = "mGroup-topic-test-cgtest-185933-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-185933-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5987895964889645030" {
  display_name = "mGroup-topic-test-hrresize-040025-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040025-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5993069313030614448" {
  display_name = "mGroup-topic-test-hrresize-040025-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040025-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--5994218622097484812" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic2-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic2-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6020797580846016711" {
  display_name = "mGroup-topic-test-cgtest-190134-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190134-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6023073074805283885" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic31-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic31-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6029452095081266158" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic5-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic5-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6029712682244716999" {
  display_name = "mGroup-topic-test-dailysanity-090109-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-090109-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6030672758593579199" {
  display_name = "mGroup-topic-test-hrrandpubs-020025-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020025-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6033725124271872173" {
  display_name = "mGroup-topic-test-highrates-144417-topic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-highrates-144417-topic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6035864354645790177" {
  display_name = "mGroup-topic-test-dailysanity-163041-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-163041-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6045226252797587809" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic23-26-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic23-26-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6060048595574413819" {
  display_name = "mGroup-topic-test-hrresize-020039-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020039-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--606760247679844015" {
  display_name = "mGroup-topic-test-hrresize-040028-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040028-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6096931296784405941" {
  display_name = "mGroup-topic-test-cgtest-111147-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-111147-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6105063182773052469" {
  display_name = "mGroup-topic-test-longmsgs-195117-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-195117-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6112204303956014541" {
  display_name = "mGroup-topic-test-dailysanity-185933-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-185933-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6124622217233892569" {
  display_name = "mGroup-topic-test-cnsmrgrptest-030056-cgtopic-testcg1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-030056-cgtopic-testcg1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6126722688138093779" {
  display_name = "mGroup-topic-test-cgtest-190136-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190136-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6127682822187055653" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic28-17-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic28-17-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--613117408600271192" {
  display_name = "mGroup-topic-test-cgtest-190140-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190140-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6138253504747479703" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic27-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic27-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6142587971728884726" {
  display_name = "mGroup-topic-test-dailysanity-071646-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-071646-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6144128951863243211" {
  display_name = "mGroup-topic-test-cgtest-190126-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190126-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6172328661981647423" {
  display_name = "mGroup-topic-test-dailysanity-073908-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-073908-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6172937555063565617" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190158-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190158-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6176712007208493945" {
  display_name = "mGroup-topic-test-highrates-082944-topic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-highrates-082944-topic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--617684946725734450" {
  display_name = "mGroup-topic-test-cgtest-190202-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190202-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6193252185295679973" {
  display_name = "mGroup-topic-test-serverrestart-234526-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-serverrestart-234526-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6202647253396287697" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic23-20-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic23-20-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6213166354863467326" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic47-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic47-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6214751352310927526" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic20-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic20-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6218476925946730059" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190137-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190137-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6224783882302728495" {
  display_name = "mGroup-topic-test-shortlongmsgs-131745-short-msg-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-131745-short-msg-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6233084771962008639" {
  display_name = "mGroup-topic-test-hrresize-040034-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040034-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6233991412957440979" {
  display_name = "mGroup-topic-test-shortlongmsgs-131745-short-msg-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-131745-short-msg-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6234882602788193183" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic7-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic7-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6237917010692177934" {
  display_name = "mGroup-topic-test-dailysanity-121450-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-121450-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6242969127040883082" {
  display_name = "mGroup-topic-test-cgtest-190129-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6243204408237521509" {
  display_name = "mGroup-topic-test-hrresize-040025-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040025-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6248341009712879044" {
  display_name = "mGroup-topic-test-dailysanity-190130-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190130-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6270828038259010702" {
  display_name = "mGroup-topic-test-cgtest-121442-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-121442-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6277765444289868008" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic48-4-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic48-4-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6278743683555941572" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic45-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic45-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6289014698036081304" {
  display_name = "mGroup-topic-test-longmsgs-230030-longtopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-230030-longtopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6291539180585600426" {
  display_name = "mGroup-topic-test-cgtest-071717-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-071717-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6295372983648493154" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic46-8-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic46-8-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6296087274191027325" {
  display_name = "mGroup-topic-test-dailysanity-121450-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-121450-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6297347017475525443" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic16-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic16-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6299062940963712036" {
  display_name = "mGroup-topic-test-hrresize-040158-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040158-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6306911555209540846" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic50-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic50-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6308385941438072346" {
  display_name = "mGroup-topic-test-highrates-112740-topic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-highrates-112740-topic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6340840862837976350" {
  display_name = "mGroup-topic-test-longmsgs-111152-longtopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-111152-longtopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6343517532266867341" {
  display_name = "mGroup-topic-test-hrrandpubs-020026-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020026-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6355558271809826189" {
  display_name = "mGroup-topic-test-hrresize-040027-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040027-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6355631575864445648" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic41-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic41-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--636226156709936642" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic31-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic31-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6372426359139562727" {
  display_name = "mGroup-topic-test-dailysanity-190130-sanitytopic2-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190130-sanitytopic2-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6372966922315029688" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic32-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic32-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--637424104242658914" {
  display_name = "mGroup-topic-test-hrrandpubs-020027-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020027-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6378458851176208701" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic20-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic20-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6379741488526239174" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic9-32-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic9-32-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6385071298653224415" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic49-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic49-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6396306987241270246" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic40-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic40-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6405209865303315374" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic28-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic28-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6405964332013240355" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190138-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190138-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6422673403469090902" {
  display_name = "mGroup-topic-test-cgtest-040105-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-040105-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6441393377904254834" {
  display_name = "mGroup-topic-test-cgtest-130109-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130109-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6442125021000966043" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic36-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic36-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6453573312238577665" {
  display_name = "mGroup-topic-test-hrresize-040027-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040027-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--646272354813693190" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic25-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic25-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6463731731757707412" {
  display_name = "mGroup-topic-test-cgtest-190129-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6471083123755490177" {
  display_name = "mGroup-topic-test-cgtest-090105-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090105-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6474079086956567598" {
  display_name = "mGroup-topic-test-dailysanity-190137-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190137-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6476361074841243957" {
  display_name = "mGroup-topic-test-cgtest-190127-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190127-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6501408808884398122" {
  display_name = "mGroup-topic-test-dailysanity-040110-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-040110-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6549730754319247083" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic32-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic32-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6552335387708610093" {
  display_name = "mGroup-topic-test-dailysanity-154450-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-154450-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--655810208839881608" {
  display_name = "mGroup-topic-test-hrsanity-233759-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-233759-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6559967450294674256" {
  display_name = "mGroup-topic-test-clustermanytopics-012932-topic34-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012932-topic34-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6565051978205352967" {
  display_name = "mGroup-topic-test-longmsgs-230037-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-230037-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6579344074606754648" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic18-4-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic18-4-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6579571938696897318" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic10-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic10-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--658570535920154442" {
  display_name = "mGroup-topic-test-dailysanity-090103-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-090103-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6587829487785363104" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic9-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic9-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6595746043426363322" {
  display_name = "mGroup-topic-test-multi-topics-200819-topic26-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-200819-topic26-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6599216198966141763" {
  display_name = "mGroup-topic-test-cgtest-221639-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-221639-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6602171761811921530" {
  display_name = "mGroup-topic-test-cgtest-130147-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130147-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6615733109214099531" {
  display_name = "mGroup-topic-test-dailysanity-190130-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190130-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--661683184033673831" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic35-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic35-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6645908655312510880" {
  display_name = "mGroup-topic-test-dailysanity-190122-sanitytopic2-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190122-sanitytopic2-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6656133014272693755" {
  display_name = "mGroup-topic-test-hrresize-040027-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040027-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6657829380368638725" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic24-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic24-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6672905957428543074" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic13-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic13-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6672950984216245883" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic50-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic50-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6673239015842020901" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190158-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190158-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6677129446188280644" {
  display_name = "mGroup-topic-test-hrresize-040158-nostalgy-5-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040158-nostalgy-5-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6683027066401019260" {
  display_name = "mGroup-topic-test-hrresize-040027-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040027-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--671037659315258975" {
  display_name = "mGroup-topic-test-dailysanity-130104-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-130104-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6710458211950620011" {
  display_name = "mGroup-topic-test-longmsgs-123907-longtopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-123907-longtopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6728177065342882030" {
  display_name = "mGroup-topic-test-longmsgs-190116-longtopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-190116-longtopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6732503805898258738" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic21-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic21-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6733921706605705793" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic30-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic30-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6737699868167141702" {
  display_name = "mGroup-topic-test-hrrandpubs-020026-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020026-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6751883011153679667" {
  display_name = "mGroup-topic-test-hrrandpubs-020028-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020028-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6763634109596817813" {
  display_name = "mGroup-topic-test-longmsgs-230101-longtopic1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-230101-longtopic1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6782869558739258363" {
  display_name = "mGroup-topic-test-dailysanity-190131-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190131-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6795457299817405531" {
  display_name = "row-staging-topicserver-online-cassandra"
  filter       = "resource.metadata.name=has_substring(\"row-staging-topicserver-online-cassandra\")"
  is_cluster   = true
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--679571725238738813" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic44-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic44-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6801704900856776501" {
  display_name = "mGroup-topic-test-hrrandpubs-020025-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020025-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6821880445274627677" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic11-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic11-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6825010593343443586" {
  display_name = "mGroup-topic-test-cgtest-190140-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190140-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6840882186545605995" {
  display_name = "mGroup-topic-test-hrrandpubs-114016-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-114016-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6842070804791607351" {
  display_name = "mGroup-topic-test-hrsanity-160655-nostalgy-hrs-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-160655-nostalgy-hrs-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6843287448243875120" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic19-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic19-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6849236021397437914" {
  display_name = "mGroup-topic-test-hrrandpubs-020024-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020024-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6855505899501238010" {
  display_name = "mGroup-topic-test-hrresize-040028-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040028-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6857338145006470420" {
  display_name = "mGroup-topic-test-hrrandpubs-020025-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020025-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6865497481533944014" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190138-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190138-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6874399922619746812" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190158-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190158-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6876685047704692495" {
  display_name = "mGroup-topic-test-cgtest-070114-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-070114-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6887277261806246865" {
  display_name = "mGroup-topic-test-cnsmrgrptest-165059-cgtopic-testcg2-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-165059-cgtopic-testcg2-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6903399060759495650" {
  display_name = "mGroup-topic-test-hrrandpubs-020024-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020024-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6903854447605819668" {
  display_name = "mGroup-topic-test-cgtest-190126-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190126-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6927805472510469883" {
  display_name = "mGroup-topic-test-cgtest-190116-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190116-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6936573100663707659" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic9-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic9-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--693840660600336517" {
  display_name = "mGroup-topic-test-hrresize-040028-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040028-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6961852411021464675" {
  display_name = "mGroup-topic-test-dailysanity-190130-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190130-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6963854152891871286" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic28-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic28-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6973528342258796827" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic6-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic6-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--69740602285750881" {
  display_name = "mGroup-topic-test-cgtest-221639-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-221639-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--6992787454949396958" {
  display_name = "mGroup-topic-test-dailysanity-084734-sanitytopic1-pub"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-084734-sanitytopic1-pub\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7001160746230744441" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic21-27-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic21-27-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7004512537704916786" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic29-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic29-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7014807702561144250" {
  display_name = "mGroup-topic-test-cgtest-190135-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190135-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7024700772980164549" {
  display_name = "mGroup-topic-test-dailysanity-190157-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190157-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--702819658403310959" {
  display_name = "mGroup-topic-test-cgtest-190138-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190138-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7031475432061362036" {
  display_name = "mGroup-topic-test-dailysanity-090109-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-090109-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7033278252400613662" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic48-31-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic48-31-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7053518143547049296" {
  display_name = "mGroup-topic-test-hrresize-040025-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040025-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7054604435478805154" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic14-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic14-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7058968746358567429" {
  display_name = "mGroup-topic-test-longmsgs-190119-longtopic1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-190119-longtopic1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7060677812726166302" {
  display_name = "mGroup-topic-test-longmsgs-230037-longtopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-230037-longtopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7088862053713962800" {
  display_name = "mGroup-topic-test-cnsmrgrptest-094641-cgtopic-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-094641-cgtopic-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7093888190550782344" {
  display_name = "mGroup-topic-test-cgtest-190126-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190126-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7108851579034883274" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic22-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic22-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7116657914083617611" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic43-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic43-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7125385463625475804" {
  display_name = "mGroup-topic-test-dailysanity-190128-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190128-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7128505059551737731" {
  display_name = "mGroup-topic-test-dailysanity-190126-sanitytopic2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190126-sanitytopic2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7144143362436874538" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic41-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic41-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7145456471011007255" {
  display_name = "mGroup-topic-test-cgtest-130109-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130109-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7146898090746912025" {
  display_name = "mGroup-topic-test-cgtest-190109-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190109-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7147103389689347086" {
  display_name = "mGroup-topic-test-cgtest-190126-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190126-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7167992789720347087" {
  display_name = "mGroup-topic-test-dailysanity-221642-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-221642-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7168608269377442964" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic41-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic41-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7179575130208683459" {
  display_name = "mGroup-topic-test-dailysanity-190128-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190128-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7196975022954412599" {
  display_name = "mGroup-topic-test-serverrestart-215150-2-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-serverrestart-215150-2-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7222052556623988520" {
  display_name = "mGroup-topic-test-dailysanity-190124-sanitytopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190124-sanitytopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--723667569738955632" {
  display_name = "mGroup-topic-test-longmsgs-190126-longtopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-190126-longtopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7239127085425135686" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic15-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic15-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7251408116739376945" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic19-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic19-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7259659236127579612" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic32-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic32-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7268640158237890" {
  display_name = "mGroup-topic-test-hrrandpubs-020025-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020025-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7273425817994945554" {
  display_name = "mGroup-topic-test-cgtest-190129-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7277588373637705044" {
  display_name = "mGroup-topic-test-cgtest-154443-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-154443-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7279855803126979114" {
  display_name = "mGroup-topic-test-cgtest-190138-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190138-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7288479552322786492" {
  display_name = "mGroup-topic-test-cgtest-090105-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090105-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7305517318905429460" {
  display_name = "mGroup-topic-test-cgtest-190129-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--730946823759038784" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic22-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic22-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--732082858170491648" {
  display_name = "mGroup-topic-test-longmsgs-030059-longtopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-030059-longtopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7332095153806665967" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7347888320603681377" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic48-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic48-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7347888320603681407" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic25-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic25-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7350629699254721881" {
  display_name = "mGroup-topic-test-hrresize-040024-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040024-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7355675761155873184" {
  display_name = "mGroup-topic-test-cgtest-190134-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190134-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7360056673521585369" {
  display_name = "mGroup-topic-test-cgtest-190122-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190122-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7364868778548158511" {
  display_name = "mGroup-topic-test-longmsgs-230041-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-230041-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7395157370839598286" {
  display_name = "mGroup-topic-test-longmsgs-190127-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-190127-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7399376399538809471" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic45-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic45-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--744149925413811267" {
  display_name = "mGroup-topic-test-cgtest-090105-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090105-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7449782739860357372" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic43-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic43-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7456711893512252313" {
  display_name = "mGroup-topic-test-hrsanity-143408-nostalgy-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-143408-nostalgy-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7466081973830209796" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic15-28-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic15-28-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--748169135395136712" {
  display_name = "mGroup-topic-test-dailysanity-090103-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-090103-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7491876599325666135" {
  display_name = "mGroup-topic-test-hrrandpubs-020030-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020030-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7502633069551959761" {
  display_name = "mGroup-topic-test-clustermanytopics-012932-topic38-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012932-topic38-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--750319782430282436" {
  display_name = "mGroup-topic-test-shortlongmsgs-190133-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-190133-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7510144413269682921" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic15-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic15-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7515781408058043507" {
  display_name = "mGroup-topic-test-cgtest-040105-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-040105-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7515827548284479563" {
  display_name = "mGroup-topic-test-hrresize-040025-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040025-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7522976013662219320" {
  display_name = "mGroup-topic-test-cgtest-190134-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190134-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7540675768411049210" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic18-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic18-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7545570123154542460" {
  display_name = "mGroup-topic-test-cgtest-190140-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190140-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7551862662000501898" {
  display_name = "mGroup-topic-test-cgtest-121442-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-121442-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--756254647076009084" {
  display_name = "mGroup-topic-test-cgtest-181648-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-181648-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7562898856881280802" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic23-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic23-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--756306501643332351" {
  display_name = "mGroup-topic-test-dailysanity-190133-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190133-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7571134634672191146" {
  display_name = "mGroup-topic-test-dailysanity-190130-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190130-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7572338509306311375" {
  display_name = "mGroup-topic-test-cgtest-090101-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090101-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7581479109730470376" {
  display_name = "mGroup-topic-test-serverrestart-215054-2-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-serverrestart-215054-2-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7603163548979527989" {
  display_name = "mGroup-topic-test-dailysanity-190121-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190121-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7607869931397675075" {
  display_name = "mGroup-topic-test-cgtest-190202-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190202-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7612054070225062502" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic30-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic30-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7615862449232188469" {
  display_name = "mGroup-topic-test-clustermanytopics-012932-topic13-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012932-topic13-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7629325964591974650" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190144-cgtopic-testcg3-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190144-cgtopic-testcg3-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7644420138688017174" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic49-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic49-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7659554874004416474" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic28-29-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic28-29-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7663839086579179991" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic34-7-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic34-7-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7676335407869599778" {
  display_name = "mGroup-topic-test-dailysanity-095320-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-095320-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7684625506656955424" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic19-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic19-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7717736667133609499" {
  display_name = "mGroup-topic-test-cgtest-040105-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-040105-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7721980091279934291" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic21-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic21-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7722403393599211031" {
  display_name = "mGroup-topic-test-cgtest-221639-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-221639-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7726761455263910127" {
  display_name = "mGroup-topic-test-hrresize-040026-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040026-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7730358696915536977" {
  display_name = "mGroup-topic-test-hrresize-040025-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040025-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7733923718783536562" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7734240577923668649" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic17-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic17-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7769345457100854193" {
  display_name = "mGroup-topic-test-clustermanytopics-012932-topic19-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012932-topic19-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7780692142627990567" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic20-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic20-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7806468980129388444" {
  display_name = "mGroup-topic-test-dailysanity-190117-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190117-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7806791989918572144" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic11-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic11-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7812571068929337414" {
  display_name = "mGroup-topic-test-hrsanity-233955-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-233955-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7816129611935839565" {
  display_name = "mGroup-topic-test-cgtest-130105-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130105-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7825571989510577347" {
  display_name = "mGroup-topic-test-hrrandpubs-020028-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020028-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7825622391569296304" {
  display_name = "mGroup-topic-test-cgtest-190129-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7838823310124442024" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic43-27-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic43-27-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7851406937057737288" {
  display_name = "mGroup-topic-test-dailysanity-190121-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190121-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7851464267301495051" {
  display_name = "mGroup-topic-test-dailysanity-100654-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-100654-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7857852582316392490" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic18-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic18-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7860722658308757172" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic46-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic46-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--786090957533163050" {
  display_name = "mGroup-topic-test-manypubs-204226-manypubs1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-manypubs-204226-manypubs1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7865542228338268021" {
  display_name = "mGroup-topic-test-cgtest-163042-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-163042-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7874309876514089072" {
  display_name = "mGroup-topic-test-cgtest-190136-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190136-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7889376385799088663" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic49-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic49-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7889748149320950478" {
  display_name = "mGroup-topic-test-shortlongmsgs-131745-long-msg-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-131745-long-msg-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7889860147724597715" {
  display_name = "mGroup-topic-test-shortlongmsgs-230117-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-230117-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7901488222604232768" {
  display_name = "mGroup-topic-test-cgtest-154443-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-154443-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7906787894850111272" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic14-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic14-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--790714950198372340" {
  display_name = "mGroup-topic-test-dailysanity-090103-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-090103-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7911349140031843966" {
  display_name = "mGroup-topic-test-cgtest-130105-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130105-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7917713881144289978" {
  display_name = "mGroup-topic-test-longmsgs-190141-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-190141-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7923871467901032456" {
  display_name = "mGroup-topic-test-cgtest-154443-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-154443-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7930665557346960648" {
  display_name = "mGroup-topic-test-dailysanity-071646-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-071646-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7931724922186947420" {
  display_name = "mGroup-topic-test-dailysanity-215651-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-215651-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7935064312797911526" {
  display_name = "mGroup-topic-test-dailysanity-130146-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-130146-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7941808405819653009" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic39-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic39-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7947560455123676260" {
  display_name = "mGroup-topic-test-dailysanity-185933-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-185933-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7950576402732821460" {
  display_name = "mGroup-topic-test-dailysanity-190108-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190108-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7952019672607432953" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic40-13-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic40-13-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--795429877642907295" {
  display_name = "mGroup-topic-test-cgtest-190126-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190126-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7970164730445364716" {
  display_name = "mGroup-topic-test-cgtest-154443-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-154443-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--7986470320890886916" {
  display_name = "mGroup-topic-test-cgtest-190129-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8015653358706920563" {
  display_name = "mGroup-topic-test-hrrandpubs-020025-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020025-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--802423337691093600" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic10-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic10-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8024526981720342049" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic21-5-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic21-5-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8032917857704354423" {
  display_name = "mGroup-topic-test-dailysanity-190136-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190136-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8037183684468260044" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190135-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190135-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8039629979636708609" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic4-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic4-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--806420315144510748" {
  display_name = "mGroup-topic-test-cgtest-190129-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8064952102632948814" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic31-25-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic31-25-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8069201698216668197" {
  display_name = "mGroup-topic-test-hrresize-020033-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020033-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8070841697312799884" {
  display_name = "mGroup-topic-test-cgtest-130102-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130102-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8082963137247786389" {
  display_name = "mGroup-topic-test-dailysanity-190111-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190111-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8084263404321850982" {
  display_name = "mGroup-topic-test-shortlongmsgs-030101-long-msg-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-030101-long-msg-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8103396880483443526" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic33-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic33-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8115008904483627181" {
  display_name = "bbb"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-084734-sanitytopic1-sub\")"
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8115108950732334662" {
  display_name = "mGroup-topic-test-hrrandpubs-020028-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020028-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8123797950527336484" {
  display_name = "mGroup-topic-test-longmsgs-190129-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-190129-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8124346099667241532" {
  display_name = "mGroup-topic-test-multi-topics-200819-topic17-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-200819-topic17-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8126173180935509518" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic42-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic42-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8141313764049677941" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic18-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic18-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8141727354752716671" {
  display_name = "mGroup-topic-test-cgtest-121442-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-121442-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--814296722015952275" {
  display_name = "mGroup-topic-test-shortlongmsgs-130059-short-msg-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-130059-short-msg-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8153213112684655146" {
  display_name = "mGroup-topic-test-dailysanity-190139-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190139-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8155041903210678989" {
  display_name = "mGroup-topic-test-cnsmrgrptest-030053-cgtopic-testcg2-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-030053-cgtopic-testcg2-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8162774326159758322" {
  display_name = "mGroup-topic-test-cgtest-090105-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090105-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8184153403317103821" {
  display_name = "mGroup-topic-test-dailysanity-190136-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190136-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--81884690919939152" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic4-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic4-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8203514910624019373" {
  display_name = "mGroup-topic-test-cgtest-190134-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190134-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8207539170198897464" {
  display_name = "mGroup-topic-test-dailysanity-171605-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-171605-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8218109437293018152" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic29-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic29-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8223182644533251551" {
  display_name = "mGroup-topic-test-clustermanytopics-012932-topic48-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012932-topic48-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8227976381380733565" {
  display_name = "mGroup-topic-test-dailysanity-104158-sanitytopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-104158-sanitytopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--823627624934232012" {
  display_name = "mGroup-topic-test-dailysanity-190131-sanitytopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190131-sanitytopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8238947334890292498" {
  display_name = "mGroup-topic-test-cgtest-163042-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-163042-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8241492148785126857" {
  display_name = "mGroup-topic-test-hrresize-040059-nostalgy-5-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040059-nostalgy-5-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8259922266665041929" {
  display_name = "mGroup-topic-test-clustermanytopics-012503-topic13-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-012503-topic13-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--826207540935039119" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic11-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic11-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8269785804670805726" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic41-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic41-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8270574746004075089" {
  display_name = "mGroup-topic-test-longmsgs-185936-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-185936-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8271308945224396267" {
  display_name = "mGroup-topic-test-hrresize-040026-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040026-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8277571120819592703" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic21-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic21-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8280152708245281246" {
  display_name = "mGroup-topic-test-longmsgs-030108-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-030108-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8282500047438654022" {
  display_name = "mGroup-topic-test-dailysanity-190136-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190136-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8285128258300927320" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190138-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190138-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8292463451207867173" {
  display_name = "mGroup-topic-test-dailysanity-084456-sanitytopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-084456-sanitytopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8306600259748509201" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic18-4-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic18-4-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8307054993834267776" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic44-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic44-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8308005290567913875" {
  display_name = "mGroup-topic-test-cgtest-130108-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130108-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8310453146444006203" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190135-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190135-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8324198273328318956" {
  display_name = "mGroup-topic-test-cgtest-190127-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190127-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8325834273091560014" {
  display_name = "mGroup-topic-test-hrrandpubs-020035-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020035-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8337193627500699312" {
  display_name = "mGroup-topic-test-dailysanity-190105-sanitytopic2-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190105-sanitytopic2-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--833986827144030125" {
  display_name = "mGroup-topic-test-dailysanity-190115-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190115-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8341275395402942341" {
  display_name = "mGroup-topic-test-hrresize-020033-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-020033-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8355192158832723668" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic46-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic46-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8368926588676829067" {
  display_name = "mGroup-topic-test-hrresize-040026-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040026-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8372189483595453022" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic37-10-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic37-10-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8382722477058948077" {
  display_name = "mGroup-topic-test-cgtest-130147-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130147-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8383609161691125558" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic23-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic23-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8388117266944626639" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic5-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic5-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8388631702017414912" {
  display_name = "mGroup-topic-test-cgtest-130102-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-130102-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8409546070883498174" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic27-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic27-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8412850603217159522" {
  display_name = "mGroup-topic-test-hrrandpubs-020030-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020030-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8428545903703316732" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic50-14-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic50-14-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8428560325057611774" {
  display_name = "mGroup-topic-test-cgtest-181648-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-181648-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8434791700977099990" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic23-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic23-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8435654121121997284" {
  display_name = "mGroup-topic-test-dailysanity-173655-sanitytopic2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-173655-sanitytopic2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8465050009874835142" {
  display_name = "mGroup-topic-test-longmsgs-190130-longtopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-190130-longtopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8470666240724936497" {
  display_name = "mGroup-topic-test-dailysanity-190135-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190135-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8481205497168754180" {
  display_name = "mGroup-topic-test-dailysanity-145959-sanitytopic1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-145959-sanitytopic1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8487624351572956316" {
  display_name = "mGroup-topic-test-serverrestart-020612-2-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-serverrestart-020612-2-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8488" {
  display_name = "Staging"
  filter       = "resource.metadata.name=has_substring(\"staging\")"
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8488531310103684607" {
  display_name = "mGroup-topic-test-cgtest-070114-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-070114-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8489" {
  display_name = "Db"
  filter       = "resource.metadata.name=has_substring(\"db\")"
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8489051211036982590" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic26-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic26-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8490960692155758464" {
  display_name = "mGroup-topic-test-cgtest-190125-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190125-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8500898166258944997" {
  display_name = "mGroup-topic-test-hrrandpubs-020027-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020027-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8515933237550305057" {
  display_name = "mGroup-topic-test-cgtest-144642-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-144642-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8526720604575640704" {
  display_name = "mGroup-topic-test-cgtest-070114-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-070114-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8534446456336707553" {
  display_name = "mGroup-topic-test-cgtest-185933-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-185933-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8543141100432297809" {
  display_name = "mGroup-topic-test-hrrandpubs-020028-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020028-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8544972001969243059" {
  display_name = "mGroup-topic-test-cgtest-190120-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190120-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8546851804306714508" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic48-8-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic48-8-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8577111924249810466" {
  display_name = "mGroup-topic-test-hrresize-023311-nostalgy-4-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-023311-nostalgy-4-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8579876438678848773" {
  display_name = "mGroup-topic-test-hrrandpubs-020028-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020028-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8582980459855352418" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic28-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic28-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8585089926373421557" {
  display_name = "mGroup-topic-test-longmsgs-052042-longtopic1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-052042-longtopic1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8597145676010238213" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic42-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic42-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8607883911226322926" {
  display_name = "mGroup-topic-test-hrsanity-102506-nostalgy-hrs-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-102506-nostalgy-hrs-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8637493386500494430" {
  display_name = "mGroup-topic-test-cgtest-190125-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190125-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8638100373074393344" {
  display_name = "mGroup-topic-test-cgtest-185933-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-185933-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8639773698472151466" {
  display_name = "mGroup-topic-test-dailysanity-181646-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-181646-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--865249278034875997" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic7-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic7-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8661382658749093453" {
  display_name = "mGroup-topic-test-longmsgs-190135-longtopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-190135-longtopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--86722759659620523" {
  display_name = "mGroup-topic-test-cgtest-190136-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190136-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8678931596187315362" {
  display_name = "mGroup-topic-test-hrrandpubs-020024-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020024-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8679734403439623555" {
  display_name = "mGroup-topic-test-hrsanity-180021-nostalgy-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-180021-nostalgy-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8686971842414418929" {
  display_name = "mGroup-topic-test-hrrandpubs-020030-nostalgy-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020030-nostalgy-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8693155264491786005" {
  display_name = "mGroup-topic-test-cgtest-190120-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190120-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8710718397508217278" {
  display_name = "mGroup-topic-test-cgtest-190134-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190134-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8715778889822827463" {
  display_name = "mGroup-topic-test-cgtest-190119-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190119-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8717061138563465887" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190129-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190129-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8722645682505937807" {
  display_name = "mGroup-topic-test-multi-topics-201101-topic12-31-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201101-topic12-31-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8737477221144151501" {
  display_name = "mGroup-topic-test-manypubs-203811-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-manypubs-203811-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8744235573980599713" {
  display_name = "mGroup-topic-test-dailysanity-071646-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-071646-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8754668799296262153" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic13-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic13-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8761705384602511661" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic18-31-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic18-31-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8771355733123444384" {
  display_name = "mGroup-topic-test-dailysanity-181646-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-181646-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8778534463906219182" {
  display_name = "mGroup-topic-test-manypubs-005714-manypubs1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-manypubs-005714-manypubs1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8779823773971677499" {
  display_name = "mGroup-topic-test-hrresize-040028-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040028-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8793980996315166089" {
  display_name = "mGroup-topic-test-hrresize-023311-nostalgy-5-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-023311-nostalgy-5-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8795057190470184982" {
  display_name = "mGroup-topic-test-cnsmrgrptest-112849-cgtopic-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-112849-cgtopic-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8797724971999061232" {
  display_name = "mGroup-topic-test-cgtest-190132-cgtopic-testcg1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190132-cgtopic-testcg1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8805886273971792893" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic27-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic27-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8807809211334947074" {
  display_name = "mGroup-topic-test-hrresize-040028-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040028-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8812018549182290230" {
  display_name = "mGroup-topic-test-hrresize-040025-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040025-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8822921682111967338" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8848463830481890329" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic47-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic47-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8857498890388934718" {
  display_name = "mGroup-topic-test-longmsgs-190131-longtopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-longmsgs-190131-longtopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8876493512734448551" {
  display_name = "mGroup-topic-test-multi-topics-201053-topic37-16-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201053-topic37-16-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8876719620459389890" {
  display_name = "mGroup-topic-test-hrresize-040026-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040026-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8899056243294332148" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic46-19-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic46-19-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8910101540092815988" {
  display_name = "mGroup-topic-test-dailysanity-090105-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-090105-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8912948827784087350" {
  display_name = "mGroup-topic-test-hrrandpubs-020025-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020025-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8916682671973883239" {
  display_name = "mGroup-topic-test-dailysanity-040118-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-040118-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8919878897112776496" {
  display_name = "mGroup-topic-test-hrsanity-133916-nostalgy-hrs-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-133916-nostalgy-hrs-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8922935750868578190" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic46-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic46-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8938898640856295218" {
  display_name = "mGroup-topic-test-hrrandpubs-020031-nostalgy-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020031-nostalgy-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8940727943031674451" {
  display_name = "mGroup-topic-test-hrrandpubs-020039-nostalgy-hrs-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020039-nostalgy-hrs-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8954129038823571915" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic30-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic30-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--895451257272437844" {
  display_name = "mGroup-topic-test-clustermanytopics-211107-topic28-6-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211107-topic28-6-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8968492279825220081" {
  display_name = "mGroup-topic-test-cgtest-121442-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-121442-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8969823429557372930" {
  display_name = "mGroup-topic-test-multi-topics-002718-topic28-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002718-topic28-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8988303612872710383" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic48-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic48-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8991245131637228483" {
  display_name = "mGroup-topic-test-cgtest-190125-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190125-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--8996930443212938004" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190135-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190135-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9019479378994722384" {
  display_name = "mGroup-topic-test-dailysanity-190134-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190134-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9027192865811008412" {
  display_name = "mGroup-topic-test-highrates-082944-topic1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-highrates-082944-topic1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9038479326844600482" {
  display_name = "mGroup-topic-test-dailysanity-190119-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190119-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9043680813630634960" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic9-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic9-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9054646531636558208" {
  display_name = "mGroup-topic-test-hrrandpubs-020026-nostalgy-2-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrrandpubs-020026-nostalgy-2-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9060358204686676842" {
  display_name = "mGroup-topic-test-cgtest-070114-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-070114-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9060593048817829940" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic25-11-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic25-11-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9066106577999722988" {
  display_name = "mGroup-topic-test-dailysanity-190128-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190128-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9069016597886743107" {
  display_name = "mGroup-topic-test-hrresize-040158-nostalgy-4-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrresize-040158-nostalgy-4-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9085867923779925714" {
  display_name = "mGroup-topic-test-multi-topics-201044-topic12-5-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201044-topic12-5-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9096836247718130896" {
  display_name = "mGroup-topic-test-cgtest-190120-cgtopic-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190120-cgtopic-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9097342323291119971" {
  display_name = "mGroup-topic-test-dailysanity-163041-sanitytopic1-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-163041-sanitytopic1-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9105912755602542915" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic11-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic11-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9124796794400892465" {
  display_name = "mGroup-topic-test-manysubs-190110-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-manysubs-190110-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9130104053039579773" {
  display_name = "mGroup-topic-test-dailysanity-190130-sanitytopic1-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190130-sanitytopic1-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9137107415198103699" {
  display_name = "mGroup-topic-test-multi-topics-002826-topic4-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-002826-topic4-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9141212067892720629" {
  display_name = "staging-mmc"
  filter       = "resource.metadata.name=starts_with(\"staging-mmc\")"
  is_cluster   = true
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9143006275861051519" {
  display_name = "mGroup-topic-test-dailysanity-190131-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-dailysanity-190131-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9182411307364010598" {
  display_name = "mGroup-topic-test-cgtest-090101-cgtopic-testcg3-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-090101-cgtopic-testcg3-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--9201888677931776324" {
  display_name = "mGroup-topic-test-cnsmrgrptest-190144-1-topicserver"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cnsmrgrptest-190144-1-topicserver\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--927214957538079235" {
  display_name = "mGroup-topic-test-hrsanity-060024-nostalgy-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-hrsanity-060024-nostalgy-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--927352320545303824" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic10-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic10-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--930286580996518855" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic12-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic12-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--933161529140001026" {
  display_name = "mGroup-topic-test-cgtest-070114-cgtopic-testcg2-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-070114-cgtopic-testcg2-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--938166404436196373" {
  display_name = "mGroup-topic-test-cgtest-190129-cgtopic-1-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-cgtest-190129-cgtopic-1-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--95268265744353830" {
  display_name = "mGroup-topic-test-shortlongmsgs-130854-short-msg-subscriber"
  filter       = "resource.metadata.name=has_substring(\"topic-test-shortlongmsgs-130854-short-msg-subscriber\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--963940307436614317" {
  display_name = "mGroup-topic-test-clustermanytopics-211131-topic35-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211131-topic35-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--965337485298063231" {
  display_name = "mGroup-topic-test-clustermanytopics-211143-topic36-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-211143-topic36-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--973609200453438448" {
  display_name = "mGroup-topic-test-multi-topics-131044-topic23-3-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-131044-topic23-3-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--982274692197859139" {
  display_name = "mGroup-topic-test-clustermanytopics-013119-topic43-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-clustermanytopics-013119-topic43-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_group" "projects--waze_development--groups--982279410590497563" {
  display_name = "mGroup-topic-test-multi-topics-201159-topic3-1-publisher"
  filter       = "resource.metadata.name=has_substring(\"topic-test-multi-topics-201159-topic3-1-publisher\") AND resource.type = \"gce_instance\""
  is_cluster   = false
  project      = "waze-development"
}

resource "google_monitoring_notification_channel" "projects--waze_development--notificationChannels--10651237603377009513" {
  enabled = true

  labels = {
    email_address = "waze-ads-backend@google.com"
  }

  project     = "waze-development"
  type        = "email"
  user_labels = {}
}

resource "google_monitoring_notification_channel" "projects--waze_development--notificationChannels--12740028274668241490" {
  enabled = true

  labels = {
    email_address = "rotemartzi@google.com"
  }

  project     = "waze-development"
  type        = "email"
  user_labels = {}
}

resource "google_monitoring_notification_channel" "projects--waze_development--notificationChannels--13714067585895474016" {}

resource "google_monitoring_notification_channel" "projects--waze_development--notificationChannels--14843605170521580243" {}

resource "google_monitoring_notification_channel" "projects--waze_development--notificationChannels--17904979677041311299" {}

resource "google_monitoring_notification_channel" "projects--waze_development--notificationChannels--2642538331486778778" {}

resource "google_monitoring_notification_channel" "projects--waze_development--notificationChannels--3392062229024280053" {}

resource "google_monitoring_notification_channel" "projects--waze_development--notificationChannels--3447124078559702958" {
  enabled = true

  labels = {
    email_address = "waze-alerts@google.com"
  }

  project     = "waze-development"
  type        = "email"
  user_labels = {}
}

resource "google_monitoring_notification_channel" "projects--waze_development--notificationChannels--431916785291802427" {}

resource "google_monitoring_notification_channel" "projects--waze_development--notificationChannels--6380234528628896522" {}

resource "google_monitoring_notification_channel" "projects--waze_development--notificationChannels--6652406645113239179" {}

resource "google_monitoring_notification_channel" "projects--waze_development--notificationChannels--8670135148204309587" {}

resource "google_monitoring_uptime_check_config" "projects--waze_development--uptimeCheckConfigs--biz_gcp_wazestg_com" {}
