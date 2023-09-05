terraform {
  required_providers {
    newtest = {
      source  = "testing/newtest"
      version = "0.4.2"
    }
  }
}

provider "newtest" {
  # Configuration options
}