package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
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
func NewItemModel(url, db string) ItemModel {
	conn := mon.MustNewModel(url, db, ItemCollectionName)
	ensureIndexes(context.Background(), new(Item), conn)
	return &customItemModel{
		defaultItemModel: newDefaultItemModel(conn),
	}
}
