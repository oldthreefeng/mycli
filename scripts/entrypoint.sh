#!/usr/bin/env bash

if [ -n "$ak" ] && [ -n "$sk" ]; then
	echo "huawei ak and sk is check ok"
else
  echo "huawei ak and sk is needed"
  exit -1
fi

# shellcheck disable=SC2068
mycli $@