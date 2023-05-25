package array

import (
	"fmt"
)

func Array() {

	/////////////////////////////////////////////////
	// Example 1: Assign and access value
	var country [3]string
	country[0] = "India"  // Assign a value to the first element
	country[1] = "Canada" // Assign a value to the second element
	country[2] = "Japan"  // Assign a value to the third element

	fmt.Println(country[0]) // Access the first element value
	fmt.Println(country[1]) // Access the second element value
	fmt.Println(country[2]) // Access the third element value

	/////////////////////////////////////////////////
	// Example 2: Initializing an array with array literal
	x := [5]int{10, 20, 30, 40, 50}   // Intialized with values
	var y [5]int = [5]int{10, 20, 30} // Partial assignment

	fmt.Println(x)
	fmt.Println(y)
}