package main

import (
	"fmt"
)

type struct Person {
	name string
	addr string
	phone string
}

func main() {
	
	map_hash_table()
}

func strcture_analysis(){
	p1 := new(Person)
	p1 := Person(name: "joe", addr: "lol", phone: "123")
	p1.name = "joe"
	p1.addr = "lol"
	p1.phone = "123"
}



func map_hash_table() {
	//key type value type
	// var id_map map[string]int
	// id_map = make(map[string]int)
	id_map := map[string]int{
		"joe": 123}

	id_map["jane"] = 456
	fmt.Println(id_map, id_map["jane"])
	id, p := id_map["joe"]
	fmt.Println(id, p)

	for key, val := range id_map {
		fmt.Println(key, val)
	}

}

func append_func() {
	sli_1 := make([]int, 10, 10)
	sli_2 := append(sli_1, 100)
	fmt.Println((sli_2))

}

func make_function() {
	//type, lenght, capacity
	sli_1 := make([]int, 10)
	sli_2 := make([]int, 10, 15)
	fmt.Println(sli_1, sli_2)
}

func slice_initialization() {
	//This is a slice because the brackets are empty
	a_1 := []int{1, 2, 3}
	fmt.Println(a_1)
	//return a_1
}

func show_slicing_behaviour() {
	a_1 := [...]int{1, 2, 3, 4, 5}
	sli_1 := a_1[:3]
	sli_2 := a_1[2:]
	fmt.Println(sli_1, sli_2)
	sli_1[2] = 10
	fmt.Println(sli_1, sli_2, a_1)
}
