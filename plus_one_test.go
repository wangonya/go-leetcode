package goleetcode

import "testing"

func testPlusOne(in []int, out []int, t *testing.T) {
	res := plusOne(in)

	for i, n := range out {
		if res[i] != n {
			failMessage(in, out, res, t)
		}
	}
}

func TestPlusOne(t *testing.T) {
	testPlusOne([]int{1,2,3},[]int{1,2,4},t)
	testPlusOne([]int{4,3,2,1}, []int{4,3,2,2},t)
	testPlusOne([]int{9},[]int{1,0},t)
	testPlusOne([]int{9,9},[]int{1,0,0},t)
	testPlusOne([]int{8,9},[]int{9,0},t)
}
