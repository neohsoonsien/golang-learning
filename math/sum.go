package math

import "fmt"

func Sum(number int) int {

	// using range in for-loop iteration
	fruits := []string{"apple", "orange", "banana"}
	for i, fruit := range fruits {
		fmt.Println(fruit, i)
	}

	// using map in for-loop iteration
	students := map[string]string{
		"I24": "John",
		"I54": "Antoinne",
	}
	for key, value := range students {
		fmt.Println(key, value)
	}

	// using the basic for-loop in summation
	var sum int
	for i := 0; i < number; i++ {
		sum += i
	}
	return sum
}
