package pointer

import (
	"fmt"
)

func getString(param *string) string {
	return *param
}

func String() {
	value := "Hello"
	fmt.Println(getString(&value))
}