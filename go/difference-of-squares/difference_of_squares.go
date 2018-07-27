package diffsquares

func Difference(num int) int {
    return SquareOfSums(num) - SumOfSquares(num)
}


func SumOfSquares(num int) int {
	total := 0
	for i := 1; i <= num; i++ {
		total += i*i
	}
	return total
}

func SquareOfSums(num int) int {
	total := 0
	for i := 1; i <= num; i++ {
		total += i
	}
	return total * total
}