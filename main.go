package main

import (
	"fmt"
	"net/http"
	"strconv"

	chordTransposer "github.com/halilcagriakkuzu/go-chord-transposer"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Chord struct {
	mgm.DefaultModel `bson:",inline"`
	Body             string `json:"body" bson:"body"`
}

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

	// Routes
	e.GET("/chords/:id/:t", getChord)

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
