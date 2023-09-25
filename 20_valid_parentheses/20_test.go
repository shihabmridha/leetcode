package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	res := isValid("()")
	if !res {
		t.Logf("Output: %t, Expected: true", res)
		t.Fail()
	}

	res = isValid("([]{})")
	if !res {
		t.Logf("Output: %t, Expected: true", res)
		t.Fail()
	}

	res = isValid("(]")
	if res {
		t.Logf("Output: %t, Expected: false", res)
		t.Fail()
	}

	res = isValid("))")
	if res {
		t.Logf("Output: %t, Expected: false", res)
		t.Fail()
	}
}
