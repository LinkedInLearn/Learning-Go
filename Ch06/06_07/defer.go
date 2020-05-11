package main

import "fmt"

func main() {
	defer fmt.Println("Close the File")
	fmt.Println("Open the File")
	defer fmt.Println("1")
	defer fmt.Println("2")
	myFunc()
	defer fmt.Println("3")
	defer fmt.Println("4")
	fmt.Println("5")

	x := 1000
	defer fmt.Println("Value of x:", x)
	x++;
	fmt.Println("Value of x after incrementing:", x)
}

func myFunc() {
	defer fmt.Println("Deffered in the function")
	fmt.Println("Not Deffered in the function")
}
