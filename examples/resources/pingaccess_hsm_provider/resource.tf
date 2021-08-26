resource "pingaccess_hsm_provider" "example" {
  class_name    = "com.pingidentity.pa.hsm.cloudhsm.plugin.AwsCloudHsmProvider"
  name          = "example"
  configuration = <<EOF
  {
    "user": "bob",
    "password": "top_secret",
    "partition": "p1"
  }
  EOF
}
