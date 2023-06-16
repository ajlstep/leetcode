package addbinary

import (
	"strconv"
)

func addBinary2(a string, b string) string {
	ab, _ := strconv.ParseUint(a, 2, 64)
	bb, _ := strconv.ParseUint(b, 2, 64)
	rez := ab + bb
	str := strconv.FormatUint(uint64(rez), 2)
	return str
}

func addBinary(a string, b string) string {
	getO := func(s int) string {
		ret := ""
		for i := 0; i < s; i++ {
			ret += "0"
		}
		return ret
	}

	lena := len(a)
	lenb := len(b)
	if lena > lenb {
		b = getO(lena-lenb) + b
	} else {
		a = getO(lenb-lena) + a
	}

	sum := ""
	carry := 0
	for i := len(a) - 1; i >= 0; i-- {
		digitA := int(a[i] - '0')
		digitB := int(b[i] - '0')

		// Suma cifrelor și a cifrei de transport
		digitSum := digitA + digitB + carry

		// Calcularea cifrei rezultat și a noii cifre de transport
		if digitSum == 0 || digitSum == 1 {
			sum = strconv.Itoa(digitSum) + sum
			carry = 0
		} else if digitSum == 2 {
			sum = "0" + sum
			carry = 1
		} else if digitSum == 3 {
			sum = "1" + sum
			carry = 1
		}
	}

	// Adăugarea cifrei de transport finală, dacă există
	if carry == 1 {
		sum = "1" + sum
	}

	return sum
}

// func addBinary(a string, b string) string {
// 	getO := func(s int) string {
// 		ret := ""
// 		for i := 0; i < s; i++ {
// 			ret += "0"
// 		}
// 		return ret
// 	}
// 	a, b = func(a, b string) (string, string) {
// 		lena := len(a)
// 		lenb := len(b)
// 		if lena > lenb {
// 			b = getO(lena-lenb) + b
// 		} else {
// 			a = getO(lenb-lena) + a
// 		}
// 		return a, b
// 	}(a, b)

// 	summ := func(a, b, rest string) (string, string) {
// 		switch a + b + rest {
// 		case "000":
// 			return "0", "0"
// 		case "001":
// 			return "1", "0"
// 		case "010":
// 			return "1", "0"
// 		case "011":
// 			return "1", "0"
// 		case "100":
// 			return "1", "0"
// 		case "101":
// 			return "1", "0"
// 		case "110":
// 			return "0", "1"
// 		case "111":
// 			return "1", "1"
// 		}
// 		return "", ""
// 	}
// 	aa := strings.Split(a, "")
// 	bb := strings.Split(b, "")
// 	ret := ""
// 	rest := "0"
// 	for i := len(aa) - 1; i >= 0; i-- {
// 		var sm string
// 		sm, rest = summ(aa[i], bb[i], rest)
// 		ret = sm + ret
// 	}
// 	if rest == "1" {
// 		ret = "1" + ret
// 	}
// 	return ret
// }
