---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingaccess_certificate Data Source - terraform-provider-pingaccess"
subcategory: ""
description: |-
  Use this data source to get certificate information in the PingAccess instance.
---

# pingaccess_certificate (Data Source)

Use this data source to get certificate information in the PingAccess instance.

## Example Usage

```terraform
data "pingaccess_certificate" "example" {
  alias = "example"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `alias` (String) The alias for the certificate.

### Read-Only

- `expires` (Number) The date at which the certificate expires as the number of milliseconds since January 1, 1970, 00:00:00 GMT.
- `id` (String) The ID of this resource.
- `issuer_dn` (String) The issuer DN for the certificate.
- `md5sum` (String) The MD5 sum for the certificate. The value will be set to "" when in FIPS mode.
- `serial_number` (String) The serial number for the certificate.
- `sha1sum` (String) The SHA1 sum for the certificate.
- `signature_algorithm` (String) The algorithm used to sign the certificate.
- `status` (String) A high-level status for the certificate.
- `subject_cn` (String) The subject CN for the certificate.
- `subject_dn` (String) The subject DN for the certificate.
- `valid_from` (Number) The date at which the certificate is valid from as the number of milliseconds since January 1, 1970, 00:00:00 GMT.
