// Declaring an Array

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var intArray [5]int
	var strArray [5]string

	fmt.Println(reflect.ValueOf(intArray).Kind())
	fmt.Println(reflect.ValueOf(strArray).Kind())
}

//------------------------------------------------------------------------------------------------------------------------
Assign and Access Values

package main

import "fmt"

func main() {
	var theArray [3]string
	theArray[0] = "India"  // Assign a value to the first element
	theArray[1] = "Canada" // Assign a value to the second element
	theArray[2] = "Japan"  // Assign a value to the third element

	fmt.Println(theArray[0]) // Access the first element value
	fmt.Println(theArray[1]) // Access the second element valu
	fmt.Println(theArray[2]) // Access the third element valu
}

//------------------------------------------------------------------------------------------------------------------------
Initializing an Array with an Array Literal

package main

import "fmt"

func main() {
	x := [5]int{10, 20, 30, 40, 50}   // Intialized with values
	var y [5]int = [5]int{10, 20, 30} // Partial assignment

	fmt.Println(x)
	fmt.Println(y)
}

//------------------------------------------------------------------------------------------------------------------------
Initializing an Array with ellipses...
When we use ... instead of specifying the length. The compiler can identify the length of an array, based on the elements specified in the array declaration.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := [...]int{10, 20, 30}

	fmt.Println(reflect.ValueOf(x).Kind())
	fmt.Println(len(x))
}

//------------------------------------------------------------------------------------------------------------------------
Initializing Values for Specific Elements
When an array declare using an array literal, values can be initialize for specific elements.

A value of 10 is assigned to the second element (index 1) and a value of 30 is assigned to the fourth element (index 3).

package main

import "fmt"

func main() {
	x := [5]int{1: 10, 3: 30}
	fmt.Println(x)
}



