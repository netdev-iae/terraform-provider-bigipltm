terraform {
  required_providers {
    bigipcustum = {
      version = "0.0.1"
      source = "terraform.lab.local/net/bigipcustum"
    }
  }
}


provider "bigipcustum" {
  address  = "172.16.0.129"
  username = "admin"
  password = "123***sss"
} 


resource "bigipcustum_monitor_custum" "monitor" {
  name       = "/Common/terraform_monitor"
  parent     = "/Common/http"
  send       = "GET /some/path\r\n"
  timeout    = "15"
  interval   = "46"
}


resource "bigipcustum_monitor_custum" "monitor" {
  name       = "/Common/terraform_monitor"
  parent     = "/Common/"
  send       = "GET /some/path\r\n"
  timeout    = "15"
  interval   = "46"
}


