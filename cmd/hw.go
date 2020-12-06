/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/oldthreefeng/mycli/huawei"
	"github.com/spf13/cobra"
	"os"
)

var (
	ak               = getDefaultPathname("ak", "")
	sk               = getDefaultPathname("sk", "")
	eip              bool
	count            int32
	eipBandwidth     int32
	rootVolume       int32
	serverId         string
	eipId            string
	adminPass        string
	SubnetId         string
	Vpcid            string
	ImageRef         string
	FlavorRef        string
	AvailabilityZone string
	keyName          string
	projectId        string
)

func init() {

	rootCmd.AddCommand(NewHuaweiCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// huaweiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// huaweiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func HuaweiCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("huawei command caled")
}

func PreRunHuaWei(cmd *cobra.Command, args []string) {
	if ak == "" || sk == "" {
		fmt.Println("ak or sk blank is not allow")
		os.Exit(-1)
	}
}

func NewHuaweiCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "huawei",
		Aliases: []string{"hw"},
		Short:   "huawei ecs create",
		Run:     HuaweiCmdFunc,
		PreRun:  PreRunHuaWei,
	}
	cmd.AddCommand(NewHuaweiCreateCmd())
	cmd.AddCommand(NewHuaweiListCmd())
	cmd.AddCommand(NewHuaweiDeleteCmd())
	return cmd
}

func NewHuaweiCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "create",
		Short:  "create ecs in sgp",
		Run:    HuaweiCreateCmdFunc,
		PreRun: PreRunHuaWei,
	}
	cmd.Flags().BoolVar(&eip, "eip", false, "create huawei ecs with eip or not")
	cmd.Flags().Int32VarP(&count, "count", "c", 1, "Specify huawei ecs count")
	cmd.Flags().Int32VarP(&eipBandwidth, "eipBandwidth", "", 100, "Specify huawei eip Bandwidth")
	cmd.Flags().Int32VarP(&rootVolume, "rootVolume", "", 40, "Specify huawei ecs root volume 40G")
	cmd.Flags().StringVar(&adminPass, "adminPass", "Louishong4168@123", "huawei root pass")
	cmd.Flags().StringVar(&SubnetId, "SubnetId", "b5ea4e5d-de19-442b-ac32-3998100e4854", "huawei subnet id")
	cmd.Flags().StringVar(&Vpcid, "Vpcid", "a55545d8-a4cb-436d-a8ec-45c66aff725c", "huawei Vpcid ")
	cmd.Flags().StringVar(&ImageRef, "ImageRef", "456416e6-1270-46a4-975e-3558ac03d4cd", "huawei image id , default is centos 7.6")
	cmd.Flags().StringVar(&FlavorRef, "FlavorRef", "kc1.large.2", "huawei falvor id , default is  2C 4G")
	// huawei cloud only allow one method to auth
	cmd.Flags().StringVar(&keyName, "keyName", "", "ssh key name, when use key, the admin passwd have no effect")

	cmd.Flags().StringVar(&projectId, "projectId", "06b275f705800f262f3bc014ffcdbde1", "huawei project id")
	cmd.Flags().StringVar(&AvailabilityZone, "Zone", "ap-southeast-3a", "huawei AvailabilityZone , default is xin jia po")
	return cmd
}

func HuaweiCreateCmdFunc(cmd *cobra.Command, args []string) {
	hc := huawei.GetDefaultHAuth(ak, sk, projectId, AvailabilityZone)
	hc.GenerateEipServer(count, eipBandwidth, rootVolume, eip, FlavorRef, ImageRef, Vpcid, SubnetId, adminPass, keyName)
}

func NewHuaweiListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "list",
		Short:  "list ecs in sgp",
		Run:    HuaweiListCmdFunc,
		PreRun: PreRunHuaWei,
	}
	cmd.Flags().StringVar(&serverId, "id", "", "huawei ecs server id")
	cmd.Flags().StringVar(&projectId, "projectId", "06b275f705800f262f3bc014ffcdbde1", "huawei project id")
	cmd.Flags().StringVar(&AvailabilityZone, "FlavorRef", "ap-southeast-3a", "huawei AvailabilityZone , default is centos xin jia po")
	cmd.AddCommand(NewHuaweiIpCmd())
	return cmd
}

func HuaweiListCmdFunc(cmd *cobra.Command, args []string) {
	hc := huawei.GetDefaultHAuth(ak, sk, projectId, AvailabilityZone)
	if serverId != "" {
		hc.Show(serverId)
	} else {
		hc.ListServer()
	}
}

func NewHuaweiIpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "ip",
		Short:  "ip ecs in sgp",
		Run:    HuaweiListIpCmdFunc,
		PreRun: PreRunHuaWei,
	}
	cmd.Flags().StringVar(&projectId, "projectId", "06b275f705800f262f3bc014ffcdbde1", "huawei project id")
	cmd.Flags().StringVar(&AvailabilityZone, "FlavorRef", "ap-southeast-3a", "huawei AvailabilityZone , default is centos xin jia po")
	return cmd
}

func HuaweiListIpCmdFunc(cmd *cobra.Command, args []string) {
	hc := huawei.GetDefaultHAuth(ak, sk, projectId, AvailabilityZone)
	hc.ListIps()
}

func NewHuaweiDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "delete",
		Short:  "delete ecs in sgp",
		Run:    HuaweiDeleteCmdFunc,
		PreRun: PreRunHuaWei,
	}
	cmd.Flags().StringVar(&serverId, "id", "", "huawei ecs server id")
	cmd.Flags().BoolVar(&eip, "eip", false, "delete eip or not")
	cmd.Flags().StringVar(&projectId, "projectId", "06b275f705800f262f3bc014ffcdbde1", "huawei project id")
	cmd.Flags().StringVar(&AvailabilityZone, "FlavorRef", "ap-southeast-3a", "huawei AvailabilityZone , default is centos xin jia po")
	return cmd
}

func HuaweiDeleteCmdFunc(cmd *cobra.Command, args []string) {
	hc := huawei.GetDefaultHAuth(ak, sk, projectId, AvailabilityZone)
	if serverId != "" {
		hc.DeleteServer(serverId, eip)
	}
}
