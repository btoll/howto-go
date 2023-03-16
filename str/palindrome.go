package str

import "fmt"

func Palindrome(p string) bool {
	fmt.Printf("Is `%s` a palindrome?\n", p)
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		if p[i] != p[j] {
			return false
		}
	}
	return true
}
