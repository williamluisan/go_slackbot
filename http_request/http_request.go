package http_request

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Get() {
	
}

func Send(content_type, method, endpoint string, body io.Reader) (string, string) {
	var url = os.Getenv("URL") + endpoint
	var client = &http.Client{}
	var bearer = "Bearer " + os.Getenv("BOT_USER_OAUTH_TOKEN")

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
	}

	/* HEADER */
	if content_type != "" {
		req.Header.Set("Content-Type", content_type)
	} else {
		req.Header.Set("Content-Type", "application/json")
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