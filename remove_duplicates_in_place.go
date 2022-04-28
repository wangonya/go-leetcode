// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

package goleetcode

func removeDuplicates(nums []int) int {
	// https://go.dev/play/p/ewM6YevmPD
	for i:=len(nums)-1;i>0;i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i],nums[i+1:]...)
		}
	}
	return len(nums)
}

// faster
func removeDuplicates2(a []int) int {
    // 1 1 2 2 3 4 4
    // j   i
    j := 0
    for i := 1; i < len(a); i++ {
        if a[i] != a[j] {
            j++
            a[j] = a[i]
        }
    }
    return j+1
}
