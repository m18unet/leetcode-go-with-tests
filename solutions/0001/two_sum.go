package two_sum

func twoSum(nums []int, target int) []int {
	records := make(map[int]int)

	for i, v := range nums {
		complement := target - v

		if res, ok := records[complement]; ok {
			return []int{res, i}
		}

		records[v] = i
	}

	return []int{}
}
