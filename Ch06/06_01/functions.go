package main

import "fmt"

func main() {
	run()
	fmt.Println(addEverything(1,20,25,23))
}

func run(){
	fmt.Println("WE RAN BOY!")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func addEverything(values ...int) int{
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}