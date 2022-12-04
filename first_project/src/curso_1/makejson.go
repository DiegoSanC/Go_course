package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	map_placeholder := make(map[string]string)
	for i := 0; i < 2; i++ {
		entered_info, info_type := ask_info(i)
		if entered_info == "unk" {
			fmt.Println("something not expected has happened. Exiting...")
			break
		}
		map_placeholder = fill_map(map_placeholder, entered_info, info_type)
	}

	json_content, err := json.Marshal(map_placeholder)

	if err == nil {
		fmt.Println("JSON byte representation of your info:")
		fmt.Println(json_content)
		fmt.Println("JSON human readable representation:")
		fmt.Println(string(json_content))
	} else {
		fmt.Println("something happened when creating json")
	}
}

func fill_map(dict map[string]string, entered_info, info_type string) map[string]string {

	dict[info_type] = entered_info
	return dict
}

func ask_info(type_ int) (string, string) {

	switch type_ {
	case 0:
		fmt.Println("Please enter your name:")
		in := bufio.NewReader(os.Stdin)
		x, _ := in.ReadString('\n')
		return x, "name"
	case 1:
		fmt.Println("Please enter your address:")
		in := bufio.NewReader(os.Stdin)
		x, _ := in.ReadString('\n')
		return x, "address"
	}
	return "unk", "unk"
}
