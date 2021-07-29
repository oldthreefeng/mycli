package redis

import (
	"context"
	"crypto/tls"

	"github.com/go-redis/redis/v8"
	"github.com/wonderivan/logger"
)

type RedisCfg struct {
	host     string
	sni      string
	password string
	DB       int
	Client   *redis.Client
}

func NewRedisDb(host, password, sni string, db int) *RedisCfg {
	cfg := &RedisCfg{
		host: host,
		password: password,
		sni: sni,
		DB: db,
	}
	opts := &redis.Options{
		Addr:     cfg.host,
		Password: cfg.password,
		DB:       cfg.DB,
	}
	if cfg.sni != "" {
		opts.TLSConfig = &tls.Config{
			ServerName: cfg.sni,
		}
	}
	cfg.Client = redis.NewClient(opts)
	return cfg
}

func (c *RedisCfg) DeleteByPrefix(ctx context.Context, prefix string) {
	var foundedRecordCount int = 0
	iter := c.Client.Scan(ctx, 0, prefix, 0).Iterator()
	logger.Info("YOUR SEARCH PATTERN= %s", prefix)
	for iter.Next(ctx) {
		c.deleteByKey(ctx, iter.Val())
		foundedRecordCount++
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	logger.Info("Deleted Count %d", foundedRecordCount)
	c.Client.Close()
}

func (c *RedisCfg) deleteByKey(ctx context.Context, key string) {
	logger.Info("Deleted key = %s", key)
	c.Client.Del(ctx, key)
}

func (c *RedisCfg) DeleteKeys(ctx context.Context, keys []string) {
	for _, k := range keys {
		c.deleteByKey(ctx, k)
	}
}
