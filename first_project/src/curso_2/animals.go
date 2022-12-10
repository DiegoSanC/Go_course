package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

type Cow struct {
	Animal
}
type Bird struct {
	Animal
}
type Snake struct {
	Animal
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func splitter(str string) []string {
	numbers := strings.Fields(str)
	return numbers
}

func clean_input(input string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	result := re_leadclose_whtsp.ReplaceAllString(input, "")
	result = re_inside_whtsp.ReplaceAllString(result, " ")
	return result
}

func main() {
	cow := Cow{
		Animal{
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		},
	}

	bird := Bird{
		Animal{
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		},
	}

	snake := Snake{
		Animal{
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss",
		},
	}

	dict_of_animals := map[string]Animal{
		"cow":   cow.Animal,
		"bird":  bird.Animal,
		"snake": snake.Animal,
	}

	for {
		fmt.Println("Please introduce request > ")
		in := bufio.NewReader(os.Stdin)
		input_string, _ := in.ReadString('\n')

		splitted_string := splitter(input_string)

		if len(splitted_string) == 1 {
			fmt.Printf("Only partial information provided: %s\n", input_string)
			fmt.Println("Please, provide name of animal and an action (eat, move or speak)")
			continue
		}

		animal := splitted_string[0]
		action := splitted_string[1]

		if len(splitted_string) > 2 {
			fmt.Println("Introduced request is longer than 2 words")
			fmt.Printf("only %s and %s are going to be used\n", animal, action)
		}

		animal = clean_input(animal)
		action = clean_input(action)

		selected_animal, ok := dict_of_animals[animal]

		if !ok {
			fmt.Printf("Selected animal %s is not one of cow, bird or snake. Try it again\n", animal)
			continue
		}

		switch action {
		case "eat":
			selected_animal.Eat()
		case "move":
			selected_animal.Move()
		case "speak":
			selected_animal.Speak()
		}
	}
}
