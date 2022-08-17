package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {

	return 69
}

func main() {
	defer writer.Flush()

	printf("%d", longestUnivaluePath(&TreeNode{}))
}
