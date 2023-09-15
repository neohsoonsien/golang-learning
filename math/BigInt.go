package math

import (
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
