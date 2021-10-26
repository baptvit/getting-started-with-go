package main

import (
    "fmt"
    "strings"
)

func main(){
	var word_str string

	fmt.Println("Please enter a word")
	fmt.Scan(&word_str)

	switch {
	case strings.Contains(strings.ToLower(word_str), "i") && strings.Contains(strings.ToLower(word_str), "a") && strings.Contains(strings.ToLower(word_str), "n"):
		if strings.Index(strings.ToLower(word_str), "i") == 0 && strings.ToLower(word_str[len(word_str)-1:]) == "n" {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	default:
		fmt.Println("Not Found!")
	}
}
