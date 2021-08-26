resource "pingaccess_acme_server" "example" {
  name = "example"
  url  = "https://acme-staging-v02.api.letsencrypt.org/directory"
}
