# this is an example using a dynamic type in terraform allowing for the json
# configuration structure to be written in HCL
resource "pingaccess_site_authenticator" "structured_example" {
  name       = "structured_example"
  class_name = "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"
  configuration = {
    "username" = "example"
    "password" = {
      "value" = "top_5ecr37"
    }
  }
}

# legacy example of using heredoc to pass in the json configuration string
resource "pingaccess_site_authenticator" "json_example" {
  name          = "json_example"
  class_name    = "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"
  configuration = <<EOF
    {
      "username": "example",
      "password": {
        "value": "top_5ecr37"
      }
    }
    EOF
}
