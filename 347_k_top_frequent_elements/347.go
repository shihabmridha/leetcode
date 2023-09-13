package ktopfrequentelements

import "sort"

func topKFrequent(nums []int, k int) []int {
	ln := len(nums)
	mp := make(map[int]int)

	for i := 0; i < ln; i++ {
		value, ok := mp[nums[i]]

		if ok {
			mp[nums[i]] = value + 1
		} else {
			mp[nums[i]] = 1
		}
	}

	un := []int{}
	for k := range mp {
		un = append(un, k)
	}

	sort.Slice(un, func(i, j int) bool { return mp[un[i]] > mp[un[j]] })

	return un[:k]
}
