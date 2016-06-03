package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"

	"github.com/benkelaar/team-health-check/models"
	"github.com/julienschmidt/httprouter"
)

type (
	// CheckController provides endpoints for Check
	CheckController struct {
		session *mgo.Session
	}
)

// NewCheckController creates new CheckController
func NewCheckController(s *mgo.Session) *CheckController {
	return &CheckController{s}
}

// GetCheck gets a check
func (cc CheckController) GetCheck(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var id = p.ByName("id")
	fmt.Printf("Call to get check %s\n", id)
	// Stub an example user
	c := models.Check{
		ID:   id,
		Name: "Bart Enkelaar",
		Team: "Team15b",
		Health: map[string]models.State{
			"Ownership": models.State{
				Level:     models.Orange,
				Direction: models.Up,
			},
		},
		Timestamp: now(),
	}

	cj, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", cj)
}

// PostCheck posts a new check, check with ID is returned
func (cc CheckController) PostCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Stub an user to be populated from the body
	c := models.Check{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&c)

	// Add an Id
	c.ID = "foo"
	c.Timestamp = now()
	fmt.Println(c)

	// Marshal provided interface into JSON structure
	cj, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", cj)
}

func now() int64 {
	return time.Now().UnixNano() / 1000000
}
