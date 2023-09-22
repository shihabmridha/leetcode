package productofarrayexceptself

func abs(a int) int {
	if a < 0 {
		a = a * -1
	}

	return a
}

func productExceptSelf(nums []int) []int {
	preMult := []int{}
	sufMult := []int{}
	res := []int{}

	prev := 1
	ln := len(nums)

	for i := 0; i < ln; i++ {
		if nums[i] != 0 {

			preMult = append(preMult, abs(prev*nums[i]))
		}

		if nums[ln-(1+i)] != 0 {
			sufMult = append(sufMult, abs(prev*nums[ln-(1+i)]))
		}
	}

	// for i := 0; i < ln; i++ {
	//     if
	// }

	return []int{}
}

// [1, 1, 0, 3, 3]			[1,  2,  3,  4]

// [1, 1, 0, 0, 0]			[1,  2,  6,  24]
// [0, 0, 0, 9, 3]			[24, 24, 12, 4]

// [0, 0, 9, 0, 0]			[24, 12, 8,  6]
