---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingaccess_keypair Data Source - terraform-provider-pingaccess"
subcategory: ""
description: |-
  Use this data source to get keypair information in the PingAccess instance.
---

# pingaccess_keypair (Data Source)

Use this data source to get keypair information in the PingAccess instance.

## Example Usage

```terraform
data "pingaccess_keypair" "admin" {
  alias = "Generated: ADMIN"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **alias** (String) A unique alias name to identify the key pair. Special characters and spaces are allowed.

### Read-Only

- **chain_certificates** (Set of Object) The complete list of certificates in the key pair certificate chain. (see [below for nested schema](#nestedatt--chain_certificates))
- **csr_pending** (Boolean) True if a CSR is generated for this key pair.
- **expires** (Number) The date at which the certificate expires as the number of milliseconds since January 1, 1970, 00:00:00 GMT.
- **hsm_provider_id** (Number) The HSM Provider ID. The default value is 0 indicating an HSM is not used for this key pair.
- **id** (String) The ID of this resource.
- **issuer_dn** (String) The issuer DN for the certificate.
- **md5sum** (String) The MD5 sum for the certificate. The value will be set to "" when in FIPS mode.
- **serial_number** (String) The serial number for the certificate.
- **sha1sum** (String) The SHA1 sum for the certificate.
- **signature_algorithm** (String) The algorithm used to sign the certificate.
- **status** (String) A high-level status for the certificate.
- **subject_cn** (String) The subject CN for the certificate.
- **subject_dn** (String) The subject DN for the certificate.
- **valid_from** (Number) The date at which the certificate is valid from as the number of milliseconds since January 1, 1970, 00:00:00 GMT.

<a id="nestedatt--chain_certificates"></a>
### Nested Schema for `chain_certificates`

Read-Only:

- **alias** (String)
- **expires** (Number)
- **issuer_dn** (String)
- **md5sum** (String)
- **serial_number** (String)
- **sha1sum** (String)
- **signature_algorithm** (String)
- **status** (String)
- **subject_cn** (String)
- **subject_dn** (String)
- **valid_from** (Number)