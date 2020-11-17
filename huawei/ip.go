package huawei

import (
	"encoding/json"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	eip "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2/model"
)

func (h *HClient) ListIps() error {

	auth := basic.NewCredentialsBuilder().
		WithAk(h.Ak).
		WithSk(h.Sk).
		WithProjectId("").
		Build()

	client := eip.NewEipClient(
		eip.EipClientBuilder().
			WithEndpoint(h.VpcEndpoint).
			WithCredential(auth).
			Build())

	request := &model.NeutronListFloatingIpsRequest{}

	response, err := client.NeutronListFloatingIps(request)

	if err == nil {
		date, _ := json.MarshalIndent(response.Floatingips, "", "    ")
		fmt.Println(string(date))
	}
	return err
}

func (h *HClient) DeleteIp(ipId string) error {
	auth := basic.NewCredentialsBuilder().
		WithAk(h.Ak).
		WithSk(h.Sk).
		WithProjectId("").
		Build()

	client := eip.NewEipClient(
		eip.EipClientBuilder().
			WithEndpoint(h.VpcEndpoint).
			WithCredential(auth).
			Build())

	request := &model.NeutronDeleteFloatingIpRequest{}
	request.FloatingipId = ipId

	response, err := client.NeutronDeleteFloatingIp(request)

	if err == nil {
		fmt.Printf("%+v\n", response)
	}
	return err
}
