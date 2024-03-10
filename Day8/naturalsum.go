package naturalsum

func NaturalSum(a int) int {
    sum := 0
    for i := 1; i <= a; i++ {
        sum += i
    }
    return sum
}