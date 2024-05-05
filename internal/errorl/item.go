package errorl

import (
	"go_zero_example/pkg/errorx"
)

const (
	CodeItemModelInsertFailed = iota + 40200
	CodeItemModelDeleteFailed
	CodeItemModelUpdateFailed
	CodeItemModelFindFailed
	CodeItemModelListFailed
	CodeItemNotFound
)

func ItemModelInsertFailed(opts ...errorx.Option) error {
	return errorx.New(CodeItemModelInsertFailed, "写入item库失败", opts...)
}

func ItemModelDeleteFailed(opts ...errorx.Option) error {
	return errorx.New(CodeItemModelDeleteFailed, "从item库删除元素失败", opts...)
}

func ItemModelUpdateFailed(opts ...errorx.Option) error {
	return errorx.New(CodeItemModelUpdateFailed, "更新item库元素失败", opts...)
}

func ItemModelFindFailed(opts ...errorx.Option) error {
	return errorx.New(CodeItemModelFindFailed, "从item库查询元素失败", opts...)
}

func ItemModelListFailed(opts ...errorx.Option) error {
	return errorx.New(CodeItemModelListFailed, "从item库列举元素失败", opts...)
}

func ItemNotFound(opts ...errorx.Option) error {
	return errorx.New(CodeItemNotFound, "item不存在", opts...)
}
