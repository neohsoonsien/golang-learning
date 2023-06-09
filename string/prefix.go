package string

import (
	"fmt"
	"strings"
)

func Prefix(originalString string, prefix string) {

	hasPrefix := strings.HasPrefix(originalString, prefix)

	if hasPrefix {
		fmt.Printf("The \"%v\" is the prefix of \"%v\". \n", prefix, originalString)
	} else {
		fmt.Printf("The \"%v\" is not a prefix of \"%v\". \n", prefix, originalString)
	}
}