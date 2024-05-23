package countingbits

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	res := countBits(2)
	if !reflect.DeepEqual(res, []int{0, 1, 1}) {
		fmt.Printf("output: %d, expected: [0,1,1]\n", res)
		t.Fail()
	}

	res = countBits(5)
	if !reflect.DeepEqual(res, []int{0, 1, 1, 2, 1, 2}) {
		fmt.Printf("output: %d, expected: [0,1,1,2,1,2]\n", res)
		t.Fail()
	}
}
