package main

import (
	"fmt"
	"os"
	"strconv"
)

func squareDigits(num int) int {
	numStr := strconv.Itoa(num)
	result := ""

	for _, digit := range numStr {
		d, _ := strconv.Atoi(string(digit))
		result += strconv.Itoa(d * d)
	}

	finalResult, _ := strconv.Atoi(result)
	return finalResult
}

func main() {
	var num int
	fmt.Fscan(os.Stdin, &num)
	fmt.Println(squareDigits(num))
}
