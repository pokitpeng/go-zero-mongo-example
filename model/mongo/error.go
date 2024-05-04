package model

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrNotFound        = mon.ErrNotFound
	ErrInvalidObjectId = errors.New("invalid objectId")
)

const (
	TagMongo   = "mongo"
	TagJson    = "json"
	TagBson    = "bson"
	TagUnique  = "unique"
	TagUniques = "unique-"
	TagIndex   = "index"
)

func ensureIndexes(ctx context.Context, model any, collection *monc.Model) {
	name, idxes := ParseModel(model)
	_, err := collection.Indexes().CreateMany(ctx, idxes)
	if err != nil {
		panic(fmt.Errorf("create indexes for model %s failed, %w", name, err))
	}
}

func ParseModel(model any) (modelName string, indexes []mongo.IndexModel) {
	rf := reflect.ValueOf(model).Elem().Type()
	modelName = rf.Name()

	unionUniques := make(map[string][]string)
	for i, max := 0, rf.NumField(); i < max; i++ {
		field := rf.Field(i)

		// 查询mongo标记
		mongoTag, ok := field.Tag.Lookup(TagMongo)
		if !ok {
			continue
		}
		// 查询索引关联标记
		index, unique, uniques := parseTagMongo(mongoTag)
		if !index && !unique && len(uniques) == 0 {
			continue
		}
		// 有索引标记
		fieldName := jsonField(field)

		if unique {
			indexes = append(indexes, mongo.IndexModel{
				Keys:    bson.D{{Key: fieldName, Value: 1}},
				Options: options.Index().SetUnique(true),
			})
		} else if index {
			indexes = append(indexes, mongo.IndexModel{
				Keys: bson.D{{Key: fieldName, Value: 1}},
			})
		}

		// 联合唯一
		for _, unique := range uniques {
			unionUniques[unique] = append(unionUniques[unique], fieldName)
		}
	}
	for _, fields := range unionUniques {
		var keys bson.D
		for _, field := range fields {
			keys = append(keys, bson.E{Key: field, Value: 1})
		}
		indexes = append(indexes, mongo.IndexModel{
			Keys:    keys,
			Options: options.Index().SetUnique(true),
		})
	}
	return
}

func jsonField(field reflect.StructField) string {
	val := field.Tag.Get(TagJson)
	fields := strings.Split(val, ",")
	if fields[0] == "" {
		panic("json field is empty")
	}
	return fields[0]
}

func parseTagMongo(val string) (index, unique bool, uniques []string) {
	idx := make(map[string]bool)
	for _, field := range strings.Split(val, ",") {
		field = strings.TrimSpace(field)
		if field == "" {
			continue
		}
		if field == TagIndex {
			index = true
		} else if field == TagUnique {
			unique = true
		} else if strings.HasPrefix(field, TagUniques) {
			idx[field] = true
		}
	}
	for key := range idx {
		uniques = append(uniques, key)
	}
	return
}
