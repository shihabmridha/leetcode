package main

import (
	"bufio"
	"fmt"
	"os"
	solution "shihabmridha/leetcode/3_longest_substring_without_repeating_characters"
)

// var reader *bufio.Reader = bufio.NewReader(os.Stdin)
// func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	defer writer.Flush()

	x := solution.LengthOfLongestSubstring("1324")

	printf("%d\n", x)
}
