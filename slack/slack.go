package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"
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

/*
Pretty JSON string
*/
func PrettyString(str string) (string, error) {
    var prettyJSON bytes.Buffer
    if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
        return "", err
    }
    return prettyJSON.String(), nil
}

func AppMention(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	var postBody map[string]interface{}
	err = json.Unmarshal([]byte(string(body)), &postBody)
	if err != nil {
		log.Fatalf("Error happened in JSON unmarshal. Err: %s", err)
	}
	
	/*
	check postBody.type if url_verification or event_callback
	if url_verification, return what they want
	*/
	// ...

	resp := postBody
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
	
	return
}