package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloworld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world from market"))
	return
}
