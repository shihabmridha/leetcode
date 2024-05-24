### Template
```
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

func main() {
	defer writer.Flush()
}

```

### Usage
I create a new directory with the problem number and problem name using the following format `{problem_number}_{problem_name}`. Create a file named `{problem_number}`. I create a `test` for each problem to test the code.


### Helpers
- Use `%v` to print array. Example: `fmt.Printf("%v", projects)`
- Create empty map `make(map[int]int)`, `map[rune]bool{}`
- Create empty rune `make([]rune, 5) // with size`, `map[rune]bool{}`
- Sort list/slice `slices.Sort(mySlice)`.
- Sort with comparer `sort.Slice(mySlice, func(i, j int) bool { return mySlice[i] > mySlice[j] })`
- Deep compare array `reflect.DeepEqual([]int{0, 1, 1}, []int{0, 1, 1})`
