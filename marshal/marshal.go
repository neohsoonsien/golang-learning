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

type Data struct {
	Environment	string	`json:"environment"`
	Value		[]User	`json:"value"`
}

func main() {
	// define array to collect the json object
	array := []User{}
	array = append(array, User{"John", 29, true})
	array = append(array, User{"James", 32, true})

	// encode any type into a byte slice
	value, err := json.Marshal(array)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(value))

	// output to json file
	_ = ioutil.WriteFile("test.json", value, 0644)

	// read in from the json file
	content, _ := ioutil.ReadFile("./test.json")

	newArray := []*User{}
	err = json.Unmarshal(content, &newArray)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(newArray); i++ {
		fmt.Println(newArray[i].Name)
		fmt.Println(newArray[i].Age)
		fmt.Println(newArray[i].Active)
	}
}