package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fileName := "./hello.txt"	

	content, err := ioutil.ReadFile(fileName)
	checkError(err)
	result := string(content)
	fmt.Println("Content of file in bytes", content)
	fmt.Println("Content of file in string:", result)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}