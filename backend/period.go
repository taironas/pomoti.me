package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
)

type context struct {
	name string
	w    http.ResponseWriter
}

// Periods hold the information of an activity
//
type Period struct {
	Type  string `json:"type"` // pomodoro or rest
	Start time.Time
	End   time.Time
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

func (c context) wrongMethodPOST() {
	developerMessage := fmt.Sprintf("%s only POST requests are supported for this route.", c.name)
	userMessage := "Oops, something went wrong, we are unable to save your data right now."
	response := createStandardResponse(400, developerMessage, userMessage)
	c.w.WriteHeader(http.StatusBadRequest)
	if err := renderJson(c.w, response); err != nil {
		log.Println(err)
	}
}

func (c context) wrongMethodGET() {
	developerMessage := fmt.Sprintf("%v: only GET requests are supported for this route.", c.name)
	userMessage := "Oops, something went wrong, we are unable to get your data right now."
	response := createStandardResponse(400, developerMessage, userMessage)
	c.w.WriteHeader(http.StatusBadRequest)
	if err := renderJson(c.w, response); err != nil {
		log.Println(err)
	}
	return
}

func (c context) emptyParam(name string) {
	developerMessage := fmt.Sprintf("%s, parameter %v is empty, unable to create period.", c.name, name)
	userMessage := "Oops, something went wrong, we are unable to save your data right now."
	response := createStandardResponse(400, developerMessage, userMessage)
	c.w.WriteHeader(http.StatusBadRequest)
	if err := renderJson(c.w, response); err != nil {
		log.Println(err)
	}

}

func (c context) wrongParamValue(name string) {
	developerMessage := fmt.Sprintf("%s, parameter %v has wrong value, unable to create period.", c.name, name)
	userMessage := "Oops, something went wrong, we are unable to save your data right now."
	response := createStandardResponse(400, developerMessage, userMessage)
	c.w.WriteHeader(http.StatusBadRequest)
	if err := renderJson(c.w, response); err != nil {
		log.Println(err)
	}

}

// createPeriod handler lets you create a period obect.
//
func createPeriod(w http.ResponseWriter, r *http.Request) {
	c := context{name: "createPeriod", w: w}

	if r.Method != "POST" {
		c.wrongMethodPOST()
		return
	}

	var strType string
	if strType = r.FormValue("type"); len(strType) == 0 {
		c.emptyParam("type")
		return
	}

	if strType != "pomodoro" && strType != "rest" {
		c.wrongParamValue("type")
		return
	}

	var start time.Time
	var strStart string
	const shortForm = "2006-01-02 15:04:05.000"
	var err error

	if strStart = r.FormValue("start"); len(strStart) == 0 {
		c.emptyParam("start")
		return
	}
	if start, err = time.Parse(shortForm, strStart); err != nil {
		c.wrongParamValue("start")
		log.Printf("%+v\n", err)
		return
	}

	var end time.Time
	var strEnd string
	if strEnd = r.FormValue("end"); len(strEnd) == 0 {
		c.emptyParam("end")
		return
	}
	if end, err = time.Parse(shortForm, strEnd); err != nil {
		c.wrongParamValue("end")
		log.Printf("%+v\n", err)
		return
	}

	uri := getMongoURI()

	log.Println("uri found: ", uri)

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
	p := Period{
		Type:  strType,
		Start: start,
		End:   end,
	}

	if err = collection.Insert(&p); err != nil {
		log.Fatal(err)
	}
	log.Println("insert period entity done")
	sendCreatePeriodResponse(w)
}

func sendCreatePeriodResponse(w http.ResponseWriter) {
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
	c := context{name: "getPeriods"}
	if r.Method != "GET" {
		c.wrongMethodGET()
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
	sendGetPeriodsResponse(w, results)
}

func sendGetPeriodsResponse(w http.ResponseWriter, results Periods) {

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
