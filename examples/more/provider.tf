terraform {
  required_providers {
    bigipltm = {
      version = "0.0.3"
      source  = "terraform.lab.local/net/bigipltm" // used here as local provider for dev purpose
    }
  }
}

variable "password" {
  type      = string
  sensitive = true
}
variable "address" {
  type = string
}

provider "bigipltm" {
  address  = var.address
  username = "admin"
  password = var.password
}

resource "bigipltm_monitor" "monitor_http" {
  name     = "/Common/terraform_monitor"
  parent   = "/Common/http"
  send     = "GET /some/path\r\n"
  timeout  = "15"
  interval = "46"
}


resource "bigipltm_monitor" "monitor_smtp" {
  name        = "/Common/smtp_monitor"
  parent      = "/Common/smtp"
  destination = "*:563"
  timeout     = "15"
  interval    = "46"
}
resource "bigipltm_monitor" "monitor_https" {
  name        = "/Common/terraform_monitor_https"
  parent      = "/Common/https"
  send        = "GET /some/path\r\n"
  timeout     = "15"
  interval    = "46"
  ssl_profile = "serverssl"
}
