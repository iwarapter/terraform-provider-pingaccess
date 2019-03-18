provider "pingaccess" {
  password = "2Access2"
}

// resource "pingaccess_rule" "my-server" {
//   class_name = "com.pingidentity.pa.policy.CIDRPolicyInterceptor"
//   name = "demo test"
//   supported_destinations = [
//     "Site",
//     "Agent"
//   ]
//   configuration = <<EOF
//   {
//     "cidrNotation": "127.0.0.1/32",
//     "negate": false,
//     "overrideIpSource": false,
//     "headers": [],
//     "headerValueLocation": "LAST",
//     "fallbackToLastHopIp": true,
//     "errorResponseCode": 403,
//     "errorResponseStatusMsg": "Forbidden",
//     "errorResponseTemplateFile": "policy.error.page.template.html",
//     "errorResponseContentType": "text/html;charset=UTF-8"
//   }
//   EOF
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
  name                       = "demo_site"
  targets                    = ["localhost:1234"]
  max_connections            = -1
  max_web_socket_connections = -1
  availability_profile_id    = 1
}

resource "pingaccess_virtualhost" "acc_test_virtualhost" {
  host                         = "demo-localhost"
  port                         = 3000
  agent_resource_cache_ttl     = 900
  key_pair_id                  = 0
  trusted_certificate_group_id = 0
}

resource "pingaccess_application" "acc_test" {
  access_validator_id = 0
  application_type    = "Web"
  agent_id            = 0
  name                = "demo-app"
  context_root        = "/bar"
  default_auth_type   = "Web"
  destination         = "Site"
  site_id             = "${pingaccess_site.acc_test_site.id}"
  virtual_host_ids    = ["${pingaccess_virtualhost.acc_test_virtualhost.id}"]

  identity_mapping_ids {
    web = "${pingaccess_identity_mapping.idm_foo.id}"

    //   api = 0
  }

  // identity_mapping_ids {
  //   web = "65"
  // }
  web_session_id = "${pingaccess_websession.demo_session.id}"
}

resource "pingaccess_identity_mapping" "idm_foo" {
  class_name = "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"
  name       = "Foo"

  configuration = <<EOF
  {
    "attributeHeaderMappings": [
      {
        "subject": true,
        "attributeName": "FOO",
        "headerName": "FOO"
      }
    ],
    "headerClientCertificateMappings": []
  }
  EOF
}

resource "pingaccess_ruleset" "ruleset_one" {
  name             = "demo-ruleset"
  success_criteria = "SuccessIfAnyOneSucceeds"
  element_type     = "Rule"

  policy = [
    "${pingaccess_rule.ruleset_rule_one.id}",
  ]
}

resource "pingaccess_rule" "ruleset_rule_one" {
  class_name = "com.pingidentity.pa.policy.CIDRPolicyInterceptor"
  name       = "demo-ruleset-rule"

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

resource "pingaccess_websession" "demo_session" {
	name = "demo-session"
	audience = "woot"
	client_credentials {
		client_id = "demo-client",
    client_secret {
			value = "atat"
		}
	}
}

resource "pingaccess_application_resource" "woot" {
  name = "Root Resource"

  methods = [
    "*",
  ]

  path_prefixes = [
    "/*",
  ]

  default_auth_type_override = "Web"
  audit_level                = "OFF"
  anonymous                  = false
  enabled                    = true

  // policy {
  //   "Web": [],
  //   "API": []
  // },
  root_resource = true

  application_id = "${pingaccess_application.acc_test.id}"

  policy {
    web {
      id   = "${pingaccess_rule.ruleset_rule_one.id}"
      type = "Rule"
    }
  }
}
