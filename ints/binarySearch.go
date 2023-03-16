package ints

import "fmt"

func recursiveBinarySearch(haystack []int, p, q, needle int) int {
	if p > q {
		return -1
	}
	mid := (q + p) / 2
	if needle == haystack[mid] {
		return mid
	}
	if haystack[mid] < needle {
		return recursiveBinarySearch(haystack, mid+1, q, needle)
	}
	return recursiveBinarySearch(haystack, p, mid-1, needle)
}

func nonRecursiveBinarySearch(haystack []int, needle int) int {
	p, q := 0, len(haystack)-1

	for p <= q {
		mid := (p + q) / 2
		if haystack[mid] == needle {
			return mid
		}
		if haystack[mid] < needle {
			p = mid + 1
		} else {
			q = mid - 1
		}
	}
	return -1
}

func BinarySearch(haystack []int, needle int) int {
	fmt.Printf("If needle `%d` is in haystack `%v`, return its index.\n", needle, haystack)
	//	return recursiveBinarySearch(haystack, 0, len(haystack)-1, needle)
	return nonRecursiveBinarySearch(haystack, needle)
}
