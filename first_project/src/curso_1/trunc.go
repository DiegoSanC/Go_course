package main

import "fmt"

func main() {
	trunc()

}
func trunc() {
	var x float64
	fmt.Println("Enter a floating point number:")
	fmt.Scan(&x)
	fmt.Println("The truncated version of your number is ", int32(x))
}
