package findtheindexofthefirstoccurrenceinastring

func strStr(haystack string, needle string) int {
	hln := len(haystack)
	nln := len(needle)

	i := 0
	for i < hln {
		idx := i
		for j := 0; j < nln && i < hln; j++ {
			if haystack[i] != needle[j] {
				break
			}
			i++
		}

		if i-idx == nln {
			return idx
		}

		i = idx + 1
	}

	return -1
}
