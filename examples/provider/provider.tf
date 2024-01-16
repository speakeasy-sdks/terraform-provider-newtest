terraform {
  required_providers {
    newtest = {
      source  = "testing/newtest"
      version = "0.16.2"
    }
  }
}

provider "newtest" {
  # Configuration options
}