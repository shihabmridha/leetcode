package solution

func RomanToInt(s string) int {
	m := make(map[string]int)

	m["O"] = 0
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000

	res := 0
	chars := []rune(s)
	len := len(chars)

	for i := 0; i < len; {
		char := string(chars[i])
		nextChar := "O"

		if i <= len-2 {
			nextChar = string(chars[i+1])
		}

		val := m[char]
		nextVal := m[nextChar]

		if val < nextVal && (val*5 == nextVal || val*10 == nextVal) {
			val = nextVal - val
			i++
		}

		res += val
		i++
	}

	return res
}
