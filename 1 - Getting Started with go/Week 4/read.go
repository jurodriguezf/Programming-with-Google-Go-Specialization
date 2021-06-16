package main

/*
Author: Julio Rodriguez
Prompt: Write a program which reads information from a file and represents it
in a slice of structs. Assume that there is a text file which contains
a series of names. Each line of the text file has a first name and a last name,
in that order, separated by a single space on the line.

Your program will define a name struct which has two fields,
fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and
create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read
from the file, your program will have a slice containing one struct for each line
in the file. After reading all lines from the file, your program should iterate
through your slice of structs and print the first and last names found in each
struct.
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type names struct {
	fname string
	lname string
}

func main() {
	var fileName string
	fmt.Print("Please enter the name of the text file (must be in same folder): ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		panic("Error reading the file")
	}

	scanner := bufio.NewScanner(file)

	dataSlice := make([]names, 0)

	for scanner.Scan() {
		// Scanning each line
		currentLine := scanner.Text()
		// Spliting line with whitespaces
		separatedText := strings.Split(currentLine, " ")
		// Generating a struct with name and last name
		newStruct := names{fname: separatedText[0], lname: separatedText[1]}
		// Appending the struct into dataSlice
		dataSlice = append(dataSlice, newStruct)
	}

	// Printing the contents of the slice
	for _, s := range dataSlice {
		fmt.Printf("Name: %s. Last name: %s.\n", s.fname, s.lname)
	}
}
