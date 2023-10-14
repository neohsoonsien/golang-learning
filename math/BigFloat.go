package math

import (
	"log"
	"math/big"
)

func BigFloat(x string) *big.Float {
	value, _, _ := big.NewFloat(0).Parse(x, 10)
	return value

}

func Float64(val string) float64 {
	value, accuracy := BigFloat(val).Float64()
	log.Printf("The accuracy is %v", accuracy)

	return value
}

func FloatQuotient(x string, y string) *big.Float {
	return big.NewFloat(0).Quo(BigFloat(x), BigFloat(y))
}

func SeparateIntegerDecimal(x string) (int64, int32) {
	integer64, _ := BigFloat(x).Int64()
	integerPart := big.NewFloat(float64(integer64))
	decimalPart := big.NewFloat(0).Sub(BigFloat(x), integerPart)
	decimalPart = big.NewFloat(0).Mul(decimalPart, big.NewFloat(float64(1000000000)))
	decimal64, _ := decimalPart.Int64()
	return integer64, int32(decimal64)
}

func FloatSubtract(x string, y string) *big.Float {
	return big.NewFloat(0).Sub(BigFloat(x), BigFloat(y))
}
