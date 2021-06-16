package main

/*
Author: Julio Rodriguez
Prompt: Write a program which prompts the user to first enter a name,
and then enter an address. Your program should create a map and add
the name and address to the map using the keys “name” and “address”,
respectively. Your program should use Marshal() to create a JSON object
from the map, and then your program should print the JSON object.
*/

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string

	fmt.Print("Please enter a name: ")
	fmt.Scan(&name)

	fmt.Print("Please enter an address: ")
	fmt.Scan(&address)

	infoMap := map[string]string{"name": name, "address": address}

	barr, _ := json.Marshal(infoMap)
	fmt.Printf("\n%s", barr)
}
