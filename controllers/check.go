package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

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
	log.Printf("Call to get check %s\n", id)

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	session := cc.session.Copy()
	defer session.Close()
	oid := bson.ObjectIdHex(id)
	c := models.Check{}

	if err := connect(session).FindId(oid).One(&c); err != nil {
		w.WriteHeader(404)
		return
	}

	writeCheckResponse(w, c, 200)
}

// FindCheck finds team or user checks
func (cc CheckController) FindCheck(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	team := r.FormValue("team")
	user := r.FormValue("user")
	log.Printf("Call to find checks for team %s and user %s\n", team, user)

	if team == "" && user == "" {
		w.WriteHeader(400)
		return
	}

	session := cc.session.Copy()
	defer session.Close()
	checks := []models.Check{}
	q := bson.M{
		"name": team,
	}
	if user != "" {
		q["name"] = user
	}

	if err := connect(session).Find(q).All(&checks); err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}

	writeCheckResponse(w, checks, 200)
}

// PostCheck posts a new check, check with ID is returned
func (cc CheckController) PostCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := cc.session.Copy()
	defer session.Close()

	c := models.Check{}
	json.NewDecoder(r.Body).Decode(&c)
	c.ID = bson.NewObjectId()
	c.Timestamp = now()
	log.Println(c)

	connect(session).Insert(c)
	writeCheckResponse(w, c, 201)
}

func connect(session *mgo.Session) *mgo.Collection {
	return session.DB("teamhealth").C("checks")
}

func writeCheckResponse(w http.ResponseWriter, data interface{}, code int) {
	dj, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// Write content-type, statuscode, payload
	headers := w.Header()
	headers.Set("Content-Type", "application/json")
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Access-Control-Allow-Headers", "Content-Type")
	headers.Add("Access-Control-Max-Age", "86400")

	w.WriteHeader(code)
	fmt.Fprintf(w, "%s", dj)
}

func now() int64 {
	return time.Now().UnixNano() / 1000000
}
