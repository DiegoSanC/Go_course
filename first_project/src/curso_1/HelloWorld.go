package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	f()
}

func f() {
	i := 0
	for i < 10 {
		i++
		if i == 5 {
			continue
		}
		fmt.Printf("%d", i)
	}
}
