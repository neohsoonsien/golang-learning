package http_mock

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

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
