package models

import "gopkg.in/mgo.v2/bson"

const (
	// Up direction
	Up = "up"
	// Same or no direction
	Same = "same"
	// Down direction
	Down = "down"
	// Red level or serious problems
	Red = "red"
	// Orange level or needs a lot of improvement
	Orange = "orange"
	// Yellow level or some issues
	Yellow = "yellow"
	// Green level or absolutely fine
	Green = "green"
)

type (
	// State is value of a single check
	State struct {
		Level     string `json:"level" bson:"level"`
		Direction string `json:"direction" bson:"direction"`
	}

	// Check data object is a stored series of checks
	Check struct {
		ID        bson.ObjectId    `json:"id" bson:"_id"`
		Name      string           `json:"name" bson:"name"`
		Team      string           `json:"team" bson:"team"`
		Health    map[string]State `json:"health" bson:"health"`
		Timestamp int64            `json:"time" bson:"time"`
	}
)
