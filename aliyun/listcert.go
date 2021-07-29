package aliyun

import (
	"fmt"
	cas20200407  "github.com/alibabacloud-go/cas-20200407/client"
	openapi  "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/oldthreefeng/mycli/utils"
)


func newCasClient() *cas20200407.Client {
	fmt.Println(utils.EnvDefault("GD_Ali_Key", ""),utils.EnvDefault("GD_Ali_Secret",""))
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: tea.String(utils.EnvDefault("GD_Ali_Key", "")),
		// 您的AccessKey Secret
		AccessKeySecret: tea.String(utils.EnvDefault("GD_Ali_Secret", "")),
	  }
	  // 访问的域名
	config.Endpoint = tea.String("cas.aliyuncs.com")
	_result := &cas20200407.Client{}
	_result, _ = cas20200407.NewClient(config)
	return _result
}

func ListCert() {
	client  := newCasClient()
		/* use STS Token 
	client, err := cas.NewClientWithStsToken("cn-hangzhou", "<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/

	listUserCertificateOrderRequest := &cas20200407.ListUserCertificateOrderRequest{}

	res, err := client.ListUserCertificateOrder(listUserCertificateOrderRequest)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", res)
}