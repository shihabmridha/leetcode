package main

import (
	"bufio"
	"fmt"
	"os"
	solution "shihabmridha/leetcode/12_integer_to_roman"
)

// var reader *bufio.Reader = bufio.NewReader(os.Stdin)
// func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	defer writer.Flush()

	x := solution.IntToRoman(1994)

	printf("%s\n", x)
}
