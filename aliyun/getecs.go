package aliyun

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/oldthreefeng/mycli/utils"
	"github.com/wonderivan/logger"
)

type AliEcs struct {
	InstanceId string `json:"instance_id"`
	InternalIp string `json:"internal_ip"`
	NickName   string `json:"nick_name"`
	Expire     string `json:"expire"`
	IsProd     bool   `json:"is_prod"`
}

func newClient() *ecs.Client {
	cli, _ := ecs.NewClientWithAccessKey("cn-hangzhou", utils.EnvDefault("LT_Ali_Key", ""), utils.EnvDefault("LT_Ali_Secret", ""))
	return cli
}

func getEcsInstanceList() ([]string, []AliEcs) {
	client := newClient()
	// request := ecs.CreateDescribeInstanceAutoRenewAttributeRequest()
	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"

	// request.RenewalStatus = "AutoRenewal"

	request.PageSize = requests.NewInteger(100)

	response, err := client.DescribeInstances(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	var instanceList []string
	var al []AliEcs
	for _, v := range response.Instances.Instance {
		var a AliEcs
		a.InstanceId = v.InstanceId
		a.NickName = v.Description
		a.Expire = v.ExpiredTime
		if len(v.VpcAttributes.PrivateIpAddress.IpAddress) > 0 {
			a.InternalIp = v.VpcAttributes.PrivateIpAddress.IpAddress[0]
		}
		for _, k := range v.Tags.Tag {
			if strings.Contains(k.TagValue, "prod") && !strings.Contains(k.TagValue, "non_prod") {
				a.IsProd = true
			}
		}
		if strings.Contains(a.NickName, "prod") && !strings.Contains(a.NickName, "non_prod") {
			a.IsProd = true
		}
		al = append(al, a)
		instanceList = append(instanceList, v.InstanceId)
	}
	return instanceList, al
}

func newAliEcs(instanceId string, client *ecs.Client) AliEcs {
	var a AliEcs
	request := ecs.CreateDescribeInstanceAttributeRequest()
	request.Scheme = "https"

	request.InstanceId = instanceId

	response, err := client.DescribeInstanceAttribute(request)
	if err != nil {
		return a
	}
	if len(response.VpcAttributes.PrivateIpAddress.IpAddress) > 0 {
		a.InternalIp = response.VpcAttributes.PrivateIpAddress.IpAddress[0]
	}
	a.InstanceId = instanceId
	a.Expire = response.ExpiredTime
	a.NickName = response.Description
	return a
}

func (a *AliEcs) getEcsInnerIp() string {
	return a.InternalIp
}

func (a *AliEcs) getEcsNickName() string {
	return a.NickName
}

// func ALlEcs() []AliEcs {
// 	list, _ := getEcsInstanceList()
// 	client := newClient()
// 	var al []AliEcs
// 	var wg sync.WaitGroup
// 	for _, v := range list {
// 		wg.Add(1)
// 		go func(id string) {
// 			defer wg.Done()
// 			al = append(al, newAliEcs(id, client))
// 		}(v)
// 	}
// 	wg.Wait()

// 	return al
// }

var homeFile = getHome() + "/.ecs.yaml"

func getHome() string {
	if home, _ := os.UserHomeDir(); home != "" {
		return home
	}
	return ""
}

func Dump(prod bool) {
	dump(prod)
}

func dump(prod bool) {
	_, al := getEcsInstanceList()
	// fmt.Println(al)
	var pl []AliEcs
	if prod {
		for _, v := range al {
			//fmt.Println(v.InternalIp)
			if v.IsProd {
				pl = append(pl, v)
			}
		}
		if len(pl) == 0 {
			logger.Warn("you dont have any prod ecs")
		} else {
			al = pl
		}
	}

	b, err := json.Marshal(al)
	if err != nil {
		return
	}
	fmt.Println(string(b))
	ioutil.WriteFile(homeFile, b, 0644)
}
func load() []AliEcs {
	var al []AliEcs
	b, err := ioutil.ReadFile(homeFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &al)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return al
}

func getProd(loaded bool) []AliEcs {
	var al, pl []AliEcs
	if loaded {
		al = load()
	} else {
		_, al = getEcsInstanceList()
	}
	for _, v := range al {
		//fmt.Println(v.InternalIp)
		if strings.HasPrefix(v.InternalIp, "192.168.110") || strings.HasPrefix(v.InternalIp, "192.168.32") {

			pl = append(pl, v)
		}
	}
	return pl
}
