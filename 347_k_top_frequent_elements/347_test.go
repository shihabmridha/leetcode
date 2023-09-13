package ktopfrequentelements

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	var res []int
	var expected []int

	expected = []int{1, 2}
	res = topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	sort.Ints(expected)
	sort.Ints(res)

	if !reflect.DeepEqual(res, expected) {
		fmt.Printf("output: %d, expected: [1, 2]\n", res)
		t.Fail()
	}

	res = topKFrequent([]int{1}, 1)
	if !reflect.DeepEqual(res, []int{1}) {
		fmt.Printf("output: %d, expected: [1]\n", res)
		t.Fail()
	}

	res = topKFrequent([]int{1, 1, 2, 3, 3, 4}, 2)
	expected = []int{1, 3}
	sort.Ints(expected)
	sort.Ints(res)

	if !reflect.DeepEqual(res, expected) {
		fmt.Printf("output: %d, expected: [1, 3]\n", res)
		t.Fail()
	}

	res = topKFrequent([]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 10)
	expected = []int{1, 2, 5, 3, 6, 7, 4, 8, 10, 11}
	sort.Ints(expected)
	sort.Ints(res)

	if !reflect.DeepEqual(res, expected) {
		fmt.Printf("output: %d, expected: [1,2,5,3,6,7,4,8,10,11]\n", res)
		t.Fail()
	}
}
