package main

import (
	"io"
	"log"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from flexible-mq"))
	return
}

func helloworld2fisher(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://fisher:8080/hello")
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

func helloworld2account(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://account:8080/hello")
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

func helloworld2market(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://market:8080/hello")
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

func helloworld2metrics(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://metrics:8080/hello")
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

func helloworld2shop(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://shop:8080/hello")
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

func helloworld2trading(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://trading:8080/hello")
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

func helloworld2userTesting(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://user-testing:8080/hello")
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
