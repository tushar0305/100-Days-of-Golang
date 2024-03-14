// Golang Program to convert a string into Uppercase

package main
import (
   "fmt"
   "unicode"
)

func ToUpper(s string) string {
   a := []rune(s)
   for i, c := range a {
      if unicode.IsLower(c) {
         a[i] = unicode.ToUpper(c)
      }
   }
   return string(a)
}

func main() {
   var input string = "convert this string to uppercase"
   fmt.Println("The given string is:\n", input)
   var output string = ToUpper(input)
   fmt.Println()
   fmt.Println("The string obtained by converting above string to uppercase is:\n", output)
}