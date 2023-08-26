terraform {
  required_providers {
    newtest = {
      source  = "testing/newtest"
      version = "0.2.0"
    }
  }
}

provider "newtest" {
  # Configuration options
}