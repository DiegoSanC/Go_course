package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	println("Enter a string:")
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	str = strings.ToLower(str)
	len_str := len(str)
	fmt.Println("length string before adaptation: ", len_str)
	str = str[:len_str-2]
	len_str = len(str)
	fmt.Println("length string after adaptation: ", len_str)

	if len_str < 3 {
		fmt.Println("Not Found!")
		return
	}

	is_valid_first_char := str[0] == 105 || str[0] == (105-32)
	is_valid_last_char := (str[len_str-1] == 110) || (str[len_str-1] == (110 - 32))
	are_valid := is_valid_first_char && is_valid_last_char

	for i := 1; i < len_str-1; i++ {
		if str[i] == 97 || str[i] == (97-32) {
			if are_valid {
				fmt.Println("Found!")
				return
			}
		}
	}
	fmt.Println("Not Found!")
	return
}
