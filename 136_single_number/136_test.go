package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	res := singleNumber([]int{2, 2, 1})

	if res != 1 {
		t.Logf("Output: %d, Expected: 1", res)
		t.Fail()
	}

	res = singleNumber([]int{4, 1, 2, 1, 2})
	if res != 4 {
		t.Logf("Output: %d, Expected: 4", res)
		t.Fail()
	}

	res = singleNumber([]int{1})
	if res != 1 {
		t.Logf("Output: %d, Expected: 1", res)
		t.Fail()
	}
}
