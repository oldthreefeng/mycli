## huawei cloud for sealos test

```bash
$ go build -o mycli main.go
$ ./mycli -h

Usage:
  mycli [command]

Available Commands:
  completion  Output shell completion code for the specified shell (bash or zsh)
  help        Help about any command
  huawei      huawei ecs create

Flags:
      --config string   config file (default is $HOME/.mycli.yaml)
  -h, --help            help for mycli
  -t, --toggle          Help message for toggle

Use "mycli [command] --help" for more information about a command.

```

## 创建

```bash
$ ./mycli huawei create  --help
create ecs in sgp

Usage:
  mycli huawei create [flags]

Flags:
      --FlavorRef string   huawei falvor id , default is centos 7.6 (default "kc1.large.2")
      --ImageRef string    huawei image id , default is 2C 4G (default "456416e6-1270-46a4-975e-3558ac03d4cd")
      --SubnetId string    huawei subnet id (default "b5ea4e5d-de19-442b-ac32-3998100e4854")
      --Vpcid string       huawei Vpcid  (default "a55545d8-a4cb-436d-a8ec-45c66aff725c")
      --Zone string        huawei AvailabilityZone , default is centos xin jia po (default "ap-southeast-3a")
      --adminPass string   huawei root pass (default "Louishong4168@123")
  -c, --count int32        Specify huawei ecs count (default 1)
      --eip                create huawei ecs with eip or not
  -h, --help               help for create
      --keyName string     ssh key name
      --projectId string   huawei project id (default "06b275f705800f262f3bc014ffcdbde1")

Global Flags:
      --config string   config file (default is $HOME/.mycli.yaml)

```

## 删除

```bash
$ ./mycli huawei delete  --help
delete ecs in sgp

Usage:
  mycli huawei delete [flags]

Flags:
      --FlavorRef string   huawei AvailabilityZone , default is centos xin jia po (default "ap-southeast-3a")
      --eipId string       huawei ecs server public eipid
  -h, --help               help for delete
      --id string          huawei ecs server id
      --projectId string   huawei project id (default "06b275f705800f262f3bc014ffcdbde1")

Global Flags:
      --config string   config file (default is $HOME/.mycli.yaml)

```

## 获取信息

```bash
$ ./mycli huawei list  --help  
list ecs in sgp

Usage:
  mycli huawei list [flags]
  mycli huawei list [command]

Available Commands:
  ip          ip ecs in sgp

Flags:
      --FlavorRef string   huawei AvailabilityZone , default is centos xin jia po (default "ap-southeast-3a")
  -h, --help               help for list
      --id string          huawei ecs server id
      --projectId string   huawei project id (default "06b275f705800f262f3bc014ffcdbde1")

Global Flags:
      --config string   config file (default is $HOME/.mycli.yaml)

Use "mycli huawei list [command] --help" for more information about a command.

```

## example

此命令会默认创建1台2核4g主机. `--eip` 默认带外网的. 默认按流量计费, 50M宽带

```bash
$ mycli hw create -c 1 --eip | tee  -a  InstanceId.json
  
$  cat InstanceId.json

$ ID=$(jq -r '.serverIds[0]' < InstanceId.json)

## 获取instanceId list
$ mycli hw list --id $ID > info.json

## 获取ip list
$ mycli hw list ip > ip.json

$ FIPID=$(jq -r ".[0].id" < ip.json)
```

删除主机及eip. 单独删除主机不会连带删除eip.

```bash
$ ./mycli huawei delete  --id $ID  --eipId $FIPID
```