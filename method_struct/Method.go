package method_struct

import (
	"fmt"
)

type Employee struct {  
    name     string
    salary   int
    currency string
}

type Rectangle struct {
    width, height 	int
}

// methods can be defined for either pointer or value receiver types
// this area method has a receiver type of *Rectangle
func (rect *Rectangle) area() int {
    return rect.width * rect.height
}

// this perimeter mathod has a value receiver
func (rect Rectangle) perimeter() int {
    return 2 * rect.width + 2 * rect.height
}

// displaySalary() method has Employee as the receiver type
func (e Employee) displaySalary() {  
    fmt.Printf("Salary of %s is %s%d.\n", e.name, e.currency, e.salary)
}

func Method() { 
	// Example 1: Employee.displaySalary()
    employee := Employee {
        name:     "Jackie",
        salary:   5000,
        currency: "$",
    }
	// calling displaySalary() method of Employee type
    employee.displaySalary()

	// Example 2: Rectangle.area() & Rectangle.perimeter()
	rectangle := Rectangle { width: 10, height: 5 }
	// Go automatically handles conversion between values and pointers for method calls.
	fmt.Println("Area: ", rectangle.area())
    fmt.Println("Perimeter:", rectangle.perimeter())
	// a pointer receiver type can:
	//    1) avoid copying on method calls 
	//	  2) allow the method to mutate the receiving struct
	rectangleCopy := &rectangle
	fmt.Println("Area: ", rectangleCopy.area())
    fmt.Println("Perimeter:", rectangleCopy.perimeter())
}