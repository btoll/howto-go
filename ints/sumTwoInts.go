package ints

import "fmt"

func SumTwoInts(nums []int, n int) []int {
	fmt.Printf("What nums in `%v` sum to `%d`?\n", nums, n)
	m := map[int]bool{}
	for _, d := range nums {
		m[d] = true
	}
	for _, d := range nums {
		_, ok := m[n-d]
		if ok {
			return []int{d, n - d}
		}
	}
	return nil
}
