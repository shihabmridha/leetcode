package romantointeger

import "testing"

func TestRomanToInt(t *testing.T) {
	res := romanToInt("MCMXCIV")

	if res != 1994 {
		t.Fail()
	}
}
