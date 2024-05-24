package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual(res, []int{0, 1}) {
		t.Logf("Output %d, Expected: [0, 1]", res)
		t.Fail()
	}

	res = twoSum([]int{3, 2, 4}, 6)
	if !reflect.DeepEqual(res, []int{1, 2}) {
		t.Logf("Output %d, Expected: [1, 2]", res)
		t.Fail()
	}

	res = twoSum([]int{3, 3}, 6)
	if !reflect.DeepEqual(res, []int{0, 1}) {
		t.Logf("Output %d, Expected: [0, 1]", res)
		t.Fail()
	}

	res = twoSum([]int{-1, -2, -3, -4, -5}, -8)
	if !reflect.DeepEqual(res, []int{4, 2}) {
		t.Logf("Output %d, Expected: [2,4] / [4,2]", res)
		t.Fail()
	}
}
