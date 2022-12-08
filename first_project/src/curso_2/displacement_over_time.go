package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func read_user_input() string {
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	return str
}

func convert_validator(variable, str string) (float64, error) {
	float_num, err := strconv.ParseFloat(str, 32)

	if err != nil {
		fmt.Printf("Something happened when reading %s variable. Input value was %s. Read error trace: \n", variable, str)
		fmt.Println(err)
	}
	return float_num, err
}

// func continous_reader(variable, str string) (float_num float64) {
// 	well_read := false
// 	var err error
// 	for !well_read {
// 		float_num, err = strconv.ParseFloat(str, 32)

// 		if err != nil {
// 			fmt.Printf("Something happened when reading %s variable. Read error trace: \n", variable)
// 			fmt.Println(err)
// 			fmt.Printf("Please introduce again value of %s: ", variable)

// 		} else {
// 			well_read = true
// 		}
// 	}
// 	return

// }

func prompt_reader(variable string) (float_num float64) {
	var err error
	conversion_error := false
	for !conversion_error {
		fmt.Printf("Please introduce value of %s: ", variable)
		str := read_user_input()
		str = clean_input(str)
		float_num, err = convert_validator(variable, str)

		if err == nil {
			conversion_error = true
		}
	}
	return

}

func clean_input(input string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	result := re_leadclose_whtsp.ReplaceAllString(input, "")
	result = re_inside_whtsp.ReplaceAllString(result, " ")
	return result
}

func GenDisplaceFn(a, v_0, s_0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (a*(math.Pow(t, 2)*0.5) + (v_0 * t) + (s_0))
	}
	return fn
}

func print_setted_values(setted_values map[string]float64) {
	fmt.Println("Setted values are:")
	for key, value := range setted_values {
		fmt.Printf("%s: %f \n", key, value)
	}

}
func main() {
	variables_prompt := map[int]string{
		0: "acceleration => a",
		1: "initial velocity => v_0",
		2: "initial displacement => s_0",
		3: "time value => t",
	}

	variables_translator := map[string]string{
		"acceleration => a":           "a",
		"initial velocity => v_0":     "v_0",
		"initial displacement => s_0": "s_0",
		"time value => t":             "t",
	}

	variables := map[string]float64{
		"a":   0,
		"v_0": 0,
		"s_0": 0,
		"t":   0,
	}

	si := make([]int, 0, len(variables_prompt))
	for key, _ := range variables_prompt {
		si = append(si, key)
	}
	sort.Ints(si)
	for _, v := range si {
		value := variables_prompt[v]
		variables[variables_translator[value]] = prompt_reader(value)
	}

	print_setted_values(variables)
	time_func := GenDisplaceFn(variables["a"], variables["v_0"], variables["s_0"])
	fmt.Printf("Displacement of %f", time_func(variables["t"]))

}
