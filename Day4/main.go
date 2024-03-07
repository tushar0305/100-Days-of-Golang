// Golang Functions Returning Multiple Values
package main

import "fmt"

func rectangle(l int, b int) (area int, para int) {
	area = l * b
	para = 2 * (l + b)
	return
}

func main() {
	var l int
	var b int
	fmt.Print("Enter the length and breadth of the rectangle: ")
	fmt.Scan(&l, &b)
	a, p := rectangle(10, 20)
	fmt.Println("The area of the rectangle is:", a)
	fmt.Println("The perimeter of the rectangle is:", p)
}
