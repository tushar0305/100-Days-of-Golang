package main

import (
	"github.com/tushar0305/100-Days-of-Golang/Day8/naturalsum"
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter a positive Integer: ")
	fmt.Scanln(&num)
	fmt.Print("Sum: ", naturalsum.NaturalSum(num)) // Corrected function name
}
