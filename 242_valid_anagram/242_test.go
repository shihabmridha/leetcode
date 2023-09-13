package valid_anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	res := isAnagram("anagram", "nagaram")

	if res == false {
		t.Fail()
	}
}
