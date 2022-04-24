package goleetcode

import "testing"

func testValidParentheses(in string, out bool, t *testing.T) {
	res := validParentheses(in)

	if res != out {
		failMessage(in, out, res, t)
	}
}

func TestValidParentheses(t *testing.T){
	testValidParentheses("()", true, t)
	testValidParentheses("()[]{}", true, t)
	testValidParentheses("(]", false, t)
	testValidParentheses("{[]}", true, t)
}
