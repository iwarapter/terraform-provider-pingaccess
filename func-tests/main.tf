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
