package main

import (
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"main/protobufs"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", helloworld)
	http.HandleFunc("/login", login)
	http.HandleFunc("/createAccount", createAccount)

	var err error
	for true {
		time.Sleep(time.Millisecond * 250)
		err = InitDB()
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
	w.Write([]byte("Hello World from account!"))
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	account := protobufs.Account{}
	err = proto.Unmarshal(raw, &account)
	if err != nil {
		log.Fatal(err)
	}

	Login(&account)

	raw, err = proto.Marshal(&account)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(raw)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	account := protobufs.Account{}
	err = proto.Unmarshal(raw, &account)
	if err != nil {
		log.Fatal(err)
	}

	err = CreateAccount(account.Username, account.Email, account.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(""))
	}
	return
}
