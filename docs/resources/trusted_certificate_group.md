---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingaccess_trusted_certificate_group Resource - terraform-provider-pingaccess"
subcategory: ""
description: |-
  Provides configuration for Trusted Certificate Groups within PingAccess.
---

# pingaccess_trusted_certificate_group (Resource)

Provides configuration for Trusted Certificate Groups within PingAccess.

## Example Usage

```terraform
resource "pingaccess_trusted_certificate_group" "example" {
  name                        = "example"
  use_java_trust_store        = true
  skip_certificate_date_check = false
  cert_ids                    = [pingaccess_certificate.example.id]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of trusted certificate group.

### Optional

- `cert_ids` (List of String) The IDs of the certificates that are in the trusted certificate group.
- `ignore_all_certificate_errors` (Boolean, Deprecated) This field is read-only and is only set to true for the Trust Any certificate group.
- `skip_certificate_date_check` (Boolean) This field is true if certificates that have expired or are not yet valid but have passed the other certificate checks should be trusted.
- `use_java_trust_store` (Boolean) This field is true if the certificates in the group should also include all certificates in the Java Trust Store.

### Read-Only

- `id` (String) The ID of this resource.
- `system_group` (Boolean) This field is read-only and indicates the trusted certificate group cannot be modified.

## Import

Import is supported using the following syntax:

```shell
terraform import pingaccess_trusted_certificate_group.example 123
```
