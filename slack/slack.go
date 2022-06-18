package slack

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"log"
	"net/http"
	_"os"
)

func LinkIntegrationTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	var postBody EventSubsTest
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
}