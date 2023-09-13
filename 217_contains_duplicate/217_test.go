package contains_duplicate

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	res := ContainsDuplicate([]int{1, 2, 3, 4})

	if res != false {
		t.Fail()
	}
}
