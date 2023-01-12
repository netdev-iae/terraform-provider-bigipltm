# Overview

A [Terraform](terraform.io) provider for F5 BigIP LTM Monitor ( for this version Only SMTP monitor is added to). This repo is a fork of the offical [repo](https://github.com/F5Networks/terraform-provider-bigip)
with the addition of a resource provider to configure BigIP resources

# Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.11.x / 0.12.x /0.13.x
- [Go](https://golang.org/doc/install) 1.19 (to build the provider plugin)

# F5 BigIP LTM requirements

- This provider uses the iControlREST API, make sure that it is installed and enabled on your F5 device before proceeding.

These BIG-IP versions are supported in these Terraform versions.

| BIG-IP version | Terraform 1.x | Terraform 0.13 | Terraform 0.12 | Terraform 0.11 |
| -------------- | ------------- | -------------- | -------------- | -------------- |
| BIG-IP 16.x    | X             | X              | X              | X              |
| BIG-IP 15.x    | X             | X              | X              | X              |
| BIG-IP 14.x    | X             | X              | X              | X              |
| BIG-IP 12.x    | X             | X              | X              | X              |
| BIG-IP 13.x    | X             | X              | X              | X              |

# Documentation

Below is an example of how you can use the provider to create GTM/DNS resources

```
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
```

# Using the Provider

You can download the binary from the releases [section](https://github.com/anesh/terraform-provider-bigip/releases) of this repo and follow the instructions of installing terraform plugins.
