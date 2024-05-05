package errorl

import (
	"go_zero_example/pkg/errorx"
)

const (
	CodeWriteSwaggerFileFailed = iota + 1000
	CodeCopyStructFailed
	CodeInvalidObjectId
)

func WriteSwaggerFileFailed(opts ...errorx.Option) error {
	return errorx.New(CodeWriteSwaggerFileFailed, "生成swagger文件失败", opts...)
}

func CopyStructFailed(opts ...errorx.Option) error {
	return errorx.New(CodeCopyStructFailed, "拷贝结构体失败", opts...)
}
func InvalidObjectId(opts ...errorx.Option) error {
	return errorx.New(CodeInvalidObjectId, "无效的ObjectId", opts...)
}
