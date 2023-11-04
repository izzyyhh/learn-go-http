package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ChuckNorrisApiResponse struct {
	Value string `json:"value"`
}

const RANDOM_CHUCK_NORRIS_API_URL = "https://api.chucknorris.io/jokes/random"

var client = &http.Client{}

func main() {
	body, _ := getBodyFromGetRequest()

	// from json to struct
	var chuckNorrisResponse ChuckNorrisApiResponse
	err := json.Unmarshal(body, &chuckNorrisResponse)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(chuckNorrisResponse.Value)

	// from struct to json
	marshalled, _ := json.Marshal(chuckNorrisResponse)
	fmt.Println(string(marshalled))

}

func getBodyFromGetRequest() ([]byte, error) {
	resp, err := client.Get(RANDOM_CHUCK_NORRIS_API_URL)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return body, nil
}
