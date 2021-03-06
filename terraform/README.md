# terraform-provider-graylog

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/go-graylog/terraform)
[![Build Status](https://travis-ci.org/suzuki-shunsuke/go-graylog.svg?branch=master)](https://travis-ci.org/suzuki-shunsuke/go-graylog)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/go-graylog/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/go-graylog)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/go-graylog)](https://goreportcard.com/report/github.com/suzuki-shunsuke/go-graylog)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-graylog.svg)](https://github.com/suzuki-shunsuke/go-graylog)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/go-graylog.svg)](https://github.com/suzuki-shunsuke/go-graylog/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-graylog/master/LICENSE)

terraform provider for [Graylog](https://www.graylog.org/).

This is sub project of [go-graylog](https://github.com/suzuki-shunsuke/go-graylog).

## Motivation

http://docs.graylog.org/en/2.4/pages/users_and_roles/permission_system.html

The Graylog permission system is extremely flexible but you can't utilize this flexibility from Web UI.
By using this provider, you can utilize this flexibility and manage the infrastructure as code.

## Install

[Download binary](https://github.com/suzuki-shunsuke/go-graylog/releases) and install under `~/.terraform.d/plugins`.

https://www.terraform.io/docs/configuration/providers.html#third-party-plugins

```
$ wget https://github.com/suzuki-shunsuke/go-graylog/releases/download/v0.1.4/terraform-provider-graylog_v0.1.4_darwin_amd64.gz
$ gzip -d terraform-provider-graylog_v0.1.4_darwin_amd64.gz
$ mkdir -p ~/.terraform.d/plugins
$ mv terraform-provider-graylog_v0.1.4_darwin_amd64 ~/.terraform.d/plugins/terraform-provider-graylog_v0.1.4
$ chmod +x ~/.terraform.d/plugins/terraform-provider-graylog_v0.1.4
```

## Example

```
provider "graylog" {
  web_endpoint_uri = "${var.web_endpoint_uri}"
  auth_name = "${var.auth_name}"
  auth_password = "${var.auth_password}"
}

// Role my-role-2
resource "graylog_role" "my-role-2" {
  name = "my-role-2"
  permissions = ["users:edit"]
  description = "Created by terraform"
}
```

## Variables

name | Environment variable | description
--- | --- | ---
web_endpoint_uri | GRAYLOG_WEB_ENDPOINT_URI |
auth_name | GRAYLOG_AUTH_NAME |
auth_password | GRAYLOG_AUTH_PASSWORD |

## Resources

* [role](docs/role.md)
* [user](docs/user.md)
* [input](docs/input.md)
* [index_set](docs/index_set.md)
* [stream](docs/stream.md)
