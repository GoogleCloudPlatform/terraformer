terraform {
  required_providers {
    auth0 = {
      source  = "alexkappa/auth0"
      version = "0.26.1"
    }
  }
}

provider "auth0" {
  # Configuration options
}
