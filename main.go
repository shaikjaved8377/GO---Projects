package main

import (
	"fmt"

	"go-projetcs/calc"
)

func main() {
	fmt.Println("Hello, World!")
	a, b := 5, 3
	fmt.Println("Addition:", calc.Add(a, b))
	fmt.Println("Subtraction:", calc.Subtract(a, b))
	fmt.Println("Multiplication:", calc.Multiply(a, b))
	fmt.Println("Division:", calc.Divide(a, b))
	fmt.Println("Division by zero:", calc.Divide(a, 0))
}
