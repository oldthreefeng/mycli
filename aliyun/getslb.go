package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/oldthreefeng/mycli/utils"
	"github.com/wonderivan/logger"
)

func newSlbClient() *slb.Client {
	client, _ := slb.NewClientWithAccessKey("cn-hangzhou", utils.EnvDefault("LT_Ali_Key",""), utils.EnvDefault("LT_Ali_Secret", ""))
	return client
}

func getSlbList() []string {
	var LoadBalancerId []string
	client := newSlbClient()
	request := slb.CreateDescribeLoadBalancersRequest()
	request.Scheme = "https"

	response, err := client.DescribeLoadBalancers(request)
	if err != nil {
		return LoadBalancerId
	}

	for _, v := range response.LoadBalancers.LoadBalancer {
		logger.Info(v.Address, v.LoadBalancerName) 
		LoadBalancerId = append(LoadBalancerId, v.LoadBalancerId)
	}
	return LoadBalancerId
}

func getSlb(loadBalanceId string , client *slb.Client) *slb.DescribeLoadBalancerAttributeResponse {
	request := slb.CreateDescribeLoadBalancerAttributeRequest()
	request.Scheme = "https"

	request.LoadBalancerId = loadBalanceId

	response, err := client.DescribeLoadBalancerAttribute(request)
	if err != nil {
		return nil
	}
	return response

}
