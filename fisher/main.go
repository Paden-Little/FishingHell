package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloworld)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from fisher!"))
	return
}
