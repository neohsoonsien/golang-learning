package marshal

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func Marshal() {
	/***********
	Marshal : object type / interface ---> []byte
	***********/
	// define array to collect the json object
	array := []User{}
	array = append(array, User{"John", 29, true})
	array = append(array, User{"James", 32, true})

	// encode any type into []byte
	value, err := json.Marshal(array)
	if err != nil {
		panic(err)
	}
	// stringify []byte to string
	fmt.Println(string(value))

	// output to json file
	_ = os.WriteFile("test.json", value, 0644)
}
