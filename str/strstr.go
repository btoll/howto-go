package str

import "fmt"

//func sub(haystack, needle string) bool {
//	for i := 0; i < len(needle); i++ {
//		if haystack[i] != needle[i] {
//			return false
//		}
//	}
//	return true
//}

func sub(haystack, needle string) bool {
	i := 0
	for i < len(needle) {
		if haystack[i] != needle[i] {
			return false
		}
		i++
	}
	return true
}

func Strstr(haystack, needle string) int {
	fmt.Printf("Return the first occurrence of needle in `%s` in haystack `%s`.\n", haystack, needle)
	for i := 0; i < len(haystack); i++ {
		if len(needle)+i <= len(haystack) && needle[0] == haystack[i] {
			//			if sub(haystack[i:], needle) {
			if sub(haystack[i:i+len(needle)], needle) {
				return i
			}
		}
	}
	return -1
}
