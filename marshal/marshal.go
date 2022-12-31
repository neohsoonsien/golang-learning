package main

import (
    "fmt"
    "encoding/json"
	"io/ioutil"
)

type User struct {
    Name 		string	`json:"name"`
    Age 		int		`json:"age"`
    Active 		bool	`json:"active"`
}

func main() {
	// encode any type into a byte slice
	value, err := json.Marshal(User{"John", 29, true})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(value))

	// output to json file
	_ = ioutil.WriteFile("test.json", value, 0644)
}