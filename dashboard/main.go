package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public/")))

	//Admin dashboard inputs
	http.HandleFunc("/getFlexible", flexibleMQ)
	http.HandleFunc("/getFisher", fisher)
	http.HandleFunc("/getAccount", account)
	http.HandleFunc("/getMarket", market)
	http.HandleFunc("/getMetrics", metrics)
	http.HandleFunc("/getShop", shop)
	http.HandleFunc("/getTrading", trading)
	http.HandleFunc("/getUserTesting", userTesting)

	//Fisher inputs
	http.HandleFunc("/fish", fish)

	//Account inputs
	http.HandleFunc("/login", login)
	http.HandleFunc("/createAccount", createAccount)

	fmt.Println("Serving on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
