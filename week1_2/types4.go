package main

import "fmt"

type Counter int

func main() {
	var counter Counter
	fmt.Printf("counter value (of type %T) is %d", counter, counter)
}