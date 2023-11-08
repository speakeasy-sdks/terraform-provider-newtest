terraform {
  required_providers {
    newtest = {
      source  = "testing/newtest"
      version = "0.12.0"
    }
  }
}

provider "newtest" {
  # Configuration options
}