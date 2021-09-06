---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingaccess_trusted_certificate_group Data Source - terraform-provider-pingaccess"
subcategory: ""
description: |-
  Use this data source to get the ID of a trusted certificate group in Ping Access, you can reference it by name without having to hard code the IDs as input.
---

# pingaccess_trusted_certificate_group (Data Source)

Use this data source to get the ID of a trusted certificate group in Ping Access, you can reference it by name without having to hard code the IDs as input.

## Example Usage

```terraform
data "pingaccess_trusted_certificate_group" "trust_any" {
  name = "Trust Any"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) The name of trusted certificate group.

### Read-Only

- **cert_ids** (List of String) The IDs of the certificates that are in the trusted certificate group.
- **id** (String) The ID of this resource.
- **ignore_all_certificate_errors** (Boolean) This field is read-only and is only set to true for the Trust Any certificate group.
- **skip_certificate_date_check** (Boolean) This field is true if certificates that have expired or are not yet valid but have passed the other certificate checks should be trusted.
- **system_group** (Boolean) This field is read-only and indicates the trusted certificate group cannot be modified.
- **use_java_trust_store** (Boolean) This field is true if the certificates in the group should also include all certificates in the Java Trust Store.