// https://leetcode.com/problems/majority-element/

package goleetcode

func majorityElement(nums []int) int {
	m := map[int]int{}

	for _, n := range nums {
		m[n] += 1
	}

	max := 0

	for n := range m {
		if m[n] > m[max] {
			max = n
		}
	}

	return max
}

// faster
// func majorityElement(nums []int) int {
// 	m := make(map[int]int)
// 	n := len(nums) / 2
// 	for _, v := range nums {
// 		m[v]++
// 		if m[v] > n {
// 			return v
// 		}
// 	}
// 	return -1
// }
