package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloworld)
	http.HandleFunc("/helloFisher", helloworld2fisher)
	http.HandleFunc("/helloAccount", helloworld2account)
	http.HandleFunc("/helloMarket", helloworld2market)
	http.HandleFunc("/helloMetrics", helloworld2metrics)
	http.HandleFunc("/helloShop", helloworld2shop)
	http.HandleFunc("/helloTrading", helloworld2trading)
	http.HandleFunc("/helloUserTesting", helloworld2userTesting)

	http.HandleFunc("/fish", fish)

	http.HandleFunc("/login", login)
	http.HandleFunc("/createAccount", createAccount)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
