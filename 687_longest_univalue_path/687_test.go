package longestunivaluepath

import "testing"

func TestLongestUnivaluePath(t *testing.T) {
	inp1 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
		},
	}

	res := longestUnivaluePath(&inp1)

	if res != 3 {
		t.Fail()
	}

	// var input = solution.TreeNode{
	// 	Val: 5,
	// 	Left: &solution.TreeNode{
	// 		Val: 4,
	// 		Left: &solution.TreeNode{
	// 			Val: 1,
	// 		},
	// 		Right: &solution.TreeNode{
	// 			Val: 1,
	// 		},
	// 	},
	// 	Right: &solution.TreeNode{
	// 		Val: 5,
	// 		Right: &solution.TreeNode{
	// 			Val: 5,
	// 		},
	// 	},
	// }

	// var input = solution.TreeNode{
	// 	Val: 1,
	// 	Left: &solution.TreeNode{
	// 		Val: 4,
	// 		Left: &solution.TreeNode{
	// 			Val: 4,
	// 		},
	// 		Right: &solution.TreeNode{
	// 			Val: 4,
	// 		},
	// 	},
	// 	Right: &solution.TreeNode{
	// 		Val: 5,
	// 		Right: &solution.TreeNode{
	// 			Val: 5,
	// 		},
	// 	},
	// }

	// var input = solution.TreeNode{
	// 	Val: 1,
	// 	Left: &solution.TreeNode{
	// 		Val: 1,
	// 		Left: &solution.TreeNode{
	// 			Val: 1,
	// 		},
	// 		Right: &solution.TreeNode{
	// 			Val: 3,
	// 			Left: &solution.TreeNode{
	// 				Val: 3,
	// 				Left: &solution.TreeNode{
	// 					Val: 4,
	// 				},
	// 				Right: &solution.TreeNode{
	// 					Val: 3,
	// 					Left: &solution.TreeNode{
	// 						Val: 5,
	// 					},
	// 					Right: &solution.TreeNode{
	// 						Val: 3,
	// 					},
	// 				},
	// 			},
	// 			Right: &solution.TreeNode{
	// 				Val: 3,
	// 			},
	// 		},
	// 	},
	// 	Right: &solution.TreeNode{
	// 		Val: 1,
	// 		Right: &solution.TreeNode{
	// 			Val: 5,
	// 			Left: &solution.TreeNode{
	// 				Val: 6,
	// 			},
	// 			Right: &solution.TreeNode{
	// 				Val: 5,
	// 			},
	// 		},
	// 	},
	// }
}
