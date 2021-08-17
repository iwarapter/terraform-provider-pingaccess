resource "pingaccess_site_authenticator" "demo_site_authenticator" {
  name       = "demo-site-authenticator"
  class_name = "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"

  configuration = <<EOF
    {
      "username": "demo",
      "password": {
        "value": "top_5ecr37"
      }
    }
    EOF
}


resource "pingaccess_site_authenticator" "demo_2" {
  name       = "demo-2"
  class_name = "com.pingidentity.pa.siteauthenticators.BasicAuthTargetSiteAuthenticator"

  configuration = {
    "username" = "demo"
    "password" = {
      "value" = "top_5ecr37"
    }
  }
}

resource "pingaccess_site_authenticator" "issue_72_json" {
  name          = "issue 72 jsomn"
  class_name    = "com.pingidentity.pa.siteauthenticators.MutualTlsSiteAuthenticator"
  configuration = <<EOF
  {
    "keyPairId": "${pingaccess_keypair.demo_keypair.id}"
  }
  EOF
}

resource "pingaccess_site_authenticator" "issue_72" {
  name       = "issue 72"
  class_name = "com.pingidentity.pa.siteauthenticators.MutualTlsSiteAuthenticator"
  configuration = {
    "keyPairId" = pingaccess_keypair.demo_keypair.id
  }
}
