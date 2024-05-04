package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	CacheRedis cache.CacheConf
	MongoUri   string `json:",default=mongodb://user:passwd@127.0.0.1:27017/"`
	Database   string `json:",default=go_zero_example"`
	Auth       struct {
		AccessSecret string `json:""`
		AccessExpire int64  `json:""`
	} `json:""`
}
