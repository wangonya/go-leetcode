package goleetcode

import "testing"

func test(in []int, target int, out []int, t *testing.T) {
	twosum := twoSum(in, target)
	for i, v := range twosum {
		if v != out[i] {
			t.Fatalf("Input: %v, expected: %v, got: %v", in, out, twosum)
		}
	}
}

func TestTwoSum(t *testing.T) {
	test([]int{2, 7, 11, 15}, 9, []int{0, 1}, t)
	test([]int{3, 2, 4}, 6, []int{1, 2}, t)
	test([]int{3, 3}, 6, []int{0, 1}, t)
}
