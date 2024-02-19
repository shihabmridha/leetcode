package laststoneweight

import (
	"slices"
)

func lastStoneWeight(stones []int) int {
	ln := len(stones)

	if ln == 1 {
		return stones[0]
	}

	slices.Sort(stones)

	for stones[ln-2] != 0 {
		stones[ln-1] = stones[ln-1] - stones[ln-2]
		stones[ln-2] = 0

		slices.Sort(stones)
	}

	return stones[ln-1]
}
