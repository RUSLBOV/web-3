package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	result := ""

	for i, char := range input {
		result += string(char)
		if i < len(input)-1 {
			result += "*"
		}
	}

	fmt.Println(result)
}
