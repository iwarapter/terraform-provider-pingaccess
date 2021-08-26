resource "pingaccess_engine_listener" "example" {
  name   = "example"
  port   = 443
  secure = true
}
