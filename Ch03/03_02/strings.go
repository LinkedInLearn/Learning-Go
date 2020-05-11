package main

import "fmt"

func main() {
	str1 := "This is a string"
	fmt.Printf("Implicitly typed string: %v : %T \n",str1, str1)

	var str2 string = "This is a string2"
	fmt.Printf("Explicitly typed string: %v : %T \n",str2, str2)
}
