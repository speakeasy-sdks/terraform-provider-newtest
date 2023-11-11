terraform {
  required_providers {
    newtest = {
      source  = "testing/newtest"
      version = "0.13.1"
    }
  }
}

provider "newtest" {
  # Configuration options
}