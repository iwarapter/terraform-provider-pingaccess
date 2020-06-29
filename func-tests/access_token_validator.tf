resource "pingaccess_access_token_validator" "demo_access_token_validator" {
  class_name = "com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"
  name       = "demo"

  configuration = <<EOF
	{
		"description": null,
		"path": "/bar",
		"subjectAttributeName": "demo",
		"issuer": null,
		"audience": null
	}
	EOF
}
