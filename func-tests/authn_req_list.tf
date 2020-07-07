resource "pingaccess_authn_req_list" "demo" {
  name = "demo"
  authn_reqs = [
    "one",
    "two",
  ]
}
