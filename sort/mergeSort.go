package sort

import "fmt"

func merge(l1, l2 []int) []int {
	fmt.Println("l1, l2", l1, l2)
	i, j := 0, 0
	merged := []int{}
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			merged = append(merged, l1[i])
			i += 1
		} else {
			merged = append(merged, l2[j])
			j += 1
		}
	}
	for i < len(l1) {
		merged = append(merged, l1[i])
		i += 1
	}
	for j < len(l2) {
		merged = append(merged, l2[j])
		j += 1
	}
	return merged
}

func MergeSort(nums []int) []int {
	//	fmt.Println(merge([]int{1, 4, 6, 8}, []int{0, 2, 7, 9}))
	if len(nums) == 1 {
		return nums
	}
	mid := int(len(nums) / 2)
	l1 := MergeSort(nums[:mid])
	l2 := MergeSort(nums[mid:])

	// []int{4, 1, 8, 6}

	// {4}
	// {4, 1}

	// {8}
	// {8, 6}
	return merge(l1, l2)
}
