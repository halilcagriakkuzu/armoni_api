package model

import "github.com/kamva/mgm/v3"

type Chord struct {
	mgm.DefaultModel `bson:",inline"`
	Body             string `json:"body" bson:"body"`
}
