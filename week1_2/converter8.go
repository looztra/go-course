package main

import (
	"fmt"
	"github.com/looztra/go-course/ftoc"
)

func main() {

	var tempInFahrenheit float64
	fmt.Println("Please input a temp in Fahrenheit")
	fmt.Scanf("%f", &tempInFahrenheit)
	tempInCelcius := ftoc.FtoC(tempInFahrenheit)

	// do not use a function as we are not supposed to know how it works by now
	fmt.Printf("Converted in Celcius : %.2f", tempInCelcius)
}
