package climbingstars

import "math/big"

func climbStairs(n int) int {
	if n < 4 {
		return n
	}
	sum := 0
	numbersOfOne := n % 2
	for numbersOfTwo := n / 2; numbersOfTwo >= 0; numbersOfTwo -= 1 {
		sum += factOfThree(numbersOfOne, numbersOfTwo)
		numbersOfOne += 2
	}
	return sum
}

func factOfThree(numbersOfOne, numbersOfTwo int) int {
	a := factorial(numbersOfOne + numbersOfTwo)
	b := factorial(numbersOfOne)
	c := factorial(numbersOfTwo)
	mul := new(big.Int).Mul(b, c)
	div := new(big.Int).Div(a, mul)
	return int(div.Int64())
}

func factorial(n int) *big.Int {
	factVal := big.NewInt(1)
	if big.NewInt(int64(n)).Cmp(big.NewInt(2)) == -1 {
		return factVal
	} else {
		m := big.NewInt(0)
		for i := 1; i <= n; i++ {
			factVal = m.Mul(factVal, big.NewInt(int64(i)))
		}

	}
	return factVal
}
