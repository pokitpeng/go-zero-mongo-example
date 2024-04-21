package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
)

const UserCollectionName = "user"

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the mongo.
func NewUserModel(url, db string) UserModel {
	conn := mon.MustNewModel(url, db, UserCollectionName)
	ensureIndexes(context.Background(), new(User), conn)
	return &customUserModel{
		defaultUserModel: newDefaultUserModel(conn),
	}
}
