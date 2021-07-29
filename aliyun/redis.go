package aliyun

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	r_kvstore20150101 "github.com/alibabacloud-go/r-kvstore-20150101/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RedisClient struct {
	*r_kvstore20150101.Client
}


func CreateredisClient(accessKeyId string, accessKeySecret string) (*RedisClient, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String("r-kvstore.aliyuncs.com")
	redisClient, err := r_kvstore20150101.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &RedisClient{redisClient}, nil
}

  
func DescribeRedislist () (rlist []string) {
	// 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := CreateredisClient("","")
	if _err != nil {
	  return nil
	}
  
	describeInstancesRequest := &r_kvstore20150101.DescribeInstancesRequest{
	  VpcId: tea.String("prod-vpc-id"),
	}
	runtime := &util.RuntimeOptions{}

	list, _err := client.DescribeInstancesWithOptions(describeInstancesRequest, runtime)
	if _err != nil {
	  return nil
	}
	for _,v := range list.Body.Instances.KVStoreInstance {
		// fmt.Println(*v.InstanceId)
		rlist = append(rlist, *v.InstanceId)

	}
	return
}

func doredisAd() {
	list := DescribeRedislist()
	for _, v := range list {
		fmt.Println(v)
		AppentRedisIps(v)
	}

}

  
func AppentRedisIps (redisid string) (_err error){
	client, _err := CreateredisClient("","")
	if _err != nil {
	  return _err
	}

	modifySecurityIpsRequest := &r_kvstore20150101.ModifySecurityIpsRequest{
		InstanceId: tea.String(redisid),
		SecurityIps: tea.String("ip"),
		ModifyMode: tea.String("Append"),
		SecurityIpGroupName: tea.String("gd_feilian"),
	  }
	  runtime := &util.RuntimeOptions{}
	  tryErr := func()(_e error) {
		defer func() {
		  if r := tea.Recover(recover()); r != nil {
			_e = r
		  }
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.ModifySecurityIpsWithOptions(modifySecurityIpsRequest, runtime)
		if _err != nil {
		  return _err
		}
	
		return nil
	  }()
	
	  if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
		  error = _t
		} else {
		  error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
		  return _err
		}
	  }
	  return _err
}