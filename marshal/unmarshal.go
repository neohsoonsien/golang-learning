package marshal

import (
	"encoding/json"
	"fmt"
	"os"
)

func Unmarshal() {
	/***********
	Unmarshal : object type / interface ---> []byte
	***********/

	// read in from the json file
	// content in []byte
	content, _ := os.ReadFile("./test.json")

	newUsers := []User{}
	err := json.Unmarshal(content, &newUsers)
	if err != nil {
		panic(err)
	}

	for _, user := range newUsers {
		fmt.Println(user.Name)
		fmt.Println(user.Age)
		fmt.Println(user.Active)
	}
}
