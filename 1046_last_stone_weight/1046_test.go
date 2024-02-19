package laststoneweight

import "testing"

func TestLastStoneWeight(t *testing.T) {
	res := lastStoneWeight([]int{2, 7, 4, 1, 8, 1})
	if res != 1 {
		t.Logf("Output: %d, Expected: 1", res)
		t.Fail()
	}

	res = lastStoneWeight([]int{31, 26, 33, 21, 40})
	if res != 9 {
		t.Logf("Output: %d, Expected: 9", res)
		t.Fail()
	}

	res = lastStoneWeight([]int{1})
	if res != 1 {
		t.Logf("Output: %d, Expected: 1", res)
		t.Fail()
	}
}
