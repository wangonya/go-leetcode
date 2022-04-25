package goleetcode

import "testing"

func testRemoveDuplicates(in []int, out int, t *testing.T){
	res := removeDuplicates(in)

	if res != out {
		failMessage(in,out,res,t)
	}
}

func TestRemoveDuplicates(t *testing.T){
	testRemoveDuplicates([]int{1,1,2},2,t)
	testRemoveDuplicates([]int{0,0,1,1,1,2,2,3,3,4},5,t)
}
