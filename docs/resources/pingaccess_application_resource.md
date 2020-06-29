#Resource: pingaccess_application_resource

Provides a application resource.

## Example Usage
```terraform
{!../func-tests//application_resource.tf!}
```

## Argument Attributes

The following arguments are supported:


- [`anonymous`](#anonymous) - True if the resource is anonymous.

- [`application_id`](#application_id) - The id of the associated application. This field is read-only.

- [`audit_level`](#audit_level) - ['ON' or 'OFF']: Indicates if audit logging is enabled for the resource.

- [`default_auth_type_override`](#default_auth_type_override) - ['Web' or 'API']: For Web + API applications (dynamic) defaultAuthType selects the processing mode when a request: does not have a token (web session, OAuth bearer) or has both tokens. default_auth_type_override overrides the defaultAuthType at the application level for this resource. A value of null indicates the resource should not override the defaultAuthType.

- [`enabled`](#enabled) - True if the resource is enabled.

- [`methods`](#methods) - An array of HTTP methods configured for the resource

- [`name`](#name) - The name of the resource

- [`path_patterns`](#path_patterns) - A list of one or more request path-matching patterns

- [`path_prefixes`](#path_prefixes) - An array of path prefixes for the resource (DEPRECATED - to be removed in a future release; please use 'pathPatterns' instead)

- [`policy`](#policy) - A map of policy items associated with the resource. The key is 'Web' or 'API' and the value is a list of PolicyItems.

- [`root_resource`](#root_resource) - True if the resource is the root resource for the application

- [`unprotected`](#unprotected) - True if the resource is unprotected.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The application s's ID.

## Import

PingAccess applications can be imported using the id, e.g.

```bash
$ terraform import pingaccess_application.demo_application 123
```
