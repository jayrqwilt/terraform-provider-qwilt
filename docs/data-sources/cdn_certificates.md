---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "qwilt_cdn_certificates Data Source - qwilt"
subcategory: ""
description: |-
  Retrieves the certificates uploaded to Qwilt CDN by your organization and the associated metadata.
---

# qwilt_cdn_certificates (Data Source)

Retrieves the certificates uploaded to Qwilt CDN by your organization and the associated metadata.

## Example Usage

```terraform
data "qwilt_cdn_certificates" "certificates_list" {
  filter = {
    cert_id = "all"
  }
}

data "qwilt_cdn_certificates" "detail" {
  filter = {
    cert_id = 25
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (Attributes) Data source filter attributes. (see [below for nested schema](#nestedatt--filter))

### Read-Only

- `certificate` (Attributes List) List of certificates. (see [below for nested schema](#nestedatt--certificate))

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Optional:

- `cert_id` (Number) The ID of the specific certificate you want to retrieve.


<a id="nestedatt--certificate"></a>
### Nested Schema for `certificate`

Required:

- `certificate` (String) A single X.509 signed PEM certificate, encoded in Base64.
- `certificate_chain` (String) An ordered concatenation of PEM-encoded signed certificates. The first is the signer of the imported certificate, and the last is an intermediate CA signed by a well known Root CA. The whole string must be Base64 encoded.

Optional:

- `description` (String) The certificate description.

Read-Only:

- `cert_id` (Number) The unique identifier of the certificate. The certId will be needed when you add the certificate configuration and when you assign it to a site.
- `domain` (String) The site host domain the certificate is linked to.
- `pk_hash` (String) A unique identifier for the private key that does not expose the actual key itself.
- `status` (String) The status of the certificate:["ISSUED",
          "ACTIVE",
          "EXPIRED",
          "REVOKED"].
- `tenant` (String) The organization your user is assigned to.
- `type` (String) The certificate's type.