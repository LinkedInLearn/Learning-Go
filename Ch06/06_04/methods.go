package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) Speak3Times() {
	fmt.Println(d.Sound)
}

func main() {
	poodle := Dog{Breed: "poodle",
		Weight: 10,
		Sound:  "woof"}
	fmt.Println(poodle)
	poodle.Speak()

	poodle.Sound = "BARK BITCH"
	poodle.Speak()

}
