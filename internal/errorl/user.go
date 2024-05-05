package errorl

import (
	"go_zero_example/pkg/errorx"
)

const (
	CodeUserAlreadyExist = iota + 40100
	CodeUserNotFound
	CodeUsernameOrPasswordWrong
	CodeUserModelInsertFailed
	CodeGenerateTokenFailed
)

func UserAlreadyExist(opts ...errorx.Option) error {
	return errorx.New(CodeUserAlreadyExist, "用户名已存在", opts...)
}

func UserNotFound(opts ...errorx.Option) error {
	return errorx.New(CodeUserNotFound, "用户不存在", opts...)
}

func UsernameOrPasswordWrong(opts ...errorx.Option) error {
	return errorx.New(CodeUsernameOrPasswordWrong, "用户名或密码错误", opts...)
}

func UserModelInsertFailed(opts ...errorx.Option) error {
	return errorx.New(CodeUserModelInsertFailed, "写入用户库失败", opts...)
}

func GenerateTokenFailed(opts ...errorx.Option) error {
	return errorx.New(CodeGenerateTokenFailed, "生成token失败", opts...)
}
