# this is an example using a dynamic type in terraform allowing for the json
# configuration structure to be written in HCL
resource "pingaccess_access_token_validator" "dynamic" {
  class_name = "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"
  name       = "demo_two"

  configuration = {
    "path"                 = "/bar"
    "subjectAttributeName" = "demo"
  }
}

# legacy example of using heredoc to pass in the json configuration string
resource "pingaccess_access_token_validator" "heredoc" {
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
