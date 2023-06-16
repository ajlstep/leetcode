package polindromenumbers

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x > -1 && x < 10 {
		return true
	}
	xStr := strconv.Itoa(x)
	ret := true
	for i := 0; i < len(xStr)/2; i++ {
		if xStr[i] != xStr[len(xStr)-(i+1)] {
			ret = false
			break
		}
	}
	return ret
}
