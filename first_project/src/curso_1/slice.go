package main

import (
	"fmt"
	"sort"
	"strconv"
)

// func main() {
// 	user_input := ask_user_input()
// 	int_value, err := strconv.Atoi(user_input)
// 	if err == nil {
// 		fmt.Println("No error")
// 	}
// 	fmt.Println(int_value, err)
// }

func main() {
	slice := make([]int, 3)
	// slice[0] = math.MaxInt64
	// slice[1] = math.MaxInt64
	// slice[2] = math.MaxInt64
	counter := 0
	original_len := len(slice)

	for {

		user_input := ask_user_input()
		if user_input == "X" {
			break
		}

		int_value, err := strconv.Atoi(user_input)

		if err != nil {
			fmt.Println("Input is not int. Try it again please")
			continue
		}

		if counter < original_len {
			slice[original_len-1-counter] = int_value
			counter = counter + 1
		} else {
			slice = append(slice, int_value)

		}

		sort.Ints(slice)

		fmt.Println(slice)

	}

}
func ask_user_input() string {
	var x string
	fmt.Println("Enter an int number or enter X in case you want to exit:")
	fmt.Scan(&x)
	return x
}
