package main

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// mongo handler returns a json file with a mongo message.
//
func mongo(w http.ResponseWriter, r *http.Request) {

	var uri string
	uri = getMongoURI()

	log.Println("uri found: ", uri)

	var err error
	var session *mgo.Session

	log.Println("start dial")
	if session, err = mgo.Dial(uri); err != nil {
		panic(err)
	}
	log.Println("dial done")

	defer session.Close()
	log.Println("setting session mode")
	session.SetMode(mgo.Monotonic, true)

	log.Println("trying to get pomotime database")
	collection := session.DB("").C("period")

	log.Println("start insert to period entity")
	var p1, p2 Period
	p1 = Period{Type: "pomodoro"}
	p2 = Period{Type: "rest"}

	if err = collection.Insert(&p1, &p2); err != nil {
		log.Fatal(err)
	}

	var result1, result2 Period

	log.Println("start search for pomodoro object")
	if err = collection.Find(bson.M{"type": "pomodoro"}).One(&result1); err != nil {
		log.Fatal(err)
	}

	log.Println("start search for rest object")
	if err = collection.Find(bson.M{"type": "rest"}).One(&result2); err != nil {
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

	log.Println("sending response")
	if err := renderJson(w, data); err != nil {
		log.Println(err)
	}
}
