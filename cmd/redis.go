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

	"github.com/oldthreefeng/mycli/pkg/redis"
	"github.com/oldthreefeng/mycli/utils"
	"github.com/spf13/cobra"
	"github.com/wonderivan/logger"
)

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("redis called")
	},
}

var (
	redisHost, redisPasswd, sni, prefix string
	db                                  int
)

func init() {
	redisCmd.AddCommand(newDelCmd())
	redisCmd.PersistentFlags().StringVarP(&redisHost, "host", "", "", "redis host")
	redisCmd.PersistentFlags().StringVarP(&redisPasswd, "password", "a", "", "redis password")
	redisCmd.PersistentFlags().StringVarP(&sni, "sni", "", "", "redis sni if support")
	redisCmd.PersistentFlags().StringVarP(&prefix, "prefix", "", "", "redis key prefix")

	redisCmd.PersistentFlags().IntVarP(&db, "db", "", 0, "redis db.")

	rootCmd.AddCommand(redisCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// redisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// redisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func newDelCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del",
		Short: "del an key by go-redis",
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()
			rc := redis.NewRedisDb(redisHost, redisPasswd, sni, db)

			if prefix == "" {
				return
			}

			if !Force {
				prompt := fmt.Sprintf("del will delete you redis keys prefix %s (y/n)?", prefix)
				result := utils.Confirm(prompt)
				if !result {
					logger.Info("del %s  is skip, Exit", prefix)
					os.Exit(-1)
				}
			}
			rc.DeleteByPrefix(ctx, prefix)

		},
	}
	cmd.Flags().BoolVarP(&Force, "force", "f", false, "force")
	return cmd
}
