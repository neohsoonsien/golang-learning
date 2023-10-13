package decimal128

import (
	"fmt"

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

func ConvertDecimal128ToString(decimal primitive.Decimal128) string {
	return decimal.String()
}
