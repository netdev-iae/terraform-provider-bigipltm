terraform {
  required_providers {
    bigipcustum = {
      version = "0.0.2"
      source = "terraform.lab.local/net/bigipcustum"
    }
  }
}


provider "bigipcustum" {
  address  = "172.16.0.129"
  username = "admin"
  password = "123***sss"
} 


resource "bigipcustum_monitor_custum" "monitor_http" {
  name       = "/Common/terraform_monitor"
  parent     = "/Common/http"
  send       = "GET /some/path\r\n"
  timeout    = "15"
  interval   = "46"
}


resource "bigipcustum_monitor_custum" "monitor_smtp" {
  name       = "/Common/smtp_monitor"
  parent     = "/Common/smtp"
  destination  = "*:563"
  timeout    = "15"
  interval   = "46"
}