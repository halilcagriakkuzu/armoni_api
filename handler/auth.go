package handler

import (
	"armoni/auth"
	"armoni/model"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Load our "test" user.
		storedUser := model.LoadTestUser()
		// Initiate a new User struct.
		u := new(model.User)
		u.Email = c.FormValue("email")
		u.Password = c.FormValue("password")
		// Parse the submitted data and fill the User struct with the data from the SignIn form.
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		// Compare the stored hashed password, with the hashed version of the password that was received
		if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password)); err != nil {
			// If the two passwords don't match, return a 401 status
			return echo.NewHTTPError(http.StatusUnauthorized, "Password is incorrect")
		}
		accessToken, refreshToken, err := auth.GenerateTokens(storedUser)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
		}

		return c.JSON(http.StatusOK, echo.Map{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		})
	}
}

func Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*auth.Claims)
		email := claims.Email
		return c.String(http.StatusOK, "signup: "+email+"!")
	}
}
