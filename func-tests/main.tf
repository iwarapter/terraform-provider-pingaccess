terraform {
  required_providers {
    pingaccess = {
      source  = "iwarapter/pingaccess"
      version = "0.0.1-ci" #for functional testing
    }
  }
}

provider "pingaccess" {
  password = "2Access"
}

data "pingaccess_trusted_certificate_group" "trust_any" {
  name = "Trust Any"
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
  count = var.pa6 ? 1 : 0
  name  = "foo"
  url   = "https://host.docker.internal:14000/dir"
}

resource "pingaccess_pingfederate_runtime" "pa6_demo" {
  count                        = var.pa6 ? 1 : 0
  description                  = "demo"
  issuer                       = "https://pingfederate:9031"
  sts_token_exchange_endpoint  = "https://foo/bar"
  skip_hostname_verification   = true
  use_slo                      = false
  trusted_certificate_group_id = data.pingaccess_trusted_certificate_group.trust_any.id
  use_proxy                    = false
}

resource "pingaccess_pingfederate_runtime" "pa5_demo" {
  count                        = var.pa6 ? 0 : 1
  host                         = "pingfederate"
  port                         = 9031
  skip_hostname_verification   = true
  use_slo                      = false
  trusted_certificate_group_id = data.pingaccess_trusted_certificate_group.trust_any.id
  use_proxy                    = true
}
