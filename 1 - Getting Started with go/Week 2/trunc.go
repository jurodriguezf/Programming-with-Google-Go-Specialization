package main

/*
Author: Julio Rodriguez
Prompt: Write a program which prompts the user to enter
a floating point number and prints the integer which is
a truncated version of the floating point number that was
entered. Truncation is the process of removing the digits
to the right of the decimal place.
*/

import "fmt"

func main() {
	var num float32
	fmt.Printf("Please enter a float number: ")
	fmt.Scan(&num)

	// Casting the number works
	var truncNum = int(num)
	fmt.Printf("Truncated number is %d", truncNum)
}
