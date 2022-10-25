package easy

import "sort"

func TwoSum(nums []int, target int) []int {
	sort.Ints(nums)

	start := 0
	end := len(nums) - 1

	for start < end {
		v1 := nums[start]
		v2 := nums[end]

		if v2 == (v1 + target) {
			return []int{start, end}
		} else if v2 > (v1 + target) {
			end--
		} else {
			start++
		}
	}

	return []int{}

}
