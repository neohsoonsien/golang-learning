package errors_handler

import (
	"errors"
	"log"
	"math"
)

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}

func Handler() {
	// test case 1: negative value
	result, err := Sqrt(-1)
	if err != nil {
		log.Print(err)
	} else {
		log.Print(result)
	}

	// test case 2: positive value
	result, err = Sqrt(9)
	if err != nil {
		log.Print(err)
	} else {
		log.Print(result)
	}
}
