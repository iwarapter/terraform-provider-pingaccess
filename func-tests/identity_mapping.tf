resource "pingaccess_identity_mapping" "demo_identity_mapping" {
  class_name = "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"
  name       = "demo_identity_mapping"

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
