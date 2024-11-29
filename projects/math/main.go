package main

import (
	"fmt"
	"math"
)

var k, p, v float64

func M() float64 {
	// Масса m = p * v
	return p * v
}

func W() float64 {
	// Циклическая частота w = sqrt(k / m)
	return math.Sqrt(k / M())
}

func T() float64 {
	// Период колебаний t = 6 / w
	return 6 / W()
}
func main() {
	fmt.Print("Введите значение k (жесткость пружины): ")
	fmt.Scan(&k)

	fmt.Print("Введите значение p (плотность): ")
	fmt.Scan(&p)

	fmt.Print("Введите значение v (объем): ")
	fmt.Scan(&v)

	fmt.Println("Период колебаний t =", T())
}
