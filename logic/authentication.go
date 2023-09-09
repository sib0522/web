package logic

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

type context struct {
	echo.Context
}

func JWTJson(c echo.Context) error {
	ctx := context{c}
	err, token := ctx.getJwtToken()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, token)
}

func JWTCheck(c echo.Context) error {
	ctx := context{c}
	err, _ := ctx.getJwtToken()
	if err != nil {
		return err
	}
	return nil
}

func (c *context) getJwtToken() (error, *TokenData) {
	var ok bool
	var token *jwt.Token
	var claims *TokenData

	token, ok = c.Get("user").(*jwt.Token)
	if !ok {
		err := errors.New("JWT token missing or invalid")
		return err, nil
	}

	claims, ok = token.Claims.(*TokenData)
	if !ok {
		err := errors.New("failed to cast claims as jwt.MapClaims")
		return err, nil
	}
	return nil, claims
}
