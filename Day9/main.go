// Golang Program to Count no. of Numerical Digits in a String

package main

import "fmt"

func main(){
	mystr := "helloalexa234567playthesong"
   	fmt.Println("The string created here is:", mystr)
   	count := 0
	for _, char := range mystr{
		if char >= '0' && char <= '9'{
			count++
		}
	}
	fmt.Println("Number of numerical digits:", count)
}