// Initialize Slice with values using a Slice Literal

package main

import "fmt"

func main(){
	var intSlice = []int{10, 20, 30, 40, 50}
	var strSlice = []string{"India", "Israel", "Russia"}

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
}

