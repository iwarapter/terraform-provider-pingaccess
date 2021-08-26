resource "pingaccess_rule" "example" {
  class_name    = "com.pingidentity.pa.policy.CIDRPolicyInterceptor"
  name          = "example"
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
