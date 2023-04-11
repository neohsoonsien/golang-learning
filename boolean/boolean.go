package boolean

import (
	"fmt"
	"strconv"
)

func Boolean() {
	var str1, str2, str3 string
	str1 = "fish"
	str2 = str1
	str3 = "fruit"
	
	fmt.Println(strconv.FormatBool(str1 == str2))
	fmt.Println(strconv.FormatBool(str1 == str3))
}