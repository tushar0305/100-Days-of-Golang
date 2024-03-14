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
	a, p := rectangle(l, b)
	fmt.Println("The area of the rectangle is:", a)
	fmt.Println("The perimeter of the rectangle is:", p)
}

// Golang Passing Address to a Function
package main

import "fmt"

func update(a *int, b *string) {
	*a = *a + 5
	*b = "Hello" + *b

	return
}

func main() {
	var a int = 10
	var b string = "World"

	fmt.Println("Before : ", a, b)

	update(&a, &b)

	fmt.Println("After : ", a, b)
}

// GO Program to Swap Number Without Using Temporary Variables

package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Printf("Enter the first number: ")
	fmt.Scanln(&a)

	fmt.Printf("Enter the second number: ")
	fmt.Scanln(&b)

	a = a - b    
	b = a + b        
	a = b - a
	fmt.Printf("After Swapping: %d %d\n", a, b)
}
