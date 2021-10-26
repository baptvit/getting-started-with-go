package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputNum float64
	fmt.Print("Enter a floating point number ")
	fmt.Scan(&inputNum)
	inputNumString := fmt.Sprintf("%f", inputNum)
	leftSide := strings.Split(inputNumString, ".")[0]
	fmt.Println(leftSide)
}
