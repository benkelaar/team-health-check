package controllers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// UserController provides endpoints for Check
	UserController struct {
		session *mgo.Session
	}
)

// NewUserController creates new CheckController
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetTeams gets all available teams
func (uc UserController) GetTeams(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session := uc.session.Copy()
	defer session.Close()
	var teams []string

	if err := connect(session).Find(nil).Distinct("team", &teams); err != nil {
		w.WriteHeader(404)
		log.Fatal(err)
		return
	}

	writeCheckResponse(w, teams, 200)
}

// GetMembers gets the membets of a specific team
func (uc UserController) GetMembers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	team := p.ByName("team")
	session := uc.session.Copy()
	defer session.Close()
	var members []string

	if err := connect(session).Find(bson.M{"team": team}).Distinct("name", &members); err != nil {
		w.WriteHeader(404)
		log.Fatal(err)
		return
	}

	writeCheckResponse(w, members, 200)
}
