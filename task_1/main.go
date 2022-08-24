package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Gender string
}

func (h *Human) Talk() {
	fmt.Println("Hello, My name is", h.Name, "I'm", h.Age, "years old")
}

type Action struct {
	Human
	Place string
}

func (a *Action) walk() {
	fmt.Println("I'm going to", a.Place)
}

func main() {
	a := Action{
		Place: "Home",
		Human: Human{
			Name:   "Kolya",
			Age:    19,
			Gender: "male",
		},
	}

	a.Talk()
	a.walk()

}
