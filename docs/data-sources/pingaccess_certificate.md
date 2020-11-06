# Data Source: pingaccess_certificate

Use this data source to get the ID of a certificate in Ping Access, you can reference it by alias without having to hard code the IDs as input.

## Example Usage
```hcl
data "pingaccess_certificate" "demo_certificate" {
  alias = "amazon_root_ca1"
}
```
## Argument Attributes
The following arguments are supported:

- [`alias`](#alias) - (required) The alias for the certificate.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The certificates's ID.

- [`expires`](#expires) - The date at which the certificate expires as the number of milliseconds since January 1, 1970, 00:00:00 GMT.

- [`issuer_dn`](#issuer_dn) - The issuer DN for the certificate.

- [`md5sum`](#md5sum) - The MD5 sum for the certificate.

- [`serial_number`](#serial_number) - The serial number for the certificate.

- [`sha1sum`](#sha1sum) - The SHA1 sum for the certificate.

- [`signature_algorithm`](#signature_algorithm) -  The algorithm used to sign the certificate.

- [`subject_cn`](#subject_cn) - The subject CN for the certificate.

- [`subject_dn`](#subject_dn) - The subject DN for the certificate.

- [`valid_from`](#valid_from) - The date at which the certificate is valid from as the number of milliseconds since January 1, 1970, 00:00:00 GMT.
