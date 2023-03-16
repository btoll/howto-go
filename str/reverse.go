package str

import (
	"fmt"
	"strings"
)

func swap(s []string, i, j int) {
	a := s[i]
	s[i] = s[j]
	s[j] = a
}

func Reverse(p string) string {
	fmt.Printf("Reverse string `%s`.\n", p)
	a := strings.Split(p, "")
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		swap(a, i, j)
	}
	return strings.Join(a, "")
}
