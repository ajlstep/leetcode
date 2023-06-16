package romantointeger

import (
	"strconv"
	"strings"
)

// func romanToInt2(s string) int {
// 	sum := 0

// 	return sum
// }

func getRom(ch string) int {
	switch ch {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		if len(ch) == 2 {
			chArr := strings.Split(ch, "")
			switch chArr[0] {
			case "2":
				return 2 * getRom(chArr[1])
			case "3":
				return 3 * getRom(chArr[1])
			}
		} else {
			return 0
		}
	}
	return 0
}

func romanToInt(s string) int {
	sum := 0

	multiplierStr := func(i int) string {
		if i == 1 {
			return ""
		} else {
			return strconv.Itoa(i)
		}
	}

	combine := func(val1, val2 string) (int, bool) {
		v1 := getRom(val1)
		v2 := getRom(val2)
		if v2 > v1 && v1 != 0 && v2 != 0 {
			return v2 - v1, false
		} else {
			return v2 + v1, true
		}
	}

	val := strings.Split(s, "")
	var valComresed = make([]string, 0, len(val))
	numNow := ""
	multiplier := 1
	for _, v := range val {
		if numNow == "" {
			numNow = v
			continue
		}
		if numNow == v {
			multiplier += 1
		} else {
			valComresed = append(valComresed, multiplierStr(multiplier)+numNow)
			numNow = v
			multiplier = 1
		}
	}
	valComresed = append(valComresed, multiplierStr(multiplier)+numNow)

	if len(valComresed)%2 != 0 {
		valComresed = append([]string{"0"}, valComresed...)
	}

	for i := len(valComresed); i > 0; i -= 2 {
		nowSum, _ := combine(valComresed[i-2], valComresed[i-1])
		if i-3 >= 0 {
			_, nowIf := combine(valComresed[i-3], valComresed[i-2])
			if !nowIf {
				sum += getRom(valComresed[i-1])
				valComresed = append([]string{"0"}, valComresed...)
				i += 2
			} else {
				sum += nowSum
			}
		} else {
			sum += nowSum
		}
	}
	return sum
}

func romanToInt2(s string) int {
	// Definirea map-ului pentru a atribui valorile numerice romanei
	romanNumerals := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Inițializarea variabilelor
	result := 0
	prevValue := 0

	// Parcurgerea string-ului în ordine inversă
	for i := len(s) - 1; i >= 0; i-- {
		currentValue := romanNumerals[s[i]]

		// Verificarea regulilor de scădere și adunare
		if currentValue >= prevValue {
			result += currentValue
		} else {
			result -= currentValue
		}

		// Actualizarea valorii anterioare pentru următorul caracter
		prevValue = currentValue
	}

	return result
}
