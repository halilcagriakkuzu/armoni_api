package auth

import (
	"armoni/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/golang-jwt/jwt"
)

const (
	jwtSecretKey        = "some-secret-key"
	jwtRefreshSecretKey = "some-refresh-secret-key"
)

func GetJWTSecret() string {
	return jwtSecretKey
}

func GetRefreshJWTSecret() string {
	return jwtRefreshSecretKey
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateTokens(user *model.User) (string, string, error) {
	accessToken, err := generateAccessToken(user)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := generateRefreshToken(user)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func generateAccessToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

func generateRefreshToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	return generateToken(user, expirationTime, []byte(GetRefreshJWTSecret()))
}

func generateToken(user *model.User, expirationTime time.Time, secret []byte) (string, error) {
	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JWTErrorChecker(err error, c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status": 403,
		"error":  "Token is not valid",
	})
}

func User(c echo.Context) *Claims {
	if c.Get("user").(*jwt.Token).Valid {
		return c.Get("user").(*jwt.Token).Claims.(*Claims)
	}
	return nil
}
