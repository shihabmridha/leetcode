package searchinsertposition

import "testing"

func TestSearchInsert(t *testing.T) {
	res := searchInsert([]int{1, 3, 5, 6}, 5)
	if res != 2 {
		t.Logf("Output: %d, Expected: 2", res)
		t.Fail()
	}

	res = searchInsert([]int{1, 3, 5, 6}, 2)
	if res != 1 {
		t.Logf("Output: %d, Expected: 1", res)
		t.Fail()
	}

	res = searchInsert([]int{1, 3, 5, 6}, 7)
	if res != 4 {
		t.Logf("Output: %d, Expected: 4", res)
		t.Fail()
	}
}
