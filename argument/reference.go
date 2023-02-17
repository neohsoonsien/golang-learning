package argument

import (
	"fmt"
)

func Reference(a *int) {
	*a = 50
	*a += 60
}

func Output() {
	var value int
	value = 10
	Reference(&value)
	fmt.Printf("The value after passing the parameter as reference: %v", value)
}