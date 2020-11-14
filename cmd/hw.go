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
	ak       = getDefaultPathname("ak", "")
	sk       = getDefaultPathname("sk", "")
	eip      bool
	count    int32
	serverId string
	eipId     string
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

}

func NewHuaweiCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "huawei",
		Aliases: []string{"hw"},
		Short:   "huawei ecs create",
		Run:     HuaweiCmdFunc,
		PreRun: func(cmd *cobra.Command, args []string) {
			if ak == "" || sk == "" {
				fmt.Println("ak or sk blank is not allow")
				os.Exit(-1)
			}
		},
	}
	cmd.AddCommand(NewHuaweiCreateCmd())
	cmd.AddCommand(NewHuaweiListCmd())
	cmd.AddCommand(NewHuaweiDeleteCmd())
	return cmd
}

func NewHuaweiCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create ecs in sgp",
		Run:   HuaweiCreateCmdFunc,
	}
	cmd.Flags().BoolVar(&eip, "eip", false, "create with eip or not")
	cmd.Flags().Int32VarP(&count, "count", "c", 1, "Specify ecs count")
	return cmd
}

func HuaweiCreateCmdFunc(cmd *cobra.Command, args []string) {
	hc := huawei.GetDefaultHAuth(ak, sk)
	hc.GenerateEipServer(count, eip)
}

func NewHuaweiListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list ecs in sgp",
		Run:   HuaweiListCmdFunc,
	}
	cmd.Flags().StringVar(&serverId, "id", "", "ecs server id")
	cmd.AddCommand(NewHuaweiIpCmd())
	return cmd
}

func HuaweiListCmdFunc(cmd *cobra.Command, args []string) {
	hc := huawei.GetDefaultHAuth(ak, sk)
	if serverId != "" {
		hc.Show(serverId)
	} else {
		hc.ListServer()
	}
}

func NewHuaweiIpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ip",
		Short: "ip ecs in sgp",
		Run:   HuaweiListIpCmdFunc,
	}
	return cmd
}

func HuaweiListIpCmdFunc(cmd *cobra.Command, args []string) {
	hc := huawei.GetDefaultHAuth(ak, sk)
	hc.ListIps()
}

func NewHuaweiDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete ecs in sgp",
		Run:   HuaweiDeleteCmdFunc,
	}
	cmd.Flags().StringVar(&serverId, "id", "", "ecs server id")
	cmd.Flags().StringVar(&eipId, "eipId", "", "ecs server public ipid")
	return cmd
}

func HuaweiDeleteCmdFunc(cmd *cobra.Command, args []string) {
	hc := huawei.GetDefaultHAuth(ak, sk)
	if serverId != "" {
		hc.DeleteServer(serverId)
	}
	if eipId != "" {
		hc.DeleteIp(eipId)
	}

}
