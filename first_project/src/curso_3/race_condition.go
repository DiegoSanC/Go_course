package main

import (
	"fmt"
)

func main() {

	var number int

	/*

		I have defined 3 functions that manipulate same variable in different ways. One is going to set the variable as 1,
		other as 2 and last one as 3. We run them as go routines so due to the uncertainty of the Go scheduler and the
		interleaving process, in each execution the printed result can be different. Acting like this is generating a race
		condition around variable number, which is going to be shared among goroutines.

		It can be 0, if the print function is executed before any of the goroutines or it can be 1, 2 or 3 depending of
		which of the goroutines is executed just before the print function is executed.

	*/
	go set_one(&number)
	go set_two(&number)
	go set_three(&number)

	fmt.Println(number)
}

func set_one(n *int) {
	*n = 1
}

func set_two(n *int) {
	*n = 2
}
func set_three(n *int) {
	*n = 3
}
