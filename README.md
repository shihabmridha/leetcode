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
I create a new directory with the problem number and problem name using the following format `{problem_number}_{problem_name}`. Create a file named `{problem_number}`. For all the solutions I use the package name `solution` so that when I use it in the `main.go` file I can call the function like this `solution.FunctionName`. Make sure to export the function and necessary types so that I can call it from `main()`.

In the `main.go` file, I import the package and call the function.  I prepare the necessary input in the `main()` and pass it to the solution function.