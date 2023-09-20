package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	MONGO_DB_USER = "root"
	MONGO_DB_PASS = "example"
	MONGO_DB_NAME = "e-commerce"
)

func main() {
	http.HandleFunc("/products", requestHandler)
	http.ListenAndServe(":8080", nil)
}

func requestHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var respMessage map[string]interface{}
	var respErr error

	ctx := context.Background()

	isCached, productsCache, err := getFromCache(ctx)

	if err != nil {

		respErr = err

	} else {

		if isCached == true {

			respMessage = productsCache
			respMessage["_source"] = "Redis Cache"

		} else {

			respMessage, err = getFromDb(ctx)

			if err != nil {
				respErr = err
			}

			err = addToCache(ctx, respMessage)

			if err != nil {
				respErr = err
			}

			respMessage["_source"] = "MongoDB database"
		}
	}

	if respErr != nil {

		fmt.Fprintf(w, respErr.Error())

	} else {

		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")

		if err := enc.Encode(respMessage); err != nil {
			fmt.Fprintf(w, err.Error())
		}

	}
}
