package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/taironas/route"
)

var root = flag.String("root", "app", "file system path")

func init() {
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
}

func main() {
	r := new(route.Router)

	r.HandleFunc("/api/hello", helloWorld)

	r.AddStaticResource(root)

	log.Println("Listening on " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// helloWorld handler returns a json file with a helloworld message.
//
func helloWorld(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string
	}{
		"hello world",
	}

	if err := renderJson(w, data); err != nil {
		log.Println(err)
	}
}

// renderJson renders data to json and writes it to response writer
func renderJson(w http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}
