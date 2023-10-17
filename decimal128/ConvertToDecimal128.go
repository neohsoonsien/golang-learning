package decimal128

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertToDecimal128(units int64, nanos int32) (primitive.Decimal128, error) {
	// compare sign to make sure units and nanos sign always match
	if units != 0 && nanos != 0 && (units>>63) != int64(nanos>>31) {
		nan, _ := primitive.ParseDecimal128("NaN")
		return nan, fmt.Errorf("units and nanos sign not match")
	}
	if units < 0 && nanos < 0 {
		nanos *= -1
	}

	// special case handling for value like -0.XXXXXXXXX
	if units == 0 && nanos < 0 {
		return primitive.ParseDecimal128(fmt.Sprintf("-0.%09d", -nanos))
	}
	return primitive.ParseDecimal128(fmt.Sprintf("%d.%09d", units, nanos))
}

func ConvertBigIntToDecimal128(bigInt *big.Int) (primitive.Decimal128, error) {
	decimal, err := ConvertToDecimal128(bigInt.Int64(), int32(0))
	if err != nil {
		return ConvertToDecimal128(int64(0), int32(0)) // return 0 to DB if encounter error
	}
	return decimal, nil
}

func ConvertDecimal128ToBig(decimal primitive.Decimal128) (*big.Float, *big.Int, error) {
	// convert decimal128 to bigInt and its exponential
	bigInt, exp, _ := decimal.BigInt()

	// convert to *big.Float, and multiply with the exponential
	bigFloat := big.NewFloat(float64(bigInt.Int64()))
	expFloat64 := math.Pow10(exp)
	bigFloat = big.NewFloat(0).Mul(bigFloat, big.NewFloat(float64(expFloat64)))

	// truncate the *big.Float to *big.Int
	integer64, _ := bigFloat.Int64()
	bigInt = big.NewInt(integer64)

	// first response is *big.Float, second response is *big.Int
	return bigFloat, bigInt, nil
}

func ConvertDecimalToDecimal128(x decimal.Decimal) (primitive.Decimal128, error) {
	fmt.Printf("*big.Int is %v", x.BigInt())
	fmt.Printf("exp is %v", x.Exponent())

	// these are the fields for decimal
	bigInt := x.BigInt()
	exponent := int64(x.Exponent())

	// these are the fields for Decimal128
	integerPart := bigInt.Int64()
	decimalPart := int32(0)

	multiplier := big.NewInt(10)
	i := int64(1)
	for i < IntAbs(exponent) {
		multiplier = multiplier.Mul(multiplier, big.NewInt(10))
		i += 1
	}

	unsignedBigInt := big.NewInt(0).Abs(bigInt)
	if exponent > 0 {
		bigInt = bigInt.Mul(bigInt, multiplier)
	} else if exponent < 0 {
		if unsignedBigInt.Cmp(multiplier) == 1 {
			modulus, remainder := bigInt.DivMod(unsignedBigInt, multiplier, big.NewInt(1))
			integerPart = big.NewInt(0).Mul(modulus, big.NewInt(int64(bigInt.Sign()))).Int64()
			trailingZeros := big.NewInt(10)
			i := int64(1)
			for i < int64(9)+exponent {
				trailingZeros = trailingZeros.Mul(trailingZeros, big.NewInt(10))
				i += 1
			}
			decimalPart = int32(big.NewInt(0).Mul(remainder, trailingZeros).Int64())
		}
	}

	decimal128, _ := ConvertToDecimal128(integerPart, decimalPart)

	return decimal128, nil
}

func ConvertStringToDecimal128(str string) (primitive.Decimal128, error) {
	var decimal primitive.Decimal128
	var err error
	if !strings.Contains(str, ".") {
		value, _ := strconv.ParseInt(str, 10, 64)
		decimal, err = ConvertToDecimal128(value, int32(0))
		if err != nil {
			return ConvertToDecimal128(int64(0), int32(0)) // return 0 to DB if encounter error
		}
	} else {
		value, _ := strconv.ParseFloat(str, 64)
		integerPart := int64(value)
		decimalPart := value - float64(integerPart)
		decimal, err = ConvertToDecimal128(integerPart, int32(decimalPart*float64(1000000000)))
		if err != nil {
			return ConvertToDecimal128(int64(0), int32(0)) // return 0 to DB if encounter error
		}
	}
	return decimal, nil
}

func ConvertDecimal128ToString(decimal primitive.Decimal128) string {
	return decimal.String()
}

func IntAbs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
