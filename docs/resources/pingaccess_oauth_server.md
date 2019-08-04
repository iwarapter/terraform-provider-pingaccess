Provides an Authorization Server configuration.

### Example Usage
```terraform
{!../pingaccess/test_cases/oauth_server.tf!}
```
### Argument Attributes
The following arguments are supported:

- [`audit_level`](#audit_level) - Enable to record requests to third-party OAuth 2.0 Authorization Server to the audit store.

- [`cache_tokens`](#cache_tokens) - Enable to retain token details for subsequent requests.

- [`client_credentials`](#client_credentials) - Specify the client credentials.

- [`description`](#description) - The description of the third-party OAuth 2.0 Authorization Server.

- [`introspection_endpoint`](#introspection_endpoint) - The third-party OAuth 2.0 Authorization Server's token introspection endpoint.

- [`secure`](#secure) - Enable if third-party OAuth 2.0 Authorization Server is expecting HTTPS connections.

- [`send_audience`](#send_audience) - Enable to send the URI the user requested as the 'aud' OAuth parameter for PingAccess to the OAuth 2.0 Authorization server.

- [`subject_attribute_name`](#subject_attribute_name) - The attribute you want to use from the OAuth access token as the subject for auditing purposes.

- [`targets`](#targets) - One or more server hostname:port pairs used to access third-party OAuth 2.0 Authorization Server.

- [`token_time_to_live_seconds`](#token_time_to_live_seconds) - Defines the number of seconds to cache the access token. -1 means no limit. This value should be less than the PingFederate Token Lifetime.

- [`trusted_certificate_group_id`](#trusted_certificate_group_id) - The group of certificates to use when authenticating to third-party OAuth 2.0 Authorization Server.

- [`use_proxy`](#use_proxy) - True if a proxy should be used for HTTP or HTTPS requests.


### Attributes Reference

No additional attributes are provided.
