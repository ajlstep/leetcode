package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{3, 1, 4, 1, 1, 3, 5, 1, 2, 2}
	nums2 := []int{4, 1, 5, 2, 1, 1, 1, 5, 3, 1, 1, 1, 2, 3, 1, 4, 3, 5, 5, 3, 1, 2, 3, 2, 4, 1, 1, 1, 5, 3}

	fmt.Println(maxUncrossedLines(nums1, nums2))
}

type Bb struct {
	a     int
	b     int
	index int
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	ret := 0
	vecBb := make([]Bb, 0, len(nums1))
	for key, value := range nums1 {
		for key2, value2 := range nums2 {
			if value == value2 {
				vecBb = append(vecBb, Bb{
					a:     key,
					b:     key2,
					index: 0,
				})
			}
		}
	}

	for i := 0; i < len(vecBb); i++ {
		pNow := &vecBb[i]
		for j := 0; j < len(vecBb); j++ {
			if j == i {
				continue
			}
			if ifTrue(*pNow, vecBb[j]) {
				pNow.index += 1
			}
		}
	}

	sort.Slice(vecBb, func(i, j int) bool {
		return vecBb[i].index > vecBb[j].index
	})

	vecA := make([]int, 0)
	vecB := make([]int, 0)

	vecAddBb := make([]Bb, 0, len(vecBb))

	for key := range vecBb {
		if key == 0 {
			ret++
			vecAddBb = append(vecAddBb, vecBb[key])
			vecA = append(vecA, vecBb[key].a)
			vecB = append(vecB, vecBb[key].b)
			continue
		}
		if ifTrueC(vecBb[key], vecAddBb) && notUse(vecBb[key].a, vecA) && notUse(vecBb[key].b, vecB) {
			vecAddBb = append(vecAddBb, vecBb[key])
			vecA = append(vecA, vecBb[key].a)
			vecB = append(vecB, vecBb[key].b)
			ret++
		} else {
			continue
		}
	}
	return ret
}

func ifTrue(p1, p2 Bb) bool {
	s1 := (p1.a < p2.a && p1.b > p2.b)
	s2 := (p1.a > p2.a && p1.b < p2.b)
	s3 := s1 || s2
	s4 := !s3
	return s4
}

func ifTrueC(p1 Bb, p2 []Bb) bool {
	ret := true
	for _, v := range p2 {
		if !ifTrue(p1, v) {
			return false
		}
	}
	return ret
}

func notUse(i int, arr []int) bool {
	ret := true
	for _, v := range arr {
		if v == i {
			return false
		}
	}
	return ret
}
