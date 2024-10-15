package main

import (
	"log"
	"main/datalayer"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", helloworld)
	time.Sleep(3 * time.Second)
	datalayer.TestDBConnection()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from fisher!"))
	return
}
