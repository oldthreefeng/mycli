package tls

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/oldthreefeng/mycli/utils"
)

const (
	// TLSCertKey is the key for tls certificates in a TLS secert.
	TLSCertKey = "tls.crt"
	// TLSPrivateKeyKey is the key for the private key field in a TLS secret.
	TLSPrivateKeyKey = "tls.key"
)


var (
	DefaultTransport *http.Transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   5*time.Second,
			KeepAlive: 20 * time.Second,
		}).Dial,
	}
	ns string
	// just like https://apollo.com/configfiles/json/apolloId/default
	ApolloUrl = utils.EnvDefault("ApolloUrl", "")
)


func deployTls(ns string) {
	namespace := ns 
	if len(ns) == 0 {
		return
	}
	openurl := fmt.Sprintf("%s/star.%s", ApolloUrl ,namespace)
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	res, err := HttpGet(ctx, openurl, nil)
	if err != nil {
		fmt.Println("获取apollo失败，",err)
		return
	}
	tl := tls{}
	json.Unmarshal(res, &tl)
	crtFile := namespace + ".crt"
	keyFile := namespace + ".key"
	if isValidCertOrKey(TLSCertKey, tl.Crt) && isValidCertOrKey(TLSPrivateKeyKey, tl.Key) {
		err = ioutil.WriteFile(crtFile, []byte(tl.Crt), 0644)
		err = ioutil.WriteFile(keyFile, []byte(tl.Key), 0644)
		if err != nil {
			return
		}
		fmt.Println(tl.Key)
		fmt.Println(tl.Crt)
		return
	}
	fmt.Println("手动更新")
}

type tls struct {
	Crt string `json:"tls.crt"`
	Key string `json:"tls.key"`
}

type ApplySecret struct {
	name    string
	tlsData map[string]string
}

func HttpGet(ctx context.Context, url string, header map[string]string) (respBody []byte, err error) {
	client := &http.Client{Transport: DefaultTransport}
	req, err := http.NewRequest("GET", url, nil)
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {

		return nil, fmt.Errorf("http dial fail (%s)", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("http request fail, url (%s), code (%d), body (%s)", url, resp.StatusCode, string(body))
	}

	return body, err
}

func HttpPost(ctx context.Context, url string, body []byte) (respBody []byte, err error) {
	client := &http.Client{Transport: DefaultTransport}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	req.Header.Add("Content-Type", "application/json")
	req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http dial fail (%s)", err.Error())
	}
	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("http request fail, url (%s), code (%d), body (%s)", url, resp.StatusCode, string(body))
	}
	return rbody, err
}

func isValidCertOrKey(key, value string) bool {

	if key != TLSCertKey && key != TLSPrivateKeyKey {
		return false
	}
	var d1 = []byte(value)
	DERBlock, rest := pem.Decode(d1)
	if DERBlock == nil {
		return false
	}
	switch key {
	case TLSCertKey:

		_, err := x509.ParseCertificate(DERBlock.Bytes)
		if err != nil {
			return false
		}
		// 中间证书
		DERChainBlock, _ := pem.Decode(rest)

		if DERChainBlock == nil {
			return false
		}
		_, err = x509.ParseCertificate(DERChainBlock.Bytes)
		if err != nil {
			return false
		}
	case TLSPrivateKeyKey:
		// 支持的私钥类型， 常见的就这几种
		var err error
		if DERBlock.Type == "RSA PRIVATE KEY" {
			//RSA PKCS1
			_, err = x509.ParsePKCS1PrivateKey(DERBlock.Bytes)
		} else if DERBlock.Type == "PRIVATE KEY" {
			//pkcs8格式的私钥解析
			_, err = x509.ParsePKCS8PrivateKey(DERBlock.Bytes)
		}
		if err != nil {
			return false
		}
	}

	return true
}
