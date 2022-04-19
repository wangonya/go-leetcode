package goleetcode

func twoSum(nums []int, target int) [2]int {
	result := [2]int{}
	m := make(map[int]int)

	for i, num := range nums {
		remainder := target - num
		n, exists := m[remainder]

		if exists {
			result = [2]int{n,i}
			break
		}

		m[num] = i
	}
	
	return result
}
