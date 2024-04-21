package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 索引标记示例
//
// 普通索引
// Name string `bson:"Name,omitempty" json:"Name,omitempty" mongo:"index"`
// 唯一索引
// Name string `bson:"Name,omitempty" json:"Name,omitempty" mongo:"unique"`
// 联合唯一索引
// Name string `bson:"Name,omitempty" json:"Name,omitempty" mongo:"unique-name"`
// DeleteAt time.Time `bson:"DeleteAt,omitempty" json:"DeleteAt,omitempty" mongo:"unique-name"`

type {{.Type}} struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"ID,omitempty"`
	// TODO: Fill your own fields
	UpdateAt *int64 `bson:"UpdateAt,omitempty" json:"UpdateAt,omitempty"`
	CreateAt *int64 `bson:"CreateAt,omitempty" json:"CreateAt,omitempty"`
	DeleteAt *int64 `bson:"DeleteAt,omitempty" json:"DeleteAt,omitempty"`
}
