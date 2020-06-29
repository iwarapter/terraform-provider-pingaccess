resource "pingaccess_virtualhost" "demo" {
  host                         = "demo"
  port                         = 4001
  agent_resource_cache_ttl     = 900
  key_pair_id                  = 0
  trusted_certificate_group_id = 0
}
