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
	"github.com/oldthreefeng/mycli/pkg/play"
	"github.com/wonderivan/logger"

	"github.com/spf13/cobra"
)

// dingdingCmd represents the dingding command
var (
	dingCmd = &cobra.Command{
		Use:     "ding",
		Aliases: []string{"d"},
		Short:   "dingding usage",
		Run:     dingCmdFunc,
	}
	K8sVersion string
	DdToken    string
)

func init() {
	rootCmd.AddCommand(dingCmd)
	dingCmd.Flags().StringVar(&K8sVersion, "version", EnvDefault("Version", ""), "k8s version")
	dingCmd.Flags().StringVar(&DdToken, "dd-token", EnvDefault("DD_TOKEN", ""), "dingding token")

}

func dingCmdFunc(cmd *cobra.Command, args []string) {
	var t play.Linking
	if K8sVersion == "" || DdToken == "" {
		return
	}
	t.MessageUrl = "https://sealyun.com"
	t.Text = fmt.Sprintf("kubernetes-arm64自动发布%s版本", K8sVersion)
	t.Title = fmt.Sprintf("kube%s-arm64 发布成功", K8sVersion)
	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", DdToken)

	//fmt.Printf("text: %s\ntitle: %s\nurl: %s\n", t.Text, t.Title, url)

	err := t.Dingding(url)
	if err != nil {
		logger.Error("Dingding error: ", err)
	}
}
