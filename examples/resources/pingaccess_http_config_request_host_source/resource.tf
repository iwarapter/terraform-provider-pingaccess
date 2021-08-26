resource "pingaccess_http_config_request_host_source" "example" {
  header_name_list    = ["Host", "X-Forwarded-Host"]
  list_value_location = "LAST"
}
