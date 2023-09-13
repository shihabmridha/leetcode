package valid_anagram

func isAnagram(s string, t string) bool {
	sln := len(s)
	if sln != len(t) {
		return false
	}

	sm := make(map[rune]int)
	tm := make(map[rune]int)

	for _, c := range s {
		value, ok := sm[c]
		if ok {
			sm[c] = value + 1
		} else {
			sm[c] = 1
		}
	}

	for _, c := range t {
		value, ok := tm[c]
		if ok {
			tm[c] = value + 1
		} else {
			tm[c] = 1
		}
	}

	for _, c := range s {
		if sm[c] != tm[c] {
			return false
		}
	}

	return true
}
