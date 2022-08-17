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

// ^ template

func lengthOfLongestSubstring(s string) int {
	flag := map[rune]bool{}
	max := 0
	count := 0

	for _, c := range s {
		if flag[c] {
			if count > max {
				max = count
			}

			count = 0
		} else {
			count++
		}
	}

	return 0
}

func main() {
	defer writer.Flush()

	printf("%d", lengthOfLongestSubstring("abcabcbb"))
}
