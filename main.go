package main

import (
	"html/template"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/benkelaar/team-health-check/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	var port = "3000"

	log.Println("Starting team health check server")
	r := httprouter.New()
	cc := controllers.NewCheckController(getSession())
	tc := controllers.NewUserController(getSession())

	r.ServeFiles("/static/*filepath", http.Dir("static/"))
	r.GET("/", serveIndex)

	r.GET("/v1/teams", tc.GetTeams)
	r.GET("/v1/teams/:team/members", tc.GetMembers)

	r.GET("/v1/checks/:id", cc.GetCheck)
	r.POST("/v1/checks", cc.PostCheck)

	r.GET("/v1/checks", cc.FindCheck)

	log.Printf("Listening on port %s\n", port)
	http.ListenAndServe("localhost:"+port, r)
	log.Printf("Closing team health check server")
}

func serveIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Panic(err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}
	return s
}
