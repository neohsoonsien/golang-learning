package pointer

import (
	"fmt"
)

func getString(param *string) string {
	return *param
}

func checkMemoryAddress(memory *string) string {

	if memory == nil {
		return "Empty memory"
	}

	return "Non-empty memory"
}

func String() {
	value := "Hello"
	fmt.Println(getString(&value))

	var key string
	fmt.Println(checkMemoryAddress(&key))
}