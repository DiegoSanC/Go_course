package main

import (
	"fmt"
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

// func a() string {
// 	return "hola"
// }

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

	fmt.Println(cow)
	fmt.Println(bird)
	fmt.Println(snake)

	fmt.Println(cow.Animal)
	fmt.Println(bird.Animal)
	fmt.Println(snake.Animal)

	fmt.Println(cow.Animal.food)
	fmt.Println(bird.Animal)
	fmt.Println(snake.Animal)

	// dict_of_animals := map[string]func()string{
	// 	"cow":   a,
	// }

	// fmt.Println("Please introduce value of animal: ")
	// in := bufio.NewReader(os.Stdin)
	// animal, _ := in.ReadString('\n')

	// switch animal, action {
	// case "cow", "eat":
	// 	cow.Eat()
	// case "cow", "move":
	// 	cow.
	// case "cow", "speak":

	// case "bird", "eat":
	// case "bird", "move":
	// case "bird", "speak":

	// case "snake", "eat":
	// case "snake", "move":
	// case "snake", "speak":

	// }
	// selected_animal := dict_of_animals[str]
	// fmt.Println("Selected animal: ")
	// selected_animal.Eat()
	// selected_animal.Move()
	// selected_animal.Speak()
}
