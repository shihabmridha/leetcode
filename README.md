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
