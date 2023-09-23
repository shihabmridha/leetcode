package findminimuminrotatedsortedarray

import "testing"

func TestFindMin(t *testing.T) {
	res := findMin([]int{3, 4, 5, 1, 2})

	if res != 1 {
		t.Logf("Output: %d, Expected: 1", res)
		t.Fail()
	}
}
