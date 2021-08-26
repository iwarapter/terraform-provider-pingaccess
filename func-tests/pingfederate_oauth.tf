resource "pingaccess_pingfederate_oauth" "demo" {
  count                  = local.isPA6 ? 0 : 1
  access_validator_id    = 1
  cache_tokens           = true
  subject_attribute_name = "san"
  name                   = "PingFederate"
  client_id              = "oauth"
  client_secret {
    value = "top_secret"
  }
  send_audience              = true
  token_time_to_live_seconds = 300
  use_token_introspection    = true
}

resource "pingaccess_pingfederate_oauth" "demo2" {
  count                  = local.isPA6 ? 1 : 0
  access_validator_id    = 1
  cache_tokens           = true
  subject_attribute_name = "san"
  name                   = "PingFederate"
  client_credentials {
    client_id = "oauth"
    client_secret {
      value = "top_secret"
    }
  }
  send_audience              = true
  token_time_to_live_seconds = 300
  use_token_introspection    = true
}
