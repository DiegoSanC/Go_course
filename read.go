package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Name struct {
	first_name string
	last_name  string
}

func main() {
	var path_to_file string
	var names []Name
	fmt.Println("Please, introduce full path to the file: ")
	fmt.Scanln(&path_to_file)

	read_file, err := os.Open(path_to_file)

	if err != nil {
		fmt.Println(err)
	}
	file_scanner := bufio.NewScanner(read_file)

	file_scanner.Split(bufio.ScanLines)

	for file_scanner.Scan() {

		//fmt.Println(file_scanner.Text())
		line_content := file_scanner.Text()
		split_content := strings.Split(line_content, " ")
		name_instance := Name{
			trim_name(split_content[0]),
			trim_name(split_content[1]),
		}
		names = append(names, name_instance)
	}

	for i, n := range names {
		fmt.Println("Line number: " + strconv.Itoa(i))
		fmt.Println("First Name: " + n.first_name)
		fmt.Println("Last Name: " + n.last_name)
	}

	read_file.Close()

}

func trim_name(n string) string {

	if len(n) > 20 {
		return string(n[0:20])
	} else {
		return n
	}
}
