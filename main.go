package main

import (
	"fmt"
	"math"
)

func isSquare(x int) bool {
	num := int(math.Sqrt(float64(x)))
	return num*num == x
}

func isCube(y int) bool {
	num2 := int(math.Round(math.Pow(float64(y), 1.0/3.0)))
	return num2*num2*num2 == y
}

func main() {
	var n int
	fmt.Println("Input Number: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		if isSquare(i) && isCube(i) {
			fmt.Println("SquareCube \n")
		} else if isSquare(i) {
			fmt.Println("Square \n")
		} else if isCube(i) {
			fmt.Println("Cube \n")
		} else {
			fmt.Println(i)
		}
	}
}
