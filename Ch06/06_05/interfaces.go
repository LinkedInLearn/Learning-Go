package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string{
	return "bark"
}

type Cat struct {
}

func (c Cat) Speak() string{
	return "Meow"
}

type Cow struct {
}

func (c Cow) Speak() string{
	return "Moo"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Cow{}}
	for _, i := range animals{
		fmt.Println(i.Speak())
	}
}
