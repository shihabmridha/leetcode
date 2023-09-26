package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	res := longestCommonPrefix([]string{
		"flower", "flow", "flight",
	})

	if res != "fl" {
		t.Logf("Output: %s, Expected: fl", res)
		t.Fail()
	}

	res = longestCommonPrefix([]string{
		"dog", "racecar", "car",
	})
	if res != "" {
		t.Logf("Output: %s, Expected: \"\"", res)
		t.Fail()
	}

	res = longestCommonPrefix([]string{
		"flower", "flow", "flight", "",
	})
	if res != "" {
		t.Logf("Output: %s, Expected: \"\"", res)
		t.Fail()
	}

	res = longestCommonPrefix([]string{
		"",
	})
	if res != "" {
		t.Logf("Output: %s, Expected: \"\"", res)
		t.Fail()
	}

	res = longestCommonPrefix([]string{
		"a",
	})
	if res != "a" {
		t.Logf("Output: %s, Expected: a", res)
		t.Fail()
	}

	res = longestCommonPrefix([]string{
		"ab", "a",
	})
	if res != "a" {
		t.Logf("Output: %s, Expected: a", res)
		t.Fail()
	}
}
