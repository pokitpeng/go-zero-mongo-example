// Code generated by goctl. DO NOT EDIT.
package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/proto"
)

var prefixUserCacheKey = "cache:user:"

type userModel interface {
	Insert(ctx context.Context, data *User) error
	FindOne(ctx context.Context, id string) (*User, error)
	Find(ctx context.Context, filter *User) ([]*User, error)
	Update(ctx context.Context, data *User) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string, realDelete bool) (int64, error)
	List(ctx context.Context, filter *User, offset, limit int64, orderBy, order string) ([]*User, int64, error)
}

type defaultUserModel struct {
	conn *monc.Model
}

func newDefaultUserModel(conn *monc.Model) *defaultUserModel {
	return &defaultUserModel{conn: conn}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
	}
	now := time.Now().Unix()
	data.CreateAt = proto.Int64(now)
	data.UpdateAt = proto.Int64(now)
	data.DeleteAt = proto.Int64(0)

	key := prefixUserCacheKey + data.ID.Hex()
	_, err := m.conn.InsertOne(ctx, key, data)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id string) (*User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data User
	key := prefixUserCacheKey + id
	err = m.conn.FindOne(ctx, key, &data, bson.M{"_id": oid, "DeleteAt": 0})
	switch err {
	case nil:
		return &data, nil
	case monc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Find(ctx context.Context, filter *User) ([]*User, error) {
	var datas []*User
	if filter == nil {
		filter = new(User)
	}
	filter.DeleteAt = proto.Int64(0)

	err := m.conn.Find(ctx, &datas, filter)
	switch err {
	case nil:
	case monc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
	return datas, err
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) (*mongo.UpdateResult, error) {
	data.UpdateAt = proto.Int64(time.Now().Unix())
	key := prefixUserCacheKey + data.ID.Hex()
	res, err := m.conn.UpdateOne(ctx, key, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultUserModel) Delete(ctx context.Context, id string, realDelete bool) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}
	if realDelete {
		key := prefixUserCacheKey + id
		res, err := m.conn.DeleteOne(ctx, key, bson.M{"_id": oid})
		return res, err
	} else {
		key := prefixUserCacheKey + id
		data, err := m.FindOne(ctx, id)
		if err != nil {
			return 0, err
		}
		data.DeleteAt = proto.Int64(time.Now().Unix())
		res, err := m.conn.UpdateOne(ctx, key, bson.M{"_id": oid}, bson.M{"$set": data})
		m.conn.DelCache(ctx, key)
		return res.ModifiedCount, err
	}
}

func (m *defaultUserModel) List(ctx context.Context, filter *User, offset, limit int64, orderBy, order string) ([]*User, int64, error) {
	var datas []*User
	if filter == nil {
		filter = new(User)
	}
	filter.DeleteAt = proto.Int64(0)
	total, err := m.conn.CountDocuments(ctx, filter)
	switch err {
	case nil:
	case monc.ErrNotFound:
		return nil, 0, ErrNotFound
	default:
		return nil, 0, err
	}
	findOptions := &options.FindOptions{
		Limit: proto.Int64(int64(limit)),
		Skip:  proto.Int64(int64(offset)),
	}
	if orderBy != "" {
		if order == "Desc" {
			findOptions.SetSort(bson.D{{orderBy, -1}})
		} else {
			findOptions.SetSort(bson.D{{orderBy, 1}})
		}
	}
	err = m.conn.Find(ctx, &datas, filter, findOptions)
	switch err {
	case nil:
	case monc.ErrNotFound:
		return nil, total, ErrNotFound
	default:
		return nil, total, err
	}
	return datas, total, err
}
