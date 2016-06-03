package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/benkelaar/team-health-check/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	var port = "3000"

	fmt.Println("Starting team health check server")
	r := httprouter.New()
	cc := controllers.NewCheckController(getSession())

	r.GET("/v1/checks/:id", cc.GetCheck)
	r.POST("/v1/checks/", cc.PostCheck)

	fmt.Printf("Listening on port %s\n", port)
	http.ListenAndServe("localhost:"+port, r)
	defer fmt.Printf("Closing team health check server")
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
