package main

import (
	"https://github.com/tushar0305/100-Days-of-Golang/Day8/naturalsum"
	"fmt"
)

func main(){
	var num int
	fmt.Print("Enter a positive Integer: ")
	fmt.Scanln(&num)
	fmt.Print("Sum: ", natural_sum(num))
}