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
	"path/filepath"
	"strings"

	"github.com/oldthreefeng/mycli/pkg/apisix"
	"github.com/oldthreefeng/mycli/utils"
	"github.com/spf13/cobra"
	"github.com/wonderivan/logger"
)

var (
	adminKey  string
	baseURL   string
	sslId     string
	updateAll bool
)

func init() {
	rootCmd.AddCommand(NewApisixCmd())
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apisixCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apisixCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func NewApisixCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "apisix",
		Aliases: []string{"a"},
		Short:   "apisix interface ",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("apisix called")
		},
	}
	cmd.AddCommand(listCmd())
	cmd.AddCommand(getCmd())
	cmd.AddCommand(updateCmd())
	return cmd
}

func listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "list ssl in apisix",
		Run:     ApisixRunListCmd,
	}
	cmd.Flags().StringVar(&adminKey, "adminKey", EnvDefault("X_API_KEY", ""), "admin key")
	cmd.Flags().StringVar(&baseURL, "baseUrl", EnvDefault("TX_BASEURL", ""), "apisix base url")
	return cmd
}

func ApisixRunListCmd(cmd *cobra.Command, args []string) {
	var o apisix.ClusterOptions
	o.AdminKey = adminKey
	o.BaseURL = baseURL
	logger.Info(baseURL, adminKey)
	c, err := apisix.NewCluster(&o)
	if err != nil {
		logger.Error(err)
		return
	}
	s := apisix.NewSslClient(c)
	ss, err := s.List(context.TODO())
	if err != nil {
		logger.Error(err)
		return
	}

	for _, v := range ss {
		fmt.Printf("%s\n", v.Sni)
	}
}

func getCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get ssl in apisix",
		Args: cobra.MaximumNArgs(1),
		Run:   ApisixRunGetCmd,
	}
	cmd.Flags().StringVar(&adminKey, "adminKey", EnvDefault("X_API_KEY", ""), "admin key")
	cmd.Flags().StringVar(&baseURL, "baseUrl", EnvDefault("TX_BASEURL", ""), "apisix base url")
	cmd.Flags().StringVar(&sslId, "sslId", EnvDefault("TX_SSLID", ""), "apisix ssl id")
	return cmd
}

func ApisixRunGetCmd(cmd *cobra.Command, args []string) {
	var o apisix.ClusterOptions
	o.AdminKey = adminKey
	o.BaseURL = baseURL
	if len(args) == 1 {
		sslId = args[0]
		logger.Debug("sslId: %s", args)
	}
	c, err := apisix.NewCluster(&o)
	if err != nil {
		return
	}
	s := apisix.NewSslClient(c)
	if len(sslId) == 0 {
		return
	}
	ss , err := s.Get(context.TODO(), sslId)
	if err != nil {
		utils.ProcessError(err)
	}
	logger.Info(ss.Sni)
}

func updateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update ssl in apisix",
		Run:   ApisixRunUpdateCmd,
	}
	cmd.Flags().StringVar(&adminKey, "adminKey", EnvDefault("X_API_KEY", ""), "admin key")
	cmd.Flags().StringVar(&baseURL, "baseUrl", EnvDefault("TX_BASEURL", ""), "apisix base url")
	cmd.Flags().StringVar(&sslId, "sslId", EnvDefault("TX_SSLID", ""), "apisix ssl id")
	cmd.Flags().BoolVar(&updateAll, "updateAll", false, "update all apisix ssl")
	return cmd
}

func ApisixRunUpdateCmd(cmd *cobra.Command, args []string) {
	var o apisix.ClusterOptions
	o.AdminKey = adminKey
	o.BaseURL = baseURL
	c, err := apisix.NewCluster(&o)
	if err != nil {
		return
	}
	s := apisix.NewSslClient(c)
	if updateAll {
		ss, err := s.List(context.TODO())
		if err != nil {
			logger.Error(err)
			return
		}


		for _, v := range ss {
			if !strings.Contains(v.Sni, "kupeiai.cn"){
				v.Cert, v.Key = setCertAndKeyBySniFromAll()

				// if get nothing from files . break not to update ssl
				if v.Cert == "" || v.Key == "" {
					break
				}
				newSsl, err := s.Update(context.TODO(), v)
				if err != nil {
					logger.Error("update ssl error, ssl: %s, err: %s", v.ID, err)
					break
				}
				logger.Info("update sni: ", v.Sni)
				logger.Debug(newSsl)
			} else {
				v.Cert, v.Key = setCertAndKeyBySniFromFile(v.Sni)

				// if get nothing from files . break not to update ssl
				if v.Cert == "" || v.Key == "" {
					break
				}
				newSsl, err := s.Update(context.TODO(), v)
				if err != nil {
					logger.Error("update ssl error, ssl: %s, err: %s", v.ID, err)
					break
				}
				logger.Info("update sni: ", v.Sni)
				logger.Debug(newSsl)
			}
			
		}

		return
	}

	if len(args) == 1 {
		sslId = args[0]
		logger.Debug("sslId: ", sslId)
	}

	if sslId == "" {
		return
	}

	ss, err := s.Get(context.TODO(), sslId)
	if err != nil {
		logger.Error("get ssl error by id, ssl: %s, error: %s", sslId, err)
	}

	ss.Cert, ss.Key = setCertAndKeyBySniFromFile(ss.Sni)

	s.Update(context.TODO(), ss)
}

func setCertAndKeyBySniFromFile(sni string) (certPath, keyPath string) {
	var basePath string 
	if home, _ := os.UserHomeDir(); home != "" {
		basePath = filepath.Join(home, "tmp")
	}
	basePath = basePath + "/sslcert/"
	logger.Debug("basePaht: %s", basePath)
	domain := strings.Trim(sni, "*.")
	certPath = basePath + domain + "/fullchain.cer"
	keyPath = basePath + domain + "/" + domain + ".key"
	return utils.ReadStringFromFile(certPath), utils.ReadStringFromFile(keyPath)
}

func setCertAndKeyBySniFromAll ()  (certPath, keyPath string) {
	var basePath string 
	if home, _ := os.UserHomeDir(); home != "" {
		basePath = filepath.Join(home, "tmp")
	}
	basePath = basePath + "/sslcert/all.learnta.cn"
	certPath = basePath + "/fullchain.cer"
	keyPath = basePath  + "/all.learnta.cn.key"
	return utils.ReadStringFromFile(certPath), utils.ReadStringFromFile(keyPath)
}
