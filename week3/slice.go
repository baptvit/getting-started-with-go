package main

import ("fmt"
	"strconv"
	"sort")

func main() {
	s := make([]int,3)
	for {
	fmt.Println("Please enter a int or exit with X.")
        var input string
        fmt.Scanln(&input)
        fmt.Println(input)
	i, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}else {
		s = append(s, i)
	}
	sort.Ints(s)
	fmt.Println("Sorted: ", s)
	if input == "X"{
		fmt.Println("Exit looping...")
		fmt.Println("SOrted: ", s)
		break
		}

	}
}
