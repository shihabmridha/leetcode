package validparentheses

func isValid(s string) bool {
	mp := map[string]string{"(": ")", "{": "}", "[": "]"}

	ln := len(s)
	if ln < 2 {
		return false
	}

	i := 0
	res := true
	nextToClose := make([]string, ln)

	for _, c := range s {
		curr := string(c)
		close, isOpening := mp[curr]

		if isOpening {
			nextToClose[i] = close
			i = i + 1
			continue
		}

		if i == 0 || curr != nextToClose[i-1] {
			res = false
			break
		}

		nextToClose[i] = ""
		i = i - 1
	}

	// one or more bracket pending to close
	if i > 0 {
		res = false
	}

	return res
}
