package maximumproductofthreenumbers

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	ln := len(nums)

	na := nums[ln-1] * nums[ln-2] * nums[ln-3]
	bro := nums[0] * nums[1] * nums[ln-1]

	if na > bro {
		return na
	}

	return bro
}
