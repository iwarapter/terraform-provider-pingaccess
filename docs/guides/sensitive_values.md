---
page_title: "Sensitive Values with PingAccess"
---

# Sensitive Values with PingAccess

Several of the API's in PingAccess pass configuration which contains sensitive values.

## For PingAccess 5.x -> 6.0.x

The API will return an encrypted value back.
However, the value is different for every request, irrespective of whether this has changed or not in the API.

Because of this behaviour you cannot rely on the provider to correctly detect any drift in sensitive values for any resources.

## For PingAccess 6.1.x -> Above

The API still returns an encrypted value, however the value is static for subsequent requests.

This means we can detect drift from a value which the provider _**originally sets**_, this is currently supported for only the following resources:

- `pingaccess_websession`
- `pingaccess_oauth_server`
