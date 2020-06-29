resource "pingaccess_https_listener" "demo" {
  name                          = "ADMIN"
  key_pair_id                   = 1
  use_server_cipher_suite_order = true
}
