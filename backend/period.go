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

// Periods holds an array of Period type.
//
type Periods []Period

type standardResponse struct {
	Status           int    `json:"status"`
	DeveloperMessage string `json:"developerMessage"`
	UserMessage      string `json:"userMessage"`
}

// createStandardResponse creates a standard response object with respect to the
// status, the developer message and the user message.
//
func createStandardResponse(status int, devMsg, userMsg string) standardResponse {
	return standardResponse{
		Status:           status,
		DeveloperMessage: devMsg,
		UserMessage:      userMsg,
	}
}

// createPeriod handler lets you create a period obect.
//
func createPeriod(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		developerMessage := "createPeriod only POST requests are supported for this route."
		userMessage := "Oops, something went wrong, we are unable to save your data right now."
		response := createStandardResponse(400, developerMessage, userMessage)
		w.WriteHeader(http.StatusBadRequest)
		log.Println("sending response")

		if err := renderJson(w, response); err != nil {
			log.Println(err)
		}
		return
	}

	var strType string

	if strType = r.FormValue("type"); len(strType) == 0 {
		developerMessage := "createPeriod had empty type parameter, unable to create period."
		userMessage := "Oops, something went wrong, we are unable to save your data right now."
		response := createStandardResponse(400, developerMessage, userMessage)
		w.WriteHeader(http.StatusBadRequest)
		log.Println("sending response")
		if err := renderJson(w, response); err != nil {
			log.Println(err)
		}
		return
	} else {
		if strType != "pomodoro" && strType != "rest" {
			developerMessage := "createPeriod has wrong type, unable to create period."
			userMessage := "Oops, something went wrong, we are unable to save your data right now."
			response := createStandardResponse(400, developerMessage, userMessage)
			w.WriteHeader(http.StatusBadRequest)
			log.Println("sending response")
			if err := renderJson(w, response); err != nil {
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
		developerMessage := "getPeriods: only GET requests are supported for this route."
		userMessage := "Oops, something went wrong, we are unable to get your data right now."
		response := createStandardResponse(400, developerMessage, userMessage)
		w.WriteHeader(http.StatusBadRequest)
		log.Println("sending response")
		if err := renderJson(w, response); err != nil {
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
