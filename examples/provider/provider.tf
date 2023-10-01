terraform {
  required_providers {
    newtest = {
      source  = "testing/newtest"
      version = "0.8.0"
    }
  }
}

provider "newtest" {
  # Configuration options
}