package handler

import (
	"armoni/model"
	"fmt"
	"net/http"
	"strconv"

	chordTransposer "github.com/halilcagriakkuzu/go-chord-transposer"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
)

func GetChord() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		t, err := strconv.Atoi(c.Param("t"))
		if err != nil {
			panic(err)
		}

		song := &model.Chord{}
		coll := mgm.Coll(song)
		err = coll.FindByID(id, song)
		if err != nil {
			panic(err)
		}

		resultSong := chordTransposer.TransposeChords(song.Body, t, "%v")

		return c.String(http.StatusOK, fmt.Sprintf("%+v\n", resultSong))
	}
}
