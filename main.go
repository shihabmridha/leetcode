package main

import (
	"log"
	"os"
	"strings"
)

// var reader *bufio.Reader = bufio.NewReader(os.Stdin)
// var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

// func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
// func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func getDirName(problemName string) string {
	str := strings.ToLower(problemName)
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ReplaceAll(str, " ", "_")

	return str
}

func main() {
	// defer writer.Flush()

	problem := "125. Valid Palindrome"

	dirName := getDirName(problem)

	err := os.Mkdir(dirName, 0755)
	if err != nil {
		log.Fatal(err)
	}

	fileName := strings.Split(dirName, "_")[0]
	pkg := strings.Join(strings.Split(dirName, "_")[1:], "")

	s, err := os.OpenFile(dirName+"/"+fileName+".go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	s.Write([]byte("package " + pkg))
	s.Close()

	t, err := os.OpenFile(dirName+"/"+fileName+"_test.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	t.Write([]byte("package " + pkg))
	t.Close()
}
