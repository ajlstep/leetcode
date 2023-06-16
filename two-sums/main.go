package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

type Indexes struct {
	index int
	rest  int
}

func twoSum(nums []int, target int) []int {
	var ret = make([]int, 0, 2)
	var ind = make([]Indexes, 0, len(nums))
	for key, val := range nums {
		resp := find(ind, val)
		if resp != -1 {
			return []int{resp, key}
		}
		ind = append(ind, Indexes{
			index: key,
			rest:  target - val,
		})
	}
	return ret
}

func find(ind []Indexes, val int) int {
	for _, v := range ind {
		if v.rest == val {
			return v.index
		}
	}
	return -1
}
