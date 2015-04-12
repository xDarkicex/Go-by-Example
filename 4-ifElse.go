//Branching with if and else in Go is straight-forward.

package main
import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	//You can have an if statement without an else.
	}

	if number := 9; number < 0 {
		fmt.Println(number, "is negative")
	} else if number < 10 {
		fmt.Println(number, "has 1 digit")
	} else {
		fmt.Println(number, "has multiple digits")
	}
}