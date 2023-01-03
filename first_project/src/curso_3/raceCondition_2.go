package main

import "fmt"

/**
A race condition occurs when two or more goroutines can access shared data and they try to change it at the same time.
*/

var cnt = 0
var c = make(chan int, 10)

func count1() {
	for i := 0; i < 100000; i++ {
		cnt += 1
	}
	c <- 0
}

func count2() {
	for i := 0; i < 100000; i++ {
		cnt += 1
	}
	c <- 0
}

func main() {
	go count1() // start one goroutine which increment variable cnt by 100000
	go count2() // start another goroutine which increment variable cnt by 100000
	<-c
	<-c
	fmt.Println(cnt) // because of the race condition, the final cnt value is probably a value smaller than 200000

}
