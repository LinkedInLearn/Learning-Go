package main

import (
	"fmt"
	"os"
	"errors"
)

func main() {
	f, err := os.Open("filename.ext")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}
	
	myError := errors.New("My error string")
	fmt.Println(myError)


	attendance := map[string]bool{
		"Ann": true,
		"Mike": true}

	//altenative to defining attendance
	//attendance := map[string]bool{}
	//attendance["Ann"] = true
	//attendance["Mike"] = true

	attended, ok := attendance["Mike"]
	if ok {
		fmt.Println("Mike attended?", attended)
		fmt.Println("Mike ok?", ok)
	} else {
		fmt.Println("No info for Mike")
	}

}
