package diffsquares

// SquareOfSums
// @link: https://en.wikipedia.org/wiki/1_%2B_2_%2B_3_%2B_4_%2B_%E2%8B%AF.
func SquareOfSums(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares
// @link http://www.trans4mind.com/personal_development/mathematics/series/sumNaturalSquares.htm.
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
