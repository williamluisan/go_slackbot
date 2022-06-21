package main

import (
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "os"

	"github.com/joho/godotenv"
	slack "github.com/williamluisan/go_slackbot/slack"
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

	// http.HandleFunc("/", slack.LinkIntegrationTest)
	http.HandleFunc("/event", slack.AppMention)

	port := ":80"
	fmt.Println("Server is running on port " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}