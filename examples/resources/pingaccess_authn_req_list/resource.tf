resource "pingaccess_authn_req_list" "example" {
  name       = "demo"
  authn_reqs = ["one", "two"]
}
