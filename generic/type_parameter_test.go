package generic

import (
	"testing"
)

var (
	// Initialize a map for the integer values
	ints = map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats = map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
)

func TestSumInts(t *testing.T) {
	t.Log("TestSumInts")

	// actual value
	actual := SumInts(ints)

	// expected value
	expected := int64(46)
	if actual != expected {
		t.Errorf("Expected value: %v, Actual: %v, Error: %s", expected, actual, "SumInts(m map[string]int64) does not tally.")
	}
}

func TestSumFloats(t *testing.T) {
	t.Log("TestSumFloats")

	// actual value
	actual := SumFloats(floats)

	// expected value
	expected := float64(62.97)
	if actual != expected {
		t.Errorf("Expected value: %v, Actual: %v, Error: %s", expected, actual, "SumFloats(m map[string]float64) does not tally.")
	}
}

func TestSumIntsOrFloats(t *testing.T) {
	t.Log("TestSumIntsOrFloats")

	t.Log("calling generic function with type arguments")
	// sum float64
	actualFloat := SumIntsOrFloats[string, float64](floats)
	expectedFloat := float64(62.97)
	if actualFloat != expectedFloat {
		t.Errorf("Expected value: %v, Actual: %v, Error: %s", expectedFloat, actualFloat, "SumIntsOrFloats[key comparable, number int64 | float64](m map[key]number) does not tally.")
	}
	// sum int64
	actualInt := SumIntsOrFloats[string, int64](ints)
	expectedInt := int64(46)
	if actualInt != expectedInt {
		t.Errorf("Expected value: %v, Actual: %v, Error: %s", expectedInt, actualInt, "SumIntsOrFloats[key comparable, number int64 | float64](m map[key]number) does not tally.")
	}

	t.Log("calling generic function without type arguments")
	// sum float64
	actualFloat = SumIntsOrFloats(floats)
	expectedFloat = float64(62.97)
	if actualFloat != expectedFloat {
		t.Errorf("Expected value: %v, Actual: %v, Error: %s", expectedFloat, actualFloat, "SumIntsOrFloats[key comparable, number int64 | float64](m map[key]number) does not tally.")
	}
	// sum int64
	actualInt = SumIntsOrFloats(ints)
	expectedInt = int64(46)
	if actualInt != expectedInt {
		t.Errorf("Expected value: %v, Actual: %v, Error: %s", expectedInt, actualInt, "SumIntsOrFloats[key comparable, number int64 | float64](m map[key]number) does not tally.")
	}
}

func TestSumNumbers(t *testing.T) {
	t.Log("TestSumNumbers")

	// sum float64
	actualFloat := SumIntsOrFloats(floats)
	expectedFloat := float64(62.97)
	if actualFloat != expectedFloat {
		t.Errorf("Expected value: %v, Actual: %v, Error: %s", expectedFloat, actualFloat, "SumIntsOrFloats[key comparable, number Number](m map[key]number) does not tally.")
	}

	// sum int64
	actualInt := SumIntsOrFloats(ints)
	expectedInt := int64(46)
	if actualInt != expectedInt {
		t.Errorf("Expected value: %v, Actual: %v, Error: %s", expectedInt, actualInt, "SumIntsOrFloats[key comparable, number Number](m map[key]number) does not tally.")
	}
}
