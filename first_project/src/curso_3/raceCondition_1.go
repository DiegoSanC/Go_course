// Explanation
// Race Condition: An undesirable situation that occurs when a device or system attempts to perform two or more operations at the same time.
// Here both sub-routines race(), are executing simultaneously and both tries to print the value of count at same time.
// As a result, it can't perform the incrementation of count task.

package main

import (
	"fmt"
	"time"
)

var count int

func race() {
	fmt.Println("The count is", count)
	count++
}

func main() {
	fmt.Println("The race condtition")
	count = 0
	go race()
	go race()
	time.Sleep(1 * time.Second)
}
