package findminimuminrotatedsortedarray

import "testing"

func TestFindMin(t *testing.T) {
	res := findMin([]int{3, 4, 5, 1, 2})
	if res != 1 {
		t.Logf("Output: %d, Expected: 1", res)
		t.Fail()
	}

	res = findMin([]int{4, 5, 6, 7, 0, 1, 2})
	if res != 0 {
		t.Logf("Output: %d, Expected: 0", res)
		t.Fail()
	}

	res = findMin([]int{11, 13, 15, 17})
	if res != 11 {
		t.Logf("Output: %d, Expected: 11", res)
		t.Fail()
	}
}
