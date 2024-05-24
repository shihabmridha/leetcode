package rangesumofbst

import (
	"fmt"
	"testing"
)

func TestRangeSumBST(t *testing.T) {
	bst1 := TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	res := rangeSumBST(&bst1, 7, 15)
	if res != 32 {
		fmt.Printf("Output: %d, Expected: 32\n", res)
		t.Fail()
	}

	bst2 := TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
			Left: &TreeNode{
				Val: 13,
			},
		},
	}
	res = rangeSumBST(&bst2, 6, 10)
	if res != 23 {
		fmt.Printf("Output: %d, Expected: 23\n", res)
		t.Fail()
	}
}
