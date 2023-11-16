terraform {
  required_providers {
    newtest = {
      source  = "testing/newtest"
      version = "0.13.3"
    }
  }
}

provider "newtest" {
  # Configuration options
}