package goleetcode

import "testing"

func failMessage(input interface{}, expected interface{}, got interface{}, t *testing.T) {
	t.Fatalf("Input: %v, expected: %v, got: %v", input, expected, got)
}
