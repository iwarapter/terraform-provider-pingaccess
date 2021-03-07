resource "pingaccess_access_token_validator" "demo_one" {
  class_name = "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"
  name       = "demo"

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

resource "pingaccess_access_token_validator" "demo_two" {
  class_name = "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"
  name       = "demo_two"

  configuration = {
    "path"                 = "/bar"
    "subjectAttributeName" = "demo"
  }
}
