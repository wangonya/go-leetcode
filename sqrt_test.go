package goleetcode

import "testing"

func testSqrt(in int, out int, t *testing.T) {
	result := sqrt(in)
	if result != out {
		t.Fatalf("Input: %v, expected: %v, got: %v", in, out, result)
	}
}

func TestSqrt(t *testing.T) {
	testSqrt(4, 2, t)
	testSqrt(8, 2, t)
	testSqrt(9, 3, t)
}
