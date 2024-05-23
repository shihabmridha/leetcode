package numberof1bits

import (
	"math/bits"
)

func hammingWeight(n int) int {
	x := bits.Len32(uint32(n))
	count := 0

	for i := range x {
		if (n & (1 << i)) != 0 {
			count += 1
		}
	}

	return count
}

// Only check the last bit via & operator and then shift to the right to check then next bit
// Once everything is shifte to right n becomes 0
func BetterVersion(n int) int {
	res := 0
	for n != 0 {
		if n&1 == 1 {
			res++
		}
		n >>= 1
	}

	return res
}
