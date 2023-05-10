package solution

func IntToRoman(num int) string {
	roman := ""

	for num > 0 {
		count := 0
		char := ""

		if num >= 1000 {
			count = num / 1000
			char = "M"
			num = num % 1000
		} else if num >= 900 {
			count = 1
			char = "CM"
			num -= 900
		} else if num >= 500 {
			count = num / 500
			char = "D"
			num = num % 500
		} else if num >= 400 {
			count = 1
			char = "CD"
			num -= 400
		} else if num >= 100 {
			count = num / 100
			char = "C"
			num = num % 100
		} else if num >= 90 {
			count = 1
			char = "XC"
			num -= 90
		} else if num >= 50 {
			count = num / 50
			char = "L"
			num = num % 50
		} else if num >= 40 {
			count = 1
			char = "XL"
			num -= 40
		} else if num >= 10 {
			count = num / 10
			char = "X"
			num = num % 10
		} else if num >= 9 {
			count = 1
			char = "IX"
			num -= 9
		} else if num >= 5 {
			count = num / 5
			char = "V"
			num = num % 5
		} else if num >= 4 {
			count = 1
			char = "IV"
			num -= 4
		} else {
			count = num
			char = "I"
			num = 0
		}

		for ; count > 0; count-- {
			roman += char
		}
	}

	return roman
}
