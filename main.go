package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Human struct {
	Name string
}

func (h Human) Speak() string {
	return "Hello, my name is " + h.Name
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Mew!"
}

type Robot struct{}

func (r Robot) Speak() string { return "Beep boop" }

type Person struct{}

func (p Person) Speak() string { return "Hi there!" }

type Duck struct{}

func (d Duck) Speak() string { return "Quack!" }

func Broadcast(s Speaker) {
	fmt.Println("Broadcasting:", s.Speak())
}

func Identify(s Speaker) {
	switch v := s.(type) {
	case Human:
		fmt.Println("Speaker is a Human", v.Speak())
	case Cat:
		fmt.Println("Speaker is a Cat", v.Speak())
	default:
		fmt.Println("Unknown type")
	}
}

func main() {

	// Definition & Implementation
	var s1 Speaker = Human{Name: "Johnny"}
	var s2 Speaker = Cat{}
	fmt.Println(s1.Speak(), "\t", s2.Speak())

	for x := range 37 {
		fmt.Print("-")
		if x == 36 {
			fmt.Println()
		}
		x++
	}

	// Type Switch
	Identify(Human{})
	Identify(Cat{})

	// Function Parameter
	for x := range 37 {
		fmt.Print("-")
		if x == 36 {
			fmt.Println()
		}
		x++
	}

	Broadcast(Robot{})

	// Interface Slice Collection
	for x := range 37 {
		fmt.Print("-")
		if x == 36 {
			fmt.Println()
		}
		x++
	}

	speakers := []Speaker{Person{}, Duck{}, Robot{}}
	for _, s := range speakers {
		fmt.Println(s.Speak())
	}

	// Interface Slice Collection
	for x := range 37 {
		fmt.Print("-")
		if x == 36 {
			fmt.Println()
		}
		x++
	}

}
