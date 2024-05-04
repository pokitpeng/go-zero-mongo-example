package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

const ItemCollectionName = "item"

var _ ItemModel = (*customItemModel)(nil)

type (
	// ItemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customItemModel.
	ItemModel interface {
		itemModel
	}

	customItemModel struct {
		*defaultItemModel
	}
)

// NewItemModel returns a model for the mongo.
func NewItemModel(url, db string, c cache.CacheConf, opts ...cache.Option) ItemModel {
	conn := monc.MustNewModel(url, db, ItemCollectionName, c, opts...)
	ensureIndexes(context.Background(), new(Item), conn)
	return &customItemModel{
		defaultItemModel: newDefaultItemModel(conn),
	}
}
