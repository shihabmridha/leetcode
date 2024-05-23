package countingbits

func hammingWeight(n int) int {
	res := 0
	for n != 0 {
		if n&1 == 1 {
			res++
		}
		n >>= 1
	}

	return res
}

func countBits(n int) []int {
	res := make([]int, n+1)

	for i := range n + 1 {
		res[i] = hammingWeight(i)
	}

	return res
}
