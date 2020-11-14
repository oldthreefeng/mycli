package huawei

import (
	"encoding/json"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
)

const (
	endpoint     = "https://ecs.ap-southeast-3.myhuaweicloud.com"
	VpcEndpoint = "https://vpc.ap-southeast-3.myhuaweicloud.com"
	projectID    = "06b275f705800f262f3bc014ffcdbde1"
	defaultCount = 1
)

type HClient struct {
	Ak     string
	Sk     string
	client *ecs.EcsClient
}

func GetDefaultHAuth(ak, sk string) *HClient {
	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		WithProjectId(projectID).
		Build()
	return &HClient{
		Ak: ak,
		Sk: sk,
		client: ecs.NewEcsClient(
			ecs.EcsClientBuilder().
				WithEndpoint(endpoint).
				WithCredential(auth).
				Build()),
	}
}

func (h *HClient) Show(serverid string) {

	client := h.client

	request := &model.ShowServerRequest{}
	request.ServerId = serverid

	response, err := client.ShowServer(request)
	if err == nil {
		date, _ := json.MarshalIndent(response.Server, "", "    ")
		fmt.Println(string(date))
	} else {
		fmt.Println(err)
	}
}

func (h *HClient) GenerateEipServer(count int32, eip bool) []string {

	client := h.client
	request := &model.CreatePostPaidServersRequest{}
	var listPostPaidServerNicNicsPostPaidServer = []model.PostPaidServerNic{
		{
			SubnetId: "b5ea4e5d-de19-442b-ac32-3998100e4854",
		},
	}
	var listPostPaidServerTagServerTagsPostPaidServer = []model.PostPaidServerTag{
		{
			Key:   "test",
			Value: "sealos",
		},
	}
	sizePostPaidServerEipBandwidth := int32(5)
	chargemodePostPaidServerEipBandwidth := "traffic"
	bandwidthPostPaidServerEip := &model.PostPaidServerEipBandwidth{
		Size:       &sizePostPaidServerEipBandwidth,
		Sharetype:  model.GetPostPaidServerEipBandwidthSharetypeEnum().PER,
		Chargemode: &chargemodePostPaidServerEipBandwidth,
	}
	publicipPostPaidServer := &model.PostPaidServerPublicip{}
	if eip {
		eipPostPaidServerPublicip := &model.PostPaidServerEip{
			Iptype:    "5_bgp",
			Bandwidth: bandwidthPostPaidServerEip,
		}
		publicipPostPaidServer = &model.PostPaidServerPublicip{
			Eip: eipPostPaidServerPublicip,
		}
	}
	countPostPaidServer := count
	isAutoRenamePostPaidServer := true
	keyNamePostPaidServer := "release"
	adminPassPostPaidServer := "Louishong4168#123"
	sizePostPaidServerRootVolume := int32(40)
	rootVolumePostPaidServer := &model.PostPaidServerRootVolume{
		Volumetype: model.GetPostPaidServerRootVolumeVolumetypeEnum().SSD,
		Size:       &sizePostPaidServerRootVolume,
	}
	serverCreatePostPaidServersRequestBody := &model.PostPaidServer{
		AvailabilityZone: "ap-southeast-3a",
		FlavorRef:        "kc1.large.2",
		ImageRef:         "456416e6-1270-46a4-975e-3558ac03d4cd",
		Name:             "sealos",
		Nics:             listPostPaidServerNicNicsPostPaidServer,
		Publicip:         publicipPostPaidServer,
		RootVolume:       rootVolumePostPaidServer,
		ServerTags:       &listPostPaidServerTagServerTagsPostPaidServer,
		Vpcid:            "a55545d8-a4cb-436d-a8ec-45c66aff725c",
		KeyName:          &keyNamePostPaidServer,
		AdminPass:        &adminPassPostPaidServer,
		IsAutoRename:     &isAutoRenamePostPaidServer,
		Count:            &countPostPaidServer,
	}
	request.Body = &model.CreatePostPaidServersRequestBody{
		Server: serverCreatePostPaidServersRequestBody,
	}

	response, err := client.CreatePostPaidServers(request)

	if err == nil {
		date, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(date))
		return *response.ServerIds
	} else {
		fmt.Println(err)
		return nil
	}
}

func (h *HClient) DeleteServer(serverId string) {

	client := h.client

	request := &model.DeleteServersRequest{}
	var listServerIdServersDeleteServersRequestBody = []model.ServerId{
		{
			Id: serverId,
		},
	}
	request.Body = &model.DeleteServersRequestBody{
		Servers: listServerIdServersDeleteServersRequestBody,
	}

	response, err := client.DeleteServers(request)

	if err == nil {
		date, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(date))
	} else {
		fmt.Println(err)
	}
}

func (h *HClient) ListServer() {
	client := h.client
	request := &model.ListServersDetailsRequest{}
	response, err := client.ListServersDetails(request)

	if err == nil {
		//fmt.Printf("%+v\n", response.Servers)
		date, _ := json.MarshalIndent(response.Servers, "", "    ")
		fmt.Println(string(date))
	} else {
		fmt.Println(err)
	}
}
