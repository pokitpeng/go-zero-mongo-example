package user

import (
	"go_zero_example/internal/types"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestCreateUser(t *testing.T) {
	client := resty.New()

	result := new(types.CreateUserResp)
	_, err := client.R().
		SetBody(&types.CreateUserReq{
			Username: "tom",
			Password: "123456",
		}).
		SetResult(result).
		Post("http://localhost:8888/api/user")
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatal(result.Msg)
	}
}
