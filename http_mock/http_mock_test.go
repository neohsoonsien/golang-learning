package http_mock

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

// MockClient is the mock client
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var (
	// GetDoFunc fetches the mock client's `Do` func
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func TestPost_RealDummyAPI(t *testing.T) {
	// declare and initialize the client
	client := &http.Client{}

	// prepare the request and call the https://dummyjson.com API
	item := Item{
		Title: "Stabilo colour pencil",
	}

	response, status, err := Post(client, "https://dummyjson.com/products/add", item)
	if err != nil {
		t.Errorf("Error in calling Post function")
	}
	if status != 200 {
		t.Errorf("Unhealthy response status")
	}

	// decode the Response
	var result Response
	err = json.Unmarshal(response, &result)
	if err != nil {
		t.Errorf("Error in unmarshalling response")
	}

	t.Logf("The result is %v", result)

	// verify the tests
	assert.Equal(t, result.Id, 101)
	assert.Equal(t, result.Title, "Stabilo colour pencil")
}

func TestPost_MockClient(t *testing.T) {
	// declare and initialize the client
	client := &MockClient{}

	// build response for MockClient
	fakeResponse := Response{
		Id:    202,
		Title: "Fake item",
	}
	jsonBytes, err := json.Marshal(fakeResponse)
	if err != nil {
		t.Errorf("Error in marshalling request")
	}
	// create a new reader with that JSON
	r := io.NopCloser(bytes.NewReader([]byte(jsonBytes)))
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	// prepare the request and call the Mock API
	item := Item{
		Title: "Fake item",
	}
	response, status, err := Post(client, "https://fakeurl.com", item)
	if err != nil {
		t.Errorf("Error in calling Post function")
	}
	if status != 200 {
		t.Errorf("Unhealthy response status")
	}

	// decode the Response
	var result Response
	err = json.Unmarshal(response, &result)
	if err != nil {
		t.Errorf("Error in unmarshalling response")
	}

	t.Logf("The result is %v", result)

	// verify the tests
	assert.Equal(t, result.Id, 202)
	assert.Equal(t, result.Title, "Fake item")
}
