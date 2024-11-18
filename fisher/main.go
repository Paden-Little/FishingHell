package main

import (
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"main/datalayer"
	"main/fisher"
	"main/protobufs"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", helloworld)
	http.HandleFunc("/fish", fish)
	var err error
	for true {
		time.Sleep(time.Millisecond * 250)
		err = datalayer.InitDB()
		if err != nil {
			log.Println(err)
		} else {
			break
		}
	}

	log.Printf("Serving")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from fisher!"))
	return
}

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

	fish := fisher.GetFish(arguments.GetPool(), arguments.GetPoleID(), arguments.GetBaitID())

	datalayer.SaveFish(fish, arguments.UserID)

	resp, err := proto.Marshal(fish)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := w.Write(resp); err != nil {
		log.Fatal(err)
	} else {
		return
	}
}
