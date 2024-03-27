terraform {
  required_providers {
    newtest = {
      source  = "testing/newtest"
      version = "0.21.0"
    }
  }
}

provider "newtest" {
  # Configuration options
}