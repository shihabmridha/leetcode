package findminimuminrotatedsortedarray

func findMin(nums []int) int {
	ln := len(nums)

	low := 0
	high := ln
	mid := low + ((high - low) / 2)

	for low < high {
		curr := nums[mid]
		next := nums[mid+1]
		prev := nums[mid-1]

		if curr < next {

		}

		// if nums[mid]
	}

	return 0
}
