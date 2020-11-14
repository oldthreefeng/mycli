## huawei cloud for sealos test

```bash
$ go build -o mycli main.go
$ ./mycli -h

Usage:
  mycli [command]

Available Commands:
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
  -c, --count int32   Specify ecs count (default 1)
      --eip           create with eip or not
  -h, --help          help for create

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
        --eipId string   ecs server public ipid
    -h, --help           help for delete
        --id string      ecs server id
  
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
  -h, --help        help for list
      --id string   ecs server id

Global Flags:
      --config string   config file (default is $HOME/.mycli.yaml)

Use "mycli huawei list [command] --help" for more information about a command.
```

## example

此命令会默认创建1台2核4g主机. `--eip` 默认带外网的. 

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