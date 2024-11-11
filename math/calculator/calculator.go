package main

import "fmt"

// Function to add two numbers
func Add(a, b int) int {
	return a + b
}

// Function to multiply two numbers
func Multiply(a, b int) int {
	return a * b
}

// Main function to showcase usage
func main() {
	fmt.Println("Addition: ", Add(3, 4))
	fmt.Println("Multiplication: ", Multiply(3, 4))
}
