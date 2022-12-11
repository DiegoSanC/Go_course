package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Please, input animal name and desired method after '>'. To exit press 'x'")

	for {

		fmt.Print(">")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input := scanner.Text()
			if input == "x" {
				break
			}

			fmt.Println(input)
			values := strings.Fields(input)
			animal := Init(values[0])
			animal.DoRequest(values[1])
		}
	}
}

func (animal *Animal) DoRequest(request string) {
	if request == "speak" {
		animal.Speak()
	} else if request == "eat" {
		animal.Eat()
	} else if request == "move" {
		animal.Move()
	}
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func Init(name string) Animal {
	if name == "cow" {
		return Animal{"grass", "walk", "moo"}
	}
	if name == "snake" {
		return Animal{"mice", "slither", "hsss"}
	}
	if name == "bird" {
		return Animal{"worms", "fly", "peep"}
	}
	return Animal{"", "", ""}
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
