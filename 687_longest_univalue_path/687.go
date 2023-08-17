package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result int

func solution(node *TreeNode, value int) int {
	if node == nil {
		return 0
	}

	left := solution(node.Left, node.Val)
	right := solution(node.Right, node.Val)

	var nodeValue = left + right

	if nodeValue > result {
		result = nodeValue
	}

	if node.Val == value {
		nodeValue += 1
	} else {
		return 0
	}

	if nodeValue > result {
		result = nodeValue
	}

	return nodeValue
}

func LongestUnivaluePath(root *TreeNode) int {
	result = 0

	solution(root, -9999)

	return result
}
