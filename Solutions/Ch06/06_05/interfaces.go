package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Cow struct {
}

func (c Cow) Speak() string {
	return "Moo!"
}

type Tree struct {
	
}

func main() {
	poodle := Animal(Dog{})
	fmt.Println(poodle)
	
	animals := []Animal{Dog{}, Cat{}, Cow{}}
	// here _ is the index, animal is the object at index
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	//alternative
	for animal := range animals {
		fmt.Println(animals[animal].Speak())
	}
}
