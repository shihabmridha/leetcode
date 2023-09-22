package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	ln := len(nums)

	preMult := make([]int, ln)
	sufMult := make([]int, ln)
	res := make([]int, ln)

	prev1 := 1
	prev2 := 1

	for i := 0; i < ln; i++ {
		preMult[i] = prev1 * nums[i]
		prev1 = preMult[i]

		sufMult[ln-(1+i)] = prev2 * nums[ln-(1+i)]
		prev2 = sufMult[ln-(1+i)]
	}

	for i := 0; i < ln; i++ {
		a := 1
		b := 1

		if i != 0 {
			a = preMult[i-1]
		}

		if i+1 < ln {
			b = sufMult[i+1]
		}

		res[i] = a * b
	}

	return res
}

// [1, 1, 0, 3, 3]			[1,  2,  3,  4]

// [1, 1, 0, 0, 0]			[1,  2,  6,  24]
// [0, 0, 0, 9, 3]			[24, 24, 12, 4]

// [0, 0, 9, 0, 0]			[24, 12, 8,  6]
