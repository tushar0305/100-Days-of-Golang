// Remove Duplicate Elements From An Array Using Loops

/*package main
import "fmt"
func main() {
   arr := []int{1, 2, 2, 4, 4, 5, 7, 5}
   fmt.Println("The unsorted array entered is:", arr)
   size := len(arr)
   for i := 0; i < size; i++ {
      for j := i + 1; j < size; j++ {
         if arr[i] == arr[j] {
            for k := j; k < size-1; k++ {
               arr[k] = arr[k+1]
            }
            size--
            j--
         }
      }
   }
   fmt.Println("The elements of array obtained after removing the duplicate values is:")
   for i := 0; i < size; i++ {
      fmt.Println(arr[i])
   }
}*/

//Remove Duplicate Values From An Array Of Strings Using Append()

package main

import (
	"fmt"
   "sort"
)

func removeDuplicate(strs []string) []string {
   sort.Strings(strs)
   for i := len(strs) - 1; i > 0; i-- {
      if strs[i] == strs[i-1] {
         strs = append(strs[:i], strs[i+1:]...)
      }
   }
   return strs
}

func main() {
	arr := []string{"a", "b", "b", "c", "b", "d"}
	fmt.Println("The unsorted array entered is:", arr)
	result := removeDuplicate(arr)
	fmt.Println("The array obtained after removing the duplicate values is:", result)

}
