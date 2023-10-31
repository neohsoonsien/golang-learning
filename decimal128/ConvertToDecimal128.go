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

func ConvertDecimal128ToBig(decimal128 primitive.Decimal128) (*big.Float, *big.Int, error) {
	// convert decimal128 to bigInt and its exponential
	bigInt, exp, _ := decimal128.BigInt()

	// convert to decimal, then operate to *big.Float
	value := decimal.NewFromBigInt(bigInt, int32(exp))
	valueFloat64, _ := value.Float64()
	bigFloat := big.NewFloat(valueFloat64)

	// truncate the integer part to *big.Int
	valueInt64 := value.IntPart()
	bigInt = big.NewInt(valueInt64)

	// first response is *big.Float, second response is *big.Int
	return bigFloat, bigInt, nil
}

func ConvertDecimalToDecimal128(x decimal.Decimal) (primitive.Decimal128, error) {
	str := x.String()

	decimal128, _ := ConvertStringToDecimal128(str)

	return decimal128, nil
}
func ConvertStringToDecimal128(str string) (primitive.Decimal128, error) {
	var x primitive.Decimal128
	var err error
	if !strings.Contains(str, ".") {
		value, _ := strconv.ParseInt(str, 10, 64)
		x, err = ConvertToDecimal128(value, int32(0))
		if err != nil {
			return ConvertToDecimal128(int64(0), int32(0)) // return 0 to DB if encounter error
		}
	} else {
		value, _ := decimal.NewFromString(str)
		integerPart := decimal.NewFromInt(value.IntPart())
		decimalPart := value.Sub(integerPart)
		decimalPart = decimalPart.Mul(decimal.NewFromInt32(1000000000))
		x, err = ConvertToDecimal128(value.IntPart(), int32(decimalPart.IntPart()))
		if err != nil {
			return ConvertToDecimal128(int64(0), int32(0)) // return 0 to DB if encounter error
		}
	}
	return x, nil
}

func CompareDecimal128(d1, d2 primitive.Decimal128) (int, error) {
	b1, exp1, err := d1.BigInt()
	if err != nil {
		return int(math.NaN()), err
	}
	b2, exp2, err := d2.BigInt()
	if err != nil {
		return int(math.NaN()), err
	}

	// compare sign
	sign := b1.Sign()
	if sign != b2.Sign() {
		if b1.Sign() > 0 {
			return 1, nil
		} else {
			return -1, nil
		}
	}

	// adjust for both side to match same digits length
	len1, len2 := len(b1.String()), len(b2.String())
	if len1 < len2 {
		exp1 -= len2 - len1
		expDiff := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(len2-len1)), nil)
		b1 = big.NewInt(0).Mul(b1, expDiff)
	} else if len1 > len2 {
		exp2 -= len1 - len2
		expDiff := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(len1-len2)), nil)
		b2 = big.NewInt(0).Mul(b2, expDiff)
	}

	// compare digits if same exp
	if exp1 == exp2 {
		return b1.Cmp(b2), nil
	}

	// compare exp
	if sign < 0 {
		if exp1 < exp2 {
			return 1, nil
		}
		return -1, nil
	} else {
		if exp1 < exp2 {
			return -1, nil
		}

		return 1, nil
	}
}

func IntAbs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
