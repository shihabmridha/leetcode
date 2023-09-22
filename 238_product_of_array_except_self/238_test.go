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
}
