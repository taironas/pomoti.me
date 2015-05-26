package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/taironas/route"
)

var (
	prod       = flag.Bool("prod", false, "determines production.")
	dart       = flag.Bool("dart", true, "determines if dart version should be run.")
	typescript = flag.Bool("typescript", false, "determines if typescript version should be run.")

	//typescript version
	typescriptRoot = flag.String("typescriptRoot", "app", "type script root.")

	// dart version
	// dev: use chromium
	// prod: use dart2js
	dartProductionRoot = flag.String("dartProductionRoot", "app-dart/build/web", "dart prod root.")
	dartDeveloperWeb   = flag.String("dartDeveloperWeb", "app-dart/web", "dart dev web path.")
	dartDeveloperApp   = flag.String("dartDeveloperApp", "app-dart/", "dart dev app path.")
)

func init() {
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
}

func main() {

	flag.Parse()

	r := new(route.Router)

	r.HandleFunc("/api/hello", helloWorld)

	setStaticResources(r)

	log.Println("Listening on " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func setStaticResources(r *route.Router) {

	if *typescript {
		r.AddStaticResource(typescriptRoot)
	} else if *dart {
		if *prod {
			r.AddStaticResource(dartProductionRoot)
		} else {
			r.AddStaticResource(dartDeveloperWeb)
			r.AddStaticResource(dartDeveloperApp)
		}
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
