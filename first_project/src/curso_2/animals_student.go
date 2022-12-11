package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	data := map[string]Animal{
		"cow": {
			"grass",
			"walk",
			"moo",
		},
		"bird": {
			"worms",
			"fly",
			"peep",
		},
		"snake": {
			"mice",
			"slither",
			"hsss",
		},
	}

	request := make([]string, 2)
	fmt.Print(">")
	fmt.Scan(&request[0], &request[1])

	if animal, ok := data[request[0]]; ok {
		if request[1] == "eat" {
			animal.Eat()
		} else if request[1] == "move" {
			animal.Move()
		} else if request[1] == "speak" {
			animal.Speak()
		} else {
			fmt.Println("Invalid action")
		}
	} else {
		fmt.Println("Invalid animal")
	}
}
