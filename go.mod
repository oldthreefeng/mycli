module github.com/oldthreefeng/mycli

go 1.16

require (
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1022
	github.com/go-redis/redis/v8 v8.11.0
	github.com/huaweicloud/huaweicloud-sdk-go-v3 v0.0.21-beta
	github.com/pkg/errors v0.9.1
	github.com/pkg/sftp v1.12.0
	github.com/spf13/cobra v1.1.3
	github.com/wonderivan/logger v1.0.0
	go.uber.org/multierr v1.6.0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	k8s.io/api v0.21.0
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v0.21.0
)

replace github.com/wonderivan/logger => ./pkg/logger
