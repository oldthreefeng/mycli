package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"unicode"
)

func Do(domain, uri, method, accessToken string, post []byte) error {
	_, err := do(domain, uri, method, accessToken, nil, post)
	return err
}
func DoBody(domain, uri, method, accessToken string, post []byte) ([]byte, error) {
	return do(domain, uri, method, accessToken, nil, post)
}
func DoBodyAddHeader(domain, uri, method, accessToken string, headers map[string]string, post []byte) ([]byte, error) {
	return do(domain, uri, method, accessToken, headers, post)
}
func do(domain, uri, method, accessToken string, headers map[string]string, post []byte) ([]byte, error) {
	req, err := http.NewRequest(method, domain+uri, bytes.NewReader(post))
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-agent", "apisix")
	req.Header.Set("X-API-KEY", accessToken)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		//logger.Error("response error is %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
		return nil, fmt.Errorf("respone status is not correct")
	} else {
		var out map[string]interface{}
		_ = json.Unmarshal(body, &out)
		if code, ok := out["code"].(float64); ok && code != 200 {
			return nil, fmt.Errorf(out["message"].(string))
		}
	}
	return body, nil
}

func IsUrl(u string) (url.URL, bool) {
	if uu, err := url.Parse(u); err == nil && uu != nil && uu.Host != "" {
		return *uu, true
	}
	return url.URL{}, false
}

func httpGetBodyByte(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return ReadByteFrom(resp.Body)
}

var (
	jsonPrefix = []byte("{")
)

// GetByteStream is read file/url/stdin to []byte stream
func GetByteStream(s string) (bt []byte, err error) {
	if s == "-" {
		// read stdin
		return ReadByteFrom(os.Stdin)
	} else if _, u := IsUrl(s); u {
		// read url
		return httpGetBodyByte(s)
	} else if FileExist(s) {
		// read file
		return ioutil.ReadFile(s)
	} else {
		// other case
		return
	}
}
func GetFile(s string) (bt []byte, err error) {
	if _, u := IsUrl(s); u {
		// read url
		return httpGetBodyByte(s)
	} else if FileExist(s) {
		// read file
		return ioutil.ReadFile(s)
	} else {
		// other case
		return nil, errors.New("fetch file is error")
	}
}

func ReadByteFrom(in io.Reader) (bt []byte, err error) {
	var b bytes.Buffer
	_, err = b.ReadFrom(in)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func HasJSONPrefix(buf []byte) bool {
	return hasPrefix(buf, jsonPrefix)
}
func hasPrefix(buf []byte, prefix []byte) bool {
	trim := bytes.TrimLeftFunc(buf, unicode.IsSpace)
	return bytes.HasPrefix(trim, prefix)
}

func JsonConvert(from interface{}, to interface{}) error {
	var data []byte
	var err error
	if data, err = json.Marshal(from); err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(json.Unmarshal(data, to))
}

func ReadStringFromFile(s string) string {
	buf, err := ioutil.ReadFile(s)
	if err != nil {
		return ""
	}
	return string(buf)
}

func EnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex || val == "" {
		return defVal
	}
	return val
}

// IsEnvSet returns true if an environment variable is set
func IsEnvSet(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}
