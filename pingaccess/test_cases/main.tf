# tag::pingaccess_certificate[]   
resource "pingaccess_certificate" "demo_certificate" {
  alias     = "amazon_root_ca1"
  file_data = "${base64encode(file("amazon_root_ca1.pem"))}"
}

# end::pingaccess_certificate[]

# tag::pingaccess_site[]   
resource "pingaccess_site" "demo_site" {
  name    = "demo-site"
  targets = ["localhost:1234"]
}

# end::pingaccess_site[]

# tag::pingaccess_virtualhost[]   
resource "pingaccess_virtualhost" "demo_virtualhost" {
  host = "localhost"
  port = 1234
}

# end::pingaccess_virtualhost[]

# tag::pingaccess_site_authenticator[]
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

  hidden_fields = ["password"]
}

# end::pingaccess_site_authenticator[]

# tag::pingaccess_ruleset[]
resource "pingaccess_ruleset" "demo_ruleset" {
  name             = "demo_ruleset"
  success_criteria = "SuccessIfAllSucceed"
  element_type     = "Rule"

  policy = [
    "${pingaccess_rule.demo_rule.id}",
    "${pingaccess_rule.second_demo_rule.id}",
  ]
}

# end::pingaccess_ruleset[]

# tag::pingaccess_identity_mapping[]
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

# end::pingaccess_identity_mapping[]

# tag::pingaccess_rule[]
resource "pingaccess_rule" "demo_rule" {
  class_name = "com.pingidentity.pa.policy.CIDRPolicyInterceptor"
  name       = "demo_rule"

  supported_destinations = [
    "Site",
    "Agent",
  ]

  configuration = <<EOF
  {
    "cidrNotation": "127.0.0.1/32",
    "negate": false,
    "overrideIpSource": false,
    "headers": [],
    "headerValueLocation": "LAST",
    "fallbackToLastHopIp": true,
    "errorResponseCode": 404,
    "errorResponseStatusMsg": "Forbidden",
    "errorResponseTemplateFile": "policy.error.page.template.html",
    "errorResponseContentType": "text/html;charset=UTF-8",
    "rejectionHandler": null,
    "rejectionHandlingEnabled": false
  }
  EOF
}

# end::pingaccess_rule[]

# tag::pingaccess_third_party_service[]
resource "pingaccess_third_party_service" "demo_third_party_service" {
  name = "demo"

  targets = [
    "server.domain:1234",
  ]
}

# end::pingaccess_third_party_service[]

