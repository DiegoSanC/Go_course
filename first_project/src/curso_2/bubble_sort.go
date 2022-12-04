package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func splitter(str string) []string {
	numbers := strings.Fields(str)
	return numbers
}

func length_checker() {}
func is_num_checker(numbers []string) {
	for _, num := range numbers {
		fmt.Println(num)
	}
}

func main() {
	println("Please enter a sequence of 10 numbers separated by spaces:")
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	numbers := splitter(str)
	is_num_checker(numbers)

}

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	if err := scanner.Err(); err != nil {
// 		log.Println(err)
// 	}
// 	fmt.Println("Enter X to exit...")
// 	fmt.Printf("Please enter a list of integers (space-separated): ")

// 	for scanner.Scan() {
// 		var strVal string = scanner.Text()
// 		if strVal == "X" {
// 			os.Exit(0)
// 		}
// 		fmt.Printf(strVal)
// 	}
// 	fmt.Println("out of loop")
// }
