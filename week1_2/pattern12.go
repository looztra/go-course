package main

import "fmt"

func main() {

	const chain = "O"
	output := chain

	for i := 0; i < 10; i++ {
		fmt.Println(output)
		output += chain
	}

}
