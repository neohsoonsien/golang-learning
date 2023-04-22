package flagging

import (
	"flag"
	"fmt"
)

var (
	colour = flag.String("color", "blue", "to define the color")
)

func DefineColor() {

	flag.Parse() // flag.Parse() is important, otherwise the "color" variable could not be re-defined
	fmt.Printf("The color is %v.\n", *colour)
}
