package findtheindexofthefirstoccurrenceinastring

import "testing"

func TestStrStr(t *testing.T) {
	res := strStr("sadbutsad", "sad")

	if res != 0 {
		t.Logf("Output: %d, Expected: 0", res)
		t.Fail()
	}

	res = strStr("leetcode", "leeto")
	if res != -1 {
		t.Logf("Output: %d, Expected: -1", res)
		t.Fail()
	}

	res = strStr("a", "a")
	if res != 0 {
		t.Logf("Output: %d, Expected: 0", res)
		t.Fail()
	}

	res = strStr("mississippi", "issi")
	if res != 1 {
		t.Logf("Output: %d, Expected: 1", res)
		t.Fail()
	}

	res = strStr("mississippi", "issip")
	if res != 4 {
		t.Logf("Output: %d, Expected: 1", res)
		t.Fail()
	}

	res = strStr("mississippi", "issipi")
	if res != -1 {
		t.Logf("Output: %d, Expected: -1", res)
		t.Fail()
	}

	res = strStr("shihab", "shihabmridha")
	if res != -1 {
		t.Logf("Output: %d, Expected: -1", res)
		t.Fail()
	}
}
