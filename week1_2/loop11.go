package main

import "fmt"

func main() {

	fmt.Println("looping with a 'for'")
	for i := 0; i < 5; i++ {
		fmt.Println("i is", i)
	}

	fmt.Println("looping with a 'goto'")

	j := 0
Start:
	if j < 5 {
		fmt.Println("j is", j)
		j++
		goto Start
	}
}
