package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	res := lengthOfLongestSubstring("asdfapqrtuv")

	if res != 10 {
		t.Logf("output: %d, expected: 10", res)
		t.Fail()
	}
}
