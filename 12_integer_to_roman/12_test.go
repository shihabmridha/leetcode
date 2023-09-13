package integertoroman

import "testing"

func TestIntToRoman(t *testing.T) {
	res := intToRoman(1994)

	if res != "MCMXCIV" {
		t.Fail()
	}
}
