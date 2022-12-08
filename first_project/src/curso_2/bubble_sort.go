package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	list_size int = 10
)

func splitter(str string) []string {
	numbers := strings.Fields(str)
	return numbers
}

func length_checker(int_numbers []int) []int {
	if len(int_numbers) > list_size {
		fmt.Printf("the length of the list of numbers is bigger than %d", list_size)
		fmt.Println()
		fmt.Printf("Used list will be: ")
		fmt.Println(int_numbers[:list_size])
		return int_numbers[:list_size]
	} else {
		return int_numbers
	}
}

func is_num_checker(numbers []string) []int {
	var slice_of_numbers []int
	for _, str_num := range numbers {
		int_num, err := strconv.Atoi(str_num)

		if err != nil {
			fmt.Printf("String %s is not a number so it is not going to be included in 10 numbers list", str_num)
			fmt.Println()
		} else {
			slice_of_numbers = append(slice_of_numbers, int_num)
		}
	}
	return slice_of_numbers

}

func swap(numbers []int, i int) {
	temp_num := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = temp_num
}

func BubbleSort(int_numbers []int) {
	loop_again := true
	for loop_again {
		loop_again = false
		for i := 0; i < len(int_numbers)-1; i++ {
			if int_numbers[i] > int_numbers[i+1] {
				swap(int_numbers, i)
				loop_again = true //If we go over all the slice and never swap, it means the slice is already sorted
			}
		}
	}
}

func main() {
	println("Please enter a sequence of 10 numbers separated by spaces:")
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	numbers := splitter(str)
	int_numbers := is_num_checker(numbers)
	final_int_numbers := length_checker(int_numbers)
	fmt.Println("Unordered list:")
	fmt.Println(final_int_numbers)
	BubbleSort(final_int_numbers)
	fmt.Println("Sorted list:")
	fmt.Println(final_int_numbers)

}

// prueba := []int{8, 1, 2, 3, 4}
// fmt.Println(prueba)
// swap(prueba, 0)
// fmt.Println(prueba)
