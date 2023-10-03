package math

import (
	"log"
	"math/big"
)

func BigInt(val string) *big.Int {

	bigInt := new(big.Int)
	bigInt, ok := bigInt.SetString(val, 10)
	if !ok {
		return big.NewInt(0)
	}
	return bigInt
}

func AddBigInt(a *big.Int, b *big.Int) *big.Int {
	return a.Add(a, b)
}

func BigFloat(val string) *big.Float {
	return new(big.Float).SetInt(BigInt(val))
}

func Float64(val string) float64 {
	value, accuracy := BigFloat(val).Float64()
	log.Printf("The accuracy is %v", accuracy)

	return value
}

func FloatQuotient(x string, y string) *big.Float {
	return big.NewFloat(0).Quo(BigFloat(x), BigFloat(y))
}

func FloatSubtract(x string, y string) *big.Float {
	return big.NewFloat(0).Sub(BigFloat(x), BigFloat(y))
}
