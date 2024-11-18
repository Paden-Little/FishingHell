package main

import (
	"bytes"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"main.go/protobufs"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://account:8080/login", "", bytes.NewBuffer(raw))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	raw, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	account := protobufs.Account{}
	if err := proto.Unmarshal(raw, &account); err != nil {
		log.Fatal(err)
	}

	w.Write(raw)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://account:8080/createAccount", "", bytes.NewBuffer(raw))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	raw, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(raw)
}
