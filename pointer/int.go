// Golang program to demonstrate the declaration
// and initialization of pointers
package pointer

import "fmt"

func Int() {

	// taking a normal variable
	var x int = 5748

	// declaration of pointer
	var p *int

	// initialization of pointer
	p = &x

		// displaying the result
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p)
	fmt.Println("Value of x retrieved from p = ", *p)
}
