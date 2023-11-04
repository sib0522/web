package ApiUserLogin

import "github.com/labstack/echo/v4"

type Request struct {
	Uuid string
}

type Response struct {
	Uuid  string
	Level uint32
	Exp   uint64
}

func (r Request) Unmarshal(c echo.Context) {}
func (r Response) Marshal(c echo.Context)  {}

// Requestから送られてきたuuidが空だったら、新しいユーザーデータを作成し、それをResponseとして送る
// Requestから値が入っているuuidが送られてきたら、サーバー側でuuidのデータを取得し、それをResponseとして送る
