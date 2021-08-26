resource "pingaccess_availability_profile" "example" {
  name          = "example"
  class_name    = "com.pingidentity.pa.ha.availability.ondemand.OnDemandAvailabilityPlugin"
  configuration = <<EOF
  {
    "connectTimeout": 10000,
    "pooledConnectionTimeout": -1,
    "readTimeout": -1,
    "maxRetries": 2,
    "retryDelay": 250,
    "failedRetryTimeout": 60,
    "failureHttpStatusCodes": []
  }
  EOF
}
