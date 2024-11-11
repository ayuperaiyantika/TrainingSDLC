package main

import "testing"

// Integration test to check Add and Multiply together
func TestAddAndMultiplyIntegration(t *testing.T) {
	// Step 1: Add two numbers
	sum := Add(3, 4)

	// Step 2: Multiply the result of addition with another number
	result := Multiply(sum, 2)

	// Expected result: (3 + 4) * 2 = 14
	expected := 14

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("Integration test failed. Got %d; expected %d", result, expected)
	}
}
