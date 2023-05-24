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
	fmt.Println(sum(slice...))
	fmt.Println(slice[1:3])

	// Example 3: append all the trailing elements into a slice
	slice2 := []int{1, 7, 3}
	fmt.Println(append(slice, slice2...))

	// Example 4: pop and shift functions for the slice
	var value int
	value, shifted := shift(slice...)
	fmt.Printf("The first element is %v, and the slice after shift operation is %v\n", value, shifted)
	value, popped := pop(slice...)
	fmt.Printf("The last element is %v, and the slice after pop operation is %v\n", value, popped)

	// Example 5: Initialize a slice
	var intSlice = make([]int, 10)
	intSlice[0] = 0
	intSlice[9] = 9
	fmt.Printf("The intSlice is %v\n", intSlice)
}

// Variadic function parameters
func sum(array ...int) int {
	var total int
	total = 0
	for _, elem := range array {
		total += elem
	}
	return total
}

func shift(array ...int) (int, []int) {
	first, array := array[0], array[1:]
	return first, array
}

func pop(array ...int) (int, []int) {
	last, array := array[len(array)-1], array[:len(array)-1]
	return last, array
}