package main

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

// Period hold the information of an activity
type Period struct {
	Type string // pomodoro or rest
}

type Periods []Period

// createPeriod handler lets you create a period obect.
//
func createPeriod(w http.ResponseWriter, r *http.Request) {

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
	c := session.DB("").C("period")

	log.Println("start insert to period entity")
	var p Period
	p = Period{"pomodoro"}

	if err = c.Insert(&p); err != nil {
		log.Fatal(err)
	}

	data := struct {
		Status           int    `json:"status"`
		DeveloperMessage string `json:"developerMessage"`
		UserMessage      string `json:"userMessage"`
	}{
		Status:           200,
		DeveloperMessage: "app correctly stored period in database.",
		UserMessage:      "your data was stored correctly.",
	}

	log.Println("sending response")
	if err := renderJson(w, data); err != nil {
		log.Println(err)
	}
}

// getPeriods handler returns a list of periods.
//
func getPeriods(w http.ResponseWriter, r *http.Request) {

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

	log.Println("trying database")
	c := session.DB("").C("period")

	var results []Period

	log.Println("start search for all period objects")
	if err = c.Find(nil).All(&results); err != nil {
		log.Fatal(err)
	}

	data := struct {
		Status           int     `json:"status"`
		DeveloperMessage string  `json:"developerMessage"`
		UserMessage      string  `json:"userMessage"`
		Periods          Periods `json:"periods"`
	}{
		Status:           200,
		DeveloperMessage: "returning all periods from database.",
		UserMessage:      "",
		Periods:          results,
	}

	log.Println("sending response")
	if err := renderJson(w, data); err != nil {
		log.Println(err)
	}
}
