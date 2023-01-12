terraform {
  required_providers {
    bigipltm = {
      version = "0.0.3"
      source  = "terraform.lab.local/net/bigipltm"
    }
  }
}

variable "password" {
  type = string
}
provider "bigipltm" {
  address  = "172.16.0.129"
  username = "admin"
  password = var.password
}


resource "bigipltm_monitor" "monitor_smtp" {
  name        = "/Common/smtp_monitor"
  parent      = "/Common/smtp"
  destination = "*:563"
  timeout     = "15"
  interval    = "46"
}
