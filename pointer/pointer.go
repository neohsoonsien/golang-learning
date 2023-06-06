package pointer

import (
	"fmt"
)

func getString(param *string) string {
	return *param
}

func Pointer() {
	value := "Hello"
	fmt.Println(getString(&value))
}