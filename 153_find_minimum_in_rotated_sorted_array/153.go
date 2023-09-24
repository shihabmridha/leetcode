package findminimuminrotatedsortedarray

func findMin(nums []int) int {
	ln := len(nums)

	low := 0
	high := ln - 1

	res := nums[0]

	lowVal := nums[low]
	highVal := nums[high]

	for low < high {
		// everything is sorten in this range
		if lowVal < highVal {
			return lowVal
		}

		mid := low + ((high - low) / 2)

		curr := nums[mid]
		next := nums[mid+1]
		if mid+1 <= ln-1 {
			next = nums[mid+1]
		}

		prev := nums[mid]
		if mid-1 >= 0 {
			prev = nums[mid-1]
		}

		if curr > next {
			res = next
			break
		}

		if curr < prev {
			res = curr
			break
		}

		if curr < lowVal {
			// break should be on the left so go left
			high = mid - 1
			highVal = nums[high]
		}

		if curr > lowVal {
			// break should be on the right so go right
			low = mid + 1
			lowVal = nums[low]
		}
	}

	return res
}
