package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	maxDigit := '0'

	for _, char := range input {
		if char > maxDigit {
			maxDigit = char
		}
	}

	fmt.Println(string(maxDigit))
}
