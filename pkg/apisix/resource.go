package apisix

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/wonderivan/logger"
)

type getResponse struct {
	Item item `json:"node"`
}

type listResponse struct {
	Count string `json:"count"`
	Node  node   `json:"node"`
}

type createResponse struct {
	Action string `json:"action"`
	Item   item   `json:"node"`
}

type updateResponse = createResponse

type node struct {
	Key   string `json:"key"`
	Items items  `json:"nodes"`
}

type items []item

// items implements json.Unmarshaler interface.
// lua-cjson doesn't distinguish empty array and table,
// and by default empty array will be encoded as '{}'.
// We have to maintain the compatibility.
func (items *items) UnmarshalJSON(p []byte) error {
	if p[0] == '{' {
		if len(p) != 2 {
			return errors.New("unexpected non-empty object")
		}
		return nil
	}
	var data []item
	if err := json.Unmarshal(p, &data); err != nil {
		return err
	}
	*items = data
	return nil
}

// ssl decodes item.Value and converts it to v1.Ssl.
func (i *item) ssl() (*Ssl, error) {
	logger.Debug("got ssl: %s", string(i.Value))
	var ssl Ssl
	if err := json.Unmarshal(i.Value, &ssl); err != nil {
		return nil, err
	}

	list := strings.Split(i.Key, "/")
	id := list[len(list)-1]
	ssl.ID = id
	// ssl.Group = clusterName
	// ssl.FullName = id
	return &ssl, nil
}

type item struct {
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}
