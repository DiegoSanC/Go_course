package main

import "fmt"

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * t * t) + (v0 * t) + s0
	}
}

func main() {
	var a, v0, s0, t float64
	fmt.Print("Enter acceleration, initial velocity and initial displacement: ")
	fmt.Scan(&a, &v0, &s0)
	var fn = GenDisplaceFn(a, v0, s0)
	for {
		fmt.Print("Enter time (or -1 to exit): ")
		fmt.Scan(&t)
		if t == float64(-1) {
			break
		}
		fmt.Println("displacement: ", fn(t))
	}
}
