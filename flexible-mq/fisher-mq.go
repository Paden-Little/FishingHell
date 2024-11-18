package main

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"main.go/protobufs"
	"net/http"
)

func fish(w http.ResponseWriter, r *http.Request) {
	message, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	arguments := protobufs.StartFish{}
	err = proto.Unmarshal(message, &arguments)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://fisher:8080/fish", "", bytes.NewBuffer(message))
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fish := &protobufs.Fish{}
	err = proto.Unmarshal(body, fish)
	if err != nil {
		return
	}

	message, err = proto.Marshal(fish)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
