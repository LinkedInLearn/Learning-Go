package main

import "fmt"

func main() {
	
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true
	stringLength, _ := fmt.Println(str1,str2,str3 )

	//if err == nil {
		fmt.Println("String Length",stringLength)
	//}
	fmt.Printf("ValueI : %v\n", aNumber)
	fmt.Printf("ValueF : %.2f\n", float64(aNumber))
	fmt.Printf("ValueB : %v\n", isTrue)
	fmt.Printf("Type : %T\n%T\n%T\n", isTrue, aNumber, str3)
}
