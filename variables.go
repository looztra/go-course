package main

import "fmt"

func main() {
	var aString string
	var anInt int
	var aBoolean bool

	fmt.Println("var aString valued to", aString)
	fmt.Println("var anInt valued to", anInt)
	fmt.Println("var aBoolean valued to", aBoolean)

	anotherString := "I am a string"
	anotherInt := 733
	anotherBoolean := true

	fmt.Println("var anotherString valued to", anotherString)
	fmt.Println("var anotherInt valued to", anotherInt)
	fmt.Println("var anotherBoolean valued to", anotherBoolean)

	piAsFloat := float32(3.14)

	fmt.Printf("My pi is %.2f", piAsFloat)
}
