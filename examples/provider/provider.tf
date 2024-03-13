terraform {
  required_providers {
    newtest = {
      source  = "testing/newtest"
      version = "0.19.1"
    }
  }
}

provider "newtest" {
  # Configuration options
}