module github.com/pulumi/pulumi-grafana

go 1.12

replace (
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
)

require (
	cloud.google.com/go/logging v1.0.0 // indirect
	github.com/aws/aws-sdk-go v1.25.3 // indirect
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/hashicorp/go-hclog v0.9.2 // indirect
	github.com/hashicorp/go-plugin v1.0.1 // indirect
	github.com/hashicorp/go-version v1.2.0 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/hashicorp/terraform v0.12.9
	github.com/hashicorp/terraform-config-inspect v0.0.0-20191115094559-17f92b0546e8 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.6.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20191202134852-87cfb4dc8ae1
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.5 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/terraform-providers/terraform-provider-grafana v1.5.0
	golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586 // indirect
	golang.org/x/net v0.0.0-20191009170851-d66e71096ffb // indirect
	golang.org/x/sys v0.1.0 // indirect
	google.golang.org/grpc v1.24.0 // indirect
)
