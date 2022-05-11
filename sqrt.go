// https://leetcode.com/problems/sqrtx/

package goleetcode

func sqrt(x int) int {
	// repeated subtraction method
	count := 0

	for odds := 1; x > 0; odds += 2 {
		x -= odds
		if x >= 0 {
			count += 1
		}
	}

	return count
}
