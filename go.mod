module github.com/oldthreefeng/mycli

go 1.15

require (
	github.com/huaweicloud/huaweicloud-sdk-go-v3 v0.0.21-beta
	github.com/pkg/sftp v1.12.0
	github.com/spf13/cobra v1.1.1
	github.com/wonderivan/logger v1.0.0
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
)

replace github.com/wonderivan/logger => ./pkg/logger
