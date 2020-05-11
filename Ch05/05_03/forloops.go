package main

import "fmt"

func main() {

	sum := 1
	fmt.Println("Sum:", sum)
	
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := range colors{
		fmt.Println(colors[i])
	}
	
}
