package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	flag := map[rune]bool{}
	max := 0
	count := 0

	for _, c := range s {
		if flag[c] {
			if count > max {
				max = count
			}

			count = 0
		} else {
			count++
		}
	}

	return count
}
