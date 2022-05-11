// https://leetcode.com/problems/plus-one/

package goleetcode

func plusOne(digits []int) []int {
	digitsLen := len(digits)

	for i := digitsLen - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			if i == 0 {
				return append([]int{1}, digits...)
			}
		} else {
			digits[i] += 1
			return digits
		}
	}

	return []int{0}
}
