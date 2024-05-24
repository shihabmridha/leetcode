package twosum

import "slices"

func twoSum(nums []int, target int) []int {
	mp := make(map[int][2]int)
	for i, v := range nums {
		ar, set := mp[v]
		if set {
			ar[1] = i
		} else {
			ar[0] = i
		}

		mp[v] = ar
	}

	res := []int{}

	slices.Sort(nums)
	ln := len(nums)

	for i := range ln {
		a := nums[i]
		b := target - a
		_, found := slices.BinarySearch(nums, b)

		if found {
			res = append(res, mp[a][0])

			if a != b {
				res = append(res, mp[b][0])
			} else {
				res = append(res, mp[b][1])
			}

			break
		}
	}

	return res
}
