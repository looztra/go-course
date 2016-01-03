package main
import "fmt"

func main() {

	const first = 722
	const second int = 1024

	const a, b = 44, 365
	//result := 44 * 365
	result := a * b

	fmt.Printf("const a is %d, const b is %d, multiply them together and you get %d\n", a, b, result)
}
