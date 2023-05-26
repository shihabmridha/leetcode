package main

import (
	"bufio"
	"fmt"
	"os"
	solution "shihabmridha/leetcode/687_longest_univalue_path"
)

// var reader *bufio.Reader = bufio.NewReader(os.Stdin)
// func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	defer writer.Flush()

	var input = solution.TreeNode{
		Val: 5,
		Left: &solution.TreeNode{
			Val: 4,
			Left: &solution.TreeNode{
				Val: 1,
			},
			Right: &solution.TreeNode{
				Val: 1,
			},
		},
		Right: &solution.TreeNode{
			Val: 5,
			Right: &solution.TreeNode{
				Val: 5,
			},
		},
	}

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

	x := solution.LongestUnivaluePath(&input)

	printf("%d\n", x)
}
