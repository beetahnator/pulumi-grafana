module github.com/pulumi/pulumi-grafana

go 1.12

replace (
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
)

require (
	github.com/frankban/quicktest v1.4.1 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/hashicorp/terraform v0.12.9
	github.com/hashicorp/terraform-plugin-sdk v1.4.0
	github.com/hashicorp/vault/api v1.0.5-0.20190909201928-35325e2c3262 // indirect
	github.com/kr/pty v1.1.3 // indirect
	github.com/pierrec/lz4 v2.2.6+incompatible // indirect
	github.com/pulumi/pulumi v1.6.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20191202134852-87cfb4dc8ae1
	github.com/pulumi/pulumi-terraform-bridge v1.5.0
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/terraform-providers/terraform-provider-grafana v1.5.0
	golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // indirect
)
