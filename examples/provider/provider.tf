terraform {
  required_providers {
    newtest = {
      source  = "testing/newtest"
      version = "0.12.1"
    }
  }
}

provider "newtest" {
  # Configuration options
}