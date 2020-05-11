package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	
	content := "-,...,-"

	file, err := os.Create("./fromString.txt")
	checkError(err)
	defer file.Close()


	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Println(ln)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}