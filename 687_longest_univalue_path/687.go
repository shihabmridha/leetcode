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

	if node.Left != nil && node.Right != nil {
		if node.Left.Val == node.Val && node.Right.Val == node.Val {
			var nodeValue = left + right

			if node.Val == value {
				nodeValue += 1
			}

			if nodeValue > result {
				result = nodeValue
			}

			return nodeValue
		}
	}

	if node.Left != nil && node.Left.Val == node.Val {
		if left+1 > result {
			result = left + 1
		}

		return left + 1
	}

	if node.Right != nil && node.Right.Val == node.Val {
		if right+1 > result {
			result = right + 1
		}

		return right + 1
	}

	if node.Val == value {
		return 1
	}

	return 0
}

func LongestUnivaluePath(root *TreeNode) int {
	result = 0

	solution(root, -9999)

	return result
}
