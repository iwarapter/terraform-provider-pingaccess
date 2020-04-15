#Resource: pingaccess_websession

Provides a web session.

!!! tip
    The PingAccess API does not provider repeatable means of querying a sensitive value, we are unable to detect configuration drift of any sensitive fields in the `client_credentials.client_secret` block.

## Example Usage
```terraform
{!../pingaccess/test_cases/websession.tf!}
```

## Argument Attributes

The following arguments are supported:

- [`audience`](#audience) -  Enter a unique identifier between 1 and 32 characters that defines who the PA Token is applicable to.

- [`cache_user_attributes`](#cache_user_attributes) -  Specify if PingAccess should cache user attribute information for use in policy decisions. When disabled, this data is encoded and stored in the session cookie.

- [`client_credentials`](#client_credentials) - Specify the client credentials.
  
    - [`client_credentials.client_id`](#client_credentials-client_id) - Specify the client ID.
    
    - [`client_credentials.client_secret`](#client_credentials-client_secret) - Specify the client secret.

- [`cookie_domain`](#cookie_domain) -  The domain where the cookie is stored--for example, corp.yourcompany.com.

- [`cookie_type`](#cookie_type) - ['Encrypted' or 'Signed']:  Specify an Encrypted JWT or a Signed JWT web session cookie.

- [`enable_refresh_user`](#enable_refresh_user) -  Specify if you want to have PingAccess periodically refresh user data from PingFederate for use in policy decisions.

- [`http_only_cookie`](#http_only_cookie) -  Enable the HttpOnly flag on cookies that contain the PA Token.

- [`idle_timeout_in_minutes`](#idle_timeout_in_minutes) -  The length of time you want the PingAccess Token to remain active when no activity is detected.

- [`name`](#name) -  The name of the web session.
- [`oidc_login_type`](#oidc_login_type) - ['Code' or 'POST' or 'x_post']:  The web session token type.

- [`pfsession_state_cache_in_seconds`](#pfsession_state_cache_in_seconds) -  Specify the number of seconds to cache PingFederate Session State information.

- [`refresh_user_info_claims_interval`](#refresh_user_info_claims_interval) -  Specify the maximum number of seconds to cache user attribute information when the Refresh User is enabled.

- [`request_preservation_type`](#request_preservation_type) - ['None' or 'POST' or 'All']:  Specify the types of request data to be preserved if the user is redirected to an authentication page when submitting information to a protected resource.

- [`request_profile`](#request_profile) - Specifies whether the default scopes ('profile', 'email', 'address', and 'phone') should be specified in the access request. (DEPRECATED - to be removed in a future release; please use 'scopes' instead)

- [`scopes`](#scopes) - The list of scopes to be specified in the access request. If not specified, the default scopes ('profile', 'email', 'address', and 'phone') will be used.

- [`secure_cookie`](#secure_cookie) -  Specify whether the PingAccess Cookie must be sent using only HTTPS connections.

- [`send_requested_url_to_provider`](#send_requested_url_to_provider) -  Specify if you want to send the requested URL as part of the authentication request to the OpenID Connect Provider.

- [`session_timeout_in_minutes`](#session_timeout_in_minutes) -  The length of time you want the PA Token to remain active. Once the PA Token expires, an authenticated user must re-authenticate.

- [`validate_session_is_alive`](#validate_session_is_alive) -  Specify if PingAccess should validate sessions with the configured PingFederate instance during request processing.

- [`web_storage_type`](#web_storage_type) - ['SessionStorage' or 'LocalStorage']:  Specify the type of web storage to use for request preservation data. Default is SessionStorage.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The web session's ID.

## Import

PingAccess web session can be imported using the id, e.g.

```bash
$ terraform import pingaccess_websession.demo_websession 123
```
