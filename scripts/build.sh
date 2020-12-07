#!/bin/bash
# Author:louisehong4168
# Blog:https://fenghong.tech
# Time:2020-12-07 11:52:58
# Name:build.sh
# Version:V1.0
# Description:This is a production script.
COMMIT_SHA1=$(git rev-parse --short HEAD || echo "0.0.0")
BUILD_TIME=$(date "+%F %T")
goldflags="-s -w -X 'github.com/oldthreefeng/mycli/cmd.Version=$1' -X 'github.com/oldthreefeng/sealos/cmd.Githash==${COMMIT_SHA1}' -X 'github.com/oldthreefeng/sealos/cmd.Buildstamp=${BUILD_TIME}' -X github.com/oldthreefeng/mycli/cmd.Author=oldthreefeng"
go build -o mycli -ldflags "$goldflags" main.go && command -v upx &> /dev/null && upx mycli
