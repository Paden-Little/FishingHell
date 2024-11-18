package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve files from the publish directory
	fs := http.FileServer(http.Dir("./FishingHellFrontend/FishingHellFrontend.Browser/bin/Release/net8.0-browser/publish/"))
	http.Handle("/", fs)

	// Start the server
	log.Println("Serving on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
