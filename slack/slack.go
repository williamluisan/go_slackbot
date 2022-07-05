package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"
	_ "reflect"

	"github.com/williamluisan/go_slackbot/helper"
	"github.com/williamluisan/go_slackbot/http_request"

	crufter "github.com/crufter/nested"
)

func AppMention(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// print json string passed by slack
	fmt.Println(helper.PrettyString(string(body)))
	fmt.Println("")

	var postBody map[string]interface{}
	err = json.Unmarshal([]byte(string(body)), &postBody)
	if err != nil {
		log.Fatalf("Error happened in JSON unmarshal. Err: %s", err)
	}

	request_type := postBody["type"]
	// bot_token := postBody["token"]

	/*
	check postBody.type if url_verification or event_callback
	if url_verification, return what they want
	*/
	if request_type == "url_verification" {
		resp := postBody
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonResp)
	} else {
		event_type, err := crufter.Get(postBody, "event.type")
		if ! err {
			log.Fatalf("Failed using crufter package")
		}

		if event_type == "app_mention" {
			text_mentioned, _ := crufter.Get(postBody, "event.blocks[0].elements[0].elements[1].text")
			if text_mentioned == " hi" {
				fmt.Println('a')
				status, response := http_request.Send("POST", "chat.postMessage", bytes.NewBuffer([]byte(`{
				
				}`)))
				fmt.Println(status + ": " + response)
			}
		}
	}
	
	return
}