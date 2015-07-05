package main

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

// Period hold the information of an activity
type Period struct {
	Type string `json:"type"` // pomodoro or rest
}

type Periods []Period

// createPeriod handler lets you create a period obect.
//
func createPeriod(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		data := struct {
			Status           int    `json:"status"`
			DeveloperMessage string `json:"developerMessage"`
			UserMessage      string `json:"userMessage"`
		}{
			Status:           400,
			DeveloperMessage: "createPeriod only POST requests are supported for this route.",
			UserMessage:      "Oops, something went wrong, we are unable to save your data right now.",
		}

		log.Println("sending response")
		if err := renderJson(w, data); err != nil {
			log.Println(err)
		}
		return
	}

	var strType string

	if strType = r.FormValue("type"); len(strType) == 0 {
		log.Println("empty type")
		data := struct {
			Status           int    `json:"status"`
			DeveloperMessage string `json:"developerMessage"`
			UserMessage      string `json:"userMessage"`
		}{
			Status:           400,
			DeveloperMessage: "createPeriod had empty type parameter, unable to create period.",
			UserMessage:      "Oops, something went wrong, we are unable to save your data right now.",
		}

		log.Println("sending response")
		if err := renderJson(w, data); err != nil {
			log.Println(err)
		}
		return
	} else {
		if strType != "pomodoro" && strType != "rest" {
			log.Println("empty type")
			data := struct {
				Status           int    `json:"status"`
				DeveloperMessage string `json:"developerMessage"`
				UserMessage      string `json:"userMessage"`
			}{
				Status:           400,
				DeveloperMessage: "createPeriod has wrong type, unable to create period.",
				UserMessage:      "Oops, something went wrong, we are unable to save your data right now.",
			}

			log.Println("sending response")
			if err := renderJson(w, data); err != nil {
				log.Println(err)
			}
			return
		}
	}

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
	var p Period
	p = Period{strType}

	if err = collection.Insert(&p); err != nil {
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

	if r.Method != "GET" {
		data := struct {
			Status           int    `json:"status"`
			DeveloperMessage string `json:"developerMessage"`
			UserMessage      string `json:"userMessage"`
		}{
			Status:           400,
			DeveloperMessage: "getPeriods: only GET requests are supported for this route.",
			UserMessage:      "Oops, something went wrong, we are unable to get your data right now.",
		}

		log.Println("sending response")
		if err := renderJson(w, data); err != nil {
			log.Println(err)
		}
		return
	}

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
	collection := session.DB("").C("period")

	var results []Period

	log.Println("start search for all period objects")
	if err = collection.Find(nil).All(&results); err != nil {
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
