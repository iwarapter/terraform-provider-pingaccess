PingAccess Terraform Provider
==================

- Website: https://iwarapter.github.io/terraform-provider-pingaccess/
- [![Gitter](https://badges.gitter.im/iwarapter/terraform-provider-pingaccess.svg)](https://gitter.im/iwarapter/terraform-provider-pingaccess?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
  [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=github.com.iwarapter.terraform-provider-pingaccess&metric=coverage)](https://sonarcloud.io/dashboard?id=github.com.iwarapter.terraform-provider-pingaccess) 
  [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=github.com.iwarapter.terraform-provider-pingaccess&metric=alert_status)](https://sonarcloud.io/dashboard?id=github.com.iwarapter.terraform-provider-pingaccess)
  [![Build Status](https://travis-ci.org/iwarapter/terraform-provider-pingaccess.svg?branch=master)](https://travis-ci.org/iwarapter/terraform-provider-pingaccess)
  ![GitHub release (latest by date)](https://img.shields.io/github/v/release/iwarapter/terraform-provider-pingaccess)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.12+
- [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)
- [Docker](https://www.docker.com/products/docker-desktop) latest

Developing the Provider
---------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (please check the [requirements](https://github.com/iwarapter/terraform-provider-pingaccess#requirements) before proceeding).

*Note:* This project uses [Go Modules](https://blog.golang.org/using-go-modules) making it safe to work with it outside of your existing [GOPATH](http://golang.org/doc/code.html#GOPATH). The instructions that follow assume a directory in your home directory outside of the standard GOPATH (i.e `$HOME/development/terraform-providers/`).

Clone repository to: `$HOME/development/terraform-providers/`

```sh
$ git clone git@github.com:iwarapter/terraform-provider-pingaccess.git
...
```

To compile the provider, run `make build`. This will build the provider and put the provider binary in the local directory.

```sh
$ make build
...
$ terraform-provider-pingaccess
...
```

Using the Provider
----------------------

To use a released provider in your Terraform environment, download the latest version from https://github.com/iwarapter/terraform-provider-pingaccess/releases and follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

Testing the Provider
---------------------------

In order to test the provider, you can run `make test`.

```sh
$ make test
```

This will run the acceptance tests by initializing a local docker container to execute the functional tests against.
