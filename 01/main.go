package main

import "fmt"

type Human struct {
	Name string
}

func (h Human) greeting() string {
	return fmt.Sprintf("Hi, I'm %v!", h.Name)
}

type Action struct {
	Human
	Motion string
}

func main() {
	act := Action{
		Human: Human{
			Name: "Vlad",
		},
		Motion: "waving",
	}

	fmt.Println(act.greeting())
	fmt.Printf("I can %v!", act.Motion)
}
