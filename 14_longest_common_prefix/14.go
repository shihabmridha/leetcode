package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	res := []byte{}
	mp := map[int]int{}
	ln := len(strs)

	for a, str := range strs {
		mp[a] = len(str) - 1

		if mp[a] < 0 {
			return ""
		}
	}

	exit := false

	t := len(strs[0])

	for i := 0; i < t && !exit; i++ {
		k := strs[0][i]

		for j := 1; j < ln && !exit; j++ {
			v := mp[j]

			if i > v {
				exit = true
				continue
			}

			if i > v {
				exit = true
				continue
			}

			if k != strs[j][i] {
				exit = true
			}
		}

		if !exit {
			res = append(res, k)
		}
	}

	return string(res)
}
