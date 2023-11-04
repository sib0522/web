package ApiUserObatain

import "github.com/labstack/echo/v4"

type Request struct {
	Uuid   string
	Reason uint32
	Value  uint32
}

type Response struct {
	GetExp   uint64
	GetMoney uint64
}

func (r Request) Unmarshal(c echo.Context) {}
func (r Response) Marshal(c echo.Context)  {}
