package main

import (
	"armoni/auth"
	"armoni/handler"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// MongoDb
	err := mgm.SetDefaultConfig(nil, "armoni", options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		panic(err)
	}

	println("Deneme")

	e.GET("/login", handler.Login())
	e.POST("/signup", handler.Signup())

	e.GET("/chords/:id/:t", handler.GetChord())

	myGroup := e.Group("/my")
	myGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &auth.Claims{},
		SigningKey: []byte(auth.GetJWTSecret()),
	}))

	// Attach jwt token refresher.
	//myGroup.Use(auth.TokenRefresherMiddleware)

	myGroup.GET("/profile", handler.MyProfile())
	myGroup.GET("/favorites", handler.MyFavorites())

	e.Logger.Fatal(e.Start(":1337"))
}
