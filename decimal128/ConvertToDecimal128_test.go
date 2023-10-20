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

func CompareDecimal128(d1, d2 primitive.Decimal128) (int, error) {
	b1, exp1, err := d1.BigInt()
	if err != nil {
		return 0, err
	}
	b2, exp2, err := d2.BigInt()
	if err != nil {
		return 0, err
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
