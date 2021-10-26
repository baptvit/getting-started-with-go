package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

var first_char = "I"
var middle_char_array = []string {"A"}
var last_char = "N"

func main() {

	// var input_string string
	fmt.Println("Welcome! This program will search your input for Ian. Ian could be hiding anywhere, in lowercase or uppercase. It's my job to find Ian and show him to you.\nLet me show you.")
	fmt.Printf("Enter a string to be searched: ")
	
	// Read from the console. The fmt.Scan function doesn't work with spaces.
 	consoleReader := bufio.NewReader(os.Stdin)
 	input_string, _ := consoleReader.ReadString('\n')

	// Find out if the first character matches
	if strings.Index(input_string, strings.ToUpper(first_char)) != 0 &&
	   strings.Index(input_string, strings.ToLower(first_char)) != 0 {
	   	print_failure()
	}

	// Get the last index of the input string
	last_index := len(input_string) - 2

	// Find out if the last character matches the last index
	if strings.Index(input_string, strings.ToUpper(last_char)) != last_index &&
	   strings.Index(input_string, strings.ToLower(last_char)) != last_index {
	   	print_failure()
	}

	// Find out if any of the middle characters are contained in the input string
	for i := 0; i < len(middle_char_array); i++ {

		middle_char := middle_char_array[i]

		if strings.Contains(input_string, strings.ToUpper(middle_char)) ||
		   strings.Contains(input_string, strings.ToLower(middle_char)) {
			print_success()
		}
	}

	// We didn't get a match for any of our middle characters, so print failure.
	print_failure()

}

func print_success() {
	fmt.Println("Found!")
	os.Exit(0)
}

func print_failure() {
	fmt.Println("Not Found!")
	os.Exit(0)
}