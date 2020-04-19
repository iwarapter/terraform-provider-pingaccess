resource "pingaccess_pingfederate_runtime" "demo" {
  description                  = "foo"
  issuer                       = "https://localhost:9031"
  skip_hostname_verification   = true
  use_slo                      = false
  trusted_certificate_group_id = 2
  use_proxy                    = true
}
