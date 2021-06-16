package main

/*
Author: Julio Rodriguez
Prompt: Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters
‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered
string starts with the character ‘i’, ends with the character ‘n’,
and contains the character ‘a’. The program should print “Not Found!”
otherwise. The program should not be case-sensitive, so it does not
matter if the characters are upper-case or lower-case.
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Please enter a string: ")

	scanner := bufio.NewReader(os.Stdin)
	input, _ := scanner.ReadString('\n')

	text := strings.ToLower(input)
	text = strings.TrimSpace(text)

	if strings.HasPrefix(text, "i") && strings.HasSuffix(text, "n") && strings.Contains(text, "a") {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}
