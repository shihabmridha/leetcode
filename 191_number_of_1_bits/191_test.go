package numberof1bits

import "testing"

func TestHammingWeight(t *testing.T) {
	res := hammingWeight(11)
	if res != 3 {
		t.Logf("Output: %d, Expected: 3", res)
		t.Fail()
	}

	res = hammingWeight(128)
	if res != 1 {
		t.Logf("Output: %d, Expected: 1", res)
		t.Fail()
	}

	res = hammingWeight(2147483645)
	if res != 30 {
		t.Logf("Output: %d, Expected: 30", res)
		t.Fail()
	}
}
