package isisomorphic

import (
	"strconv"
	"strings"
)

func isIsomorphic(s string, t string) bool {
	numMask := func(str string) string {
		arrstr := strings.Split(str, "")
		i := 0
		ret := ""
		var mp = make(map[string]string, len(str))
		for _, v := range arrstr {
			if value, found := mp[v]; found {
				ret += value + "|"
			} else {
				stri := strconv.Itoa(i)
				mp[v] = stri
				ret += stri + "|"
				i += 1
			}
		}
		return ret
	}
	ns := numMask(s)
	nt := numMask(t)
	return ns == nt
}
