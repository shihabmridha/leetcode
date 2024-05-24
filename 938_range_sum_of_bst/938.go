package rangesumofbst

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	return findSum(root, low, high)
}

func findSum(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}

	x := findSum(root.Left, low, high)
	y := findSum(root.Right, low, high)

	if root.Val >= low && root.Val <= high {
		return x + y + root.Val
	}

	return x + y
}
