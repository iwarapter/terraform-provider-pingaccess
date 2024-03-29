# Proxied Token Provider (PingFederate Runtime Application) Example
resource "pingaccess_pingfederate_runtime" "example_proxied" {
  targets                      = ["localhost:9031"]
  audit_level                  = "ON"
  skip_hostname_verification   = true
  use_slo                      = false
  trusted_certificate_group_id = 2
  use_proxy                    = true
  application {
    primary_virtual_host_id = 1
  }
  back_channel_secure = true
}

# Standard Token Provider Example
resource "pingaccess_pingfederate_runtime" "example_standard" {
  description                  = "foo"
  issuer                       = "https://localhost:9031"
  skip_hostname_verification   = true
  use_slo                      = false
  trusted_certificate_group_id = 2
  use_proxy                    = true
}
