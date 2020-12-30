/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"fmt"
	"os"

	"github.com/oldthreefeng/mycli/k8s"
	"github.com/oldthreefeng/mycli/utils"
	"github.com/spf13/cobra"
	"github.com/wonderivan/logger"
	v1 "k8s.io/api/core/v1"
)

var Force bool
var example = ` 
	# delete an tx namespace
	mycli cns tx
	# delete an tx namespace by force
	mycli cns tx --force
	# delete some namespaces like tx, staging
	mycli cns tx staging
	# delete some namespaces like tx, staging , qa by force
	mycli cns tx staging qa --force
	`

// cnsCmd represents the cns command
var cnsCmd = &cobra.Command{
	Use:   "cns",
	Short: "clean namespace",
	Example: example,
	Long: `kubernetes namespace delete, when the namespace is terminating and cannot delete by kubectl`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := k8s.NewClient(nil)
	    if err != nil {
			logger.Fatal("kubernetes config is not config")
		}
		ctx := context.Background()
		if len(args) == 0 {
			logger.Warn("clean namespace is empty,skip clean.")
			os.Exit(0)
		}

		for i:= range args {
			namespace := &v1.Namespace{}
			namespace.Name = args[i]
			promptFlag := false
			if !Force {
				prompt := fmt.Sprintf("clean namespace %s in this cluster, continue clean (y/n)?", args[i])
				promptFlag = utils.Confirm(prompt)
			}

			if promptFlag || Force {
				if err := k8s.DeleteNamespace(ctx, client, args[i], namespace) ; err != nil {
					logger.Error("delete namespace error: ", err)
				} else {
					logger.Info("delete namespace success %s", namespace.Name)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cnsCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cnsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
