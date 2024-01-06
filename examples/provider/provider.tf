terraform {
  required_providers {
    newtest = {
      source  = "testing/newtest"
      version = "0.16.0"
    }
  }
}

provider "newtest" {
  # Configuration options
}