package http_request

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Send(method, endpoint, body string) (string, string) {
	var url = os.Getenv("URL") + endpoint
	var client = &http.Client{}
	var bearer = "Bearer " + os.Getenv("OAUTH_TOKEN")
	
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", bearer)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error when sending request to the server")
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	return resp.Status, string(responseBody)
}