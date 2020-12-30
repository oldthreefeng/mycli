package apisix

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/oldthreefeng/mycli/utils"
	"github.com/wonderivan/logger"
)

var (
	domain      = utils.EnvDefault("TX_BASEURL", "")
	accessToken = utils.EnvDefault("X_API_KEY", "")
)

type sslNode struct {
	routeKey string
	sslCert
}

type sslCert struct {
	Key  string `json:"key"`
	Cert string `json:"cert"`
	Sni  string `json:"sni"`
}

func getApisixAllssl() {
	uri := "/ssl"
	body, err := utils.DoBody(domain, uri, "GET", accessToken, nil)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}

func initRoutes() []sslNode {
	m := map[string]string{
		"00000000000000000556": "*.t2.ilearnta.com",
		// "00000000000000000559": "*.t2.kupeiai.com",
		// "00000000000000000558": "*.t3.kupeiai.com",
		// "00000000000000000560": "*.t1.kupeiai.com",
		// "00000000000000000555": "*.t3.kupei.cn",
		// "00000000000000000624": "*.t3.learnta.work",
		// "00000000000000000644": "*.beta.learnta.cn",
		"00000000000000000622": "*.t1.learnta.work",
		// "00000000000000000645": "*.staging.learnta.cn",
		"00000000000000000623": "*.t2.learnta.work",
		// "00000000000000000638": "*.beta.ilearnta.com",
		// "00000000000000000554": "*.t2.kupei.cn",
		// "00000000000000000553": "*.t1.kupei.cn",
		"00000000000000000557": "*.t3.ilearnta.com",
		// "00000000000000000243": "*.t3.learnta.work",
		// "00000000000000000420": "*.t3.kupei.cn",
		// "00000000000000000536": "*.t1.ilearnta.com",

		// "00000000000000000245": "*.t3.kupeiai.cn",
		// "00000000000000000246": "*.t2.kupeiai.cn",
		// "00000000000000000249": "*.t1.learnta.cn",
		// "00000000000000000250": "*.t2.learnta.cn",
		// "00000000000000000247": "*.t1.kupeiai.cn",
		// "00000000000000000251": "*.t3.learnta.cn",
	}
	var ss []sslNode
	var basePath string = "/home/louis/tmp/sslcert/"
	for k, v := range m {
		domain := strings.Trim(v, "*.")
		certPath := basePath + domain + "/fullchain.cer"
		keyPath := basePath + domain + "/" + domain + ".key"

		var s sslNode
		s.routeKey = k
		s.sslCert.Sni = v
		s.sslCert.Cert = utils.ReadStringFromFile(certPath)
		if s.sslCert.Cert == "" {
			break
		}
		s.sslCert.Key = utils.ReadStringFromFile(keyPath)
		if s.sslCert.Key == "" {
			break
		}
		logger.Info(domain)
		ss = append(ss, s)
	}

	return ss
}

func getSSLrouteByKeyId(r string) (s sslNode) {
	ss := initRoutes()
	for _, v := range ss {
		if v.routeKey == r {
			fmt.Println(v)
			return v
		}
	}
	return
}

func (s sslNode) setApisixByKey() {
	uri := "/ssl/" + s.routeKey
	post, err := json.Marshal(s.sslCert)
	if err != nil {
		logger.Error(err)
		return
	}
	fmt.Println(string(post))
	body, err := utils.DoBody(domain, uri, "PATCH", accessToken, post)
	if err != nil {
		logger.Error(err)
		return
	}
	fmt.Println(string(body))
}

func UpdateSSL() {
	ss := initRoutes()
	for _, v := range ss {
		v.setApisixByKey()
	}
}
