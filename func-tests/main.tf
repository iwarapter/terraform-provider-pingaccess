provider "pingaccess" {
  password = "2FederateM0re"
}

resource "pingaccess_site" "demo" {
  name                       = "demo"
  targets                    = ["localhost:4321"]
  max_connections            = -1
  max_web_socket_connections = -1
  availability_profile_id    = 1
}

resource "pingaccess_virtualhost" "demo" {
  host                         = "demo"
  port                         = 4001
  agent_resource_cache_ttl     = 900
  key_pair_id                  = 0
  trusted_certificate_group_id = 0
}

resource "pingaccess_application" "demo" {
  access_validator_id = 0
  application_type    = "API"
  agent_id            = 0
  name                = "demo"
  context_root        = "/"
  default_auth_type   = "API"
  destination         = "Site"
  site_id             = pingaccess_site.demo.id
  virtual_host_ids    = [pingaccess_virtualhost.demo.id]

  policy {
    api {
      type = "Rule"
      id   = pingaccess_rule.demo_1.id
    }
    api {
      type = "Rule"
      id   = pingaccess_rule.demo_2.id
    }
  }
}

resource "pingaccess_rule" "demo_1" {
  class_name = "com.pingidentity.pa.policy.CIDRPolicyInterceptor"
  name       = "demo_1"
  supported_destinations = [
    "Site",
    "Agent"
  ]
  configuration = <<EOF
		{
			"cidrNotation": "127.0.0.1/32",
			"negate": false,
			"overrideIpSource": false,
			"headers": [],
			"headerValueLocation": "LAST",
			"fallbackToLastHopIp": true,
			"errorResponseCode": 403,
			"errorResponseStatusMsg": "Forbidden",
			"errorResponseTemplateFile": "policy.error.page.template.html",
			"errorResponseContentType": "text/html;charset=UTF-8",
			"rejectionHandler": null,
			"rejectionHandlingEnabled": false
		}
		EOF
}

resource "pingaccess_rule" "demo_2" {
  class_name = "com.pingidentity.pa.policy.CIDRPolicyInterceptor"
  name       = "demo_2"
  supported_destinations = [
    "Site",
    "Agent"
  ]
  configuration = <<EOF
  {
    "cidrNotation": "127.0.0.${pingaccess_site.demo.id}/32",
    "negate": false,
    "overrideIpSource": false,
    "headers": [],
    "headerValueLocation": "LAST",
    "fallbackToLastHopIp": true,
    "errorResponseCode": 403,
    "errorResponseStatusMsg": "Forbidden",
    "errorResponseTemplateFile": "policy.error.page.template.html",
    "errorResponseContentType": "text/html;charset=UTF-8",
    "rejectionHandler": null,
    "rejectionHandlingEnabled": false
  }
  EOF
}

resource "pingaccess_acme_server" "acc_test" {
  name = "foo"
  url  = "https://host.docker.internal:14000/dir"
}
