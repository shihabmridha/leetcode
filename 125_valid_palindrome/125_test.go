package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	res := isPalindrome("A man, a plan, a canal: Panama")
	if !res {
		t.Logf("Output: %t, Expected: true", res)
		t.Fail()
	}

	res = isPalindrome("race a car")
	if res {
		t.Logf("Output: %t, Expected: false", res)
		t.Fail()
	}
}
