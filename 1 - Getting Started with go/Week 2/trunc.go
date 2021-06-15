package main

import "fmt"

func main() {
	var num float32
	fmt.Printf("Please enter a float number: ")
	fmt.Scan(&num)

	// Casting the number works
	var truncNum = int(num)
	fmt.Printf("Truncated number is %d", truncNum)
}
