package main

import "os"

// getMongoURI reads the MONGOLAB_URI if app is running in production,
// else returns localhost.
//
func getMongoURI() (uri string) {

	if isAppRunningInProduction() {
		uri = os.Getenv("MONGOLAB_URI")
	}

	if len(uri) == 0 {
		uri = "127.0.0.1"
	}
	return
}

func isAppRunningInProduction() bool {
	return *prod
}
