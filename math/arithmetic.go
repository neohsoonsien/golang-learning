package math

import (
	"fmt"
	"math/big"
)

func main() {
	var a int = 2
	a += 2
	fmt.Println(a)

	a -= 2
	fmt.Println(a)

	b := big.NewInt(12)
	b = b.Add(b, big.NewInt(5))
	fmt.Println(b)

	list := [4]string{"one", "two", "three", "four"}
	for i, item := range list {
		fmt.Println(i, item)
	}
}
