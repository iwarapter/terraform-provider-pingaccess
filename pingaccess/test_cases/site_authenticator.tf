resource "pingaccess_site_authenticator" "demo_site_authenticator" {
  name       = "demo-site-authenticator"
  class_name = "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"

  configuration = <<EOF
    {
      "username": "cheese",
      "password": {
        "value": "top_5ecr37"
      }
    }
    EOF
}
