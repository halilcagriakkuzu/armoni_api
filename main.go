package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	chordTransposer "github.com/halilcagriakkuzu/go-chord-transposer"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type jwtCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type Chord struct {
	mgm.DefaultModel `bson:",inline"`
	Body             string `json:"body" bson:"body"`
}

const jwtSecretString string = "TopSecret"

func main() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	err := mgm.SetDefaultConfig(nil, "armoni", options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		panic(err)
	}

	// Public Routes
	e.POST("/login", login)
	e.POST("/signup", signup)
	e.GET("/chords/:id/:t", getChord)

	// /my-profile/....
	r := e.Group("/my-profile")
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(jwtSecretString),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", myProfile)
	r.GET("/favorites", myFavorites)

	// Start server
	e.Logger.Fatal(e.Start(":1337"))
}

func getChord(c echo.Context) error {
	id := c.Param("id")
	t, err := strconv.Atoi(c.Param("t"))
	if err != nil {
		panic(err)
	}

	song := &Chord{}
	coll := mgm.Coll(song)
	err = coll.FindByID(id, song)
	if err != nil {
		panic(err)
	}

	resultSong := chordTransposer.TransposeChords(song.Body, t, "%v")

	return c.String(http.StatusOK, fmt.Sprintf("%+v\n", resultSong))
}

func login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Throws unauthorized error
	if email != "test@email.com" || password != "123" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtSecretString))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func signup(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	email := claims.Email
	return c.String(http.StatusOK, "signup: "+email+"!")
}

func myProfile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	email := claims.Email
	return c.String(http.StatusOK, "profil: "+email+"!")
}

func myFavorites(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	email := claims.Email
	return c.String(http.StatusOK, "favoriler: "+email+"!")
}
