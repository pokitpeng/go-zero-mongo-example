package svc

import (
	"go_zero_example/internal/config"
	model "go_zero_example/model/mongo"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	ItemModel model.ItemModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	UserModel := model.NewUserModel(c.MongoUri, c.Database, c.CacheRedis)
	ItemModel := model.NewItemModel(c.MongoUri, c.Database, c.CacheRedis)
	return &ServiceContext{
		Config:    c,
		UserModel: UserModel,
		ItemModel: ItemModel,
	}
}
