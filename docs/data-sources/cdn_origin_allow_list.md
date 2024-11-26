---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "qwilt_cdn_origin_allow_list Data Source - qwilt"
subcategory: ""
description: |-
  Retrieves the device ip's to be added to origin allow iist.
---

# qwilt_cdn_origin_allow_list (Data Source)

Retrieves the device ip's to be added to origin allow iist.

## Example Usage

```terraform
data "qwilt_cdn_origin_allow_list" "origin_allow_list" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `create_time_millis` (Number) The time this instance of the IP address list was generated.
- `ip_data` (Attributes Map) A dictionary structure where each key is a network name, and the value is an object comprised of two arrays; one for the IPv4 addresses and one for the IPv6 addresses in the network that the Qwilt CDN may use to request content from your origin. (see [below for nested schema](#nestedatt--ip_data))
- `md5` (String) A unique identifier for this instance of the IP address list.

<a id="nestedatt--ip_data"></a>
### Nested Schema for `ip_data`

Read-Only:

- `ipv4` (List of String) The IPv4 addresses in the network that Qwilt CDN may use to request content from your origin.
- `ipv6` (List of String) The IPv6 addresses in the network that Qwilt CDN may use to request content from your origin.