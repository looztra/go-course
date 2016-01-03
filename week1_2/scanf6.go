package main

import "fmt"

func main() {

	var input int
	fmt.Println("Please enter an int to be double")
	fmt.Scanf("%d", &input)

	output := input * 2

	fmt.Printf("Here is your input (%d) doubled: %d", input, output)

}
