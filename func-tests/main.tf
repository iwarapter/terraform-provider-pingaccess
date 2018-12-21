provider "pingaccess" {}

// resource "pingaccess_rule" "my-server" {
//   class_name = "com.pingidentity.pa.policy.OAuthAttributeValuePolicyInterceptor"
//   name = "demo rule"
//   supported_destinations = [
//     "Site",
//     "Agent",
//   ]
//   configuration {
//     cidrNotation = "192.168.0.1/32"
//     negate = false
//     overrideIpSource = false
//     headers = false
//     headerValueLocation = "LAST"
//     fallbackToLastHopIp = true
//     errorResponseCode = 403
//     errorResponseStatusMsg = "Forbidden"
//     errorResponseTemplateFile = "policy.error.page.template.html"
//     errorResponseContentType = "text/html;charset=UTF-8"
//   }
// }

// resource "pingaccess_virtualhost" "localhost_3000" {
//   host                         = "localhost"
//   port                         = 3000
//   agent_resource_cache_ttl     = 900
//   key_pair_id                  = 0
//   trusted_certificate_group_id = 0
// }

// resource "pingaccess_site" "bar" {
//   name                       = "bar"
//   targets                    = ["localhost:1234"]
//   max_connections            = -1
//   max_web_socket_connections = -1
//   availability_profile_id    = 1
// }

resource "pingaccess_site" "acc_test_site" {
  name                       = "acc_test_site"
  targets                    = ["localhost:1234"]
  max_connections            = -1
  max_web_socket_connections = -1
  availability_profile_id    = 1
}

resource "pingaccess_virtualhost" "acc_test_virtualhost" {
  host                         = "localhost"
  port                         = 3000
  agent_resource_cache_ttl     = 900
  key_pair_id                  = 0
  trusted_certificate_group_id = 0
}

resource "pingaccess_application" "acc_test" {
  access_validator_id = 0
  application_type    = "Web"
  agent_id            = 0
  name                = "bar"
  context_root        = "/bar"
  default_auth_type   = "Web"
  destination         = "Site"
  site_id             = "${pingaccess_site.acc_test_site.id}"
  virtual_host_ids    = ["${pingaccess_virtualhost.acc_test_virtualhost.id}"]
}

resource "pingaccess_application_resource" "woot" {
  name = "woot"
  methods = [
    "*"
  ]
  path_prefixes = [
    "/woot"
  ]
  default_auth_type_override = "Web"
  audit_level = "OFF"
  anonymous = false
  enabled = true
  // policy {
  //   "Web": [],
  //   "API": []
  // },
  root_resource = false
  application_id = "${pingaccess_application.acc_test.id}"
}