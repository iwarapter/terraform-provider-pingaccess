resource "pingaccess_acme_server" "test" {
  count = local.isPA6 ? 1 : 0
  name  = "example"
  url   = "https://acme-staging-v02.api.letsencrypt.org/directory"
}
