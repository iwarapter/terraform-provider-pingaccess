## 0.2.0 (Unreleased)

NOTES:

* This release add support for Terraform 0.12.x
* Continuous Integration was setup with ([Travis](https://travis-ci.org/iwarapter/terraform-provider-pingaccess))
* Code Quality checks setup with ([SonarQube](https://sonarcloud.io/dashboard?id=github.com.iwarapter.terraform-provider-pingaccess))

FEATURES:

* **New DataSource:** `pingaccess_trusted_certificate_group` ([#3](https://github.com/iwarapter/terraform-provider-pingaccess/issues/3))
* **New DataSource:** `pingaccess_keypair` ([#4](https://github.com/iwarapter/terraform-provider-pingaccess/issues/4))
* **New Resource:** `pingaccess_access_token_validator`
* **New Resource:** `pingaccess_keypair` ([#4](https://github.com/iwarapter/terraform-provider-pingaccess/issues/4))
* * **New Resource:** `pingaccess_auth_token_management` ([#11](https://github.com/iwarapter/terraform-provider-pingaccess/issues/11))
 
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