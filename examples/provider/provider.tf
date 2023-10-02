terraform {
  required_providers {
    newtest = {
      source  = "testing/newtest"
      version = "0.8.1"
    }
  }
}

provider "newtest" {
  # Configuration options
}