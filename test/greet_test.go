package test

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
    name := "James"
    fmt.Println("The function Hello returns ", Greet(name))
}