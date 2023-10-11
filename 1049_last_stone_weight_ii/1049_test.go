package laststoneweightii

import "testing"

func TestLastStoneWeightII(t *testing.T) {
	res := lastStoneWeightII([]int{2, 7, 4, 1, 8, 1})
	if res != 1 {
		t.Logf("Output: %d, Expected: 1", res)
		t.Fail()
	}

	res = lastStoneWeightII([]int{31, 26, 33, 21, 40})
	if res != 5 {
		t.Logf("Output: %d, Expected: 5", res)
		t.Fail()
	}
}
