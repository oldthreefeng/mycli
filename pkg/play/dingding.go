package play

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	TextTemplate = `{
    "msgtype": "text", 
    "text": {
        "content": "%s"
    }, 
    "at": {
        "atMobiles": [
            "%s"
        ], 
        "isAtAll": false
    }
}`
	LinkTemplate = `{
    "msgtype": "link", 
    "link": {
        "text": "%s",
        "title": "%s",
        "picUrl": "",
        "messageUrl": "%s"
    }
}`
	MarkTemplate = `{
     "msgtype": "markdown",
     "markdown": {
         "title":"%s",
         "text": "%s"
     },
    "at": {
        "atMobiles": [
            "%s"
        ], 
        "isAtAll": false
    }
}`
)

type Alarm interface {
	Dingding(Dingdingurl string) error
}

type MarkDowning struct {
	Title     string `json:"title"`
	Text      string `json:"text"`
	AtMobiles string `json:"atMobiles"` //应该是[]string,图方便,改成这个

}

type Linking struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type Text struct {
	MarkDowning
}

func (m Text) Dingding(DingDingUrl string) error {
	baseBody := fmt.Sprintf(TextTemplate, m.Text, m.AtMobiles)
	return dingding(DingDingUrl, baseBody)
}

func (m MarkDowning) Dingding(DingDingUrl string) error {
	baseBody := fmt.Sprintf(MarkTemplate, m.Title, m.Text, m.AtMobiles)
	return dingding(DingDingUrl, baseBody)
}

func (m Linking) Dingding(DingDingUrl string) error {
	baseBody := fmt.Sprintf(LinkTemplate, m.Title, m.Text, m.MessageUrl)
	return dingding(DingDingUrl, baseBody)
}

func dingding(DingDingUrl, baseBody string) error {
	req, err := http.NewRequest("POST", DingDingUrl, strings.NewReader(baseBody))
	if err != nil {
		return err
	}
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-agent", "firefox")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return nil
}
