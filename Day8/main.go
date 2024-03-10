package main

import (
	"fmt"
)

func NaturalSum(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}

func main() {
	var num int
	fmt.Print("Enter a positive Integer: ")
	fmt.Scanln(&num)
	fmt.Print("Sum: ", NaturalSum(num)) // Corrected function name
}
