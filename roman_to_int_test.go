package goleetcode

import "testing"

func testRomanToInt(in string, out int, t *testing.T) {
	result := romanToInt(in)
	if result != out {
		t.Fatalf("Input: %v, expected: %v, got: %v", in, out, result)
	}
}

func TestRomanToInt(t *testing.T) {
	testRomanToInt("III", 3, t)
	testRomanToInt("LVIII", 58, t)
	testRomanToInt("MCMXCIV", 1994, t)
}
