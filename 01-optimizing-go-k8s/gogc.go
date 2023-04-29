package main

import (
	"fmt"
	"os"
)

func main() {
	// Get current GOGC value
	gogc := os.Getenv("GOGC")
	fmt.Println("Current GOGC value:", gogc);

	// Set GOGC value to 50
	os.Setenv("GOGC", "50")

	// Run your Golang application
	// ...
}