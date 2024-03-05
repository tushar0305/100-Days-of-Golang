// Sort An Array Of Integers Using User-Defined Functions
package main

import "fmt"

func SelectionSort(arr []int) []int{
	for i:= 0; i < len(arr); i++{
		for j := i+1; j < len(arr); j++{
			if arr[i] > arr[j]{
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}

func BubbleSort(arr []int) []int{

	for i:=0; i<=len(arr)-1; i++{
	   for j:=0; j<len(arr)-1-i; j++{
		  if arr[j]> arr[j+1]{
			 arr[j], arr[j+1] = arr[j+1], arr[j]
		  }
	   }
	}
	return arr
 }

func main(){
	arr := []int{50, 30, 20, 10, 40, 5, 2, 100}
   	fmt.Println("The unsorted array entered is:", arr)
   	result := BubbleSort(arr)
	fmt.Println("The sorted array is:", result)
	fmt.Println()
}

//--------------------------------------------------------------------------------------------------------------------------------

// GO Program to Generate Fibonacci Sequence Up to a Certain Number

package main

import "fmt"

func main(){
	var n int

	t1 := 0
	t2 := 1
	nextTerm := 0

	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&n)

	fmt.Println("Fibnocci Series is : ")
	for i := 1; i < n; i++{
		if i == 1{
			fmt.Println(t1)
		}
		if i == 2{
			fmt.Println(t2)
		}
		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
		fmt.Println(nextTerm)
	}
}




