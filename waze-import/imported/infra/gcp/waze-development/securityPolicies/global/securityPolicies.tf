provider "google" {
  project = ""
  region  = ""
}

resource "google_compute_security_policy" "allow_google_corp" {
  name    = "allow-google-corp"
  project = "waze-development"

  rule = {
    action      = "allow"
    description = "google-corp"

    match = {
      config = {
        src_ip_ranges = ["74.125.56.128/29", "74.125.116.0/22", "72.14.224.0/21", "74.125.120.0/22", "74.125.56.129/32"]
      }

      versioned_expr = "SRC_IPS_V1"
    }

    preview  = false
    priority = "5"
  }

  rule = {
    action      = "allow"
    description = "google-corp"

    match = {
      config = {
        src_ip_ranges = ["54.195.80.117/32", "54.197.192.34/32", "74.125.56.132/31", "54.195.80.116/32", "74.125.61.227/32"]
      }

      versioned_expr = "SRC_IPS_V1"
    }

    preview  = false
    priority = "6"
  }

  rule = {
    action      = "allow"
    description = "google-corp"

    match = {
      config = {
        src_ip_ranges = ["31.154.8.66/32", "66.102.14.24/30", "31.154.8.65/32", "31.154.8.68/30", "66.102.14.16/30"]
      }

      versioned_expr = "SRC_IPS_V1"
    }

    preview  = false
    priority = "4"
  }

  rule = {
    action      = "allow"
    description = "google-corp"

    match = {
      config = {
        src_ip_ranges = ["216.239.55.189/32", "216.239.55.188/31", "216.239.45.0/24", "216.239.33.60/30", "216.239.35.0/24"]
      }

      versioned_expr = "SRC_IPS_V1"
    }

    preview  = false
    priority = "2"
  }

  rule = {
    action      = "allow"
    description = "google-corp"

    match = {
      config = {
        src_ip_ranges = ["216.239.55.42/31", "31.154.2.234/32", "216.239.55.85/32", "216.239.55.84/31", "31.154.8.64/28"]
      }

      versioned_expr = "SRC_IPS_V1"
    }

    preview  = false
    priority = "3"
  }

  rule = {
    action      = "allow"
    description = "google-corp"

    match = {
      config = {
        src_ip_ranges = ["104.132.52.0/24", "212.179.82.66/31", "212.179.82.64/28", "212.179.82.74/32", "104.132.34.64/27"]
      }

      versioned_expr = "SRC_IPS_V1"
    }

    preview  = false
    priority = "1"
  }

  rule = {
    action      = "deny(403)"
    description = "default rule"

    match = {
      config = {
        src_ip_ranges = ["*"]
      }

      versioned_expr = "SRC_IPS_V1"
    }

    preview  = false
    priority = "2147483647"
  }

  rule = {
    action      = "allow"
    description = "google-corp"

    match = {
      config = {
        src_ip_ranges = ["54.197.192.35/32"]
      }

      versioned_expr = "SRC_IPS_V1"
    }

    preview  = false
    priority = "7"
  }
}
