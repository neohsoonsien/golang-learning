package method_struct

import (
	"fmt"
)

type Employee struct {  
    name     string
    salary   int
    currency string
}

// displaySalary() method has Employee as the receiver type
func (e Employee) displaySalary() {  
    fmt.Printf("Salary of %s is %s%d.\n", e.name, e.currency, e.salary)
}

func Method() {  
    employee1 := Employee {
        name:     "Jackie",
        salary:   5000,
        currency: "$",
    }
    employee1.displaySalary() // calling displaySalary() method of Employee type
}