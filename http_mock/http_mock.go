package http_mock

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// HttpClient interface
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Item struct {
	Title string `json:"title"`
}

type Response struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

// Post sends a post request to the URL with the body
func Post(client HttpClient, url string, item Item) ([]byte, int, error) {
	// prepare request
	jsonBytes, err := json.Marshal(item)
	if err != nil {
		return []byte{}, 0, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return []byte{}, 0, err
	}
	request.Header.Add("accept", "application/json")
	request.Header.Add("content-type", "application/json")

	// retrieve response
	res, err := client.Do(request)
	if err != nil {
		log.Printf("Error in sending the GET request, err: %v", err)
		return []byte{}, 0, err
	}
	if res.StatusCode != http.StatusOK {
		log.Printf("Request failed with status code: %v", res.StatusCode)
		return []byte{}, res.StatusCode, err
	}
	defer res.Body.Close()

	// reading the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return []byte{}, 0, err
	}

	return body, res.StatusCode, nil
}
