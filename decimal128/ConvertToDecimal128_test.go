package decimal128

import (
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestConvertToDecimal128(t *testing.T) {
	t.Log("TestConvertToDecimal128")

	// valid convert
	res, err := ConvertToDecimal128(100, 500000000)
	if err != nil {
		t.Errorf("ConvertToDecimal128 with error: %v", err)
	}
	value, _ := primitive.ParseDecimal128("100.5")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}

	// valid convert
	res, err = ConvertToDecimal128(-100, -500000000)
	if err != nil {
		t.Errorf("ConvertToDecimal128 with error: %v", err)
	}
	value, _ = primitive.ParseDecimal128("-100.5")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}

	// valid convert
	res, err = ConvertToDecimal128(0, -500000000)
	if err != nil {
		t.Errorf("ConvertToDecimal128 with error: %v", err)
	}
	value, _ = primitive.ParseDecimal128("-0.5")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}

	// sign not match
	_, err = ConvertToDecimal128(-100, 500000000)
	if err == nil {
		t.Errorf("Expect error: %s", "units and nanos sign not match")
	}
	_, err = ConvertToDecimal128(100, -500000000)
	if err == nil {
		t.Errorf("Expect error: %s", "units and nanos sign not match")
	}
}

func TestConvertBigIntToDecimal128(t *testing.T) {
	t.Log("TestConvertBigIntToDecimal128")

	// valid convert
	res, err := ConvertBigIntToDecimal128(big.NewInt(89045))
	if err != nil {
		t.Errorf("ConvertBigIntToDecimal128 with error: %v", err)
	}
	value, _ := primitive.ParseDecimal128("89045")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}

	// valid convert
	res, err = ConvertBigIntToDecimal128(big.NewInt(-89045))
	if err != nil {
		t.Errorf("ConvertBigIntToDecimal128 with error: %v", err)
	}
	value, _ = primitive.ParseDecimal128("-89045")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}
}

func TestConvertDecimal128ToBig(t *testing.T) {
	t.Log("TestConvertDecimal128ToBig")

	// *big.Int - valid convert
	expectedBigInt, _ := ConvertToDecimal128(int64(0), int32(0))
	_, bigInt, err := ConvertDecimal128ToBig(expectedBigInt)
	if err != nil {
		t.Errorf("ConvertDecimal128ToBigInt with error: %v", err)
	}
	if big.NewInt(0).Cmp(bigInt) != 0 {
		t.Errorf("Expected value: %v, Actual: %v, Error: *big.Int values do not match", big.NewInt(15), bigInt)
	}

	// *big.Int - valid convert
	expectedBigInt, _ = primitive.ParseDecimal128("-15")
	_, bigInt, err = ConvertDecimal128ToBig(expectedBigInt)
	if err != nil {
		t.Errorf("ConvertDecimal128ToBigInt with error: %v", err)
	}
	if big.NewInt(-15).Cmp(bigInt) != 0 {
		t.Errorf("Expected value: %v, Actual: %v, Error: *big.Int values do not match", big.NewInt(-15), bigInt)
	}

	// *big.Float - valid convert
	expectedBigFloat, _ := primitive.ParseDecimal128("47.35")
	bigFloat, _, err := ConvertDecimal128ToBig(expectedBigFloat)
	if err != nil {
		t.Errorf("ConvertDecimal128ToBigFloat with error: %v", err)
	}
	if big.NewFloat(47.35).Cmp(bigFloat) != 0 {
		t.Errorf("Expected value: %v, Actual: %v, Error: *big.Float values do not match", big.NewFloat(47.35), bigFloat)
	}

	// *big.Float - valid convert
	expectedBigFloat, _ = primitive.ParseDecimal128("-13.35")
	bigFloat, _, err = ConvertDecimal128ToBig(expectedBigFloat)
	if err != nil {
		t.Errorf("ConvertDecimal128ToBigFloat with error: %v", err)
	}
	if big.NewFloat(-13.35).Cmp(bigFloat) != 0 {
		t.Errorf("Expected value: %v, Actual: %v, Error: *big.Float values do not match", big.NewFloat(47.35), bigFloat)
	}

	// *big.Float - valid convert
	expectedBigFloat, _ = primitive.ParseDecimal128("0.0565")
	bigFloat, _, err = ConvertDecimal128ToBig(expectedBigFloat)
	if err != nil {
		t.Errorf("ConvertDecimal128ToBigFloat with error: %v", err)
	}
	if big.NewFloat(0.0565).Cmp(bigFloat) != 0 {
		t.Errorf("Expected value: %v, Actual: %v, Error: *big.Float values do not match", big.NewFloat(47.35), bigFloat)
	}

	// *big.Float - valid convert
	expectedBigFloat, _ = primitive.ParseDecimal128("-0.0478")
	bigFloat, _, err = ConvertDecimal128ToBig(expectedBigFloat)
	if err != nil {
		t.Errorf("ConvertDecimal128ToBigFloat with error: %v", err)
	}
	if big.NewFloat(-0.0478).Cmp(bigFloat) != 0 {
		t.Errorf("Expected value: %v, Actual: %v, Error: *big.Float values do not match", big.NewFloat(47.35), bigFloat)
	}
}

func TestConvertDecimalToDecimal128(t *testing.T) {
	t.Log("TestConvertDecimalToDecimal128")

	// valid convert
	x := decimal.NewFromFloat(589)
	res, _ := ConvertDecimalToDecimal128(x)
	value, _ := primitive.ParseDecimal128("589")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}

	// valid convert
	x = decimal.NewFromFloat(54.89)
	res, _ = ConvertDecimalToDecimal128(x)
	value, _ = primitive.ParseDecimal128("54.89")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}

	// valid convert
	x = decimal.NewFromFloat(-0.079)
	res, _ = ConvertDecimalToDecimal128(x)
	value, _ = primitive.ParseDecimal128("-0.079")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}

	// valid convert
	x = decimal.NewFromFloat(0.3)
	res, _ = ConvertDecimalToDecimal128(x)
	value, _ = primitive.ParseDecimal128("0.3")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}
}

func TestConvertStringToDecimal128(t *testing.T) {
	t.Log("TestConvertStringToDecimal128")

	// valid convert
	res, err := ConvertStringToDecimal128("34.98745")
	if err != nil {
		t.Errorf("ConvertStringToDecimal128 with error: %v", err)
	}
	value, _ := primitive.ParseDecimal128("34.98745")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}

	// valid convert
	res, err = ConvertStringToDecimal128("-34.98745")
	if err != nil {
		t.Errorf("ConvertStringToDecimal128 with error: %v", err)
	}
	value, _ = primitive.ParseDecimal128("-34.98745")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}

	// valid convert
	res, err = ConvertStringToDecimal128("-0.14")
	if err != nil {
		t.Errorf("ConvertStringToDecimal128 with error: %v", err)
	}
	value, _ = primitive.ParseDecimal128("-0.14")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}

	// valid convert
	res, err = ConvertStringToDecimal128("0.85")
	if err != nil {
		t.Errorf("ConvertStringToDecimal128 with error: %v", err)
	}
	value, _ = primitive.ParseDecimal128("0.85")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}

	// valid convert
	res, err = ConvertStringToDecimal128("17.30")
	if err != nil {
		t.Errorf("ConvertStringToDecimal128 with error: %v", err)
	}
	value, _ = primitive.ParseDecimal128("17.30")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}

	// valid convert
	res, err = ConvertStringToDecimal128("4")
	if err != nil {
		t.Errorf("ConvertStringToDecimal128 with error: %v", err)
	}
	value, _ = primitive.ParseDecimal128("4")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}

	// valid convert
	res, err = ConvertStringToDecimal128("3.0")
	if err != nil {
		t.Errorf("ConvertStringToDecimal128 with error: %v", err)
	}
	value, _ = primitive.ParseDecimal128("3.0")
	if equality, err := CompareDecimal128(res, value); err != nil || equality != 0 {
		t.Errorf("Expected value: %s, Actual: %s, Error: %v", value.String(), res.String(), err)
	}
}

func TestCompareDecimal128(t *testing.T) {
	t.Log("TestCompareDecimal128")

	// compare d1 = d2
	expected := primitive.Decimal128{}
	actual := primitive.Decimal128{}
	if equality, err := CompareDecimal128(expected, actual); err != nil || equality != 0 {
		t.Errorf("Expected value: %v, Actual: %v, Error: Decimal128 values do not match", expected, actual)
	}

	expected, _ = primitive.ParseDecimal128("33")
	actual, _ = primitive.ParseDecimal128("33")
	if equality, err := CompareDecimal128(expected, actual); err != nil || equality != 0 {
		t.Errorf("Expected value: %v, Actual: %v, Error: Expect d1 = d2", expected, actual)
	}

	// compare d1 > d2
	expected, _ = primitive.ParseDecimal128("20")
	actual, _ = primitive.ParseDecimal128("-20")
	if equality, err := CompareDecimal128(expected, actual); err != nil || equality != 1 {
		t.Errorf("Expected value: %v, Actual: %v, Error: Expect d1 > d2", expected, actual)
	}

	expected, _ = primitive.ParseDecimal128("9")
	actual, _ = primitive.ParseDecimal128("8")
	if equality, err := CompareDecimal128(expected, actual); err != nil || equality != 1 {
		t.Errorf("Expected value: %v, Actual: %v, Error: Expect d1 > d2", expected, actual)
	}

	expected, _ = primitive.ParseDecimal128("5")
	actual, _ = primitive.ParseDecimal128("-1")
	if equality, err := CompareDecimal128(expected, actual); err != nil || equality != 1 {
		t.Errorf("Expected value: %v, Actual: %v, Error: Expect d1 > d2", expected, actual)
	}

	expected, _ = primitive.ParseDecimal128("5")
	actual = primitive.Decimal128{}
	if equality, err := CompareDecimal128(expected, actual); err != nil || equality != 1 {
		t.Errorf("Expected value: %v, Actual: %v, Error: Expect d1 > d2", expected, actual)
	}

	// compare d1 < d2
	expected, _ = primitive.ParseDecimal128("-20")
	actual, _ = primitive.ParseDecimal128("20")
	if equality, err := CompareDecimal128(expected, actual); err != nil || equality != -1 {
		t.Errorf("Expected value: %v, Actual: %v, Error: Expect d1 < d2", expected, actual)
	}

	expected, _ = primitive.ParseDecimal128("-3")
	actual, _ = primitive.ParseDecimal128("4")
	if equality, err := CompareDecimal128(expected, actual); err != nil || equality != -1 {
		t.Errorf("Expected value: %v, Actual: %v, Error: Expect d1 < d2", expected, actual)
	}

	expected, _ = primitive.ParseDecimal128("8")
	actual, _ = primitive.ParseDecimal128("9")
	if equality, err := CompareDecimal128(expected, actual); err != nil || equality != -1 {
		t.Errorf("Expected value: %v, Actual: %v, Error: Expect d1 < d2", expected, actual)
	}

	expected = primitive.Decimal128{}
	actual, _ = primitive.ParseDecimal128("5")
	if equality, err := CompareDecimal128(expected, actual); err != nil || equality != -1 {
		t.Errorf("Expected value: %v, Actual: %v, Error: Expect d1 < d2", expected, actual)
	}
}
