## 0.9.0 (March 9th, 2022)

NOTES:

* This release has been built for PingAccess 6.x and uses the SDK for that version, whilst the API remains mostly the same backwards compatibility with PingAccess 5.x is not being maintained.

FEATUERS:

* Support for virtual resources of applications. (#98)

BUG FIXES:

* Fix bug with state upgrade for `pingaccess_access_token_validator`. (#92)
* Fix ordering of policy rules on `pingaccess_ruleset`. (#111)
* Fix panic when using private key jwt on websession. (#116)

## 0.8.0 (August 26th, 2021)

NOTES:

* This release has been built for PingAccess 6.x and uses the SDK for that version, whilst the API remains mostly the same backwards compatibility with PingAccess 5.x is not being maintained.

FEATURES:

* Add support for dynamic configuration blocks to allow native HCL when defining plugin configuration, supported resources:
    - `pingaccess_access_token_validator`
    - `pingaccess_site_authenticator`
* Add support for proxied pingfederate runtime (PA 6.2+). (#72)
* Extend support for `websession` and `pingfederate_oauth` resource `client_credentials` block to support certificate/private key jwt auth. (#75)

BUG FIXES:

* Fix issue with site authenticator not handling interpolated value. (#74)
* Fix issue with applications not correctly tracking policies. (#48)


## 0.7.0 (November 23, 2020)

NOTES:

* This release has been built for PingAccess 6.x and uses the SDK for that version, whilst the API remains mostly the same backwards compatibility with PingAccess 5.x is not being maintained.

FEATURES:

* **New Resource:** `pingaccess_availability_profile`
* **New Resource:** `pingaccess_load_balancing_strategy`

## 0.6.0 (November 17, 2020)

NOTES:

* This release is built for PingAccess 6.x and uses the SDK for that version, whilst the API remains mostly the same backwards compatibility with PingAccess 5.x is not being maintained.

BUG FIXES:

* Add configuration validation for the provider block for any initial connection issues.
* Fix issues with importing resources, additional test cases and documentation.
* `resource/application_resource`: Fix issue with `path_patterns` attribute.

## 0.5.0 (November 6, 2020)

NOTES:

* This release is built for PingAccess 6.x and uses the SDK for that version, whilst the API remains mostly the same backwards compatibility with PingAccess 5.x is not being maintained.
* This is the first version available on the Terraform Registry https://registry.terraform.io/providers/iwarapter/pingaccess/latest

FEATURES:

* **New Data Source:** `pingaccess_keypair_csr`
* **New Resource:** `pingaccess_keypair_csr`

## 0.4.0 (April 20, 2020)

NOTES:

* This release is built for PingAccess 6.x and uses the SDK for that version, whilst the API remains mostly the same backwards compatibility with PingAccess 5.x is not being maintained.
* This release changes the way several resources handle the json configuration to mask sensitive values, the following resources are affected:
    - `pingaccess_access_token_validator`
    - `pingaccess_hsm_provider`
    - `pingaccess_rule`
    - `pingaccess_identity_mapping`
    - `pingaccess_site_authenticator`

FEATURES:

* **New Data Source:** `pingaccess_pingfederate_runtime_metadata`
* **New Data Source:** `pingaccess_acme_default`
* **New Resource:** `pingaccess_acme_server`
* **New Resource:** `pingaccess_hsm_provider`
* **New Resource:** `pingaccess_pingfederate_admin`
* resource/websession: Added support for `same_site` attribute.
* resource/keypair: Added `hsm_provider_id` attribute.

BUG FIXES:

* resource/pingaccess_rule: Fixed issue with interpreted configuration producing inconsistent final plan.

## 0.3.0 (November 13, 2019)

FEATURES:

* **New Resource:** `pingaccess_authn_req_lists` ([#16](https://github.com/iwarapter/terraform-provider-pingaccess/issues/16))
* **New Resource:** `pingaccess_https_listener` ([#15](https://github.com/iwarapter/terraform-provider-pingaccess/issues/15))
* **New Resource:** `pingaccess_engine_listener` ([#14](https://github.com/iwarapter/terraform-provider-pingaccess/issues/14))

BUG FIXES:

* resource/pingaccess_application: Fixed issued with `policy`ignoring rule order ([#17](https://github.com/iwarapter/terraform-provider-pingaccess/issues/17))
* resource/pingaccess_application_resource: Fixed issued with `policy`ignoring rule order ([#17](https://github.com/iwarapter/terraform-provider-pingaccess/issues/17))

## 0.2.0 (October 12, 2019)

NOTES:

* This release add support for Terraform 0.12.x
* Continuous Integration was setup with ([Travis](https://travis-ci.org/iwarapter/terraform-provider-pingaccess))
* Code Quality checks setup with ([SonarQube](https://sonarcloud.io/dashboard?id=github.com.iwarapter.terraform-provider-pingaccess))

FEATURES:

* **New DataSource:** `pingaccess_trusted_certificate_group` ([#3](https://github.com/iwarapter/terraform-provider-pingaccess/issues/3))
* **New DataSource:** `pingaccess_keypair` ([#4](https://github.com/iwarapter/terraform-provider-pingaccess/issues/4))
* **New Resource:** `pingaccess_access_token_validator`
* **New Resource:** `pingaccess_keypair` ([#4](https://github.com/iwarapter/terraform-provider-pingaccess/issues/4))
* **New Resource:** `pingaccess_auth_token_management` ([#11](https://github.com/iwarapter/terraform-provider-pingaccess/issues/11))

BUG FIXES:

* resource/pingaccess_application: Fixing issue with an empty `policy` block ([#7](https://github.com/iwarapter/terraform-provider-pingaccess/issues/7))
* resource/pingaccess_application_resource: Fixing issue with an empty `policy` block ([#7](https://github.com/iwarapter/terraform-provider-pingaccess/issues/7))
* resource/pingaccess_application_resource: path_prefixes is deprecated but required  ([#8](https://github.com/iwarapter/terraform-provider-pingaccess/issues/8))
* resource/pingaccess_certificate: Certificate `alias` change forces new resource ([#6](https://github.com/iwarapter/terraform-provider-pingaccess/issues/6))


## 0.0.1-BETA (August 4, 2019)

NOTES:

* This was the initial release of a stable, but still in development version of the provider.

FEATURES:

* **New DataSource:** `pingaccess_certificate`

* **New Resource:** `pingaccess_certificate`
* **New Resource:** `pingaccess_identity_mapping`
* **New Resource:** `pingaccess_rule`
* **New Resource:** `pingaccess_ruleset`
* **New Resource:** `pingaccess_virtualhost`
* **New Resource:** `pingaccess_site`
* **New Resource:** `pingaccess_application`
* **New Resource:** `pingaccess_application_resource`
* **New Resource:** `pingaccess_websession`
* **New Resource:** `pingaccess_site_authenticator`
* **New Resource:** `pingaccess_third_party_service`
* **New Resource:** `pingaccess_trusted_certificate_group`
* **New Resource:** `pingaccess_pingfederate_runtime`
* **New Resource:** `pingaccess_pingfederate_oauth`
* **New Resource:** `pingaccess_oauth_server`
* **New Resource:** `pingaccess_http_config_request_host_source`
