package aliyun

import (
	"fmt"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/oldthreefeng/mycli/utils"
	"github.com/wonderivan/logger"
)


func newOssClient(location string) *oss.Client {
	region := fmt.Sprintf("http://%s.aliyuncs.com", location) 
	client, _ := oss.New(region, utils.EnvDefault("GD_Ali_Key",""), utils.EnvDefault("GD_Ali_Secret", ""))
	return client
}

func newOssClientByKey(key, secret string) *oss.Client {
	client, _ := oss.New("http://oss-cn-hangzhou.aliyuncs.com", key, secret)
	return client
}

func listAllBucket () {
	var content string 

	content = header

	keys := make(map[string]string, 3)


	

	for k, v := range keys {
		c := listBucket(k,v)
		content += c
	}
	
	fmt.Println(content)
}


func listBucket(key,secret string) string {
	client := newOssClientByKey(key,secret)

	resllut, err  := client.ListBuckets()
	if err != nil {
		logger.Error(err)
	}
	var content,line string
	// content = header
	if key == "" {
		return ""
	}
	for _,v := range resllut.Buckets {
		var product string
		if strings.Contains(v.Name, "hlg") {
			product = "hlg"
		} else if strings.Contains(v.Name, "cri") || 
		strings.Contains(v.Name, "sre") ||  strings.Contains(v.Name, "druid") || 
		strings.Contains(v.Name, "flink") || strings.Contains(v.Name, "warehouse")  || 
		strings.Contains(v.Name, "thanos") || strings.Contains(v.Name, "kafka") ||
		 strings.Contains(v.Name, "elastic") || strings.Contains(v.Name, "loki") ||
		 strings.Contains(v.Name, "prometheus") {
			product = "sre"
		}
		line = fmt.Sprintf("| %s | %s | %s |\n", v.Name, product , v.Location)
		content += line
	}
	fmt.Println(content)
	fmt.Println("done")
	return content

}

const header = `| ossbucket Name |   业务线| 地域      |
| -------------------- | ------------------------------- | ----------------------- |
`
