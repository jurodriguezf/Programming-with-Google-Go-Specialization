# Programming with Google Go Specialization

Repository for the Coursera's "Programming with Google Go" specialization.

## Course 1: Getting Started with Go

Learn the basics of Go, an open source programming language originally developed by a team at Google and enhanced by many contributors from the open source community. This course is designed for individuals with previous programming experience using such languages as C, Python, or Java, and covers the fundamental elements of Go. Topics include data types, protocols, formats, and writing code that incorporates RFCs and JSON. Most importantly, you’ll have a chance to practice writing Go programs and receive feedback from your peers. Upon completing this course, you'll be able to implement simple Go programs, which will prepare you for subsequent study at a more advanced level.

### Week 1
**Prompt (test.go):** Download and install the Go tools on your machine. Write a Go program to print “Hello, world!” on the screen. Compile and run the program.

### Week 2
**Prompt 1 (trunc.go):** Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.

**Prompt 2 (findian.go):** Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

### Week 3
**Prompt (slice.go):** Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.