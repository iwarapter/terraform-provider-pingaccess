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
