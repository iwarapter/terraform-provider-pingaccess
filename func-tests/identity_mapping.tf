locals {
  pa62_idm_conf = <<EOF
  {
    "exclusionList": false,
    "exclusionListAttributes": [],
    "exclusionListSubject": null,
    "headerNamePrefix": null,
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

  default_idm_conf = <<EOF
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

resource "pingaccess_identity_mapping" "demo_identity_mapping" {
  class_name = "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"
  name       = "demo_identity_mapping"

  configuration = local.isPA6_2 ? local.pa62_idm_conf : local.default_idm_conf
}
