# Resource: pingaccess_application

Provides a application.

## Example Usage
```hcl
resource "pingaccess_application" "example" {
  application_type = "Web"
  name             = "demo"
  context_root     = "/"
  destination      = "Site"
  site_id          = pingaccess_site.demo.id
  virtual_host_ids = [pingaccess_virtualhost.demo.id]

  policy {
    web {
      type = "Rule"
      id   = pingaccess_rule.demo_rule.id
    }
  }
}

```

## Argument Attributes

The following arguments are supported:

- [`access_validator_id`](#access_validator_id) - Non-zero if the application is protected by an Authorization Server. Only applies to applications of type API.

- [`agent_id`](#agent_id) - The ID of the agent associated with the application or zero if none.

- [`application_type`](#application_type) - ['Web' or 'API' or 'Dynamic']: The type of application.

- [`case_sensitive_path`](#case_sensitive_path) - True if the path is case sensitive.

- [`context_root`](#context_root) -  The context root of the application.

- [`default_auth_type`](#default_auth_type) - 'Web' or 'API', For Web + API applications (dynamic) defaultAuthType selects the processing mode when a request: does not have a token (web session, OAuth bearer) or has both tokens. This setting applies to all resources in the application except where overridden with defaultAuthTypeOverride.

- [`description`](#description) - A description of the application.

- [`destination`](#destination) - 'Site' or 'Agent', The application destination type.

- [`enabled`](#enabled) - True if the application is enabled.

- [`identity_mapping_ids`](#identity_mapping_ids) - A map of Identity Mappings associated with the application. The key is 'Web' or 'API' and the value is an Identity Mapping ID.

- [`name`](#name) - The application name.

- [`policy`](#policy) - A map of policy items associated with the application. The key is 'Web' or 'API' and the value is a list of PolicyItems.

- [`realm`](#realm) - The OAuth realm associated with the application.

- [`require_https`](#require_https) - True if the application requires HTTPS connections.

- [`site_id`](#site_id) - The ID of the site associated with the application or zero if none.

- [`spa_support_enabled`](#spa_support_enabled) - Enable SPA support.

- [`virtual_host_ids`](#virtual_host_ids) - An array of virtual host IDs associated with the application.

- [`web_session_id`](#web_session_id) - The ID of the web session associated with the application or zero if none.

An ``identity_mapping_ids`` block supports the following arguments:

- [`web`](#identity_mapping_ids_web) - The identity mapping ID.
- [`api`](#identity_mapping_ids_api) - The identity mapping ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The application s's ID.

## Import

PingAccess applications can be imported using the id, e.g.

```bash
$ terraform import pingaccess_application .demo_application  123
```
