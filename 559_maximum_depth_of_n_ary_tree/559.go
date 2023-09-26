package maximumdepthofnarytree

type Node struct {
	Val      int
	Children []*Node
}

var depth int = 1

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func solve(root *Node, count int) {
	depth = max(depth, count)

	if root.Children == nil {
		return
	}

	for _, node := range root.Children {
		solve(node, count+1)
	}
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	depth = 0
	solve(root, depth+1)

	return depth
}
