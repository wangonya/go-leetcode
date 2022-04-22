package goleetcode

import "testing"

func testLongestCommonPrefix(in []string, out string, t *testing.T) {
	result := longestCommonPrefix(in)
	if result != out {
		t.Fatalf("Input: %v, expected: %v, got: %v", in, out, result)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	testLongestCommonPrefix([]string{"flower", "flow", "flight"}, "fl", t)
	testLongestCommonPrefix([]string{"dog", "racecar", "car"}, "", t)
	testLongestCommonPrefix([]string{"ab", "a"}, "a", t)
}
