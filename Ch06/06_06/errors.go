package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.ext")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println("Error:", err)
	}
	attendance := map[string]bool{
		"Ann": true,
		"Mike": true}

	attended, ok := attendance["Mike"]
	if ok {
		fmt.Println(attended)
	} else {
		fmt.Println("Did not arttend")
	}
}
