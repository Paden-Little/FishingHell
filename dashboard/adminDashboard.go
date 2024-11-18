package main

import (
	"bytes"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"main/protobufs"
	"net/http"
	"strconv"
)

var secret = []byte("6664dc12921e8c5874975ae9358eb431bb86a3881594ff17f64e276e9ddaa9e83153537c0f64514b940e2d3a533793230cbfe3c9ed03aeb28aabffb0fde64e9d14fc4bf11f94a91d8ca4f9fafd0b50593347b1bf54a62033687a9477b1846a0bdd5a61c3dbbbc466f7871673e7b61b66c2e2a1f3d5fbb7c5f03bee5b2fcc76e421d0f8c0a8218985ff1002cf924cb65093f1fca1727dd1b6734fd65c305cd6182a944a57381ecde9f9f0242b546500bd2d60eb7200cb568c9d89ce5214ccbae3f363cacb57cb814cdbf0cce7b9d4d0771e54564d4f16f38f14ecb388215ae673bcc37b16efbcdee55de09089037817f48b115be390e4f98ac41770cd5e6314d2")

func fish(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	poleID, err := strconv.Atoi(r.FormValue("rod"))
	if err != nil {
		log.Fatal(err)
	}

	baitID, err := strconv.Atoi(r.FormValue("bait"))
	if err != nil {
		log.Fatal(err)
	}

	pb := protobufs.StartFish{
		Pool:   r.FormValue("pool"),
		PoleID: int32(poleID),
		BaitID: int32(baitID),
	}

	message, err := proto.Marshal(&pb)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://flexible-mq:8080/fish", "", bytes.NewBuffer(message))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fish := &protobufs.Fish{}
	err = proto.Unmarshal(bytes, fish)

	w.Write([]byte(fish.String()))
}

func login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	account := &protobufs.Account{
		Username: username,
		Password: password,
	}

	message, err := proto.Marshal(account)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://flexible-mq:8080/login", "", bytes.NewBuffer(message))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":     username,
		"password": password,
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(bytes.NewBuffer([]byte(tokenString)).Bytes())
	if err != nil {
		return
	}
}

//func parse(w http.ResponseWriter, r *http.Request) {
//	if err := r.ParseForm(); err != nil {
//		log.Fatal(err)
//	}
//
//	token := r.FormValue("key")
//
//	parsedToken, err := jwt.ParseWithClaims(token)
//}

func createAccount(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	account := &protobufs.Account{
		Username: username,
		Password: password,
	}

	message, err := proto.Marshal(account)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://flexible-mq:8080/createAccount", "", bytes.NewBuffer(message))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	w.Write(bytes.NewBuffer([]byte("Successfully Created Account")).Bytes())
}

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
