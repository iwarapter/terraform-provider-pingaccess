provider "pingaccess" {
  // password = "2Access2"
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

// resource "pingaccess_site" "acc_test_site" {
//   name                       = "demo_site"
//   targets                    = ["localhost:1234"]
//   max_connections            = -1
//   max_web_socket_connections = -1
//   availability_profile_id    = 1
// }

// resource "pingaccess_virtualhost" "acc_test_virtualhost" {
//   host                         = "demo-localhost"
//   port                         = 3000
//   agent_resource_cache_ttl     = 900
//   key_pair_id                  = 0
//   trusted_certificate_group_id = 0
// }

// resource "pingaccess_application" "acc_test" {
//   // access_validator_id = 0
//   application_type    = "Web"
//   agent_id            = 0
//   name                = "demo-app"
//   context_root        = "/bar"
//   default_auth_type   = "Web"
//   destination         = "Site"
//   site_id             = "${pingaccess_site.acc_test_site.id}"
//   virtual_host_ids    = ["${pingaccess_virtualhost.acc_test_virtualhost.id}"]

//   // identity_mapping_ids {
//   //   web = "${pingaccess_identity_mapping.idm_foo.id}"

//   //   //   api = 0
//   // }

//   // identity_mapping_ids {
//   //   web = "65"
//   // }
//   // web_session_id = "${pingaccess_websession.demo_session.id}"
// }

// resource "pingaccess_identity_mapping" "idm_foo" {
//   class_name = "com.pingidentity.pa.identitymappings.HeaderIdentityMapping"
//   name       = "Foo"

//   configuration = <<EOF
//   {
//     "attributeHeaderMappings": [
//       {
//         "subject": true,
//         "attributeName": "FOO",
//         "headerName": "FOO"
//       }
//     ],
//     "headerClientCertificateMappings": []
//   }
//   EOF
// }

// resource "pingaccess_ruleset" "ruleset_one" {
//   name             = "demo-ruleset"
//   success_criteria = "SuccessIfAnyOneSucceeds"
//   element_type     = "Rule"

//   policy = [
//     "${pingaccess_rule.ruleset_rule_one.id}",
//   ]
// }

// resource "pingaccess_rule" "ruleset_rule_one" {
//   class_name = "com.pingidentity.pa.policy.CIDRPolicyInterceptor"
//   name       = "demo-ruleset-rule"

//   supported_destinations = [
//     "Site",
//     "Agent",
//   ]

//   configuration = <<EOF
//   {
//     "cidrNotation": "127.0.0.1/32",
//     "negate": false,
//     "overrideIpSource": false,
//     "headers": [],
//     "headerValueLocation": "LAST",
//     "fallbackToLastHopIp": true,
//     "errorResponseCode": 404,
//     "errorResponseStatusMsg": "Forbidden",
//     "errorResponseTemplateFile": "policy.error.page.template.html",
//     "errorResponseContentType": "text/html;charset=UTF-8",
//     "rejectionHandler": null,
//     "rejectionHandlingEnabled": false
//   }
//   EOF
// }

// resource "pingaccess_websession" "demo_session" {
// 	name = "demo-session"
// 	audience = "woot"
// 	client_credentials {
// 		client_id = "demo-client",
//     client_secret {
// 			value = "atat"
// 		}
// 	}
//   request_profile = false
// }

// resource "pingaccess_application_resource" "woot" {
//   name = "Root Resource"

//   methods = [
//     "*",
//   ]

//   path_prefixes = [
//     "/*",
//   ]

//   default_auth_type_override = "Web"
//   audit_level                = "OFF"
//   anonymous                  = false
//   enabled                    = true

//   // policy {
//   //   "Web": [],
//   //   "API": []
//   // },
//   root_resource = true

//   application_id = "${pingaccess_application.acc_test.id}"

//   policy {
//     web {
//       id   = "${pingaccess_rule.ruleset_rule_one.id}"
//       type = "Rule"
//     }
//   }
// }

resource "pingaccess_trusted_certificate_group" "demo_tcg" {
  name = "test"
  cert_ids = [
    pingaccess_certificate.demo_certificate.id
  ]
}

resource "pingaccess_certificate" "demo_certificate" {
  alias     = "cert2"
  file_data = "${base64encode(file("amazon_root_ca1.pem"))}"
}

// data "pingaccess_certificate" "data_test" {
//   alias = "string"
// }

// output "cert_id" {
//   value = "${data.pingaccess_certificate.data_test.id}"
// }

// output "cert_alias" {
//   value = "${data.pingaccess_certificate.data_test.alias}"
// }

// output "cert_expires" {
//   value = "${data.pingaccess_certificate.data_test.expires}"
// }

// output "cert_issuer_dn" {
//   value = "${data.pingaccess_certificate.data_test.issuer_dn}"
// }

// output "cert_md5sum" {
//   value = "${data.pingaccess_certificate.data_test.md5sum}"
// }

// output "cert_serial_number" {
//   value = "${data.pingaccess_certificate.data_test.serial_number}"
// }

// output "cert_sha1sum" {
//   value = "${data.pingaccess_certificate.data_test.sha1sum}"
// }

// output "cert_signature_algorithm" {
//   value = "${data.pingaccess_certificate.data_test.signature_algorithm}"
// }

// output "cert_status" {
//   value = "${data.pingaccess_certificate.data_test.status}"
// }

// output "cert_subject_dn" {
//   value = "${data.pingaccess_certificate.data_test.subject_dn}"
// }

// output "cert_valid_from" {
//   value = "${data.pingaccess_certificate.data_test.valid_from}"
// }
