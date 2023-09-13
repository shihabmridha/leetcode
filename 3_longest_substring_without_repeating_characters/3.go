package longestsubstringwithoutrepeatingcharacters

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLongestSubstring(s string) int {
	flags := map[rune]bool{}

	res := 0
	count := 0

	rn := []rune(s)
	ln := len(rn)

	for i := 0; i < ln; i++ {
		c := rn[i]
		if flags[c] {
			res = max(res, count)

			flags = map[rune]bool{}

			count = 0
		} else {
			flags[c] = true
			count++
		}
	}

	res = max(res, count)

	return res
}
