resource "pingaccess_auth_token_management" "demo" {
  key_roll_enabled = true
  key_roll_period_in_hours = 24
  issuer = "PingAccessAuthToken"
  signing_algorithm = "P-256"
}
