---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bigipltm_monitor Data Source - terraform-provider-bigipltm"
subcategory: ""
description: |-
  
---

# bigipltm_monitor (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the LTM Monitor
- `partition` (String) partition of LTM Monitor

### Read-Only

- `adaptive` (String) ftp adaptive
- `adaptive_limit` (Number) Integer value
- `database` (String) the database in which your user is created
- `defaults_from` (String) Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp, /Common/gateway-icmp or /Common/tcp-half-open or /Common/smtp.
- `destination` (String) Alias for the destination
- `filename` (String) Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
- `id` (String) The ID of this resource.
- `interval` (Number)
- `ip_dscp` (Number)
- `manual_resume` (String)
- `mode` (String) Specifies the data transfer process (DTP) mode. The default value is passive.
- `receive_disable` (String) Expected response string.
- `reverse` (String)
- `time_until_up` (Number) Time in seconds
- `timeout` (Number)
- `transparent` (String)
- `username` (String) Specifies the user name if the monitored target requires authentication


