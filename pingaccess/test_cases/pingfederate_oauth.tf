resource "pingaccess_pingfederate_oauth" "demo" {
  access_validator_id    = 1
  cache_tokens           = true
  subject_attribute_name = "san"
  name                   = "foo"
  client_id              = "oauth"
  client_secret {
    value = "top_secret"
  }
  send_audience              = true
  token_time_to_live_seconds = 300
  use_token_introspection    = true
}
