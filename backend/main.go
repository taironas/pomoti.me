package main

import (
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

	r.AddStaticResource(root)

	log.Println("Listening on " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
