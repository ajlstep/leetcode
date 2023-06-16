package longestCommonPrefix

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	getCommon := func(v1 []string, v2 []string) int {
		com := 0
		for i := 0; i < min(len(v1), len(v2)); i++ {
			if v1[i] == v2[i] {
				com += 1
			} else {
				break
			}
		}
		return com
	}
	if len(strs) == 0 {
		return ""
	}
	str := strings.Split(strs[0], "")
	common := len(str) + 1
	for i := 1; i < len(strs); i++ {
		common = min(getCommon(str, strings.Split(strs[i], "")), common)
	}
	if common == len(str)+1 {
		return strs[0]
	} else {
		return strings.Join(str[0:common], "")
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
