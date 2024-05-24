package searchinsertposition

func searchInsert(nums []int, target int) int {
	up := len(nums) - 1
	low := 0
	mid := 0

	for low <= up {
		mid = (up + low) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			up = mid - 1
		}
	}

	if (mid == 0 && nums[mid] > target) || (nums[mid] > target && nums[mid-1] < target) {
		return mid
	}

	if nums[mid] > target {
		return mid - 1
	}

	return mid + 1
}
