package slice

import (
	"fmt"
)

func Slice() {
	// Example 1: update slice
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

	// Example 2: triple dots as argument / parameter in variadic function
	slice = []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(slice...))

	// Example 3: append all the trailing elements into a slice
	slice2 := []int{1, 7, 3}
	fmt.Println(append(slice, slice2...))
}

// Variadic function parameters
func Sum(array ...int) int {
	var sum int
	sum = 0
	for _, elem := range array {
		sum += elem
	}
	return sum
}