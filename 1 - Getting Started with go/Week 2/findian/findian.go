package main

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
