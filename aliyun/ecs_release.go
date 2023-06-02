package aliyun

import (
  openapi  "github.com/alibabacloud-go/darabonba-openapi/v2/client"
  openapiutil  "github.com/alibabacloud-go/openapi-util/service"
  util  "github.com/alibabacloud-go/tea-utils/v2/service"
  "github.com/alibabacloud-go/tea/tea"
)


/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
 func CreateecsClient(accessKeyId string, accessKeySecret string) (*openapi.Client, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String("ecs-cn-hangzhou.aliyuncs.com")
	Client, err := openapi.NewClient(config)
	if err != nil {
		return nil, err
	}
	return Client, nil
}


/**
 * API 相关
 * @param path params
 * @return OpenApi.Params
 */
func CreateApiInfo () (_result *openapi.Params) {
  params := &openapi.Params{
    // 接口名称
    Action: tea.String("ModifyInstanceChargeType"),
    // 接口版本
    Version: tea.String("2014-05-26"),
    // 接口协议
    Protocol: tea.String("HTTPS"),
    // 接口 HTTP 方法
    Method: tea.String("POST"),
    AuthType: tea.String("AK"),
    Style: tea.String("RPC"),
    // 接口 PATH
    Pathname: tea.String("/"),
    // 接口请求体内容格式
    ReqBodyType: tea.String("json"),
    // 接口响应体内容格式
    BodyType: tea.String("json"),
  }
  _result = params
  return _result
}

func _main () (_err error) {
  // 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
  client, _err := CreateecsClient("xx", "xxxx")
  if _err != nil {
    return _err
  }

  

  params := CreateApiInfo()
  // query params
  queries := map[string]interface{}{}
  // queries["InstanceIds"] = tea.String("[\"i-bp158t7ih7cscp0m1qt6\",\"i-bp158t7ih7cscp0m1qt5\",\"i-bp158t7ih7cscp0m1qt4\",\"i-bp158t7ih7cscp0m1qt3\",\"i-bp166fkqa3s3rn820lcp\",\"i-bp166fkqa3s3rn820lcl\",\"i-bp166fkqa3s3rn820lct\",\"i-bp166fkqa3s3rn820lcq\"]")
  // queries["InstanceIds"] = tea.String("[\"i-bp1dt116r1hhkkwmgjo7\",\"i-bp13rjb1w9j0sbj159gf\",\"i-bp1asq88xtqof33apxg6\"]")

  queries["InstanceIds"] = tea.String("[\"i-bp166fkqa3s3rn820lco\",\"i-bp166fkqa3s3rn820lcr\"]")
  queries["RegionId"] = tea.String("cn-hangzhou")
  queries["InstanceChargeType"] = tea.String("PostPaid")
  // runtime options
  runtime := &util.RuntimeOptions{}
  request := &openapi.OpenApiRequest{
    Query: openapiutil.Query(queries),
  }
  // 复制代码运行请自行打印 API 的返回值
  // 返回值为 Map 类型，可从 Map 中获得三类数据：响应体 body、响应头 headers、HTTP 返回的状态码 statusCode
  _, _err = client.CallApi(params, request, runtime)
  if _err != nil {
    return _err
  }
  return _err
}



