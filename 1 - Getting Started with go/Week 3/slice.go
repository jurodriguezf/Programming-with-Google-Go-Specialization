package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	exitCode string = "X"
	listSize int    = 3
)

func main() {
	sli := make([]int, listSize)
	var input string
	index := 0

	for {
		fmt.Print("\nPlease enter an integer (or exit typing 'X'): ")
		fmt.Scan(&input)

		// It was easier checking first if the input was 'X', rather than handling multiple
		// variable types. Then I would cast the string to an integer

		if input == exitCode {
			fmt.Print("\nGoodbye!")
			os.Exit(0)
		}

		// If the index is <=2, I can't use the append() function, as the slice would
		// append the number at index [3], leaving 3 elements containing 0 (that is [0 0 0 newNumber]).
		// I also can't simply do slice[n] = newNumber, as the slice gets sorted at each iteration.
		// So, the solution I came up with was to modify the slice (adding the number), creating
		// a copy and sorting that copy, so that I can continue using an index to modify the slice.

		num, _ := strconv.Atoi(input)
		if index <= 2 {
			sli[index] = num
			sliCopy := make([]int, listSize)
			copy(sliCopy, sli)
			sort.Ints(sliCopy)
			fmt.Printf("%v", sliCopy)
			index++
		} else {
			sli = append(sli, num)
			sort.Ints(sli)
			fmt.Printf("%v", sli)
		}
	}
}
