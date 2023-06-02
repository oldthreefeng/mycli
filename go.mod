module github.com/oldthreefeng/mycli

go 1.16

require (
	github.com/alibabacloud-go/cas-20200407 v1.0.5
	github.com/alibabacloud-go/darabonba-openapi v0.1.18
	github.com/alibabacloud-go/darabonba-openapi/v2 v2.0.2
	github.com/alibabacloud-go/openapi-util v0.1.0
	github.com/alibabacloud-go/r-kvstore-20150101/v3 v3.0.0
	github.com/alibabacloud-go/rds-20140815/v3 v3.0.3
	github.com/alibabacloud-go/tea v1.1.19
	github.com/alibabacloud-go/tea-utils/v2 v2.0.1
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1022
	github.com/aliyun/aliyun-oss-go-sdk v2.2.1+incompatible
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/go-redis/redis/v8 v8.11.0
	github.com/huaweicloud/huaweicloud-sdk-go-v3 v0.0.21-beta
	github.com/pkg/errors v0.9.1
	github.com/pkg/sftp v1.12.0
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/wonderivan/logger v1.0.0
	go.uber.org/multierr v1.6.0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	k8s.io/api v0.21.0
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v0.21.0
)

replace github.com/wonderivan/logger => ./pkg/logger
