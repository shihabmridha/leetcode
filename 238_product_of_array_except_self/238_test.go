package productofarrayexceptself

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	res := productExceptSelf([]int{1, 2, 3, 4})
	if !reflect.DeepEqual(res, []int{24, 12, 8, 6}) {
		t.Logf("Output: %d, Expected: [24, 12, 8, 6]", res)
		t.Fail()
	}

	res = productExceptSelf([]int{1, 1, 0, 3, 3})
	if !reflect.DeepEqual(res, []int{0, 0, 9, 0, 0}) {
		t.Logf("Output: %d, Expected: [0, 0, 9, 0, 0]", res)
		t.Fail()
	}

	res = productExceptSelf([]int{4, 3, 2, 1, 2})
	if !reflect.DeepEqual(res, []int{12, 16, 24, 48, 24}) {
		t.Logf("Output: %d, Expected: [12,16,24,48,24]", res)
		t.Fail()
	}

}
