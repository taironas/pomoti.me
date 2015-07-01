package main

import (
	"log"
	"net/http"
)

type Period struct {
	Type string
}

// mongo handler returns a json file with a mongo message.
//
func mongo(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string
	}{
		"mongo",
	}

	if err := renderJson(w, data); err != nil {
		log.Println(err)
	}
}
