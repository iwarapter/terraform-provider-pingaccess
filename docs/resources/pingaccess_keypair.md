#Resource: pingaccess_keypair

Provides a keypair.

!!! tip
    The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of the sensitive `password` attribute.

### Example Usage
```terraform
{!../pingaccess/test_cases/keypair.tf!}
```
### Argument Attributes
The following arguments are supported:

- [`alias`](#alias) - (required) The alias for the keypair.

- [`file_data`](#file_data) - (required) The base64-encoded data of the keypair.

- [`password`](#password) - The Password used to protect the PKCS#12 file.

- [`hsm_provider_id`](#hsm_provider_id) - The HSM Provider ID.

### Attributes Reference

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

### Import

PingAccess keypairs can be imported using the id, e.g.

```bash
$ terraform import pingaccess_keypair.demo_keypair 123
```