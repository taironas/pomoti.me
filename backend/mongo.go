package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Period struct {
	Type string
}

// mongo handler returns a json file with a mongo message.
//
func mongo(w http.ResponseWriter, r *http.Request) {

	uri := os.Getenv("MONGOLAB_URI")
	log.Println("uri found: ", uri)
	var err error
	var session *mgo.Session

	log.Println("start dial")
	if session, err = mgo.Dial(uri); err != nil {
		panic(err)
	}
	log.Println("dial done")

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("pomotime").C("period")

	var p1, p2 Period
	p1 = Period{"pomodoro"}
	p2 = Period{"rest"}

	if err = c.Insert(&p1, &p2); err != nil {
		log.Fatal(err)
	}

	var result1, result2 Period

	if err = c.Find(bson.M{"type": "pomodoro"}).One(&result1); err != nil {
		log.Fatal(err)
	}

	if err = c.Find(bson.M{"type": "rest"}).One(&result2); err != nil {
		log.Fatal(err)
	}

	data := []struct {
		Message string
		Period  Period
	}{
		{
			"mongo",
			result1,
		},
		{
			"mongo",
			result2,
		},
	}

	if err := renderJson(w, data); err != nil {
		log.Println(err)
	}
}
