package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_"os"

	"github.com/joho/godotenv"
	request "github.com/williamluisan/go_slackbot/http_request"
)

func loadConfig() {
	// load configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// load configuration
	loadConfig()

	// _, body := request.Send(http.MethodGet, "conversations.list", "")
	// log.Fatal(body)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		
		var postBody request.EventSubsTest
		err = json.Unmarshal([]byte(string(body)), &postBody)
		if err != nil {
			log.Fatalf("Error happened in JSON unmarshal. Err: %s", err)
		}
		
		resp := postBody.Challenge
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	})

	port := ":80"
	fmt.Println("Server is running on port " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}