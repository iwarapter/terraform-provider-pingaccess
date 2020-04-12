module github.com/iwarapter/terraform-provider-pingaccess

go 1.12

replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999

require (
	github.com/Microsoft/go-winio v0.4.13 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/containerd/continuity v0.0.0-20190426062206-aaeac12a7ffc // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/google/go-cmp v0.3.1
	github.com/hashicorp/terraform-plugin-sdk v1.9.0
	github.com/iwarapter/pingaccess-sdk-go v0.0.0-20200405225845-0449df2d35ec
	github.com/ory/dockertest v3.3.4+incompatible
	github.com/tidwall/gjson v1.3.2
	github.com/tidwall/sjson v1.0.4
)
