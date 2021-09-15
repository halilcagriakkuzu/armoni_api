package handler

import (
	"armoni/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MyProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "profil: "+auth.User(c).Email+"!")
	}
}

func MyFavorites() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "favorite:"+auth.User(c).Email+"!")
	}
}
