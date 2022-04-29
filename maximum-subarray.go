// https://leetcode.com/problems/maximum-subarray/

package goleetcode

import "math"

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	maxSoFar, maxEndingHere, maxElement := int(math.Inf(-1)), 0, int(math.Inf(-1))

	for _, n := range nums {
		maxEndingHere = max(maxEndingHere+n, 0)
		maxSoFar = max(maxEndingHere, maxSoFar)
		maxElement = max(maxElement, n)
	}

	if maxSoFar == 0 {
		maxSoFar = maxElement
	}

	return maxSoFar
}

// faster solution
// func maxSubArray(nums []int) int {
//     bestSum := nums[0]
//     curSum := nums[0]
    
//     for idx, val := range nums {
//         if idx == 0 {
//             continue
//         }
        
//         if curSum < 0  {
//             if curSum < val {
//                 curSum = val
//             }
//         } else {
//             curSum += val
//         }
        
//         if bestSum < curSum {
//             bestSum = curSum
//         }
//     }
    
//     return bestSum
// }

// less memory
// func maxSubArray(nums []int) int {
// 	if len(nums) < 1 {
// 		return 0
// 	}
// 	max, curr := nums[0], nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		curr = int(math.Max(float64(nums[i]), float64(curr+nums[i])))
// 		max = int(math.Max(float64(max), float64(curr)))
// 	}
// 	return max
// }
