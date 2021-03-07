---
page_title: "Structured Plugin Configuration"
---

# Structured Plugin Configuration

Several of the resources within this provider represent plugin configuration on PingAccess.
Resources like `pingaccess_access_token_validator` or `pingaccess_site_authenticator` for example.
These configuration blocks are ultimately json representing the underlying plugin descriptor, which can be seen in the example below.

### Json Configuration Style
```hcl
resource "pingaccess_access_token_validator" "example" {
  class_name    = "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"
  name          = "demo"
  configuration = <<EOF
  {
    "description": null,
    "path": "/bar",
    "subjectAttributeName": "demo",
    "issuer": null,
    "audience": null
  }
  EOF
}
```
The json structure will eventually be deprecated in favour of the structured configuration style.

### Structured Configuration Style
With the new release of the provider we can now support a dynamically structured HCL block representing the various configuration fields.

```hcl
resource "pingaccess_access_token_validator" "example" {
  class_name    = "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"
  name          = "demo"
  configuration = {
    "path"                 = "/foo"
    "subjectAttributeName" = "foo"
  }
}
```
This allows for cleaner diffs especially with the concise changes introduced in terraform 0.14 as the individual attributes will be highlighted instead of the entire json payload.

Additionally, any variables marked as `sensitive` that are used for an attribute will now only mask that specific attribute and not the entire json configuration block.
