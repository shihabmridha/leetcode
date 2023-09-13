package contains_duplicate

import "sort"

func ContainsDuplicate(nums []int) bool {
	sort.Ints(nums)
	len := len(nums)

	for i := 0; i < len-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}
