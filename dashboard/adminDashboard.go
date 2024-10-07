package main

import (
	"io"
	"log"
	"net/http"
)

func flexibleMQ(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://flexible-mq:8080/hello")
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(bytes)
	return
}

func fisher(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://flexible-mq:8080/helloFisher")
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(bytes)
	return
}

func account(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://flexible-mq:8080/helloAccount")
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(bytes)
	return
}

func market(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://flexible-mq:8080/helloMarket")
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(bytes)
	return
}

func metrics(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://flexible-mq:8080/helloMetrics")
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(bytes)
	return
}

func shop(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://flexible-mq:8080/helloShop")
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(bytes)
	return
}

func trading(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://flexible-mq:8080/helloTrading")
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(bytes)
	return
}

func userTesting(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://flexible-mq:8080/helloUserTesting")
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(bytes)
	return
}
