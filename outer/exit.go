package outer

import (
	"fmt"
)

func Exit() {

	word := ""

Exit:
	for word != "DC" {
		for _, i := range "ABCDE" {
			for _, j := range "ABCDE" {
				word = string(i) + string(j)
				fmt.Println(word)
				if word == "DC" {
					continue Exit
				}
			}
		}
	}
}