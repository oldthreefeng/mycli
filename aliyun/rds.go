package aliyun

import (
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	rds20140815 "github.com/alibabacloud-go/rds-20140815/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RdsClient struct {
	*rds20140815.Client
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
 func CreateRdsClient(accessKeyId string, accessKeySecret string) (*RdsClient, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}
	config.Endpoint = tea.String("rds.aliyuncs.com")
	rdsClient, err := rds20140815.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &RdsClient{rdsClient}, nil
}

func DescribeRDSlist () (dlist []string){
	client, _err := CreateRdsClient("","")
	if _err != nil {
	  return nil
	}
  
	describeDBInstancesRequest := &rds20140815.DescribeDBInstancesRequest{
	  VpcId: tea.String("prod-vpc-id"),
	  RegionId: tea.String("cn-hangzhou"),
	}
	runtime := &util.RuntimeOptions{}
	list , _err := client.DescribeDBInstancesWithOptions(describeDBInstancesRequest, runtime)
	if _err != nil {
	  return nil
	}
	
	for _,v := range list.Body.Items.DBInstance {
		fmt.Println(*v.DBInstanceId)
		dlist = append(dlist,*v.DBInstanceId)
	}

	return
}

func doAd() {
	list := DescribeRDSlist()
	for _, v := range list {
		// fmt.Println(v)
		AppentRdsIps(v)
	}

}


func AppentRdsIps (rdsid string) (_err error) {
  // 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
  client, _err := CreateRdsClient("","")
  if _err != nil {
    return _err
  }

  modifySecurityIpsRequest := &rds20140815.ModifySecurityIpsRequest{
    DBInstanceId: tea.String(rdsid),
    DBInstanceIPArrayName: tea.String("gd_feilian"),
    SecurityIps: tea.String("ip"),
    SecurityIPType: tea.String("IPv4"),
    WhitelistNetworkType: tea.String("VPC"),
    ModifyMode: tea.String("Append"),
  }
  runtime := &util.RuntimeOptions{
    // 超时设置，该产品部分接口调用比较慢，请您适当调整超时时间。
    ReadTimeout: tea.Int(50000),
    ConnectTimeout: tea.Int(50000),
  }
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


