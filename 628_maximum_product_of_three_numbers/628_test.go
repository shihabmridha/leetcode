package maximumproductofthreenumbers

import "testing"

func TestMaximumProduct(t *testing.T) {
	res := maximumProduct([]int{1, 2, 3})
	if res != 6 {
		t.Logf("Output %d, Expected: 6", res)
		t.Fail()
	}

	res = maximumProduct([]int{1, 2, 3, 4})
	if res != 24 {
		t.Logf("Output %d, Expected: 24", res)
		t.Fail()
	}

	res = maximumProduct([]int{-1, -2, -3})
	if res != -6 {
		t.Logf("Output %d, Expected: -6", res)
		t.Fail()
	}
}
