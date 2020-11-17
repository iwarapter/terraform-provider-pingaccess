# Resource: pingaccess_keypair

Provides a keypair.

-> The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of the sensitive `password` attribute.

## Example Usage

### Generating a Keypair
```hcl
resource "pingaccess_keypair" "demo_generate" {
  alias             = "demo2"
  city              = "London"
  common_name       = "Example"
  country           = "GB"
  key_algorithm     = "RSA"
  key_size          = 2048
  organization      = "Test"
  organization_unit = "Development"
  state             = "London"
  valid_days        = 365
}
```

### Importing a Keypair

```hcl
resource "pingaccess_keypair" "demo_keypair" {
  alias     = "demo"
  file_data = filebase64("provider.p12")
  password  = "password"
}
```

## Argument Attributes
The following arguments are supported:

- [`alias`](#alias) - (required) The alias for the keypair.
- [`hsm_provider_id`](#hsm_provider_id) - The HSM Provider ID.

### Importing KeyPair

- [`file_data`](#file_data) - (required) The base64-encoded data of the keypair.
- [`password`](#password) - The Password used to protect the PKCS#12 file.

### Generating KeyPair

- [`city`](#city) - (Required) The city or other primary location (L) where the company operates.
- [`common_name`](#common_name) - (Required) The common name (CN) identifying the certificate.
- [`country`](#country) - (Required) The country (C) where the company is based, using two capital letters.
- [`key_algorithm`](#key_algorithm) - (Required) The key algorithm to use to generate a key.
- [`key_size`](#key_size) - (Required) The number of bits used in the key. Choices depend on selected key algorithm.
- [`organization`](#organization) - (Required) The organization (O) or company name creating the certificate.
- [`organization_unit`](#organization_unit) - (Required) The specific unit within the organization (OU).
- [`state`](#state) - (Required) The state (ST) or other political unit encompassing the location.
- [`valid_days`](#valid_days) - (Required) The number of days the certificate is valid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The keypairs's ID.

- [`chain_certificates`](#chain_certificates) The complete list of certificates in the key pair certificate chain.

    - [`chain_certificates.#.expires`](#chain_certificates-expires) - The date at which the certificate expires as the number of milliseconds since January 1, 1970, 00:00:00 GMT.

    - [`chain_certificates.#.issuer_dn`](#chain_certificates-issuer_dn) - The issuer DN for the certificate.

    - [`chain_certificates.#.md5sum`](#chain_certificates-md5sum) - The MD5 sum for the certificate.

    - [`chain_certificates.#.serial_number`](#chain_certificates-serial_number) - The serial number for the certificate.

    - [`chain_certificates.#.sha1sum`](#chain_certificates-sha1sum) - The SHA1 sum for the certificate.

    - [`chain_certificates.#.signature_algorithm`](#chain_certificates-signature_algorithm) -  The algorithm used to sign the certificate.

    - [`chain_certificates.#.subject_cn`](#chain_certificates-subject_cn) - The subject CN for the certificate.

    - [`chain_certificates.#.subject_dn`](#chain_certificates-subject_dn) - The subject DN for the certificate.

    - [`chain_certificates.#.valid_from`](#chain_certificates-valid_from) - The date at which the certificate is valid from as the number of milliseconds since January 1, 1970, 00:00:00 GMT.

- [`expires`](#expires) - The date at which the keypair expires as the number of milliseconds since January 1, 1970, 00:00:00 GMT.

- [`issuer_dn`](#issuer_dn) - The issuer DN for the keypair.

- [`md5sum`](#md5sum) - The MD5 sum for the keypair.

- [`serial_number`](#serial_number) - The serial number for the keypair.

- [`sha1sum`](#sha1sum) - The SHA1 sum for the keypair.

- [`signature_algorithm`](#signature_algorithm) -  The algorithm used to sign the keypair.

- [`subject_cn`](#subject_cn) - The subject CN for the keypair.

- [`subject_dn`](#subject_dn) - The subject DN for the keypair.

- [`valid_from`](#valid_from) - The date at which the keypair is valid from as the number of milliseconds since January 1, 1970, 00:00:00 GMT.

## Import

-> This is currently only supported for generated KeyPairs.

PingAccess keypairs can be imported using the id, e.g.

```bash
$ terraform import pingaccess_keypair.example 123
```
