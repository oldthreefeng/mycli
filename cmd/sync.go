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
	"github.com/oldthreefeng/mycli/pkg/v1/gray"
	"github.com/oldthreefeng/mycli/utils"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var (
	debug bool
	syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "need to sync v1 to v2 images list",
	Run: func(cmd *cobra.Command, args []string) {
		if !Force && !debug {
			prompt := "sync image will set all your v1deployment image to v2 deployment(y/n)?"
			Force = utils.Confirm(prompt)
		}
		if Force || debug {
			gray.Sync(debug)
		}
	},
}
)

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.PersistentFlags().BoolVarP(&debug, "dry-run" , "", false, "sync image dry-run mode. not really sync the image")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
