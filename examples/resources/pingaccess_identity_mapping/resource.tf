resource "pingaccess_identity_mapping" "example" {
  class_name    = "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"
  name          = "example"
  configuration = <<EOF
  {
    "attributeHeaderMappings": [
      {
        "subject": true,
        "attributeName": "sub",
        "headerName": "sub"
      }
    ],
    "headerClientCertificateMappings": []
  }
  EOF
}
