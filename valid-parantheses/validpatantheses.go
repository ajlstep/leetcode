package validpatantheses

import "strings"

func isValid(s string) bool {
	valPharantese := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		")": -1,
		"]": -2,
		"}": -3,
	}

	summPharantese := 0
	var openedStack = make([]int, 0, len(s))
	for _, v := range strings.Split(s, "") {
		if valPharantese[v] > 0 {
			openedStack = append(openedStack, valPharantese[v])
		} else {
			if len(openedStack) == 0 {
				return false
			} else {
				if openedStack[len(openedStack)-1] != valPharantese[v]*-1 {
					return false
				} else {
					openedStack = remove(openedStack, len(openedStack)-1)
				}
			}
		}
		summPharantese += valPharantese[v]
	}
	return summPharantese == 0
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
