resource "pingaccess_load_balancing_strategy" "test" {
  name          = "example"
  class_name    = "com.pingidentity.pa.ha.lb.header.HeaderBasedLoadBalancingPlugin"
  configuration = <<EOF
  {
    "headerName": "example",
    "fallbackToFirstAvailableHost": false
  }
  EOF
}
