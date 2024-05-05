package errorx

import (
	"fmt"
	"testing"
)

var UserNotFoundExample = &Errorx{
	Code: 404,
	Msg:  "用户不存在",
	Meta: map[string]any{
		"username": "tom",
	},
}

func TestErrorx_String(t *testing.T) {
	fmt.Println(UserNotFoundExample)
}

func TestIs(t *testing.T) {
	fmt.Println(Is(UserNotFoundExample, 404))
	fmt.Println(Is(UserNotFoundExample, 405))
}
