package slice

import (
	"fmt"
)

func Slice() {
	// entries cannot be updated in slice iteration
	slice := []int{1, 1, 1}
	for _, entry := range slice {
		entry += 1
	}
	fmt.Println(slice)

	// entries have to be updated like "array" style
	for index:= range slice {
		slice[index] = slice[index] + 1
	}
	fmt.Println(slice)
}