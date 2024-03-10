// GO Program to Calculate Sum of Natural Numbers Using for.....Loop
package naturalsum

func natural_sum(a int) int{
	sum := 0
	for i := 1; i <= a; i++{
		sum += i
	}
	return sum
}