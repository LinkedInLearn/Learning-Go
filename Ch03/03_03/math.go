package main

import (
	"fmt"
	"math"
)

func main() {
	
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum: ", intSum)


	//var b1, b2, b3, bigSum big.Float
	//b1.SetFloat64(23.5)
	//b2.SetFloat64(23.5)
	//b3.SetFloat64(23.5)
	//
	//bigSum.Add(&b1, &b2).Add(&b3, &bigSum)
	//
	//fmt.Printf("Bigsum = %.10g", &bigSum)
	circleRadius := 15.5
	circumference := math.Pi * 2 * circleRadius
	fmt.Printf("Circumference: %.2f", circumference)


}
