package apisix

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/wonderivan/logger"
)

type sslClient struct {
	url     string
	cluster *cluster
}

func NewSslClient(c *cluster) SSL {
	return &sslClient{
		url:     c.baseURL + "/ssl",
		cluster: c,
	}
}

func (s *sslClient) Get(ctx context.Context, id string) (*Ssl, error) {

	// TODO Add mutex here to avoid dog-pile effection.
	url := s.url + "/" + id
	resp, err := s.cluster.getResource(ctx, url)
	if err != nil {
		return nil, err
	}
	ssl, err := resp.Item.ssl()
	if err != nil {
		logger.Warn("failed to get ssl infomation.")
		return nil, err
	}

	return ssl, nil
}

// List is only used in cache warming up. So here just pass through
// to APISIX.
func (s *sslClient) List(ctx context.Context) ([]*Ssl, error) {

	sslItems, err := s.cluster.listResource(ctx, s.url)
	if err != nil {
		logger.Error("failed to list ssl: %s", err)
		return nil, err
	}

	var items []*Ssl
	for i, item := range sslItems.Node.Items {
		ssl, err := item.ssl()
		if err != nil {
			return nil, err
		}
		items = append(items, ssl)
		logger.Info("get sni list: ", ssl.Sni)
		logger.Debug("list ssl #%d, body: %s", i, string(item.Value))
	}

	return items, nil
}

func (s *sslClient) Create(ctx context.Context, obj *Ssl) (*Ssl, error) {
	data, err := json.Marshal(Ssl{
		Sni:    obj.Sni,
		Cert:   obj.Cert,
		Key:    obj.Key,
		Status: obj.Status,
	})
	if err != nil {
		return nil, err
	}
	url := s.url + "/" + obj.ID
	resp, err := s.cluster.createResource(ctx, url, bytes.NewReader(data))
	if err != nil {
		logger.Error("failed to create ssl: %s", err)
		return nil, err
	}

	ssl, err := resp.Item.ssl()
	if err != nil {
		return nil, err
	}

	return ssl, nil
}

func (s *sslClient) Delete(ctx context.Context, obj *Ssl) error {
	url := s.url + "/" + obj.ID
	if err := s.cluster.deleteResource(ctx, url); err != nil {
		return err
	}
	return nil
}

func (s *sslClient) Update(ctx context.Context, obj *Ssl) (*Ssl, error) {

	url := s.url + "/" + obj.ID
	data, err := json.Marshal(Ssl{
		ID:     obj.ID,
		Sni:    obj.Sni,
		Cert:   obj.Cert,
		Key:    obj.Key,
		Status: obj.Status,
	})
	if err != nil {
		return nil, err
	}
	logger.Debug("updating ssl, body: %s, url: %s", string(data), url)
	// log.Debugw("updating ssl", zap.ByteString("body", data), zap.String("url", url))
	resp, err := s.cluster.updateResource(ctx, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	ssl, err := resp.Item.ssl()
	if err != nil {
		return nil, err
	}

	return ssl, nil
}
