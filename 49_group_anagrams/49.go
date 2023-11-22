package groupanagrams

func isAnagram(s map[rune]int, t map[rune]int) bool {
	if len(s) != len(t) {
		return false
	}

	for c := range s {
		if s[c] != t[c] {
			return false
		}
	}

	return true
}

func groupAnagrams(strs []string) [][]string {
	ln := len(strs)

	smp := make([]map[rune]int, ln+1)
	used := make(map[int]int, ln)
	res := [][]string{}

	for i, str := range strs {
		mp := map[rune]int{}

		for _, c := range str {
			value, ok := mp[c]
			if ok {
				mp[c] = value + 1
			} else {
				mp[c] = 1
			}
		}

		smp[i] = mp
	}

	for i := 0; i < ln; i++ {
		if used[i] == -1 {
			continue
		}

		grp := []string{strs[i]}
		for j := i + 1; j < ln; j++ {
			if used[j] == -1 {
				continue
			}

			if isAnagram(smp[i], smp[j]) {
				grp = append(grp, strs[j])
				used[j] = -1
			}
		}

		res = append(res, grp)
		used[i] = -1
	}

	return res
}
